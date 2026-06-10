# planning-with-files 任务生命周期工作流

`planning-with-files` 的价值不是创建三份 Markdown 文件，而是把 AI 协作中的任务现场落到仓库里，让任务在长会话、上下文压缩、跨天恢复或多 agent 接续时仍然可追溯。

本文给出一套更适合工程协作的使用规范：当前任务使用稳定工作区，历史任务显式归档，归档必须由用户授权触发。

## 核心结论

推荐把 `planning-with-files` 设计成任务生命周期工作流，而不是简单的文件模板：

- 当前任务固定使用 `.planning/current/`。
- 历史任务进入 `.planning/archive/YYYY-MM-DD-<topic>/`。
- 当前任务目录不强制加日期，因为它表示活跃状态，不是历史记录。
- 归档目录使用日期前缀，便于排序、审计和复盘。
- AI 不得自动归档；只有用户明确提出归档并确认后，AI 才能执行归档。
- 归档内容默认不参与当前任务上下文恢复。

## 问题背景

`planning-with-files` 常见做法是维护三份文件：

```text
task_plan.md
findings.md
progress.md
```

这套文件很适合承载一个任务的执行现场。但如果一个任务完成后继续在同一套文件里追加新任务，会出现几个问题：

- `task_plan.md` 混入多个目标，AI 难以判断当前任务到底是什么。
- `findings.md` 累积过多历史发现，旧结论可能污染新任务判断。
- `progress.md` 变成流水账，不能快速回答“当前做到哪一步”。
- AI 可能把阶段性完成误判成任务结束，提前归档或切换上下文。

因此，工程化重点不是“要不要保留三份文件”，而是明确它们的生命周期。

## 当前任务与历史任务

当前任务和历史任务的职责不同：

| 类型 | 推荐位置 | 主要作用 | AI 默认是否读取 |
| --- | --- | --- | --- |
| 当前任务 | `.planning/current/` | 保存当前目标、发现、进度和阻塞 | 是 |
| 历史任务 | `.planning/archive/YYYY-MM-DD-<topic>/` | 保存已归档任务的审计材料 | 否 |
| 最终变更记录 | `docs/changes/YYYY-MM-DD-<topic>.md` | 保存最终事实、改动、验证和后续建议 | 按需读取 |

推荐目录结构：

```text
.planning/
  current/
    task_plan.md
    findings.md
    progress.md
  archive/
    2026-06-10-planning-with-files-workflow/
      task_plan.md
      findings.md
      progress.md
docs/
  changes/
    2026-06-10-planning-with-files-workflow.md
```

`.planning/current/` 表示活跃工作区，路径应稳定，方便 AI 和脚本定位。`.planning/archive/` 表示历史记录，目录名使用日期前缀，方便按时间排序和追溯。

## 文件职责

三份规划文件应保持短期任务现场的职责，不替代长期变更记录。

| 文件 | 回答的问题 | 推荐内容 |
| --- | --- | --- |
| `task_plan.md` | 准备怎么做 | 目标、根因、阶段、验证方式、风险边界 |
| `findings.md` | 为什么这么做 | 调研发现、命令输出摘要、方案取舍、上下文依据 |
| `progress.md` | 当前做到哪一步 | checklist、当前状态、阻塞、最近验证结果 |
| `docs/changes/*.md` | 最终实际发生了什么 | 背景、根因、计划、改动、验证、后续建议 |

硬性规则：

- `.planning/current/progress.md` 仅记录当前任务执行状态；`progress.md` 不替代 `docs/changes/`。
- 如果本次会话产生任何仓库文件改动，必须在结束前自动创建或更新 `docs/changes/YYYY-MM-DD-<topic>.md`。
- 不需要等待用户额外要求“生成日期变更文档”，也不要只在最终回复中用文字描述替代文档。
- 变更文档应从 `task_plan.md`、`findings.md` 和 `progress.md` 中提炼最终事实，记录任务背景、根因定位、执行计划、变更内容、验证结果和直接相关的后续建议。

## 任务生命周期

推荐使用以下状态理解当前任务：

```text
active
-> ready_for_review
-> archive_requested
-> archived
```

各状态含义：

| 状态 | 含义 | 谁推动 |
| --- | --- | --- |
| `active` | 当前任务进行中，用户仍可能补充需求 | 用户和 AI |
| `ready_for_review` | AI 认为当前已知需求完成，等待用户反馈 | AI |
| `archive_requested` | 用户明确提出归档，AI 确认归档对象和路径 | 用户触发，AI 确认 |
| `archived` | 用户确认后，AI 执行归档 | AI 执行 |

关键规则：

- AI 可以判断“当前已知计划项已完成”。
- AI 不能单方面判断“任务生命周期已结束”。
- 用户继续补充同一目标的需求时，任务应回到 `active`。
- 用户明确提出归档前，`.planning/current/` 必须继续保留。

## 归档触发规则

归档必须由用户显式触发。AI 不得仅凭计划完成、测试通过、最终回复或任务看似结束，自动归档 `.planning/current/`。

允许触发归档的用户指令包括：

```text
归档当前任务
这个任务归档
完成并归档
开始新任务，先归档当前任务
```

如果用户只是说：

```text
做完了
继续下一个
换个任务
这个先这样
再做一个需求
```

AI 不得直接归档，应先询问是否需要归档。

AI 收到归档指令后，应先确认归档对象和路径，例如：

```text
将归档 `.planning/current/` 到 `.planning/archive/YYYY-MM-DD-<topic>/`。
确认后我再执行归档。
```

只有用户确认后，AI 才能移动文件。

归档执行时应遵守：

- 不覆盖已有归档目录。
- 如果目标目录已存在，追加短后缀，例如 `YYYY-MM-DD-<topic>-2`。
- 归档后重新初始化 `.planning/current/`。
- 后续默认只关注新的 `.planning/current/`。

## 新需求处理规则

用户可能不会一次性列完需求，因此新输入不应默认视为新任务。

推荐判断规则：

| 用户输入 | 处理方式 |
| --- | --- |
| 同一目标的补充、修正、验收反馈 | 继续当前 `.planning/current/` |
| 不同目标、不同文档、不同功能或不同问题排查 | 询问是否新建任务 |
| 用户明确要求归档 | 进入 `archive_requested` |
| 用户明确要求丢弃当前任务 | 先确认，再删除或归档为中止任务 |

示例：

```text
“刚才那套归档规则，再补一条异常处理”
```

这属于同一任务，应继续更新 `.planning/current/`。

```text
“再帮我整理 codex-skills.md 里的安装说明”
```

这可能是新任务，AI 应先确认是否需要归档当前任务或直接切换。

## 可复制的 AGENTS.md 规则

以下规则可放入仓库的 `AGENTS.md`：

````markdown
## Planning-with-files 任务生命周期工作流

使用 `planning-with-files` 时，当前任务固定使用 `.planning/current/`，不在当前任务目录名中强制加入日期。

当前任务文件固定为：

- `.planning/current/task_plan.md`
- `.planning/current/findings.md`
- `.planning/current/progress.md`

`.planning/current/progress.md` 仅记录当前任务执行状态；`progress.md` 不替代 `docs/changes/`。

AI 默认只读取和维护 `.planning/current/`。`.planning/archive/` 仅作为历史审计材料，不参与当前任务上下文恢复；除非用户明确要求追溯历史，否则 AI 不应主动读取归档规划文件。

AI 不得仅凭计划完成、测试通过、最终回复或任务看似结束，自动归档 `.planning/current/`。

归档必须由用户显式触发。允许触发归档的指令包括：

- “归档当前任务”
- “这个任务归档”
- “完成并归档”
- “开始新任务，先归档当前任务”
- 其他明确表达归档意图的指令

AI 收到归档指令后，必须先确认归档对象和归档路径，例如：

```text
将归档 `.planning/current/` 到 `.planning/archive/YYYY-MM-DD-<topic>/`。
确认后我再执行归档。
```

只有用户确认后，AI 才能执行归档。

如果用户继续补充同一目标的需求，AI 应继续更新 `.planning/current/`，不得新建任务或归档。

如果用户切换到明显不同的目标，AI 应先询问是否需要归档当前任务；未经确认不得覆盖、删除或移动 `.planning/current/`。

如果本次任务产生仓库文件改动，完成后仍必须生成或更新 `docs/changes/YYYY-MM-DD-<topic>.md`；`.planning/archive/` 不能替代变更文档。
````

## 参考依据

这套规范主要来自三个工程实践的组合：

- `planning-with-files` 的文件化上下文机制：用 `task_plan.md`、`findings.md`、`progress.md` 保存可恢复的任务现场。
- GitLab Branches 文档把分支描述为隔离工作区，并建议使用结构化命名；这对应当前任务应有稳定、清晰的工作边界。参考：[GitLab Branches](https://docs.gitlab.com/user/project/repository/branches/)。
- MADR 建议把决策记录放入独立目录，并使用 `NNNN-title-with-dashes.md` 这类可排序、可追溯的文件名；这对应历史任务应作为审计材料保存，而不是混入当前任务现场。参考：[MADR](https://adr.github.io/madr/)。
- W3C 的 ISO 日期格式说明推荐使用 `YYYY-MM-DD` 表达国际化日期；这适合归档目录前缀，降低日期歧义并便于排序。参考：[Use international date format (ISO)](https://www.w3.org/QA/Tips/iso-date)。

## 推荐结论

最终推荐采用：

```text
当前任务：.planning/current/
历史归档：.planning/archive/YYYY-MM-DD-<topic>/
长期记录：docs/changes/YYYY-MM-DD-<topic>.md
```

其中最重要的边界是：

```text
AI 负责维护当前任务现场和执行归档动作。
用户负责决定任务生命周期是否结束。
```

只有把这两个责任分开，`planning-with-files` 才能既支持长任务恢复，又避免历史上下文污染当前任务。
