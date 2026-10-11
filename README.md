# github.com/beeos-ai/sdk-go

> **Archived — use [`beeos-cloud-sdks`](https://github.com/beeos-ai/beeos-cloud-sdks) instead.** This module targeted the retired BeeOS OpenAPI gateway and is no longer maintained. The Go client now lives in the `go/` module of `beeos-cloud-sdks`. Existing `github.com/beeos-ai/sdk-go` tags stay fetchable but receive no updates.

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

## Resource Facade (Unreleased)

These additions are local development code, not in the published `v0.8.0` tag.
Existing generated APIs remain accessible through `client.API`.

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
defer cancel()

instance := client.Instance(os.Getenv("BEEOS_INSTANCE_ID"))
if _, err := instance.Mobile.WaitReady(ctx); err != nil { return err }
info, err := instance.Mobile.Info(ctx)
if err != nil { return err }
fmt.Println(info.Data.SupportedActions)

agentID := os.Getenv("BEEOS_AGENT_ID") // Must belong to this instance.
created, err := client.CreateTask(ctx, agentID, beeos.CreateTaskRequest{
    Message: "Open Settings and report the battery level without changing settings",
})
if err != nil { return err }
task := client.Task(agentID, created.Data.TaskId)
result, err := task.Wait(ctx)
if err != nil { return err }
fmt.Println(result.Data.Result)
```

- `client.Instances.Create/Get/List`, `instance.Start/Stop/Restart/Destroy` wrap
  existing lifecycle APIs. Only clean up resources your workflow created.
- `instance.Mobile` and `instance.Computer` bind the existing atomic APIs to one
  instance. Discover capabilities first; unsupported/busy errors are preserved.
- `task.Get/Wait/Cancel/Messages/Events` reuse a task ID. `Wait` requires
  `completed`; `TaskExecutionError.Snapshot` retains non-success terminal data.
  Set `WaitOptions{AllowUnsuccessful: true}` only when any terminal is expected.
  The existing `MobileClient.Run` retains its terminal-only result semantics.
- Local context cancellation/deadline does not cancel or resubmit the durable
  task. Call `task.Cancel` explicitly. Writes are not blindly replayed.
- Event consumers must cancel their context if they stop reading. `end` closes
  the stream; persist offsets and handle `backfill_truncated`. Offsets need not
  be contiguous, and an incomplete EOF frame is not a terminal event.
- `client.Files.Download(ctx, fileID, DownloadOptions{...})` resolves authorized
  metadata and returns bytes with a default 32 MiB cap. It does not use the API
  client's cookies or authentication transport for object storage. Redirects
  and non-HTTPS URLs are rejected unless `AllowInsecureHTTP` explicitly permits
  HTTP for a trusted local fixture. `ExpectedContentType` enables type checks.

These wrappers do not grant new runtime capabilities. A completed task still
requires a separate business-result check, such as PDF parsing after download.

Signed file downloads retry transient connection failures and HTTP 408/429/500/
502/503/504 up to three attempts within the same context deadline. They respect
`Retry-After`, discard partial responses, and fall back to verified HTTPS over
HTTP/1.1 after a transient transport failure. API writes, device actions, auth
failures, invalid files, and redirects are never retried by this helper. Use
`DownloadOptions{MaxAttempts: 1}` to disable download retries.

## Regenerate (maintainers)

From the monorepo: `cd sdks/openapi-sdk && ./generate.sh`

**Module:** `github.com/beeos-ai/sdk-go` — this tree is mirrored as its own
GitHub repository for third-party `go get`.
