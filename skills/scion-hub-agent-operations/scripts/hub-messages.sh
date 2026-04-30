#!/bin/bash
# List Scion Hub messages as JSON for orchestration inbox handling.

set -euo pipefail

scion --non-interactive messages --json "$@"
