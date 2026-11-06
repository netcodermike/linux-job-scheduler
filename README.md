# Linux Job Scheduler

Execute trusted local command argument arrays after bounded foreground delays.

[![Build and verify](https://github.com/netcodermike/linux-job-scheduler/actions/workflows/verify.yml/badge.svg)](https://github.com/netcodermike/linux-job-scheduler/actions/workflows/verify.yml)

**Language:** Go · **Platform:** Linux-focused · **App type:** Linux local scheduled-command runner

First-release learning implementation, with explicit scope below. This is one independently runnable app, not a placeholder in a combined repository.

## Features

- JSON job definitions
- 0–86400 second delays
- Ordered serial execution
- Child failure propagation

## Requirements

Go 1.22+ and the documented Linux facilities. No external Go modules.

## Quick start

```sh
git clone https://github.com/netcodermike/linux-job-scheduler.git
cd linux-job-scheduler
go build -o app .
./app --help
# Explicitly executes the benign example commands:
./app examples/jobs.json
```

Commands run from the repository root. Lines marked for a particular operating system apply only there. Language tools are prerequisites, not installed by the app. Close desktop windows normally; terminal/server apps stop with Ctrl+C unless their usage notes say otherwise.

## Example workflow

Inspect examples/jobs.json, then run ./app examples/jobs.json and observe the immediate greeting and delayed second job. Test a failing command in a synthetic copy.

## Architecture, configuration, and data

main.go loads 1–100 Jobs containing delay_seconds and command string arrays. Delays are relative to process start and jobs are selected by earliest scheduled delay, with stable input order for ties. It invokes executables directly without a shell and attaches output/error streams. --help prints usage. Unit tests reject invalid commands/delays; examples/jobs.json contains benign printf/true jobs. It is an explicit command runner, not an installed service.

## Verification

Build with the quick-start command where needed, then run:

```sh
go test ./...
```

Manual acceptance: Inspect examples/jobs.json, then run ./app examples/jobs.json and observe the immediate greeting and delayed second job. Test a failing command in a synthetic copy. Also test malformed inputs and an unwritable output destination.

See [validation record](docs/VALIDATION.md) for observed checks and their limits. CI results verify only what the workflow runs; they do not establish native GUI correctness or production readiness.

## Limitations and next steps

Trusted commands only. No background installation, persistence across restarts, cron syntax, recurring schedules, environment isolation, timeout, cancellation tree handling, missed-job recovery, or concurrency. A long-running job delays later ones. Commands inherit the user environment and can modify files; inspect job files before running.

## Troubleshooting

Use the exact platform/toolchain above. Preserve originals and use synthetic fixtures first. Errors are displayed in the UI or stderr. A headless machine can run the model tests but cannot verify native GUI/audio behavior. Output paths are chosen explicitly; keep backups of important data.

## Source map

- `go.mod`
- `main.go`
- `main_test.go`
- `examples/jobs.json`

## Development and license

Read [contributing](CONTRIBUTING.md) before changing the app, and [security/privacy](SECURITY.md) before supplying private data. Use synthetic fixtures for tests. No telemetry, paid API key, or account is required by the app itself. Dependencies, when present, are stated in Requirements.

[MIT license](LICENSE), copyright 2026 netcodermike.
