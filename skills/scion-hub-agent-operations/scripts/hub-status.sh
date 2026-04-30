#!/bin/bash
# Show Scion Hub connection status as JSON.

set -euo pipefail

scion --non-interactive hub status --json "$@"
