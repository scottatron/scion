/*
Copyright 2025 The Scion Authors.
*/

package dialects

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/scion/pkg/sciontool/hooks"
)

// CodexDialect parses Codex structured hook payloads and the legacy notify
// payload that older Codex builds emitted on turn completion.
type CodexDialect struct{}

func NewCodexDialect() *CodexDialect {
	return &CodexDialect{}
}

func (d *CodexDialect) Name() string {
	return "codex"
}

func (d *CodexDialect) Parse(data map[string]interface{}) (*hooks.Event, error) {
	rawName := getString(data, "type")
	if rawName == "" {
		rawName = getString(data, "event")
	}
	if rawName == "" {
		rawName = getString(data, "hook_event_name")
	}

	event := &hooks.Event{
		Name:    d.normalizeEventName(rawName),
		RawName: rawName,
		Dialect: "codex",
		Data: hooks.EventData{
			Prompt:    getString(data, "prompt"),
			Message:   firstNonEmptyString(getString(data, "title"), getString(data, "message"), getString(data, "last_assistant_message")),
			Reason:    firstNonEmptyString(getString(data, "stopReason"), getString(data, "reason")),
			Source:    getString(data, "source"),
			ToolName:  getString(data, "tool_name"),
			SessionID: getString(data, "session_id"),
			Raw:       data,
		},
	}

	// Extract tool input/output if available.
	var toolInputObject map[string]interface{}
	if val, ok := data["tool_input"]; ok {
		if str, ok := val.(string); ok {
			event.Data.ToolInput = str
		} else if m, ok := val.(map[string]interface{}); ok {
			toolInputObject = m
			event.Data.ToolInput = firstNonEmptyString(getString(m, "command"), jsonString(val))
		}
	}
	for _, key := range []string{"tool_output", "tool_response"} {
		val, ok := data[key]
		if !ok {
			continue
		}
		if str, ok := val.(string); ok {
			event.Data.ToolOutput = str
			break
		}
		event.Data.ToolOutput = jsonString(val)
		break
	}

	// Extract status fields
	if val, ok := data["success"]; ok {
		if b, ok := val.(bool); ok {
			event.Data.Success = b
		}
	}
	if val, ok := data["error"]; ok {
		if str, ok := val.(string); ok {
			event.Data.Error = str
		}
	}

	if rawName == "PermissionRequest" && event.Data.Message == "" {
		event.Data.Message = permissionRequestMessage(event.Data.ToolName, toolInputObject)
	}

	// Extract token usage
	extractTokens(data, &event.Data)

	// Extract file_path from tool_input/tool_response objects
	extractFilePath(data, &event.Data)

	return event, nil
}

func (d *CodexDialect) normalizeEventName(name string) string {
	switch name {
	case "agent-turn-complete":
		return hooks.EventResponseComplete
	case "notification":
		return hooks.EventNotification
	case "session-start", "SessionStart":
		return hooks.EventSessionStart
	case "session-end", "SessionEnd":
		return hooks.EventSessionEnd
	case "PermissionRequest":
		return hooks.EventNotification
	case "UserPromptSubmit":
		return hooks.EventPromptSubmit
	case "Stop":
		return hooks.EventResponseComplete
	case "PreToolUse":
		return hooks.EventToolStart
	case "PostToolUse":
		return hooks.EventToolEnd
	case "tool-start", "BeforeTool":
		return hooks.EventToolStart
	case "tool-end", "AfterTool":
		return hooks.EventToolEnd
	case "model-start", "BeforeModel":
		return hooks.EventModelStart
	case "model-end", "AfterModel":
		return hooks.EventModelEnd
	default:
		return name
	}
}

func firstNonEmptyString(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func jsonString(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(data)
}

func permissionRequestMessage(toolName string, toolInput map[string]interface{}) string {
	description := strings.TrimSpace(getString(toolInput, "description"))
	switch {
	case toolName != "" && description != "":
		return fmt.Sprintf("Permission requested for %s: %s", toolName, description)
	case toolName != "":
		return fmt.Sprintf("Permission requested for %s", toolName)
	case description != "":
		return fmt.Sprintf("Permission requested: %s", description)
	default:
		return "Permission requested"
	}
}
