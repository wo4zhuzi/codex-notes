# Codex Plugins

更新时间：2026-06-06。

## 结论

Codex plugin 是 Codex 的可安装分发单元，用于把 skills、apps、MCP servers 和展示资产打包成可安装、可启用、可分享的能力包。

推荐理解方式：

```text
skill = 任务触发后的工作流程说明
plugin = 可安装和分发的能力包
```

如果只是为当前仓库或本机沉淀一个流程，优先使用 skill。如果要把多个 skills、外部应用连接或 MCP 配置打包给团队复用，再封装为 plugin。

## Plugin 解决什么问题

Codex plugin 适合解决“如何分发一组能力”的问题，而不是替代 Agent 或替代单个 skill。

一个 Codex plugin 可以包含：

- `skills`：任务型工作流说明。
- `apps`：GitHub、Slack、Google Drive、Gmail 等外部应用连接。
- `mcpServers`：给 Codex 暴露外部工具或共享信息。
- 展示资产和 marketplace 元数据：用于 Codex app 或 CLI 插件目录展示。

典型场景：

- 把多个相关 skill 组成团队工具包。
- 把 skill 和 MCP server 配置一起分发。
- 把外部应用连接作为 Codex 可安装能力管理。
- 用 marketplace 给个人、项目或团队维护插件目录。

## Plugin 和 Skill 的区别

| 维度 | Codex Plugin | Codex Skill |
| --- | --- | --- |
| 核心定位 | 可安装、可分发的能力包 | 任务触发后的工作流程说明 |
| 常见内容 | skills、apps、MCP servers、资产和 marketplace 元数据 | `SKILL.md`、`references/`、`scripts/`、`assets/` |
| 使用方式 | 在 Codex app 或 CLI 插件目录安装后使用 | Codex 根据任务描述隐式触发，或用户显式调用 |
| 适合场景 | 团队工具包、多个 skill 组合、外部应用/MCP 集成 | 调试流程、写计划、验收检查、特定领域审查 |
| 生命周期 | 关注安装、启用、禁用、分发和权限 | 关注触发准确性和执行步骤质量 |

推荐取舍：

- 单一流程、单仓库规则：用 skill。
- 多个流程打包、需要团队安装：用 plugin。
- 需要外部系统数据或操作能力：优先考虑 plugin + app / MCP。
- 只是在当前任务中要求 Agent 遵循某个规则：直接写 prompt 或 `AGENTS.md`。

## 安装与管理

### 在 Codex CLI 中打开插件目录

进入 Codex CLI 后运行：

```text
/plugins
```

在插件目录中可以按 marketplace 浏览、安装、卸载插件，并启用或停用已安装插件。

### 添加 marketplace

如果要让 Codex 跟踪一个 marketplace source，使用 Codex CLI 命令：

```bash
codex plugin marketplace add owner/repo
codex plugin marketplace add owner/repo --ref main
codex plugin marketplace add https://github.com/example/plugins.git --sparse .agents/plugins
codex plugin marketplace add ./local-marketplace-root
```

查看、刷新或移除 marketplace：

```bash
codex plugin marketplace list
codex plugin marketplace upgrade
codex plugin marketplace upgrade marketplace-name
codex plugin marketplace remove marketplace-name
```

这些是 Codex 命令。不要和 Claude Code 的 `/plugin marketplace add ...` 混用。

## 创建 Codex Plugin

创建 Codex plugin 时，优先使用系统内置的 `plugin-creator` skill。最小插件结构包含：

```text
my-plugin/
└── .codex-plugin/
    └── plugin.json
```

一个最小 `plugin.json` 示例：

```json
{
  "name": "my-plugin",
  "version": "1.0.0",
  "description": "Reusable workflow for Codex",
  "skills": "./skills/"
}
```

如果只是本地试验，不需要一开始就做复杂 marketplace。先把 skill 本身打磨稳定，再决定是否封装成 plugin。

## 使用边界

- Codex plugin 和 Claude Code plugin 是不同生态的安装单元，命令不能混用。
- 第三方工具如果通过 `~/.agents/skills`、`~/.codex/skills` 或项目 `.agents/skills/` 安装到 Codex，应优先记录在 [Codex Skills 使用笔记](./codex-skills.md)。
- 不要把 plugin 当成唯一扩展方式；prompt、`AGENTS.md`、skill、MCP server 和 hooks 都有各自适用范围。
- 安装第三方 marketplace 或 plugin 前，应先检查来源、权限、hooks、MCP server 和外部应用连接风险。

## 参考来源

- OpenAI Codex Manual：`Plugins`、`Build plugins`、`Agent Skills`
- [Codex Skills 使用笔记](./codex-skills.md)
- [codex-plugin-cc 在 Claude Code 中使用](./codex-plugin-cc.md)
