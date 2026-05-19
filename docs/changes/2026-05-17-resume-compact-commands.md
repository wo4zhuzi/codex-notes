# 2026-05-17 补充 resume 与 compact 指令

## 任务背景

用户指出 `codex-core-commands.md` 缺少两个重要指令：`resume` 和 `compact`。典型场景是前一天任务未完成，关机后第二天需要通过 `codex resume --last` 回到项目继续。

## 根因定位

现有核心指令已覆盖 `/init`、`/status`、`/plan`、`/collab`、`/diff` 和 `/review`，但缺少跨会话恢复和长会话上下文压缩说明。这样会导致多天任务缺少明确恢复入口，也没有在上下文接近上限时主动压缩关键决策的规则。

## 执行计划

1. 阅读 `AGENTS.md`、`codex-core-commands.md` 和现有 `docs/changes/` 记录，确认当前工作流和文档规则。
2. 在推荐工作流中加入 `codex resume --last` 和 `/compact`。
3. 新增 `resume` 章节，说明恢复上次会话的命令、适用场景和恢复后的检查点。
4. 新增 `/compact` 章节，说明压缩上下文的触发时机、压缩前应保留的信息，以及它不能替代日期变更文档。
5. 运行本地检查，确认无敏感信息和无关改动。

AI 自检结论：计划已覆盖根因定位、修改步骤、验证方式和风险边界；本次只修改核心指令文档并新增日期变更记录，不涉及配置模板或安装命令。

## 变更内容

- `codex-core-commands.md`：在推荐工作流中加入 `codex resume --last` 和 `/compact`，并新增对应说明章节。
- `docs/changes/2026-05-17-resume-compact-commands.md`：记录本次会话背景、根因、计划、变更和验证。

追加更新：

- `codex-core-commands.md`：补充本机 Codex session 存储位置，记录 `~/.codex/sessions/YYYY/MM/DD/rollout-*.jsonl`、`session_index.jsonl`、`history.jsonl` 和本地 SQLite 状态文件的作用与注意事项。

## 验证结果

- `git status --short --untracked-files=all`：显示 `codex-core-commands.md` 已修改，`.gitignore`、`docs/changes/2026-05-16-codex-workflow.md` 和本文件为未跟踪文件；本次未回滚已有未提交内容。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中记录的检查命令示例，未发现真实密钥或遗留占位符。

追加验证：

- `find ~/.codex/sessions -maxdepth 5 -type f`：确认本机 session 文件位于 `~/.codex/sessions/YYYY/MM/DD/rollout-*.jsonl`。
- `ls -lh ~/.codex/session_index.jsonl ~/.codex/history.jsonl ~/.codex/state_5.sqlite ~/.codex/logs_2.sqlite`：确认索引、历史、状态和日志文件存在。

## 后续建议

后续跨天继续任务时，优先使用 `codex resume --last` 恢复最近会话；长任务暂停前先 `/compact`，任务产生文件改动后仍要更新 `docs/changes/` 日期记录。
