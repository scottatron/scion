# Upstream Rebase Summary (2026-04-29)

This note summarizes the upstream-only changes pulled into this fork branch while rebasing `codex/codex-hook-parity`.

- Previous fork base: `d6b5c930f8e5e710ca815c684144594d8d896656`
- New upstream head: `f1c3da8c10cc4bf81ed5cb154aa5d8c13526001e`
- Compared range: `d6b5c930f8e5e710ca815c684144594d8d896656..upstream/main`
- Scope: upstream changes only; local fork commits are excluded.
- Size: 45 commits, 131 files changed, 16,593 insertions, 2,018 deletions.

## Main Changes

- Harness configuration moved substantially toward script-based, container-script harnesses. Upstream added shared `scion_harness.py` provisioning helpers, container-script harness implementations and tests, default enablement for `allow_container_script_harnesses`, and reconciliation of the container-script bundle on each agent start.
- Universal MCP server configuration landed across the API, schema, templates, and harness config generation. The new `mcp_servers` schema is validated, staged into container-script harnesses, and applied to OpenCode and Codex configuration generation, with a Chrome DevTools MCP entry added to the web-dev template.
- Harness config management expanded with `scion harness-config install`, transfer/revision tests, upgrade helpers, and fixes for listing, grove auto-selection, duplicate config discovery, and broken harness-config directories.
- Hub, broker, and web behavior improved around agent visibility and lifecycle state. The agent list gained all/mine/shared filtering, dashboard navigation reliably shows New Grove/New Agent actions, idle agents are included in stall detection, and duplicate waiting-for-input inbox/notification messages are eliminated.
- Runtime and broker execution paths were tightened. Upstream applied the whoami-skip-su pattern to more `su` call sites, added execution-user tests, preferred API-key auth over Vertex AI during env-gather auto-detection, and logs `provision.py` output to `agent.log` for diagnostics.
- Image build tooling was refactored into pluggable builders for local Docker, local Podman, and Cloud Build. Registry handling is now optional for local builds, registry permissions are checked earlier, Cloud Build project mismatches are warned about, and a Hub image Cloud Build path was added.
- Documentation and examples were broadened. Upstream added Amp harness design and example assets, PostgreSQL strategy notes, universal MCP server design notes, hub template admin design notes, clearer quickstart container requirements, and updated custom image docs.
- Dependencies and build infrastructure were refreshed, including Astro/Starlight and PostCSS bumps, Go module updates, Node.js 22 for docs CI, rclone and Azure NTLMSSP updates, and faster Git builds using all CPU cores.

## Notable File Areas

- `.design/`: new design docs for Amp harness scripts, container builder refactor, hub template admin, PostgreSQL strategy, and template MCP servers.
- `cmd/` and `cmd/sciontool/commands/`: new harness config install and sciontool harness commands with tests.
- `pkg/api/`, `pkg/config/`, `pkg/harness/`, and `pkg/runtimebroker/`: most of the API/schema/harness/runtime policy work, including the new container-script harness path.
- `image-build/`: builder refactor, Cloud Build support, Hub image build files, registry verification, and build script cleanup.
- `examples/amp/`: new Amp harness example and supporting template files.
- `web/`: agent list scope filtering and navigation fixes.
- `docs-site/`: dependency updates and install/custom-image documentation refreshes.

## Upstream Commits

| Commit | Date | Subject |
| --- | --- | --- |
| `5b5d0130` | 2026-04-25 | feat(hub,web): add all/mine/shared scope filter to agent list view |
| `c06627dd` | 2026-04-25 | fix(hub): include idle agents in stall detection |
| `905ab397` | 2026-04-25 | fix(web): show New Grove/Agent buttons when navigating from dashboard |
| `d9514447` | 2026-04-25 | fix(broker): apply #159's whoami-skip-su pattern to remaining su call sites (#180) |
| `18f9d0b4` | 2026-04-25 | fix(docs): bump Node.js to 22 for Astro CI compatibility |
| `7e448a0f` | 2026-04-25 | build(deps): bump astro and @astrojs/starlight in /docs-site (#177) |
| `fc7315d2` | 2026-04-25 | build(deps-dev): bump postcss from 8.5.6 to 8.5.10 in /web (#187) |
| `741a8e51` | 2026-04-25 | build(deps): bump postcss from 8.5.6 to 8.5.10 in /docs-site (#186) |
| `686f2cdf` | 2026-04-25 | build(deps): bump go.opentelemetry.io/otel in /extras/scion-chat-app (#185) |
| `9719f582` | 2026-04-25 | build(deps): bump github.com/Azure/go-ntlmssp (#181) |
| `db06ffd9` | 2026-04-25 | build(deps): bump github.com/rclone/rclone from 1.72.1 to 1.73.5 (#178) |
| `707b2944` | 2026-04-25 | fix: prefer api-key over vertex-ai in env-gather auth auto-detection |
| `08d8edc2` | 2026-04-26 | chore: Update Git build to use all CPU cores (#190) |
| `57b57880` | 2026-04-26 | refactor: intial transtion to script based harness |
| `7568855c` | 2026-04-26 | internal: add docs writer template |
| `770980cb` | 2026-04-26 | docs: update quickstart with clearer container requirements (#194) |
| `e7d5a691` | 2026-04-27 | move Claude model env vars from settings.json to config.yaml |
| `db72c13e` | 2026-04-27 | feat: add 'scion harness-config install' command |
| `d07d45d6` | 2026-04-27 | test(hub): remove obsolete TestCreateAgent_HarnessFieldIgnoredWhenTemplateResolved |
| `b1c9d2a2` | 2026-04-27 | feat: add Amp harness example and design documentation |
| `afce4e16` | 2026-04-27 | docs: add postgres support strategy |
| `0a9c31d3` | 2026-04-27 | fix: eliminate duplicate notification/inbox messages for waiting_for_input |
| `d33ba561` | 2026-04-27 | refactor: pluggable container build backends via --builder |
| `12bfb3de` | 2026-04-27 | refactor: make --registry optional for local builds; fix mapfile on Bash 3.2 |
| `098e7565` | 2026-04-25 | design: add universal MCP server configuration proposal for templates |
| `87ce8860` | 2026-04-25 | design(mcp-servers): record decisions on all open questions |
| `dc9b8e92` | 2026-04-26 | design(mcp-servers): revise integration section for container-script harness model |
| `6d40893a` | 2026-04-27 | feat(api): add universal mcp_servers schema and validation |
| `34ceb9a7` | 2026-04-27 | feat(harness): add shared scion_harness.py provision helper |
| `794ec946` | 2026-04-27 | feat: stage mcp_servers for container-script harnesses |
| `7f5ef41b` | 2026-04-27 | feat(opencode): apply universal mcp_servers to opencode.json |
| `3b54cb31` | 2026-04-27 | feat(codex): apply universal mcp_servers to codex config.toml |
| `d1ef1cc7` | 2026-04-27 | template(web-dev): add chrome-devtools mcp_server entry |
| `4d998577` | 2026-04-27 | chore: add a few missing headers |
| `847150bf` | 2026-04-27 | Default allow_container_script_harnesses to true |
| `ef248a19` | 2026-04-27 | template(web-dev): drop opencode harness-config shim |
| `021208c4` | 2026-04-27 | fix: deduplicate harness configs when listing by grove ID |
| `253ab733` | 2026-04-27 | fix(cloud-build): add pre-flight registry check to catch permission errors early |
| `bdfa577e` | 2026-04-27 | fix(cloud-build): auto-detect project from --registry and warn on mismatch |
| `ffabcebe` | 2026-04-28 | fix: reload harness configs with grove filter after grove auto-selection |
| `e0495783` | 2026-04-28 | fix: FindHarnessConfigDir falls through on broken harness-config dirs |
| `61fd7a76` | 2026-04-28 | fix: activate container-script provisioner for opencode by default |
| `d03bcb36` | 2026-04-28 | fix: log provision.py output to agent.log for diagnostics |
| `2d0ba51f` | 2026-04-28 | docs: add hub template admin design doc |
| `f1c3da8c` | 2026-04-28 | fix: reconcile container-script bundle on every agent start |
