# Validation

Authoring date: November 6, 2026. Host: Windows. Target: Linux-focused.

## Local result

Target-platform source CI is recorded after publication. This host has Python, Node.js, and .NET; other native toolchains are verified through CI. No unobserved GUI/device behavior is claimed.

## Reproducible checks

The GitHub Actions workflow builds/runs the following on `ubuntu-latest`:

```sh
gofmt -w main.go main_test.go
go vet ./...
go test -race ./...
go build -o app .
```

## Manual acceptance

Inspect examples/jobs.json, then run ./app examples/jobs.json and observe the immediate greeting and delayed second job. Test a failing command in a synthetic copy. Also test malformed inputs and an unwritable output destination.

A workflow file is not evidence of a passing run. Observed CI run links and source commits are recorded after publication. Native GUI, sound, sleep/resume, and platform permissions require direct target-platform acceptance even when model tests pass.
