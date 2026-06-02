# 2026-06-02 Codex 与 Claude Code 协作文章

## 任务背景

用户希望新增一篇文章，说明 Codex 和 Claude Code 有哪些合作方式，`AGENTS.md` 与 `CLAUDE.md` 如何融合、去重和处理冲突，以及 Codex 使用 `planning-with-files` 生成的待办事项能否由 Claude Code 接续执行。

## 根因定位

Codex 与 Claude Code 协作的主要问题不是工具能力冲突，而是上下文职责未分层：

- `AGENTS.md` 和 `CLAUDE.md` 都可能承载项目规则，容易重复维护。
- 两个文件如果各自完整记录规则，会出现版本漂移和冲突。
- `task_plan.md`、`findings.md`、`progress.md` 是任务状态文件，不等同于工具指令文件。
- Claude Code 不能依赖 Codex 会话历史，接续前必须从仓库文件、计划文件和当前 diff 恢复现场。

## 执行计划

1. 阅读仓库现有 README、上下文管理和 planning-with-files 相关说明，确认文档风格和已有规则。
2. 新增根目录文章 `codex-claude-collaboration.md`。
3. 在 README 中加入新文章入口和阅读顺序。
4. 生成本次日期变更文档。
5. 运行仓库约定检查命令，确认链接和敏感内容。

## 变更内容

- 新增 `codex-claude-collaboration.md`：
  - 说明串行交接、并行分工、主从审查、双工具对照、计划与执行分离五种合作方式。
  - 推荐 `AGENTS.md` 作为共享主源，`CLAUDE.md` 通过 `@AGENTS.md` 导入共享规则并保留 Claude Code 专属补充。
  - 给出重复内容归并规则、冲突优先级和典型冲突处理方式。
  - 说明 Claude Code 接续 Codex 文件型计划前应读取的文件、运行的命令和判断标准。
  - 提供 `CLAUDE.md` 最小模板、Claude 接续提示词模板和 Codex 交接提示词模板。
- 更新 `README.md`：
  - 内容说明中补充 Codex 与 Claude Code 协作主题。
  - 目录结构中加入 `codex-claude-collaboration.md`。
  - 快速入口和使用建议中加入新文章链接。
- 新增 `docs/changes/2026-06-02-codex-claude-collaboration.md` 记录本次会话。

## 验证结果

- 已执行 `rg -n "codex-claude-collaboration|Codex 与 Claude Code" README.md codex-claude-collaboration.md`：确认 README 和新文章中均存在新文档入口与标题。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例、既有示例分支名 `project-task-branch` 和 `risk-reviewer` 名称，未发现真实密钥或遗留调试标记。
- 已执行 `git status --short`：显示本次修改 `README.md`，新增 `codex-claude-collaboration.md` 和本变更记录。

## 后续建议

如果后续真实项目中同时提交 `AGENTS.md` 和 `CLAUDE.md`，建议优先采用本文的统一源方案：`AGENTS.md` 放跨工具共享规则，`CLAUDE.md` 只做导入和 Claude Code 专属补充。
