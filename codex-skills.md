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

## 使用方式

### `planning-with-files` 使用规范

`planning-with-files` 适合需要把任务计划、过程笔记和执行状态落到文件中的长任务。它的价值不是替代 `/plan` 或 `docs/changes/`，而是让任务在中断、压缩上下文或跨天恢复时仍有清晰现场。

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

### 3. 看当前会话上下文

如果当前会话启动时出现 `Available skills` 列表，说明 Codex 已经识别到 skills。

## 常见误判

### 误判一：配置里没有 `skills = true`，所以没启用

不成立。当前配置格式没有要求添加 `skills = true`。没有关闭项时，已安装 skills 会由运行时按任务触发。

### 误判二：没有自动使用 skill，就是 skill 没打开

不一定。更常见原因是任务没有命中特定 skill 的触发条件。

例如普通代码修改、文件编辑、命令执行通常不会触发 `openai-docs`、`imagegen`、`skill-creator` 等专用 skills。

### 误判三：`memories = false` 会关闭 skills

不成立。`memories` 和 `skills` 是不同能力。

## 推荐实践

1. 需要特定 skill 时，直接在任务中点名 skill，减少触发歧义。
2. 创建自定义 skill 时，使用 `skill-creator`，不要手写零散结构。
3. 安装第三方 skill 时，使用 `skill-installer`，并确认来源可信。
4. 不建议在 `config.toml` 中添加未确认支持的 feature 开关，避免配置被忽略或产生兼容问题。

## 排查流程

如果怀疑 skill 没有生效，按以下顺序排查：

1. 确认本地是否存在 skill 文件：

```shell
find ~/.codex/skills -maxdepth 4 -type f
```

2. 确认 `~/.codex/config.toml` 中没有异常关闭项。
3. 在请求中显式点名 skill，例如 `使用 skill-creator ...`。
4. 如果仍未触发，检查该 skill 的 `SKILL.md` 触发条件是否匹配当前任务。
