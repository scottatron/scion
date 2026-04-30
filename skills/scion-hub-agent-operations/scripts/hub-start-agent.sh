#!/bin/bash
# Start a Scion agent through Hub mode and subscribe to status notifications.
#
# By default this runs from $HOME so tool-managed cwd values such as /opt/data
# do not trigger local workspace bootstrap before the Hub agent is created. Set
# SCION_HUB_PRESERVE_CWD=1 to intentionally keep the caller cwd, or set
# SCION_HUB_START_CWD to a specific safe directory.

set -euo pipefail

if [ $# -lt 2 ]; then
    echo "Usage: hub-start-agent.sh <agent-name> <task> [scion flags...]"
    echo "Env: SCION_HUB_START_CWD=/safe/path or SCION_HUB_PRESERVE_CWD=1"
    exit 1
fi

NAME="$1"
shift
TASK="$1"
shift

if [ "${SCION_HUB_PRESERVE_CWD:-0}" != "1" ]; then
    cd "${SCION_HUB_START_CWD:-$HOME}"
fi

scion --non-interactive start "$NAME" "$TASK" --notify "$@"
