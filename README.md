<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/storage-r2</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/storage-r2"><img src="https://pkg.go.dev/badge/github.com/togo-framework/storage-r2.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/storage-r2
```

<!-- /togo-header -->

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

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
