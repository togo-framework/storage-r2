# storage-r2 — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package r2store is a Cloudflare R2 driver for togo storage. R2 is S3-compatible,
so this uses the AWS S3 SDK pointed at the R2 endpoint. Implements togo.Storage
and overrides the default filesystem storage when installed.

	togo install togo-framework/storage-r2

## Install

```bash
togo install togo-framework/storage-r2
```

Set `STORAGE_DRIVER=r2`.

## Configuration

Environment variables read by this plugin (extracted from the source):

| Env var | Notes |
|---|---|
| `G` | _see provider docs_ |
| `R2_ACCESS_KEY_ID` | _see provider docs_ |
| `R2_ACCOUNT_ID` | _see provider docs_ |
| `R2_BUCKET` | _see provider docs_ |
| `R2_SECRET_ACCESS_KEY` | _see provider docs_ |

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
- README: ../README.md
