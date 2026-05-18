# Git 融入 AI 工作流

本文记录个人项目中 Git 如何融入 AI 编程工作流。目标不是建立团队级重流程，而是保护本地改动、让 AI 的修改可审查、可提交、可回滚。

## 核心结论

推荐默认流程：

```text
查看状态 -> 判断分支或 worktree -> 执行任务 -> 机械验证 -> 生成变更记录 -> diff/review -> AI 准备 commit -> 用户确认 -> AI commit -> 用户单独授权 push
```

分工原则：

- AI 负责检查状态、整理改动、生成 commit message 和提交摘要。
- 用户负责确认是否提交。
- `push` 影响远端，必须单独授权，不跟随 commit 自动执行。

## 开始任务前

AI 执行会修改文件的任务前，先运行：

```bash
git status --short
```

检查重点：

- 当前分支是否适合直接修改。
- 是否存在用户已有未提交改动。
- 本次任务是否会碰到同一批文件。
- 是否有未跟踪文件需要保留。

如果发现用户已有改动，AI 不得回滚、覆盖或混入本次提交。应先说明这些改动存在，再把本次修改限制在目标文件内。

## 什么时候直接在当前分支修改

适合直接在当前分支处理：

- 纯文档补充、错别字、链接修正。
- 小范围配置说明或示例更新。
- 当前工作区干净，任务可快速完成和回滚。
- 用户明确说“就在当前分支改”。

即使直接在当前分支改，也仍要执行验证、检查 diff，并在需要时生成 `docs/changes/YYYY-MM-DD-<topic>.md`。

## 什么时候新开分支

建议新开分支：

- 新功能、bugfix、重构、依赖升级。
- 涉及多个文件或多个模块。
- 需要 PR、代码审查或 CI 验证。
- 任务可能持续多轮对话或跨天。
- 当前分支是 `main`、`master`、`production` 等长期分支。

分支命名建议：

```text
docs/git-workflow
fix/login-error
feat/spec-mode
refactor/config-loader
```

AI 不应在未授权时主动切分支或创建分支。可以先建议分支名，等待用户确认后执行。

## 什么时候使用 git worktree

`git worktree` 适合把任务隔离到另一个目录，避免污染当前工作区。

建议使用 worktree：

- 当前分支已有未提交改动，但又要开始另一个任务。
- 需要并行比较两个实现方向。
- 夜间委托 AI 执行较大任务，希望主工作区保持不动。
- 一个任务需要长时间运行验证或保留中间状态。

不必使用 worktree：

- 当前工作区干净。
- 只是小文档、小配置或一次性修正。
- 任务明确延续当前上下文。

推荐命令形态：

```bash
git worktree add ../project-task-branch -b task-branch
```

使用 worktree 后，AI 应在新目录中重新确认 `pwd`、`git status --short` 和当前分支，避免在错误目录修改。

## AI 如何准备 commit

任务完成后，AI 应自动进入 commit 准备阶段，但不能自动执行 commit。

准备 commit 前应检查：

```bash
git status --short
git diff
```

同时确认：

- 已运行必要验证命令，并记录结果。
- 已生成或更新 `docs/changes/YYYY-MM-DD-<topic>.md`。
- 本次提交只包含当前任务相关文件。
- 没有真实 API Key、Token、私有代理地址或内部服务 URL。
- 没有混入用户已有未提交改动。

AI 汇报格式建议：

```text
准备提交以下文件：
- <path>：<改动摘要>

验证结果：
- <命令>：通过 / 失败原因

建议 commit message：
<动词 + 对象>

本次不会提交：
- <用户已有改动或无关文件>

请确认是否提交。
```

用户明确回复 `yes`、`确认提交`、`commit` 或类似授权后，AI 才能执行：

```bash
git add <本次任务相关文件>
git commit -m "<commit message>"
```

不要使用 `git add .` 作为默认做法。应显式添加本次任务相关文件，降低混入无关改动的风险。

## Commit message 规则

个人项目默认使用简短中文动宾结构：

```text
补充 Git 工作流规范
更新 AI 提交边界
修正 spec 文档链接
```

如果项目已有 Conventional Commits 或团队规范，应服从项目规范，例如：

```text
docs: update git workflow
fix: handle login timeout
```

AI 生成 commit message 时，应根据真实 diff 总结，不要照抄用户最初描述，也不要夸大未实现内容。

## Push 授权边界

AI 默认不得执行 `git push`。

允许 push 的条件：

- 用户单独明确授权，例如“push 当前分支到 origin”。
- 当前分支是任务分支，不是 `main`、`master`、`production` 等长期分支，除非用户明确要求。
- AI 已说明将推送的远端和分支。
- 工作区已确认干净，且最近 commit 是本次任务提交。

推荐授权语句：

```text
确认可以 push 到 origin 当前分支。
```

或者：

```text
请 push 当前任务分支到 origin，先不要创建 PR。
```

push 后如果需要 PR，AI 可以准备 PR 描述；是否创建 PR 也需要用户明确授权。

## PR 描述模板

```md
## 变更摘要
- 

## 影响范围
- 

## 验证结果
- 

## 风险与未覆盖项
- 

## 备注
- 是否涉及配置、迁移、发布或回滚：否/是，说明：
```

PR 描述应基于实际 diff 和验证结果，不应包含未执行的测试或未完成的功能。

## 推荐口令

任务完成后，让 AI 准备提交：

```text
请执行 commit 前检查，生成提交信息和提交摘要，我确认后你再提交。
```

确认提交：

```text
确认，按你生成的提交信息提交。
```

只提交不推送：

```text
提交本次改动，不要 push。
```

单独授权 push：

```text
确认可以 push 到 origin 当前分支。
```
