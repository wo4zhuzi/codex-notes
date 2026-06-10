# planning-with-files 任务生命周期工作流

## 任务背景

用户希望将前面对 `planning-with-files` 当前任务、归档、用户确认和 AI 执行边界的讨论整理成一篇正式文章，并作为仓库内可复用的工程规范沉淀。

## 根因定位

原有 `codex-skills.md` 已说明 `planning-with-files` 的基础用途和三份文件分工，但没有明确任务生命周期规则：

- 当前任务是否应该继续追加还是重新创建。
- 当前任务目录是否需要日期前缀。
- 归档由 AI 自动判断还是必须由用户明确触发。
- 归档内容是否参与当前任务上下文恢复。

如果缺少这些规则，AI 可能把阶段性完成误判为任务结束，也可能把历史任务上下文混入当前任务。

## 执行计划

1. 新建 `planning-with-files-workflow.md`，独立说明任务生命周期工作流。
2. 更新 `README.md`，加入新文章的内容说明、目录结构、快速入口和阅读顺序。
3. 更新 `codex-skills.md`，在 `planning-with-files` 小节中增加新文章入口。
4. 运行文档链接和敏感信息检查。

## 变更内容

- 新增 `planning-with-files-workflow.md`：
  - 明确当前任务固定使用 `.planning/current/`。
  - 明确历史任务归档到 `.planning/archive/YYYY-MM-DD-<topic>/`。
  - 定义 `active`、`ready_for_review`、`archive_requested`、`archived` 任务状态。
  - 明确 AI 不得自动归档，必须由用户显式提出归档并确认。
  - 补充新需求继续当前任务、不同目标先询问是否切换的规则。
  - 提供可复制到 `AGENTS.md` 的规则片段。
  - 补充 GitLab Branches、MADR 和 W3C ISO 日期格式作为参考依据。
- 更新 `README.md`：
  - 新增 `planning-with-files` 任务生命周期说明和快速入口。
  - 在目录结构和阅读顺序中加入新文章。
- 更新 `codex-skills.md`：
  - 在 `planning-with-files` 使用规范中增加到新文章的链接。
- 更新 `AGENTS.md`：
  - 将 `planning-with-files` 文件职责规则对齐到 `.planning/current/`。
  - 新增任务生命周期短规则，明确 `.planning/archive/` 默认不参与当前上下文恢复。
  - 明确 AI 不得自动归档，归档必须由用户显式触发并确认。
- 补强 `planning-with-files-workflow.md`：
  - 在“文件职责”中新增硬性规则，明确 `.planning/current/progress.md` 仅记录当前任务执行状态。
  - 明确 `progress.md` 不替代 `docs/changes/`。
  - 明确只要本次会话产生任何仓库文件改动，结束前必须自动创建或更新 `docs/changes/YYYY-MM-DD-<topic>.md`。

## 验证结果

已执行：

```bash
git status --short
rg -n "planning-with-files-workflow|任务生命周期|archive_requested|\\.planning/current|\\.planning/archive" README.md codex-skills.md planning-with-files-workflow.md docs/changes/2026-06-10-planning-with-files-workflow.md
rg -n "TODO|FIXME|your-api-key|sk-" .
git diff --check -- README.md codex-skills.md planning-with-files-workflow.md docs/changes/2026-06-10-planning-with-files-workflow.md
git diff --check -- AGENTS.md docs/changes/2026-06-10-planning-with-files-workflow.md
rg -n "progress.md|docs/changes|不替代|自动创建或更新|任何仓库文件改动" planning-with-files-workflow.md AGENTS.md
```

结果：

- `git status --short` 显示仅包含本次计划内改动：`README.md`、`codex-skills.md`、新增 `planning-with-files-workflow.md` 和本变更记录。
- 新文章和入口链接检查已命中 README、skills 文档、新文章和变更记录中的相关引用。
- 敏感信息扫描仅命中文档中的检查命令示例、既有 `project-task-branch` 示例分支名和 `risk-reviewer` 名称，未发现真实 API Key、Token、私有代理地址或遗留调试标记。
- `git diff --check` 未发现 Markdown 空白错误。
- `markdownlint` 未运行：当前环境 PATH 中未发现 `markdownlint` 命令。
- 已追加检查 `AGENTS.md` 与变更记录的 Markdown 空白错误，结果无异常。
- 已追加检查 `planning-with-files-workflow.md` 与 `AGENTS.md` 中 `progress.md` 和 `docs/changes` 的关键表述，确认两处规则口径一致。

## 后续建议

如果后续决定把这套规则落入仓库强约束，可单独更新 `AGENTS.md`，并保持文章作为详细说明入口。
