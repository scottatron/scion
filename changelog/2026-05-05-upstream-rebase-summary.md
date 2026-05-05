# Upstream Rebase Summary - 2026-05-05

This summarizes the upstream-only changes pulled into this fork branch when rebasing `codex/codex-hook-parity` onto the latest `GoogleCloudPlatform/scion` `main`.

## Compared Refs

- Previous fork base: `f1c3da8c10cc4bf81ed5cb154aa5d8c13526001e`
- New upstream HEAD: `2fad47cfef97e68cce6928ed442fc02fe7831366`
- Compared range: `f1c3da8c10cc4bf81ed5cb154aa5d8c13526001e..2fad47cfef97e68cce6928ed442fc02fe7831366`
- Upstream commits included: 60
- Diff stat: 122 files changed, 9,950 insertions, 937 deletions

## Main Changes

### Chat App Hardening and Command Routing

The largest functional area is the chat app. Upstream added a startup race fix, visible sent-message rendering, a default agent flow, and a split between `/scion` messaging and `/scionAdmin` administration commands. Follow-up fixes route admin commands through `slashCommand.commandId`, parse command names when the ID map misses, and keep non-instruction user messages on the notification path so harness output does not leak back into chat. Agent subscriptions and operations are now scoped by grove through space links, including a store migration that adds `grove_id` to `agent_subscriptions` primary keys.

### Hub, Broker, and Grove Isolation

The hub and broker paths now tighten colocated communication and grove scoping. Colocated broker-to-hub communication uses localhost to avoid GCE hairpin NAT issues, runtime broker delivery is scoped by grove, and `attach` now resolves agents by grove to prevent cross-grove routing. The hub also bootstraps broker subscriptions on startup, rejects unknown recipients, improves message broker tests, and falls back to the creator's GitHub token when cloning a grove.

### Metadata Server and Agent Identity

Upstream added a CLI mode system covering human, assistant, and agent modes, plus a new `scion whoami` command. The metadata server gained self-healing health checks and a diagnostic command. Related fixes keep metadata sidecar iptables rules out of the host namespace, make `agent-info.json` readable by the broker, and chown refreshed `scion-token` files so the `scion` user can read them.

### Starter Hub Deployment and GCP Auth

Starter hub scripts now include preflight validation, cloud-init waiting before first SSH, improved hub config handling, and updated env samples including `SCION_SERVER_HUB_ADMINEMAILS`. Web UI text now calls out IAM propagation delay in the service account verification dialog. Upstream also added investigation and post-mortem docs for GCP auth token failures and metadata server 502s on starter-hub VMs.

### Web and Docs Site Updates

The docs site received a standalone landing page, a new graphic asset, Node.js 24 compatible GitHub Actions updates, and Astro/Starlight dependency compatibility fixes. Web updates include the GitHub App link on the grove creation form, structured commit lists in pull-latest results, explicit GCP identity `block` mode in agent creation, the Vite entry point fix, smart file-browser hybrid search, branch selection for hub rebuilds, and an embedded-assets error page.

### Build and Image Tooling

Cloud Build templates now declare `_SHORT_SHA` and `_COMMIT_SHA` consistently, pass only substitutions referenced by the target template, and avoid overwriting Docker auth config during registry verification. The default pull-images harness list now includes `opencode` and `codex`.

### Skills and Design Documentation

Upstream added design docs for Discord chat adapters, integration environment handling, CLI modes, the chat-app race fix, and starter-hub incident investigations. The team-creation skill was streamlined and extended for extension support.

## Notable File Areas

- `.design/`: new Discord adapter, CLI modes, integration environment, chat-app race, GCP auth, and starter-hub post-mortem docs.
- `cmd/`: new CLI mode and `whoami` commands plus server/foreground dispatch updates.
- `cmd/sciontool/commands/` and `pkg/sciontool/metadata/`: metadata diagnostics, harness HOME fix, and token/readability fixes.
- `extras/scion-chat-app/`: chat command routing, event handling, notifications, adapter tests, state scoping, and install/config updates.
- `pkg/hub/` and `pkg/runtimebroker/`: grove-scoped messaging, broker subscription bootstrap, token caching, maintenance rebuild branch selection, and colocated communication fixes.
- `web/src/components/`: file browser, grove creation, agent creation, admin maintenance/config, and grove detail updates.
- `scripts/starter-hub/`: preflight validation, cloud-init wait, hub config helpers, and sample env updates.
- `docs-site/`: landing page, graphic asset, and dependency/action compatibility updates.
- `image-build/scripts/`: Cloud Build substitution and Docker auth handling fixes.

## Upstream Commit List

| Commit | Date | Subject |
| --- | --- | --- |
| `5210763e` | 2026-04-28 | fix(cloud-build): declare _SHORT_SHA and _COMMIT_SHA in all cloudbuild templates |
| `ce4a2f53` | 2026-04-28 | fix(cloud-build): only pass substitutions that the template references |
| `c7daf9d5` | 2026-04-29 | fix(cloud-build): stop overwriting Docker auth config in verify-registry |
| `d2e5ef1f` | 2026-04-27 | Add Discord chat adapter design document |
| `a69e269c` | 2026-04-28 | Resolve all open questions in Discord adapter design |
| `63eb1664` | 2026-04-29 | feat(docs-site): add standalone landing page |
| `f9127fab` | 2026-04-29 | fix(ci): upgrade GitHub Actions to Node.js 24 compatible versions |
| `60bd6bb6` | 2026-04-29 | fix(ci): upgrade pages actions to Node.js 24 native versions |
| `4c017409` | 2026-04-29 | fix(docs-site): upgrade starlight-links-validator for Zod v4 compat |
| `b5ba7103` | 2026-04-29 | fix(docs-site): upgrade astro-d2 for Astro 6.x compatibility |
| `cb288978` | 2026-04-29 | fix(broker): use localhost for co-located broker-to-hub communication |
| `600f75d7` | 2026-04-30 | fix(sciontool): resolve HOME mismatch in container-script harness provisioning |
| `c6be89f2` | 2026-04-30 | feat(web): add GitHub App link to grove creation form |
| `7b58cc67` | 2026-04-29 | docs: add SCION_SERVER_HUB_ADMINEMAILS to starter-hub env sample |
| `b3c34b0c` | 2026-04-30 | feat(starter-hub): add preflight validation to deployment pipeline |
| `03be867f` | 2026-04-30 | fix(web): add IAM propagation delay note to SA verification dialog |
| `e256b999` | 2026-04-30 | fix: include opencode and codex in default pull-images harness list |
| `d7271f36` | 2026-04-30 | refactor(skills): streamline team-creation skill and add extension support |
| `163e5a6b` | 2026-04-30 | feat(filebrowser): implement smart file list with hybrid search (phases 1 & 2) |
| `babb9aaf` | 2026-04-30 | feat(web): show structured commit list in pull-latest results |
| `87aba182` | 2026-05-01 | fix(web): send explicit GCP identity "block" mode in agent create request |
| `424c1063` | 2026-05-01 | fix(web): set main.ts as Vite entry point and import styles (#198) |
| `4baf735f` | 2026-05-01 | fix(broker): resolve GCE hairpin NAT for colocated Docker bridge containers (#206) |
| `4e692581` | 2026-05-01 | docs(investigations): add GCP auth token failure investigation |
| `ef98b80b` | 2026-05-01 | Add note to update dialog clarifying agents are unaffected |
| `bd1a73e6` | 2026-05-01 | fix(chat-app): resolve plugin startup race condition |
| `c5fbcd69` | 2026-05-01 | docs: capture chat app startup race condition design and fix |
| `d99a9f3c` | 2026-05-01 | fix(hub): resolve SPA lag from SQLite single-connection bottleneck |
| `b282b28d` | 2026-05-01 | fix(sciontool): use mode 0644 for agent-info.json so broker can read it |
| `3bc214b8` | 2026-05-01 | fix(sciontool): chown scion-token file after refresh so scion user can read it |
| `a040d1c6` | 2026-05-01 | feat(chat-app): show sent messages visibly and fix card newline rendering |
| `fdee9408` | 2026-05-02 | fix(apiclient): handle nil PageOptions in ToQuery to prevent panic |
| `a110be2f` | 2026-05-02 | fix(hub): bootstrap broker subscriptions on startup and reject unknown recipients |
| `af759428` | 2026-05-02 | fix: prevent metadata sidecar iptables rules from leaking to host namespace |
| `9652d7fe` | 2026-05-02 | docs: add post-mortem for metadata server 502 on starter-hub VMs |
| `3812beff` | 2026-05-02 | fix(chat-app): use hub GCP project for signing key auto-discovery |
| `8c57a072` | 2026-05-03 | feat(chat-app): improve messaging, fix start/stop, add default agent (#207) |
| `6635ad1d` | 2026-05-04 | feat: implement CLI mode system (human/assistant/agent) |
| `84ce7939` | 2026-05-04 | docs: add CLI mode guidance to Adding Commands section |
| `6aa006ec` | 2026-05-04 | feat: add `scion whoami` command for agent identity |
| `67e38143` | 2026-05-04 | chore: add whoami to agent CLI mode allow-list |
| `73bcb033` | 2026-05-04 | refine: tighten agent mode and remove debug logging |
| `3f1c91cc` | 2026-05-04 | update starter-hub scripts with fixed preflight |
| `8f52a45e` | 2026-05-04 | feat: add self-healing health check and diagnostic command to metadata server |
| `8c23703f` | 2026-05-04 | fix: resolve TypeScript type error in file-browser signal parameter |
| `fc07b287` | 2026-05-04 | fix(chat-app): scope all agent operations to grove via space link |
| `7110a246` | 2026-05-04 | fix(web): serve error page when web assets are not embedded |
| `550afad9` | 2026-05-04 | fix(chat-app): include grove_id in agent_subscriptions primary key |
| `42eef10a` | 2026-05-04 | feat(maintenance): support branch selection for hub rebuild operations |
| `559df610` | 2026-05-04 | fix(broker): scope agent message delivery by grove to prevent cross-grove routing |
| `f3cbfabe` | 2026-05-04 | Fix Message/MessageRaw call sites to include groveID parameter |
| `38cd4e69` | 2026-05-04 | fix(starter-hub): wait for cloud-init before first SSH in setup-repo |
| `8863fc90` | 2026-05-04 | fix(attach): scope agent lookup by grove to prevent cross-grove routing |
| `7f63c4f6` | 2026-05-04 | feat(chat-app): split /scion into messaging and /scionAdmin for admin commands |
| `7343f4c6` | 2026-05-04 | fix(hub): fall back to creator's GitHub token during grove clone |
| `7b0ce230` | 2026-05-04 | fix(chat-app): route /scionAdmin via slashCommand.commandId field |
| `00737860` | 2026-05-04 | fix(chat-app): parse command name from message text when ID map misses |
| `19ad09a9` | 2026-05-05 | fix(chat-app): only relay explicit instruction messages to chat |
| `d0c1c06e` | 2026-05-05 | fix(chat-app): route non-instruction user messages to notification path |
| `2fad47cf` | 2026-05-05 | fix(chat-app): close all paths where harness output leaks to chat |
