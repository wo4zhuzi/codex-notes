# 2026-05-18 Git 融入 AI 工作流

## 任务背景

用户希望把 Git 正式纳入 AI 编程工作流，明确什么时候新开分支、什么时候使用 `git worktree`、AI 如何准备 commit、用户如何授权 commit，以及 `push` 是否应交给 AI 执行。

## 根因定位

当前仓库已有 `codex-core-commands.md` 中的 Git / commit / PR 操作边界，但内容偏摘要，缺少以下可直接执行的规则：

- 分支和 `git worktree` 的选择条件。
- AI 准备 commit 时应该汇报哪些信息。
- 用户确认后 AI 才能执行 commit 的固定流程。
- `push` 必须单独授权的边界。
- PR 描述模板和常用触发口令。

这些缺口会导致后续 AI 执行任务时，只知道“不默认 commit”，但不知道如何安全地准备提交。

## 执行计划

1. 新增独立文档 `git-workflow.md`，作为完整 Git + AI 协作手册。
2. 在 `AGENTS.md` 中加入 Git 与提交权限硬规则摘要。
3. 更新 `README.md`、`ai-workflow.md` 和 `codex-core-commands.md`，建立入口并避免规则分散。
4. 运行文档链接、敏感信息和工作区状态检查。

AI 自检结论：本次为文档工作流补充，不涉及应用代码、构建脚本或真实凭据；风险主要是规则重复或与已有边界冲突，因此采用“AGENTS 放硬规则，git-workflow 放细节，核心命令文档放入口”的结构。

## 变更内容

- 新增 `git-workflow.md`：
  - 记录 Git 融入 AI 工作流的默认流程。
  - 补充分支、`git worktree`、commit 准备、push 授权和 PR 描述模板。
  - 明确 AI 可以准备 commit，但必须等待用户确认后才能提交。

- 更新 `AGENTS.md`：
  - 增加 Git 与提交权限硬规则。
  - 明确 AI 不得混入用户已有未提交改动。
  - 明确 `push` 必须单独授权。

- 更新 `README.md`：
  - 增加 `git-workflow.md` 到目录结构、快速入口和推荐阅读顺序。

- 更新 `ai-workflow.md`：
  - 在总纲入口中增加 Git 工作流链接。
  - 把 commit 准备阶段纳入日常默认流程。

- 更新 `codex-core-commands.md`：
  - 在推荐工作流和 Git 边界章节增加 `git-workflow.md` 入口。
  - 补充 AI 准备 commit 的汇报模板。

## 验证结果

已执行：

```bash
git status --short --untracked-files=all
rg -n "\[.*\]\(([^)]+)\)" README.md ai-workflow.md codex-core-commands.md git-workflow.md docs/changes/2026-05-18-git-workflow.md
rg -n "TODO|FIXME|your-api-key|sk-" .
```

结果：

- 链接扫描能看到 `git-workflow.md` 的新增入口和相关文档引用。
- 敏感信息扫描只命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥。
- `git status --short --untracked-files=all` 显示当前工作区仍有若干历史未跟踪文件；本次不回滚、不清理这些文件。

## 后续建议

- 如果后续要把这套规则用于新项目，可在 `/init` 后把 `AGENTS.md` 中的 Git 与提交权限段落复制到新项目。
- 如果进入团队协作场景，再补充 CI 必须通过、分支保护、PR 审批和发布回滚规则。
