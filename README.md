# github.com/beeos-ai/sdk-go

Go client for the BeeOS OpenAPI contract — served exclusively by
[`openapi-gateway`](https://github.com/beeos-ai/openagent/tree/main/backend/services/openapi-gateway).
The spec is [`backend/openapi/beeos-platform-v1.yaml`](https://github.com/beeos-ai/openagent/blob/main/backend/openapi/beeos-platform-v1.yaml);
there is no filter / subset step — the whole contract is exposed. Main BeeOS
Gateway is **not** part of this contract.

## Install

```bash
go get github.com/beeos-ai/sdk-go@latest
```

## Usage

- **Base path** — point the generated `Configuration` at your `openapi-gateway`
  host (dev `http://localhost:8095`, prod e.g. `https://openapi.beeos.ai`).
- **Auth** — pass `Authorization: Bearer <jwt>` or
  `Authorization: Bearer oag_<user-api-key>` via the `http.Client.Transport`
  wrapper you inject into the generated API client.

## Regenerate (maintainers)

From the monorepo: `cd sdks/openapi-sdk && ./generate.sh`

**Module:** `github.com/beeos-ai/sdk-go` — this tree is mirrored as its own
GitHub repository for third-party `go get`.
