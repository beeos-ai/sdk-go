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

- **Base URL** — the facade uses `https://openapi.beeos.ai` by default. Set
  `BEEOS_API_URL` or pass `BaseURL` to target another environment.
- **Auth** — set `BEEOS_API_KEY`, or pass `APIKey` explicitly. The generated
  client also supports custom authorization through its request context.

For the stable, task-focused facade:

```go
client, err := beeos.NewClient()
if err != nil { log.Fatal(err) }

agents, err := client.ListAgents(context.Background())
```

`client.API` exposes the complete generated Platform OpenAPI surface.

For phone automation across Device Agent, BeeRunner, and Redroid:

```go
mobile, err := beeos.NewMobileClient(beeos.MobileClientOptions{
    AgentID: "agent-id", InstanceID: "instance-id",
})
if err != nil { log.Fatal(err) }

if _, err = mobile.WaitReady(ctx); err != nil { log.Fatal(err) }
result, err := mobile.Run(ctx, *beeos.NewCreateTaskRequest("Open Settings"))
```

`mobile.API.MobileAPI` exposes the generated atomic-control API. BeeRunner
uses the durable task methods and does not advertise atomic control until a
trusted Portal adapter exists.

## Regenerate (maintainers)

From the monorepo: `cd sdks/openapi-sdk && ./generate.sh`

**Module:** `github.com/beeos-ai/sdk-go` — this tree is mirrored as its own
GitHub repository for third-party `go get`.
