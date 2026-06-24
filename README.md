# storage-r2

A **Cloudflare R2** driver for [togo](https://to-go.dev) storage. R2 is S3-compatible,
so this uses the AWS S3 SDK pointed at your R2 endpoint. Implements `togo.Storage` and
overrides the default filesystem storage when installed.

## Install

```bash
togo install togo-framework/storage-r2
```

## Configure (`.env`)

```ini
R2_ACCOUNT_ID=…
R2_ACCESS_KEY_ID=…
R2_SECRET_ACCESS_KEY=…
R2_BUCKET=my-bucket
```

If the R2 env isn't fully set the plugin no-ops and the default storage stays active.

MIT © togo-framework
