---
name: scion-hub-agent-operations
description: Use only for Scion Hub or Hosted mode agent operations: checking Hub connectivity, starting remote agents, listing or inspecting Hub agents, messaging agents, using notifications, and avoiding local-only Scion workflows.
---

# Scion Hub Agent Operations

Use this skill only when Scion is operating through a Hub. Treat the Hub API as the source of truth for agents, groves, notifications, and remote runtime routing.

## Non-Negotiable Hub Rules

- Always run Scion commands with `--non-interactive`.
- Prefer JSON output for inspection commands: use `--format json` where supported, or command-specific `--json`.
- Do not use `--no-hub` to work around errors.
- Do not use `--global`; Hub-mode agents operate inside a grove context or an explicit Hub grove.
- Do not use `sync` or `cdw` from inside a Hub-managed agent unless the user explicitly asks for that workflow.
- Do not assume local git worktrees. Git groves in Hub mode are provisioned by HTTPS clone on the broker.
- Use `scion look` for recent terminal state before interrupting or attaching.
- Use `--notify` when starting or messaging agents you will wait on.

## First Checks

Confirm Hub connectivity and auth before changing anything:

```bash
scion --non-interactive hub status --json
```

List available groves or brokers when routing context is unclear:

```bash
scion --non-interactive hub groves --json
scion --non-interactive hub brokers --json
```

If a command reports missing Hub authentication, stop and ask the user to authenticate or provide the intended endpoint. Do not fall back to local mode.

## Starting Agents

For the current linked grove:

```bash
scion --non-interactive start <agent-name> "task" --notify
```

For a remote Hub grove by slug:

```bash
scion --non-interactive start <agent-name> --grove <grove-slug> "task" --notify
```

For a specific template:

```bash
scion --non-interactive start <agent-name> --grove <grove-slug> "task" --type <template-name> --notify
```

When waiting on child agents from inside a Scion agent, mark yourself blocked:

```bash
sciontool status blocked "Waiting for agent <agent-name> to complete"
```

## Inspecting And Messaging

List agents:

```bash
scion --non-interactive list --format json
scion --non-interactive list --grove <grove-slug> --format json
```

Inspect an agent's recent output and terminal state:

```bash
scion --non-interactive look <agent-name>
```

Send follow-up information and stay subscribed:

```bash
scion --non-interactive message <agent-name> "message" --notify
```

Interrupt only when the new instruction must replace current work:

```bash
scion --non-interactive message <agent-name> "urgent replacement instruction" --interrupt --notify
```

## Notifications

Check notifications:

```bash
scion --non-interactive notifications --json
```

Subscribe explicitly when `--notify` was omitted or a grove-level watch is needed:

```bash
scion --non-interactive notifications subscribe --agent <agent-name>
scion --non-interactive notifications subscribe --grove <grove-slug>
```

Acknowledge handled notifications:

```bash
scion --non-interactive notifications ack <notification-id>
```

## Completion

Before reporting success, inspect the relevant agents or notifications and summarize the concrete Hub state you verified: agent name, grove, phase/activity, and any message or notification IDs available.
