# Codex Skills

## 结论

当前 Codex Skills 机制已经可用。Skills 不需要在 `~/.codex/config.toml` 中额外添加 `skills = true` 开关；是否触发取决于用户任务是否匹配某个 skill 的触发条件。

## 当前可用 Skills

当前会话识别到的个人安装 skills：

| Skill | 来源 | 适用场景 | 典型触发方式 |
| --- | --- | --- | --- |
| `brainstorming` | `obra/superpowers` | 需求澄清、方案发散、设计评审 | 创建功能、修改行为、需要技术取舍 |
| `planning-with-files` | `OthmanAdi/planning-with-files` | 文件型任务计划、过程笔记和进度跟踪 | 长任务规划、跨上下文恢复、生成 `task_plan.md` |
| `writing-plans` | `obra/superpowers` | 将已确认需求拆成可执行计划 | 根据方案生成实现计划、多步骤任务规划 |
| `systematic-debugging` | `obra/superpowers` | bug、测试失败、构建失败、异常行为排查 | 帮我排查问题、测试失败了、定位根因 |
| `react-execution` | 自定义推荐 | 执行阶段小步反馈循环 | 执行计划、接口兼容、供应商兼容、未知代码库改动 |
| `verification-before-completion` | `obra/superpowers` | 完成前强制验证，避免未验证就声称完成 | 准备说完成、修好了、测试通过前 |

当前会话识别到的系统内置 skills：

| Skill | 适用场景 | 典型触发方式 |
| --- | --- | --- |
| `imagegen` | 生成或编辑位图图片、插画、纹理、精灵图、透明背景素材 | 要求生成图片、编辑图片、创建视觉资产 |
| `openai-docs` | 查询 OpenAI 产品、模型、API、迁移和提示词官方文档 | 询问 OpenAI API、模型选择、模型升级、官方文档 |
| `plugin-creator` | 创建 Codex 插件目录和 `.codex-plugin/plugin.json` | 要求创建、脚手架化、更新 Codex plugin |
| `skill-creator` | 创建或更新 Codex skill | 要求创建新 skill、修改已有 skill、生成 skill 结构 |
| `skill-installer` | 安装 Codex skills | 要求列出可安装 skills，或从仓库安装 skill |

## 存放位置

系统 skills 位于：

```text
~/.codex/skills/.system/
```

常见结构：

```text
~/.codex/skills/.system/{skill-name}/SKILL.md
~/.codex/skills/.system/{skill-name}/scripts/
~/.codex/skills/.system/{skill-name}/references/
~/.codex/skills/.system/{skill-name}/assets/
```

`SKILL.md` 是 skill 的主说明文件，里面定义该 skill 的使用规则、触发方式和执行流程。

### 项目级 Skills

除了全局安装到 `~/.codex/skills/` 的 skills，也可以在项目内放置只对当前仓库生效的项目级 skill。

常见项目级结构：

```text
.agents/skills/{skill-name}/SKILL.md
.agents/skills/{skill-name}/scripts/
.agents/skills/{skill-name}/references/
.agents/skills/{skill-name}/assets/
skills-lock.json
```

三类 skill 的区别：

| 类型 | 典型路径 | 生效范围 |
| --- | --- | --- |
| 系统 skill | `~/.codex/skills/.system/{skill-name}/SKILL.md` | Codex 内置能力 |
| 全局 skill | `~/.codex/skills/{skill-name}/SKILL.md` | 本机所有项目 |
| 项目级 skill | `.agents/skills/{skill-name}/SKILL.md` | 当前仓库 |

项目级 skill 适合沉淀只服务当前项目的工作流，例如已有项目重设计、领域专用审查、项目专属发布检查、内部架构约束等。

### 个人自建 Skills

个人自建 skill 适合沉淀自己的固定工作流，例如把多个第三方 skills、MCP 工具、外部 CLI 和验证步骤组合成一条可复用流程。

推荐统一使用 `my-` 前缀区分个人自建 skill 和三方 skill：

```text
~/.codex/skills/my-codebase-intel/SKILL.md
~/.codex/skills/my-debug-flow/SKILL.md
.agents/skills/my-project-flow/SKILL.md
```

三方 skill 保持原名，例如 `brainstorming`、`writing-plans`、`systematic-debugging`、`understand`。不要把个人 skill 写进 `~/.codex/skills/.system/`。

完整命名、目录、元数据和组合型工作流模板见 [个人自建 Skills](./personal-skills.md)。

本仓库的个人 skill 模板集中放在 `my-skills/`，例如 `my-skills/my-codebase-intel/SKILL.md`。该目录是模板库，不是 Codex 自动发现目录；真实启用时仍需复制或链接到 `~/.codex/skills/{skill-name}/` 或 `.agents/skills/{skill-name}/`。

`skills-lock.json` 是项目级 skills 的锁文件或清单，作用类似依赖锁文件：

- 记录当前项目安装或引用了哪些 skills。
- 帮助后续会话或团队成员恢复一致的项目级 skill 状态。
- 降低只靠目录扫描导致的来源、版本或缓存不一致风险。

如果 `skills-lock.json` 由工具自动生成，通常应和 `.agents/skills/` 一起提交。但提交前要检查是否包含本机绝对路径、真实 token、私有代理地址或内部服务 URL。

`AGENTS.md` 不负责注册 skill。Codex 发现 skill 主要依赖运行时加载约定目录和锁文件；`AGENTS.md` 更适合写“本项目有哪些重要 skill、什么时候优先使用”，用于提高触发稳定性。

推荐在 `AGENTS.md` 中只写短规则，不复制完整 `SKILL.md` 内容：

```markdown
## 项目级 Skills

本项目包含项目级 skill：

- `redesign-existing-projects`：用于现有项目 UI / 交互重设计任务。

当用户要求重设计现有项目、优化已有页面体验、在不重写业务逻辑的前提下调整 UI 时，优先使用该 skill。
```

## 使用方式

### `planning-with-files` 使用规范

`planning-with-files` 适合需要把任务计划、过程笔记和执行状态落到文件中的长任务。它的价值不是替代 `/plan` 或 `docs/changes/`，而是让任务在中断、压缩上下文或跨天恢复时仍有清晰现场。

关于当前任务目录、显式归档、任务状态和 `docs/changes/` 边界，见 [planning-with-files 任务生命周期工作流](./planning-with-files-workflow.md)。

推荐在以下场景使用：

- 多步骤文档或代码改动，需要持续跟踪执行状态。
- 任务可能跨上下文、跨天或由不同 agent 接续。
- 需要把调研结论、方案取舍和执行进度拆开保存。
- 用户明确要求“把计划写到文件里”“生成任务计划和执行记录”。

显式触发示例：

```text
使用 planning-with-files skill，帮我规划并跟踪这个任务。
```

```text
用 planning-with-files 把这个长任务拆成 task_plan.md、findings.md 和 progress.md。
```

自然触发示例：

```text
这个任务可能明天继续，帮我把计划、调研笔记和当前进度都写到文件里。
```

```text
请生成可恢复执行的任务计划，并在过程中维护进度记录。
```

三个文件的推荐分工：

| 文件 | 职责 | 典型内容 | 生命周期 |
| --- | --- | --- | --- |
| `task_plan.md` | 执行前计划 | 目标、根因、执行步骤、验证方式、风险边界 | 任务开始前生成，执行中必要时更新 |
| `findings.md` | 过程上下文 | 调研发现、命令输出摘要、设计取舍、未采纳方案、注意事项 | 任务全过程维护，可作为后续参考 |
| `progress.md` | 执行中状态 | checklist、当前阻塞、下一步、已完成节点 | 当前任务执行中维护，默认不作为最终总结 |

推荐的 `progress.md` 内容保持短小：

```markdown
# Progress

- [x] 阅读项目规则和相关文档
- [x] 定位需要修改的文件
- [ ] 修改文档
- [ ] 运行验证命令
- [ ] 生成 docs/changes 记录

## 当前阻塞

无

## 下一步

修改目标文档并检查 diff。
```

和 `docs/changes/` 的边界：

- `task_plan.md` 回答“准备怎么做”。
- `findings.md` 回答“为什么这么做”。
- `progress.md` 回答“现在做到哪一步”。
- `docs/changes/YYYY-MM-DD-<topic>.md` 回答“最终实际发生了什么”。

因此，`progress.md` 只用于当前任务执行跟踪，不替代 `docs/changes/`。只要任务产生仓库文件改动，结束前仍必须从 `task_plan.md`、`findings.md` 和 `progress.md` 中提炼最终事实，生成或更新 `docs/changes/YYYY-MM-DD-<topic>.md`。

推荐采用这种分工，而不是删除 `progress.md` 或把它合并进 `docs/changes/`：

- 删除 `progress.md` 会削弱长任务恢复能力。
- 把 `progress.md` 合并进 `docs/changes/` 会让长期变更记录混入临时 checklist 和阻塞状态。
- 保留 `progress.md` 但限定为短期执行状态，能同时兼顾恢复现场和长期记录质量。

## 安装 Skills

### 推荐方式：通过 `skill-installer` 安装

Codex 内置了 `skill-installer`，用于从官方 skills 仓库或其他 GitHub 仓库安装 skill。

官方 curated skills 默认来源：

```text
https://github.com/openai/skills/tree/main/skills/.curated
```

安装后的目标目录：

```text
~/.codex/skills/{skill-name}
```

安装完成后需要重启 Codex，新的 skill 才会被当前会话识别。

### 查看官方可安装列表

在 Codex 对话中直接说：

```text
使用 skill-installer 列出可安装的 skills。
```

也可以在命令行执行：

```shell
python3 ~/.codex/skills/.system/skill-installer/scripts/list-skills.py
```

### 安装官方 curated skill

如果 skill 在官方 curated 列表中，可以使用以下命令安装：

```shell
python3 ~/.codex/skills/.system/skill-installer/scripts/install-skill-from-github.py \
  --repo openai/skills \
  --path skills/.curated/{skill-name}
```

示例：

```shell
python3 ~/.codex/skills/.system/skill-installer/scripts/install-skill-from-github.py \
  --repo openai/skills \
  --path skills/.curated/security-threat-model
```

### 安装第三方或自定义 skill

如果 skill 来自其他 GitHub 仓库，需要提供仓库和 skill 路径：

```shell
python3 ~/.codex/skills/.system/skill-installer/scripts/install-skill-from-github.py \
  --repo {owner}/{repo} \
  --path {path/to/skill}
```

也可以使用 GitHub URL：

```shell
python3 ~/.codex/skills/.system/skill-installer/scripts/install-skill-from-github.py \
  --url https://github.com/{owner}/{repo}/tree/{ref}/{path/to/skill}
```

私有仓库需要本机已有 GitHub 凭据，或配置 `GITHUB_TOKEN` / `GH_TOKEN`。

### `last30days` skill 安装示例

`last30days` 是第三方近 30 天社区研究 skill，适合调研人物、公司、产品、AI 工具趋势、竞品对比和近期舆情。

它和工程化 skills 的属性不同：`brainstorming`、`writing-plans`、`systematic-debugging` 主要服务项目执行和交付闭环；`last30days` 更偏外部信息渠道、趋势研究和个人知识库上游输入，因此单独拆成主题文档记录。

官方对 Codex、Cursor、Gemini CLI 等 Agent Skills hosts 的推荐安装方式是：

```bash
npx skills add mvanhorn/last30days-skill -g -a codex
```

如果沿用 Codex 内置 `skill-installer`，安装路径应指向仓库中的 skill 目录：

```shell
python3 ~/.codex/skills/.system/skill-installer/scripts/install-skill-from-github.py \
  --repo mvanhorn/last30days-skill \
  --path skills/last30days
```

验证：

```shell
test -f ~/.codex/skills/last30days/SKILL.md && echo "installed"
```

安装后重启 Codex。完整配置、API Key、输出目录、HTML brief 和最佳使用方式见 [last30days Skill 使用笔记](./last30days-skill.md)。

### 安装后验证

1. 确认目录存在：

```shell
find ~/.codex/skills -maxdepth 2 -type d -name "{skill-name}"
```

2. 确认主文件存在：

```shell
test -f ~/.codex/skills/{skill-name}/SKILL.md && echo "installed"
```

3. 重启 Codex。
4. 新会话中使用该 skill 的名称或匹配触发条件。

### `brainstorming` skill 安装示例

当前官方 curated skills 列表中没有发现 `brainstorming`。

因此不能直接使用以下路径安装：

```text
skills/.curated/brainstorming
```

已确认可从以下第三方仓库安装：

```text
https://github.com/obra/superpowers/blob/main/skills/brainstorming/SKILL.md
```

注意：浏览器链接指向 `SKILL.md` 文件，但安装时要使用 skill 所在目录：

```text
skills/brainstorming
```

安装命令：

```shell
python3 ~/.codex/skills/.system/skill-installer/scripts/install-skill-from-github.py \
  --repo obra/superpowers \
  --path skills/brainstorming
```

安装完成后，本地目录应为：

```shell
~/.codex/skills/brainstorming
```

验证：

```shell
test -f ~/.codex/skills/brainstorming/SKILL.md && echo "installed"
```

安装完成后重启 Codex，新的 `brainstorming` skill 才会被会话识别。

### 程序员推荐安装组合

推荐先单独安装以下 Superpowers skills，而不是一开始安装完整 Superpowers 插件：

```text
brainstorming
writing-plans
systematic-debugging
verification-before-completion
```

取舍依据：

- 覆盖日常开发最常见闭环：需求澄清、计划拆解、根因排查、完成前验证。
- 不强制引入完整 Superpowers 的重流程，例如 worktree、强制 TDD、subagent 驱动开发和分支收尾。
- 与本仓库核心工作流兼容：先计划、定位根因、执行后验证和记录上下文。

安装命令：

```shell
python3 ~/.codex/skills/.system/skill-installer/scripts/install-skill-from-github.py \
  --repo obra/superpowers \
  --path skills/brainstorming \
  --path skills/writing-plans \
  --path skills/systematic-debugging \
  --path skills/verification-before-completion
```

如已安装其中某个 skill，安装脚本可能提示目标目录已存在；这种情况可只安装缺失项。

验证：

```shell
find ~/.codex/skills -maxdepth 2 -name SKILL.md
```

当前已安装：

```text
~/.codex/skills/brainstorming/SKILL.md
~/.codex/skills/writing-plans/SKILL.md
~/.codex/skills/systematic-debugging/SKILL.md
~/.codex/skills/verification-before-completion/SKILL.md
```

### 分场景推荐

不同任务推荐显式点名不同 skill，减少自动触发歧义：

| 场景 | 推荐 skill | 用法 |
| --- | --- | --- |
| 需求澄清、方案发散、技术取舍 | `brainstorming` | 先确认目标、边界和取舍，再进入计划或实现。 |
| 长任务计划、跨上下文恢复 | `planning-with-files` | 生成 `task_plan.md`、`findings.md` 和 `progress.md`。 |
| 已确认需求拆成执行计划 | `writing-plans` | 将方案拆成可执行步骤和验证命令。 |
| bug、异常、测试失败、构建失败 | `systematic-debugging` | 先定位根因，再提出修复和验证路径。 |
| 执行计划、接口兼容、供应商兼容、未知代码库改动 | `react-execution` | 按 `Goal -> Observe -> Decide -> Act -> Verify -> Adjust -> Final` 小步推进。 |
| 准备声称完成、修好或测试通过前 | `verification-before-completion` | 先运行最新验证命令，再汇报结论。 |

`react-execution` 适合作为执行阶段 skill，而不是替代 `brainstorming`、`writing-plans` 或 `systematic-debugging`。推荐组合：

```text
复杂新功能：brainstorming -> writing-plans/planning-with-files -> react-execution -> verification-before-completion
Bug / 兼容问题：systematic-debugging -> react-execution -> verification-before-completion
普通计划执行：writing-plans -> react-execution -> verification-before-completion
```

### `agent-evaluation` 多场景 skill 设计

`agent-evaluation` 适合封装为一个全局 skill，用于在任务执行中或最终总结前判断当前结果是否具备交付证据。它的职责是做验收判断，不替代测试、`/diff` 或 `/review`：

- 测试、lint、typecheck、build 和 dry-run 负责提供机械验证。
- `agent-evaluation` 负责把验证证据组织成阶段性评估或交付评估。
- `/diff` 负责核对真实文件改动。
- `/review` 负责审查缺陷、风险和测试缺口。

推荐使用“一个 `SKILL.md` 总入口 + 多个 `references/*.md` 场景标准”的结构：

```text
~/.codex/skills/agent-evaluation/
├── SKILL.md
└── references/
    ├── docs.md
    ├── bugfix.md
    ├── feature.md
    ├── refactor.md
    ├── config.md
    ├── tool-integration.md
    ├── code-review.md
    └── high-risk.md
```

本仓库已提供可复制的 reference 模板，位置为 `docs/agent-evaluation-references/`。创建全局 skill 时，可将这些文件复制到 `~/.codex/skills/agent-evaluation/references/`，或让 `skill-creator` 参考这些文件生成对应内容。

这样做的原因：

- `SKILL.md` 保持短小，只负责判断阶段、项目风险和任务类型。
- 场景细节放入 `references/`，只在需要时读取，避免每次触发都加载全部标准。
- 不拆成多个独立 skill，避免触发边界重叠、metadata 膨胀和维护成本上升。

场景边界建议：

| 场景 | 评估边界 | 必须验证 | 停止线 |
| --- | --- | --- | --- |
| 文档任务 | 内容准确、链接存在、结构一致、无敏感信息 | `git status`、敏感信息扫描、链接目标、变更记录 | 真实密钥、链接目标不存在、内容与现有规则冲突 |
| Bugfix | 根因和修复必须能解释原始失败 | 复现失败、失败用例通过、相关测试通过 | 无法复现、根因无证据、连续 3 轮验证失败 |
| 新功能 | 满足需求边界和关键路径 | 测试、构建、关键流程验收 | 需求边界不清、核心行为无测试或验收标准 |
| 重构 | 外部行为不变，结构变化不替代行为验证 | 调用方检查、相关测试、diff 范围 | public API 意外变化、无关大面积重排 |
| 配置任务 | 格式合法、默认值清楚、兼容影响明确 | 配置解析、dry-run 或加载测试、敏感信息扫描 | 生产默认值变化未确认、真实 token 或私有地址 |
| RAG / MCP / Function Calling | 工具链可控、失败可解释、权限边界明确 | 正常路径、参数错误、工具失败、权限不足 | 工具输出和回答冲突、权限边界不清 |
| Code Review | 基于真实 diff 找行为风险 | 文件行号、严重度、测试缺口 | 未看 diff、泛泛评价、无证据结论 |
| 高风险生产变更 | 权限、安全、数据、支付、发布必须严格验收 | 测试、dry-run、回滚方案、审计或日志影响、人类确认 | 生产影响不清、回滚方案缺失、涉及真实数据但无确认 |

高风险场景不是独立替代其他场景，而是叠加标准。例如生产配置变更应同时参考 `references/config.md` 和 `references/high-risk.md`。

使用 `skill-creator` 创建该 skill 时，可使用以下提示词：

```text
使用 skill-creator，帮我创建一个个人 Codex skill：agent-evaluation。

目标：
这个 skill 用于在 AI Agent 完成任务前，根据任务阶段、项目风险和改动类型选择合适的评估标准，判断当前结果是可交付、部分完成，还是必须停止交给人。

请按以下结构创建：
- ~/.codex/skills/agent-evaluation/SKILL.md
- ~/.codex/skills/agent-evaluation/references/docs.md
- ~/.codex/skills/agent-evaluation/references/bugfix.md
- ~/.codex/skills/agent-evaluation/references/feature.md
- ~/.codex/skills/agent-evaluation/references/refactor.md
- ~/.codex/skills/agent-evaluation/references/config.md
- ~/.codex/skills/agent-evaluation/references/tool-integration.md
- ~/.codex/skills/agent-evaluation/references/code-review.md
- ~/.codex/skills/agent-evaluation/references/high-risk.md

SKILL.md 要求：
1. frontmatter name 为 agent-evaluation。
2. description 明确说明：当用户要求验收、完成前检查、评估 Agent 输出、判断是否可交付、检查验证证据、区分阶段性评估和交付评估时触发。
3. 主体保持精简，只包含：
   - 先判断当前是阶段性评估还是交付评估。
   - 再判断项目风险：个人文档、普通应用、生产核心。
   - 再判断改动类型：文档、bugfix、新功能、重构、配置、工具集成、review、高风险。
   - 根据类型读取对应 references 文件。
   - 最终输出：目标、证据、验证命令、未验证项、风险、交付状态。
4. 不要创建 README、安装说明或额外无关文件。

references 文件要求：
每个文件只写该场景的评估边界、必须验证项、停止线和最终输出要求。
默认使用简体中文。
可参考当前仓库的 docs/agent-evaluation-references/ 目录生成。
```

使用示例：

```text
使用 agent-evaluation skill，对当前任务做阶段性评估。
```

```text
使用 agent-evaluation skill，对本次任务做交付评估。
```

```text
使用 agent-evaluation skill，按 bugfix 场景评估本次修复。
```

### 切换为完整 Superpowers 插件

如果希望使用完整 Superpowers 方法论，可以安装完整 Superpowers 插件。完整插件会包含 `brainstorming`、`writing-plans`、`systematic-debugging`、`test-driven-development`、`using-git-worktrees`、`subagent-driven-development`、`requesting-code-review` 等一整套流程。

Codex CLI 安装方式：

```text
/plugins
```

在插件搜索界面搜索：

```text
superpowers
```

然后选择 `Install Plugin`。

Codex App 安装方式：

```text
Plugins -> Coding -> Superpowers -> +
```

注意事项：

- 完整 Superpowers 插件流程更重，会更积极地触发规划、TDD、review、worktree 和分支收尾等规范。
- 如果安装完整 Superpowers 插件，建议删除或停用单独安装的同名 skills，避免 `brainstorming` 等同名 skill 出现重复来源导致触发行为不稳定。
- 如果只需要轻量工程闭环，优先保留单独安装的四个 skills。

### 特殊安装：Understand Anything

`Understand Anything` 是跨平台项目理解工具。它在 Claude Code 中通过 plugin marketplace 安装；在 Codex 中推荐通过官方安装脚本把相关 skills 链接到 Codex 可发现的 skills 目录。

Codex 安装方式对应官网表格中的 `install.sh codex`：

```bash
curl -fsSL https://raw.githubusercontent.com/Lum1104/Understand-Anything/main/install.sh | bash -s codex
```

安装脚本会执行本机文件改动，建议用户先阅读脚本内容，再自行运行。AI 不应在未获得明确授权时自动执行该安装命令。

Codex 侧安装后的关键位置：

```text
~/.understand-anything/repo
~/.understand-anything-plugin
~/.agents/skills
```

其中 `~/.agents/skills` 是 Codex 可发现的用户级 skills 目录。安装脚本会把 `understand-anything-plugin/skills/` 下的能力按 skill 链接到这里。安装完成后需要重启 Codex，新的 Understand Anything skills 才会在会话中生效。

更新：

```bash
~/.understand-anything/repo/install.sh --update
```

卸载 Codex 侧安装：

```bash
~/.understand-anything/repo/install.sh --uninstall codex
```

安装后可用下面的只读命令确认链接状态：

```bash
find ~/.agents/skills -maxdepth 2 -type f -name SKILL.md | sort
```

如果安装后 Codex 没有识别到相关 skills，优先检查：

- 是否已经重启 Codex。
- `~/.agents/skills` 下是否存在 Understand Anything 相关 `SKILL.md`。
- 安装脚本是否因网络、权限或 Git 失败中断。
- 当前会话的 skills 列表是否过长，导致初始上下文只展示了部分 skills。

与 CodeGraph 搭配时，推荐分工：

- Understand Anything：先建立全局认知，例如项目图谱、业务域、模块边界和关键文件。
- CodeGraph：再做精确结构查询，例如符号定义、调用方、被调用方、影响面和 staged diff 结构审查。

不要把 Understand Anything 或 CodeGraph 的输出当成最终事实；关键结论仍要回到源码、测试、日志或运行结果验证。

### 自动触发

当用户请求明显匹配 skill 描述时，Codex 会自动使用对应 skill。

示例：

```text
帮我创建一个新的 Codex skill，用于生成 Go 项目的代码审查清单。
```

会触发 `skill-creator`。

```text
帮我查询 OpenAI 最新模型，并给出 API 调用建议。
```

会触发 `openai-docs`。

### 显式触发

也可以直接点名 skill。

示例：

```text
使用 skill-creator 创建一个用于 Spring Boot 排障的 skill。
```

```text
用 openai-docs 查一下 Responses API 的最新用法。
```

项目级 skill 也可以直接点名：

```text
使用 redesign-existing-projects skill，帮我重设计当前项目的已有页面。
```

### 使用已安装 skill

安装并重启 Codex 后，有两种使用方式：

1. 直接点名：

```text
使用 brainstorming skill，帮我为这个需求做方案发散和技术取舍。
```

2. 让任务自然命中触发条件：

```text
帮我围绕这个功能做头脑风暴，列出可选方案、风险和推荐实现路径。
```

推荐在重要任务中直接点名 skill，触发更稳定。

## 判断是否启用

### 1. 查看配置

全局配置文件：

```text
~/.codex/config.toml
```

当前配置中没有发现关闭 skills 的配置项。以下配置不会关闭 skills：

```toml
[features]
memories = false
multi_agent = false
```

说明：

- `memories = false` 只表示关闭跨会话记忆。
- `multi_agent = false` 只表示关闭多模型协作。
- 它们不影响 skills 是否可用。

### 2. 查看本地 skill 文件

```shell
find ~/.codex/skills -maxdepth 4 -type f
```

如果能看到 `.system/{skill-name}/SKILL.md`，说明本地存在对应 skill。

### 3. 查看项目级 skill 文件

```shell
find .agents/skills -maxdepth 4 -name SKILL.md
```

如果当前项目存在 `skills-lock.json`，也可以检查：

```shell
test -f skills-lock.json && echo "has project skills lock"
```

项目级 skill 文件存在，只能说明仓库中有对应 skill；是否被当前会话识别，还要看 Codex 运行时是否加载了项目级 skills。

### 4. 看当前会话上下文

如果当前会话启动时出现 `Available skills` 列表，说明 Codex 已经识别到 skills。

## 常见误判

### 误判一：配置里没有 `skills = true`，所以没启用

不成立。当前配置格式没有要求添加 `skills = true`。没有关闭项时，已安装 skills 会由运行时按任务触发。

### 误判二：没有自动使用 skill，就是 skill 没打开

不一定。更常见原因是任务没有命中特定 skill 的触发条件。

例如普通代码修改、文件编辑、命令执行通常不会触发 `openai-docs`、`imagegen`、`skill-creator` 等专用 skills。

### 误判三：`memories = false` 会关闭 skills

不成立。`memories` 和 `skills` 是不同能力。

### 误判四：项目级 skill 必须写进 `AGENTS.md` 才能被发现

不成立。`AGENTS.md` 不是 skill 注册表。项目级 skill 的发现依赖 Codex 运行时加载 `.agents/skills/` 和 `skills-lock.json` 等约定位置。

但在 `AGENTS.md` 中写明项目级 skill 的使用场景仍然有价值：它可以帮助 Agent 在任务描述不够明确时优先选择正确 skill，减少自动触发歧义。

## 推荐实践

1. 需要特定 skill 时，直接在任务中点名 skill，减少触发歧义。
2. 创建自定义 skill 时，使用 `skill-creator`，不要手写零散结构。
3. 安装第三方 skill 时，使用 `skill-installer`，并确认来源可信。
4. 不建议在 `config.toml` 中添加未确认支持的 feature 开关，避免配置被忽略或产生兼容问题。
5. 项目级 skill 的完整规则放在 `.agents/skills/{skill-name}/SKILL.md`，`AGENTS.md` 只保留短说明和触发场景。

## 排查流程

如果怀疑 skill 没有生效，按以下顺序排查：

1. 确认本地是否存在 skill 文件：

```shell
find ~/.codex/skills -maxdepth 4 -type f
```

2. 如果是项目级 skill，确认项目内是否存在 skill 文件和锁文件：

```shell
find .agents/skills -maxdepth 4 -name SKILL.md
test -f skills-lock.json && echo "has project skills lock"
```

3. 确认 `~/.codex/config.toml` 中没有异常关闭项。
4. 在请求中显式点名 skill，例如 `使用 skill-creator ...` 或 `使用 redesign-existing-projects skill ...`。
5. 如果仍未触发，检查该 skill 的 `SKILL.md` 触发条件是否匹配当前任务，必要时把 description 写得更具体。
