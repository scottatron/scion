/*
Copyright 2025 The Scion Authors.
*/

package dialects

import (
	"testing"

	"github.com/GoogleCloudPlatform/scion/pkg/sciontool/hooks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCodexDialectParse_TurnComplete(t *testing.T) {
	d := NewCodexDialect()
	event, err := d.Parse(map[string]interface{}{
		"type":  "agent-turn-complete",
		"title": "Done",
	})
	require.NoError(t, err)
	assert.Equal(t, hooks.EventResponseComplete, event.Name)
	assert.Equal(t, "Done", event.Data.Message)
	assert.Equal(t, "codex", event.Dialect)
}

func TestCodexDialectParse_FallbackEventField(t *testing.T) {
	d := NewCodexDialect()
	event, err := d.Parse(map[string]interface{}{
		"event":   "notification",
		"message": "Need approval",
	})
	require.NoError(t, err)
	assert.Equal(t, hooks.EventNotification, event.Name)
	assert.Equal(t, "Need approval", event.Data.Message)
}

func TestCodexDialect_EventMappings(t *testing.T) {
	d := NewCodexDialect()

	tests := []struct {
		rawName  string
		wantName string
	}{
		{"PreToolUse", hooks.EventToolStart},
		{"PermissionRequest", hooks.EventNotification},
		{"PostToolUse", hooks.EventToolEnd},
		{"UserPromptSubmit", hooks.EventPromptSubmit},
		{"Stop", hooks.EventResponseComplete},
		{"tool-start", hooks.EventToolStart},
		{"tool-end", hooks.EventToolEnd},
		{"model-start", hooks.EventModelStart},
		{"model-end", hooks.EventModelEnd},
		{"session-start", hooks.EventSessionStart},
		{"session-end", hooks.EventSessionEnd},
	}

	for _, tt := range tests {
		t.Run(tt.rawName, func(t *testing.T) {
			event, err := d.Parse(map[string]interface{}{"type": tt.rawName})
			require.NoError(t, err)
			assert.Equal(t, tt.wantName, event.Name)
		})
	}
}

func TestCodexDialect_TokenExtraction(t *testing.T) {
	d := NewCodexDialect()
	event, err := d.Parse(map[string]interface{}{
		"type": "model-end",
		"usage": map[string]interface{}{
			"prompt_tokens":     float64(800),
			"completion_tokens": float64(200),
		},
	})
	require.NoError(t, err)
	assert.Equal(t, int64(800), event.Data.InputTokens)
	assert.Equal(t, int64(200), event.Data.OutputTokens)
}

func TestCodexDialect_StatusFields(t *testing.T) {
	d := NewCodexDialect()
	event, err := d.Parse(map[string]interface{}{
		"type":    "tool-end",
		"success": true,
		"error":   "something failed",
	})
	require.NoError(t, err)
	assert.True(t, event.Data.Success)
	assert.Equal(t, "something failed", event.Data.Error)
}

func TestCodexDialect_CurrentHookPayloadShape(t *testing.T) {
	d := NewCodexDialect()
	event, err := d.Parse(map[string]interface{}{
		"hook_event_name": "PostToolUse",
		"session_id":      "sess-123",
		"tool_name":       "Bash",
		"tool_input": map[string]interface{}{
			"command": "git status --short",
		},
		"tool_response": map[string]interface{}{
			"exit_code": float64(0),
			"stdout":    " M file.txt",
		},
	})
	require.NoError(t, err)
	assert.Equal(t, hooks.EventToolEnd, event.Name)
	assert.Equal(t, "sess-123", event.Data.SessionID)
	assert.Equal(t, "Bash", event.Data.ToolName)
	assert.Equal(t, "git status --short", event.Data.ToolInput)
	assert.Contains(t, event.Data.ToolOutput, "\"exit_code\":0")
}

func TestCodexDialect_UserPromptAndStopFields(t *testing.T) {
	d := NewCodexDialect()

	promptEvent, err := d.Parse(map[string]interface{}{
		"hook_event_name": "UserPromptSubmit",
		"prompt":          "please investigate",
	})
	require.NoError(t, err)
	assert.Equal(t, hooks.EventPromptSubmit, promptEvent.Name)
	assert.Equal(t, "please investigate", promptEvent.Data.Prompt)

	stopEvent, err := d.Parse(map[string]interface{}{
		"hook_event_name":        "Stop",
		"last_assistant_message": "All checks passed",
	})
	require.NoError(t, err)
	assert.Equal(t, hooks.EventResponseComplete, stopEvent.Name)
	assert.Equal(t, "All checks passed", stopEvent.Data.Message)
}

func TestCodexDialect_PermissionRequestFields(t *testing.T) {
	d := NewCodexDialect()

	event, err := d.Parse(map[string]interface{}{
		"hook_event_name": "PermissionRequest",
		"session_id":      "sess-456",
		"tool_name":       "exec_command",
		"tool_input": map[string]interface{}{
			"command":     "git push",
			"description": "Push the stacked branch",
		},
	})
	require.NoError(t, err)
	assert.Equal(t, hooks.EventNotification, event.Name)
	assert.Equal(t, "sess-456", event.Data.SessionID)
	assert.Equal(t, "exec_command", event.Data.ToolName)
	assert.Equal(t, "git push", event.Data.ToolInput)
	assert.Equal(t, "Permission requested for exec_command: Push the stacked branch", event.Data.Message)
}
