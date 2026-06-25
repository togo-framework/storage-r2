# storage-r2 — documentation

togo storage driver

## Overview

Package r2store is a Cloudflare R2 driver for togo storage. R2 is S3-compatible,
so this uses the AWS S3 SDK pointed at the R2 endpoint. Implements togo.Storage
and overrides the default filesystem storage when installed.


Env: R2_ACCOUNT_ID, R2_ACCESS_KEY_ID, R2_SECRET_ACCESS_KEY, R2_BUCKET (all required).

## Install

```bash
togo install togo-framework/storage-r2
```

Set `STORAGE_DRIVER=r2`.

## Configuration

Environment variables read by this plugin (extracted from the source — see the gateway/provider docs for each value):

| Env var |
|---|
| `R2_ACCESS_KEY_ID"` |
| `R2_ACCOUNT_ID"` |
| `R2_BUCKET"` |
| `R2_SECRET_ACCESS_KEY"` |

## Usage

```go
st := k.Storage
st.Put(ctx, "path/file.txt", data)
b, _ := st.Get(ctx, "path/file.txt")
url := st.Path("path/file.txt")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/storage-r2
- Full README: ../README.md
