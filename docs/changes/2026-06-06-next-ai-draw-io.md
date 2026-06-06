# 新增 Next AI Draw.io 外部工具文档

## 任务背景

用户希望在外部工具目录中整理 `DayuanJiang/next-ai-draw-io` 的安装和使用方式，方便后续将它与 CodeGraph、Understand Anything、`my-codebase-intel` 搭配，用于生成 draw.io 架构图和流程图。

## 根因定位

仓库当前 `external-tools/` 只记录了 CodeGraph，缺少 AI 出图工具的安装、MCP 接入、提示词模板和使用边界说明。`next-ai-draw-io` 的定位不同于 CodeGraph：它负责 draw.io 图生成、编辑、预览和导出，不负责代码事实理解，因此需要单独文档说明职责边界。

## 执行计划

1. 查阅官方 README、MCP Server README 和安全索引，确认安装命令、MCP 配置、工具能力、端口、私有部署配置和版本注意事项。
2. 新增 `external-tools/next-ai-draw-io.md`，整理安装、基础使用、MCP 接入、与 CodeGraph / Understand Anything 配合方式和常见问题。
3. 更新 `external-tools/README.md`，加入工具索引。
4. 更新根目录 `README.md`，补充目录结构和使用建议。
5. 执行静态检查和敏感信息扫描。

## 变更内容

- 新增 `external-tools/next-ai-draw-io.md`：
  - 记录在线 demo、桌面应用、本地源码运行、Docker 运行和 MCP 接入方式。
  - 记录 `codex mcp add drawio -- npx -y @next-ai-drawio/mcp-server@latest` 和通用 MCP JSON 配置。
  - 说明 `PORT`、`DRAWIO_BASE_URL`、`start_session`、`create_new_diagram`、`edit_diagram`、`get_diagram`、`export_diagram` 的用途。
  - 补充 MCP 连接后的实时预览流程：先 `start_session`，再通过带 `?mcp=` 的浏览器页面预览图。
  - 明确 MCP 模式下不需要给 `next-ai-draw-io` 单独配置模型 API Key；模型推理由 Codex / Claude 负责。
  - 补充 `.env` 来源：从官方仓库根目录 `env.example` 复制，或通过 GitHub raw 地址下载。
  - 增加 Docker Compose 示例，使用 `env_file: .env` 管理 Web 应用配置，并可选启动私有 draw.io 服务。
  - 增加 MCP server 旧版本安全提示，建议使用 `@latest` 或 `0.1.19` 及以上版本。
  - 增加与 CodeGraph、Understand Anything、`my-codebase-intel` 的搭配方式和提示词。
- 更新 `external-tools/README.md`：
  - 在工具表中加入 Next AI Draw.io。
- 更新 `README.md`：
  - 在内容说明、目录结构和使用建议中加入 Next AI Draw.io。

## 验证结果

- 已确认 `external-tools/next-ai-draw-io.md` 文件存在。
- 已确认 README、外部工具目录和变更记录中包含 `next-ai-draw-io` / `drawio` 相关入口。
- 已执行敏感信息扫描，未发现真实密钥形态或遗留标记。
- 已执行 `command -v markdownlint`：本机未安装 `markdownlint`，因此未执行 Markdown lint。
- 未执行 `npx @next-ai-drawio/mcp-server@latest`、`codex mcp add ...`、`claude mcp add ...` 或 Docker 启动命令；本次只整理文档。

## 后续建议

- 如需真正接入 Codex，先由用户确认是否允许修改本机 MCP 配置，再执行 `codex mcp add drawio -- npx -y @next-ai-drawio/mcp-server@latest`。
- 如果在安全敏感环境使用，优先配置私有 draw.io 实例，并在仓库文档中只写占位域名。

## 追加变更：整理 MCP 与 Docker 边界

### 任务背景

用户反馈自己是把 `next-ai-draw-io` 装在 Docker 容器中，不是把 Codex 装在容器中，因此原文容易让人误以为 `codex mcp add drawio -- npx ...` 会自动连接容器内安装的 MCP server。用户要求重新整理 MCP 连接过程，不要和 Docker 安装混淆，并把 Docker 安装内容放到文档最下面。

### 根因定位

`external-tools/next-ai-draw-io.md` 原先在 MCP 章节之前放了较长的 Docker 运行说明。虽然文字中提到 Docker 是 Web 应用运行方式，不是 MCP 接入，但读者从顺序上仍容易把“Docker 安装 next-ai-draw-io”和“Codex 通过 stdio 启动 MCP server”混成一条链路。

MCP 注册的关键边界是：`codex mcp add <name> -- <command>` 中的 `<command>` 在 Codex 所在环境执行。如果 MCP server 只安装在容器内，必须通过 `docker exec -i` 或 `docker run --rm -i` 让 Codex 连接容器内 stdio 进程。

### 执行计划

1. 将 `external-tools/next-ai-draw-io.md` 中的 MCP 接入流程前置，先讲清 `npx`、`docker exec -i` 和 `docker run --rm -i` 三种场景。
2. 将 Docker Web 应用和 Docker Compose 运行方式移动到文档底部，并明确它们不等于 MCP 接入。
3. 同步修正 MCP 示例命令，给 `npx` 增加 `-y` 参数，避免首次运行时交互确认。
4. 执行静态检索，确认关键命令和边界说明存在。

### 变更内容

- 更新 `external-tools/next-ai-draw-io.md`：
  - 在结论中新增 MCP 接入和 Docker Web 应用的职责区分。
  - 将 Docker 运行章节移动到文档底部。
  - 在 MCP 推荐路径中新增三类接入方式：
    - Codex 所在环境有 Node / npx：`codex mcp add drawio -- npx -y @next-ai-drawio/mcp-server@latest`。
    - MCP server 装在已运行容器内：`codex mcp add drawio -- docker exec -i <容器名> npx -y @next-ai-drawio/mcp-server@latest`。
    - MCP server 通过镜像一次性启动：`codex mcp add drawio -- docker run --rm -i <镜像名>`。
  - 说明 `docker exec -i` / `docker run --rm -i` 中 `-i` 用于保持 stdio MCP 会话。
  - 将通用 MCP JSON 示例中的 `args` 调整为 `["-y", "@next-ai-drawio/mcp-server@latest"]`。

### 验证结果

- 已执行 `git status --short`：执行前工作区干净；执行后仅包含 `external-tools/next-ai-draw-io.md` 和本变更记录文件。
- 已执行关键内容检索，确认 `codex mcp add drawio`、`docker exec -i`、`docker run --rm -i`、`Docker 运行 Web 应用` 和 `MCP 接入` 等说明已写入。
- 已执行敏感信息和遗留标记扫描，未发现真实密钥形态或遗留任务标记。
- 未执行 `codex mcp add`、`npx`、`docker run` 或 `docker compose up`，本次只修改文档。
