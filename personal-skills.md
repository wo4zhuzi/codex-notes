# 个人自建 Skills

## 结论

个人自建 skill 是把自己的 AI 协作习惯、外部工具、MCP 查询和第三方 skills 组合成稳定工作流的方式。它不是三方 skill 的安装说明，也不是 plugin 分发包。

推荐统一用 `my-` 前缀区分个人自建 skill：

```text
my-codebase-intel
my-debug-flow
my-review-gate
my-doc-change-log
```

三方 skill 保持原名，例如 `brainstorming`、`writing-plans`、`systematic-debugging`、`understand`。

## 适用场景

适合创建个人自建 skill 的场景：

- 需要把多个第三方 skills 串成固定流程。
- 需要把 MCP 工具、CLI 工具和源码验证组合成稳定步骤。
- 需要复用个人偏好的审查、排障、计划或交付检查方式。
- 同类任务已经重复出现，单靠 prompt 容易漏步骤。

不适合创建个人自建 skill 的场景：

- 只是一次性任务，用 prompt 即可。
- 只是记录某个三方工具安装方式，应写到对应工具文档。
- 需要把多个 skills、MCP server 和资源打包给团队分发，优先考虑 Codex plugin。

## 命名规范

个人自建 skill 使用：

```text
my-<workflow-name>
```

命名建议：

- 使用小写短横线。
- 名称表达工作流，不只表达单个工具。
- 不使用嵌套目录，例如不要写成 `my/codebase-intel`。
- 不复用三方 skill 名称，避免触发歧义。

推荐示例：

| Skill | 用途 |
| --- | --- |
| `my-codebase-intel` | 组合 Understand Anything、CodeGraph、源码验证和 diff review。 |
| `my-debug-flow` | 组合 systematic-debugging、日志检查、测试复现和修复验证。 |
| `my-review-gate` | 组合 diff impact、代码审查、测试缺口和交付风险检查。 |
| `my-doc-change-log` | 组合文档修改、链接检查、敏感信息扫描和 `docs/changes/` 记录。 |

## 存放位置

本仓库使用 `my-skills/` 作为个人 skill 模板库：

```text
my-skills/
└── my-codebase-intel/
    ├── SKILL.md
    └── agents/
        └── openai.yaml
```

`my-skills/` 只用于保存模板和沉淀个人工作流，不是 Codex 自动发现目录。

真实启用用户级个人 skill 时，放在：

```text
~/.codex/skills/my-codebase-intel/SKILL.md
```

真实启用项目级个人 skill 时，放在：

```text
.agents/skills/my-project-flow/SKILL.md
```

取舍建议：

| 类型 | 适用场景 |
| --- | --- |
| 用户级 `~/.codex/skills/my-*` | 个人所有项目都要复用的工作流。 |
| 项目级 `.agents/skills/my-*` | 只服务当前仓库的流程、约束或领域审查。 |
| 模板库 `my-skills/my-*` | 保存可复制、可链接、可审查的个人 skill 模板。 |

不要把个人 skill 写进 `~/.codex/skills/.system/`。该目录用于系统内置 skills。

## 元数据风格

`SKILL.md` 的 frontmatter 只保留 `name` 和 `description`：

```yaml
---
name: my-codebase-intel
description: 个人工作流：当 Codex 需要理解代码库、规划代码改动、排查 bug、评估重构或审查变更时，组合使用 Understand Anything、CodeGraph、源码验证和 diff review。
---
```

如果生成 `agents/openai.yaml`，显示名建议统一使用：

```yaml
display_name: "个人 · Codebase Intel"
short_description: "用 Understand Anything + CodeGraph 固定代码库分析流程"
default_prompt: "使用 my-codebase-intel：先建立项目全局上下文，再查询目标符号、调用链和影响面，只读取必要源码，最后给出计划、验证方式和风险边界。"
```

这样在 skill 列表中可以快速区分：

```text
个人 · Codebase Intel
个人 · Review Gate
个人 · Debug Flow
```

## 组合型 Skill 设计原则

组合型 skill 不应复制多个三方 skill 的完整内容，而应固定流程边界：

- 什么时候触发。
- 先用哪个工具缩小范围。
- 每一步必须产出什么证据。
- 什么时候回到源码、测试、日志或运行结果验证。
- 哪些动作不能做。

推荐事实优先级：

```text
源码 / 测试 / 日志 / 运行结果
> CodeGraph 精确结构查询
> Understand Anything 全局知识图谱
> LLM 总结
```

不要把 Understand Anything、CodeGraph 或 LLM 总结当成最终事实。关键结论必须回到源码、测试、日志或运行结果验证。

## 示例：`my-codebase-intel`

模板位置：

```text
my-skills/my-codebase-intel/SKILL.md
my-skills/my-codebase-intel/agents/openai.yaml
```

`my-codebase-intel` 固定这条工作流：

```text
Understand Anything 建立全局认知
-> CodeGraph 查询入口、符号、调用链和影响面
-> 少量读取直接相关源码
-> 输出项目看板、项目介绍、模块介绍、分析方案或审查 findings
-> 用测试、命令、日志或手动步骤验证关键结论
```

## 使用示例

创建好 `my-codebase-intel` 后，推荐在任务开头直接点名，减少触发歧义。

日常只需要短指令，具体工具顺序和输出结构写在 `my-skills/my-codebase-intel/SKILL.md` 中。

### 项目看板

```text
使用 my-codebase-intel 帮我准备项目看板
```

### 项目介绍

```text
使用 my-codebase-intel 生成项目介绍
```

### 模块介绍

```text
使用 my-codebase-intel 介绍 demos/rag-notes 模块
```

### 代码分析

```text
使用 my-codebase-intel 分析 <目标函数或模块>
```

### Bug 排查

```text
使用 my-codebase-intel 排查 <现象或错误>
```

### 改动审查

```text
使用 my-codebase-intel 审查当前改动
```

## 创建流程

创建全新个人 skill 时优先使用 `skill-creator`。如果是从本仓库模板启用 `my-codebase-intel`，可以复制模板目录：

```bash
cp -R my-skills/my-codebase-intel ~/.codex/skills/my-codebase-intel
```

创建后执行基础检查：

```bash
test -f ~/.codex/skills/my-codebase-intel/SKILL.md
find ~/.codex/skills -maxdepth 2 -name SKILL.md | sort
```

如果要作为项目级 skill 启用，可以复制到 `.agents/skills/`：

```bash
mkdir -p .agents/skills
cp -R my-skills/my-codebase-intel .agents/skills/my-codebase-intel
```

然后检查：

```bash
find .agents/skills -maxdepth 4 -name SKILL.md
test -f skills-lock.json && echo "has project skills lock"
```

新增或修改仓库内项目级 skill 时，仍需按本仓库规则生成 `docs/changes/YYYY-MM-DD-<topic>.md`。

## 与 Plugin 的边界

个人自建 skill 关注“任务触发后的执行流程”。Codex plugin 关注“多个能力如何安装、启用和分发”。

推荐取舍：

- 单条个人工作流：用 `my-*` skill。
- 当前仓库专属流程：用项目级 `.agents/skills/my-*`。
- 多个 skills、MCP servers 和资源要打包给团队：再做 Codex plugin。
- 只是提醒 AI 遵守某条规则：优先写 prompt 或 `AGENTS.md` 短规则。
