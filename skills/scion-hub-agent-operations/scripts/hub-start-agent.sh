#!/bin/bash
# Start a Scion agent through Hub mode and subscribe to status notifications.

set -euo pipefail

if [ $# -lt 2 ]; then
    echo "Usage: hub-start-agent.sh <agent-name> <task> [scion flags...]"
    exit 1
fi

NAME="$1"
shift
TASK="$1"
shift

scion --non-interactive start "$NAME" "$TASK" --notify "$@"
