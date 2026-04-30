# Skills

Skill definitions that enable agents to manage Scion from within a session.

## Structure

- **`scion/SKILL.md`** - General Scion documentation covering local agent management, template operations, and configuration commands.
- **`scion/scripts/`** - Backing shell scripts for local/common operations (start, list, status, message).
- **`scion-hub-agent-operations/SKILL.md`** - Hub-only agent operations: status, remote start, list, look, message, and notifications.
- **`scion-hub-agent-operations/scripts/`** - Hub-safe wrappers that add `--non-interactive` and notification defaults.
- **`scion-hub-template-admin/SKILL.md`** - Hub-only template, harness-config, env, secret, grove, and orchestrator-team administration.
- **`team-creation/SKILL.md`** - General multi-agent template creation guidance.

## Usage

Use the hub-specific skills whenever the user asks about Scion Hub, Hosted mode, remote Runtime Brokers, Hub groves, Hub templates, Hub secrets, or Hub notifications. Use the general `scion` skill for local-mode workflows.
