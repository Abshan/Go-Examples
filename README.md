# Go Examples Workspace

This repository is organized by learning category so you can keep adding focused Go samples over time.

## Structure

```text
samples/
  basic/
    hello-world/
    file-read-write/
  connections/
    tcp-client/
  http-clients/
    simple-get/
  backends/
    basic-server/
```

## Quick Start

List available samples:

```bash
make list
```

Run one sample:

```bash
make run SAMPLE=basic/hello-world
make run SAMPLE=basic/file-read-write
make run SAMPLE=connections/tcp-client
make run SAMPLE=http-clients/simple-get
make run SAMPLE=backends/basic-server
```

Or use Go directly:

```bash
go run ./samples/basic/hello-world
```

## How To Add New Samples

1. Pick the right category under `samples/`.
2. Create a new folder (for example `samples/basic/loops`).
3. Add `main.go` with `package main`.
4. Add a small `README.md` for purpose and run command.

Naming guideline:
- Use kebab-case for sample folder names.
- Keep each sample self-contained and runnable with `go run ./samples/<category>/<sample>`.
