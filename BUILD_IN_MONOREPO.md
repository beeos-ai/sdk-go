# 在 monorepo 中生成/更新

`openapi-gateway` 合同 → `../beeos-ai-sdk-go/`，`module github.com/beeos-ai/sdk-go`：

```bash
cd sdks/openapi-sdk
npm install
./generate.sh
```

合同已与主 Gateway 解耦（见 `docs/adr/001-openapi-gateway-bff.md`），不再区分
「BFF 子集 / 全量」。独立发版：将本目录同步到 `github.com/beeos-ai/sdk-go`。
