#!/bin/bash
# Send a Hub-mode Scion message and subscribe to follow-up notifications.

set -euo pipefail

if [ $# -lt 2 ]; then
    echo "Usage: hub-message-agent.sh <agent-name-or-id> <message> [scion flags...]"
    exit 1
fi

AGENT="$1"
shift
MESSAGE="$1"
shift

scion --non-interactive message "$AGENT" "$MESSAGE" --notify "$@"
