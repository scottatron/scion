#!/bin/bash
# Inspect a Hub-managed Scion agent without entering an interactive session.

set -euo pipefail

if [ $# -lt 1 ]; then
    echo "Usage: hub-look-agent.sh <agent-name-or-id> [scion flags...]"
    exit 1
fi

AGENT="$1"
shift

scion --non-interactive look "$AGENT" "$@"
