---
name: scion-hub-template-admin
description: Use only for Scion Hub or Hosted mode template, harness-config, environment, secret, and multi-agent team setup workflows, especially when publishing local template or harness-config changes to Hub for remote Runtime Brokers.
---

# Scion Hub Template Admin

Use this skill for Hub-mode configuration work: templates, harness-configs, env vars, secrets, Git groves, hub-native groves, and orchestrator/worker team templates that will run through remote Runtime Brokers.

## Source Of Truth

- Hub storage is the source of truth for Hub-mode templates and harness-configs used by remote brokers.
- A live agent home is only a provisioned copy. Do not edit it to change future agents.
- Local edits are staging copies until they are pushed or synced to Hub.
- Existing agents do not automatically receive template or harness-config changes; verify with a fresh agent when behavior matters.

## Required Command Shape

- Use `scion --non-interactive ...` for every command.
- Use `--json` or `--format json` for status and inventory commands when available.
- Do not use `--no-hub`, `--global`, workspace `scion sync`, or `cdw` as shortcuts for Hub admin work.
- Prefer grove-scoped configuration unless the user explicitly asks for Hub-wide defaults.

## Template Workflow

Inspect template state:

```bash
scion --non-interactive templates list --format json
scion --non-interactive templates status --format json
```

Create or edit templates locally under `.scion/templates/<template-name>/`, then publish:

```bash
scion --non-interactive templates sync <template-name>
```

Publish all local grove templates only when the user clearly wants that broader action:

```bash
scion --non-interactive templates sync --all
```

Pull a Hub template to inspect or modify it locally:

```bash
scion --non-interactive templates pull <template-name>
```

Start a fresh agent after publishing to verify the Hub payload:

```bash
scion --non-interactive start verify-<slug> "verify template behavior" --type <template-name> --notify
```

## Harness-Config Workflow

List local and Hub harness-configs:

```bash
scion --non-interactive harness-config list --hub --format json
```

Pull the Hub copy, edit `config.yaml`, `home/...`, or `skills/...`, then publish:

```bash
scion --non-interactive harness-config pull <name>
scion --non-interactive harness-config sync <name>
```

`push` is a semantic alias for `sync`:

```bash
scion --non-interactive harness-config push <name>
```

When a template should use a specific harness-config, set `default_harness_config` in `scion-agent.yaml`, publish the template, then verify with a fresh agent.

## Env And Secrets

Use Hub env vars for non-sensitive values:

```bash
scion --non-interactive hub env list --json
scion --non-interactive hub env set --grove <grove-slug> KEY=value
```

Use Hub secrets for sensitive values:

```bash
scion --non-interactive hub secret list --json
scion --non-interactive hub secret set --grove <grove-slug> KEY value
```

For file secrets, preserve the target path separately from the source file:

```bash
scion --non-interactive hub secret set --grove <grove-slug> --type file --target '~/.codex/auth.json' CODEX_AUTH '@~/.codex/auth.json'
```

Private Git groves need a `GITHUB_TOKEN` secret with at least repository contents read access.

## Grove Setup

Create a Git-backed grove:

```bash
scion --non-interactive hub grove create https://github.com/org/repo.git --slug <grove-slug> --json
```

Start agents against that grove by slug:

```bash
cd "$HOME" && scion --non-interactive --format json --grove <grove-slug> start <agent-name> "task" --notify
```

Remember that Hub-mode Git groves clone over HTTPS on the broker; local worktrees and SSH credentials are not the provisioning path.

## Hub-Mode Team Templates

When building an orchestrator/worker team for Hub mode:

- Create exactly one orchestrator template and one or more worker templates.
- Keep role instructions in `agents.md`; do not duplicate generic Scion CLI reference material.
- Tell orchestrators to start workers with `--notify`.
- Tell orchestrators that launching workers from an inherited agent/tool cwd can accidentally bootstrap that local path before Hub creation. Prefer direct Hub dispatch; if using `scion start`, run it from `$HOME` or a tiny inert cwd rather than `/opt/data`, `/workspace`, fakeowner mounts, or large repo roots.
- Tell orchestrators to use `scion look` to collect worker output.
- Tell orchestrators to use `scion messages` as the durable worker-result channel.
- Tell workers to send final summaries and status updates back through Scion messages.
- Tell orchestrators to call `sciontool status blocked "Waiting for <reason>"` while waiting.
- Publish all templates needed by the team before starting the orchestrator.
- Start the orchestrator through Hub with `--notify`.

Minimal orchestrator reminder:

```markdown
When starting workers from inside an agent or shell tool, avoid the inherited cwd. Prefer direct Hub dispatch, or run `cd "$HOME" && scion --non-interactive --format json --grove <grove> start <name> "task" --type <template> --notify`.
When waiting, run `sciontool status blocked "Waiting for worker agents to complete"`.
Ask workers to send their final summaries through Scion messages.
Use `scion --non-interactive messages --json` and `scion --non-interactive messages --all --json` to collect worker results.
Use `scion --non-interactive look <name>` only when you need live terminal context before messaging them.
```
