# 补充 Codex 核心指令说明

## 任务背景

用户提供了一篇 Codex 新手教程，希望将其中仓库此前未记录的核心指令补充到 `codex-core-commands.md`。经确认，本次只补充 `/permissions`、`/mention`、`/side`、`/fork`、`/new` 和 `/clear`，不新增 `/resume`。

## 根因定位

现有 `codex-core-commands.md` 已覆盖 `/init`、`/status`、`/plan`、`/compact`、`/diff`、`/review` 和 `resume` 相关工作流，但缺少以下新手常用控制能力：

- 权限控制：如何查看或调整读写、命令执行和审批策略。
- 上下文约束：如何指定文件，避免 Codex 在全项目中过度搜索。
- 会话分流：如何区分临时侧聊和真正开新路线。
- 会话切换：如何在任务结束后开启新对话或清理当前会话。

## 执行计划

1. 查看 `git status --short`，确认工作区没有用户已有未提交改动。
2. 阅读 `codex-core-commands.md`，确认已有命令覆盖范围，避免重复补充 `/resume`。
3. 在现有命令说明附近插入缺失章节，保持 Markdown 层级和中文风格一致。
4. 新增本变更文档，记录本次修改背景、根因、计划、内容和验证结果。
5. 执行本地检查，确认没有敏感占位符或无关改动。

AI 自检结论：计划包含根因定位、修改步骤、验证方式和风险边界；本次仅修改目标文档和对应变更记录，不涉及应用源码、配置默认值或自动化测试。

## 变更内容

- 更新 `codex-core-commands.md`：
  - 新增 `/permissions`：说明读写、命令执行和审批策略的用途与保守授权建议。
  - 新增 `/mention`：说明指定文件或上下文的使用场景和示例。
  - 新增 `/side`：说明临时侧聊的适用边界。
  - 新增 `/fork`：说明复制当前会话开新路线的适用场景，并补充 fork 后新旧会话上下文独立发展及 `codex resume` 恢复命令边界。
  - 新增 `/new` 和 `/clear`：说明与 `/compact` 的边界。
- 新增 `docs/changes/2026-05-23-codex-core-commands.md`：
  - 记录本次文档补充过程和验证结果。

## 验证结果

- 已执行 `git status --short`：修改前工作区干净。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。
- 已执行 `git diff -- codex-core-commands.md`：确认核心文档仅新增本次指定命令说明，未新增 `/resume` 章节。
- 已读取 `docs/changes/2026-05-23-codex-core-commands.md`：确认变更记录包含任务背景、根因定位、执行计划、变更内容、验证结果和后续建议。
- 已补充检查 `codex resume --help`：确认 `codex resume` 默认显示选择器，`--last` 直接恢复最近会话，`--all` 会显示所有会话并取消当前目录过滤。
- 本轮追加 `/fork` 恢复说明前已再次执行 `git status --short`：工作区已有 `codex-core-commands.md` 修改和本变更文档未跟踪，未发现其他无关改动。

## 后续建议

- 如果后续确认本机 Codex 还支持 `/model` 等其他未记录命令，可按同样方式单独补充。
- 如需面向纯新手读者优化阅读顺序，可另行调整 `codex-core-commands.md` 的章节组织，但本次不做结构重排。
