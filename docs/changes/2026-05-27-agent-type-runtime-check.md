# 补充 agent_type 运行时查看方法

## 任务背景

用户已将自定义 subagent TOML 复制到全局注册目录，并希望在文档中补充“如何查看当前运行时有哪些 `agent_type`”的方法，避免只看注册文件而误判当前 Codex 会话是否已经加载。

## 根因定位

`subagent.md` 已说明如何注册 agent，以及如何确认 `spawn_agent` 可选类型包含自定义 agent，但缺少可复用的运行时查看方法。尤其是 `~/.codex/agents/` 或 `.codex/agents/` 中存在 TOML 只能证明配置文件存在，不能证明当前会话已经加载。

## 执行计划

- 在 `subagent.md` 的“如何确认注册是否生效”部分补充“查看运行时 agent_type”小节。
- 明确区分注册目录文件检查与运行时工具 schema 检查。
- 给出当前会话询问方式、注册文件检查命令和 `codex debug prompt-input` 辅助命令。
- 记录 `codex debug prompt-input` 可能受本地权限、沙箱和 Codex CLI 版本影响。

## 变更内容

- 更新 `subagent.md`：
  - 新增“查看运行时 agent_type”小节。
  - 补充询问当前 Codex 会话的推荐口令：`当前运行时 spawn_agent 暴露了哪些 agent_type？`
  - 补充 `find ~/.codex/agents .codex/agents ...` 命令，用于检查注册文件。
  - 补充 `codex debug prompt-input ... | rg ...` 命令，用于尝试检查模型可见的运行时工具 schema。

## 验证结果

- `git status --short`：确认当前仅包含 `subagent.md` 修改和本变更记录新增。
- `rg -n "agent_type|prompt-input|spawn_agent|~/.codex/agents" subagent.md docs/changes/2026-05-27-agent-type-runtime-check.md`：确认新增说明可检索。
- `rg -n "sk-[A-Za-z0-9_-]{20,}" subagent.md docs/changes/2026-05-27-agent-type-runtime-check.md`：未发现真实密钥形态；普通 `sk-` 会误匹配 `risk-reviewer` 中的文本，不适合用于本次局部检查。
- `command -v markdownlint`：当前环境未安装 `markdownlint`，未执行 Markdown lint。

## 后续建议

如果后续 Codex CLI 增加正式的 agent 列表命令，可再更新 `subagent.md`，将 `codex debug prompt-input` 降级为排障方式。
