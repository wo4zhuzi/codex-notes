# Codex Skills

## 结论

当前 Codex Skills 机制已经可用。Skills 不需要在 `~/.codex/config.toml` 中额外添加 `skills = true` 开关；是否触发取决于用户任务是否匹配某个 skill 的触发条件。

## 当前可用 Skills

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
