# envsk

Like `env`, but masking the value of the environment variables.

`envsk` is a simple wrapper around `env` that will always replace the value of any environment variable with `********`. This way you can check that the secret exists without leaking it to the world.

## Motivation

I was tired of having to type `env | grep SECRET` in public. I wanted a way to check that the secret existed without leaking it to the world.

Another use case is for CI/CD pipelines. You can use `envsk` to check that the secret exists before running the pipeline, but without leaking it to the logs.

## Usage

```bash
# list all secrets
$ envsk

# list secrets with starting with uppercase "^VAULT":
$ envsk --filter="^VAULT"
VAULT_ADDR=********

# list secrets with containing an uppercase V ending with an uppercase R:
$ envsk --filter=".*V.*R$"
VAULT_ADDR=********
NVM_DIR=********
```

## Installation

```bash
$ go get github.com/mdelapenya/envsk
```
