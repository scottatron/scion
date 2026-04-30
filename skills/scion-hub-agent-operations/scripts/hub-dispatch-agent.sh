#!/bin/bash
# Dispatch a Scion Hub agent directly through the Hub API without using the
# caller cwd as a local workspace.

set -euo pipefail

if [ $# -lt 2 ] || [ $# -gt 3 ]; then
    echo "Usage: hub-dispatch-agent.sh <agent-name> <task> [harness-config]"
    echo "Required env: SCION_HUB_ENDPOINT and SCION_GROVE_ID or SCION_TARGET_GROVE_ID"
    echo "Optional env: SCION_AGENT_TOKEN, SCION_TOKEN_FILE, SCION_NOTIFY, SCION_GATHER_ENV"
    exit 1
fi

NAME="$1"
TASK="$2"
HARNESS="${3:-}"

ENDPOINT="${SCION_HUB_ENDPOINT:-}"
GROVE_ID="${SCION_TARGET_GROVE_ID:-${SCION_GROVE_ID:-}}"
TOKEN_FILE="${SCION_TOKEN_FILE:-$HOME/.scion/scion-token}"
NOTIFY="${SCION_NOTIFY:-true}"
GATHER_ENV="${SCION_GATHER_ENV:-true}"

if [ -z "$ENDPOINT" ]; then
    echo "SCION_HUB_ENDPOINT is required" >&2
    exit 1
fi

if [ -z "$GROVE_ID" ]; then
    echo "SCION_GROVE_ID or SCION_TARGET_GROVE_ID is required" >&2
    exit 1
fi

case "$NOTIFY" in
    true|false) ;;
    *) echo "SCION_NOTIFY must be true or false" >&2; exit 1 ;;
esac

case "$GATHER_ENV" in
    true|false) ;;
    *) echo "SCION_GATHER_ENV must be true or false" >&2; exit 1 ;;
esac

if [ -n "${SCION_AGENT_TOKEN:-}" ]; then
    TOKEN="$SCION_AGENT_TOKEN"
else
    TOKEN="$(cat "$TOKEN_FILE")"
fi

ENDPOINT="${ENDPOINT%/}"

jq -n \
    --arg name "$NAME" \
    --arg task "$TASK" \
    --arg harness "$HARNESS" \
    --argjson notify "$NOTIFY" \
    --argjson gatherEnv "$GATHER_ENV" \
    '{
        name: $name,
        task: $task,
        notify: $notify,
        gatherEnv: $gatherEnv
    } + (if $harness == "" then {} else {harnessConfig: $harness} end)' |
    curl -fsS \
        -H "Content-Type: application/json" \
        -H "X-Scion-Agent-Token: $TOKEN" \
        -X POST "$ENDPOINT/api/v1/groves/$GROVE_ID/agents" \
        --data-binary @-
