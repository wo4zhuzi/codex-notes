# codex-notes

Codex 相关使用笔记与配置模板。

## 内容说明

本仓库主要用于沉淀 AI 编程工具的本地使用经验，包括：

- LLM 与 Agent 原理、多轮对话、工具调用和上下文边界。
- Codex Agent 运行机制、上下文发现、工具探索和任务状态恢复。
- planning-with-files 任务生命周期、当前任务工作区和归档规则。
- AI Agent 完整项目实战复盘、成功率优化和 token 成本控制。
- Codex CLI 安装、登录和基础配置。
- Agent 上下文管理、上下文生命周期和 Harness Engineering 分工。
- Agent 评估与验证、阶段性评估和交付评估。
- Codex 与 Claude Code 的协作方式、规则文件融合和跨工具接续。
- Codex plugin 的定位、安装方式，以及与 skill 的区别。
- Codex Skills 的触发、安装与 Understand Anything 的 Codex 安装方式。
- 个人自建 Skills 的命名、模板库、组合型工作流和 `my-` 前缀规范。
- codex-plugin-cc 在 Claude Code 中调用 Codex 的插件化协作流程。
- Codex Subagent 注册、调用和多 agent 审查流程。
- Function Calling 原理和订单查询 demo。
- RAG 原理和本地 Markdown 知识库问答 demo。
- Codex MCP 使用、原理和服务端设计思路。
- CodeGraph、n8n、Next AI Draw.io 等外部工具的安装、配置和 MCP 接入流程。
- CC-Switch 安装与服务商切换说明。
- 面向不同开发场景的 CC-Switch 配置模板。

## 目录结构

```text
.
├── README.md
├── ai-workflow.md
├── ai-project-checklist.md
├── ai-agent-project-practice.md
├── llm-agent-principles.md
├── codex-agent-runtime.md
├── context-management.md
├── planning-with-files-workflow.md
├── agent-evaluation.md
├── codex-claude-collaboration.md
├── codex-plugins.md
├── codex-plugin-cc.md
├── git-workflow.md
├── spec-workflow.md
├── tdd-workflow.md
├── codex-cli.md
├── codex-core-commands.md
├── codex-skills.md
├── personal-skills.md
├── my-skills/
│   └── my-codebase-intel/
├── subagent.md
├── function-calling.md
├── rag.md
├── mcp.md
├── cc-swtich.md
├── demos/
│   ├── function-calling-orders/
│   └── rag-notes/
├── external-tools/
│   ├── README.md
│   ├── codegraph.md
│   ├── n8n-ai-refactor-workflow.md
│   └── next-ai-draw-io.md
├── docs/
│   └── changes/
│       └── YYYY-MM-DD-<topic>.md
└── cc-switch-configs/
    ├── dev-main.toml
    ├── dev-debug.toml
    ├── dev-cheap.toml
    ├── dev-review.toml
    ├── dev-arch.toml
    ├── dev-subagent-review.toml
    ├── dev-subagent-bugfix.toml
    ├── dev-subagent-project.toml
    └── subagents/
        ├── pr-explorer.toml
        ├── test-impact-reviewer.toml
        ├── risk-reviewer.toml
        ├── security-reviewer.toml
        ├── compat-reviewer.toml
        └── docs-researcher.toml
```

## 快速入口

- [LLM 与 Agent 原理地图](./llm-agent-principles.md)
- [Codex Agent 运行机制](./codex-agent-runtime.md)
- [AI 编程工作流](./ai-workflow.md)
- [AI 项目落地清单](./ai-project-checklist.md)
- [AI Agent 完整项目实战复盘](./ai-agent-project-practice.md)
- [Agent 上下文管理](./context-management.md)
- [planning-with-files 任务生命周期工作流](./planning-with-files-workflow.md)
- [Agent 评估与验证](./agent-evaluation.md)
- [Codex 与 Claude Code 协作](./codex-claude-collaboration.md)
- [Codex Plugins 使用笔记](./codex-plugins.md)
- [codex-plugin-cc 在 Claude Code 中使用](./codex-plugin-cc.md)
- [Git 融入 AI 工作流](./git-workflow.md)
- [Spec 模式实践](./spec-workflow.md)
- [TDD 实践手册](./tdd-workflow.md)
- [Codex CLI 使用笔记](./codex-cli.md)
- [Codex 核心指令](./codex-core-commands.md)
- [Codex Skills 使用笔记](./codex-skills.md)
- [个人自建 Skills](./personal-skills.md)
- [Codex Subagent 使用笔记](./subagent.md)
- [Function Calling 使用笔记](./function-calling.md)
- [RAG 使用笔记](./rag.md)
- [Codex MCP 使用笔记](./mcp.md)
- [外部工具使用笔记](./external-tools)
- [AI 重构项目中 n8n 可以做什么](./external-tools/n8n-ai-refactor-workflow.md)
- [CC-Switch 使用笔记](./cc-swtich.md)
- [CC-Switch 配置模板](./cc-switch-configs)
- [会话变更记录](./docs/changes)

## 配置模板说明

`cc-switch-configs/` 下的配置文件按使用场景拆分：

- `dev-main.toml`：日常开发默认配置。
- `dev-debug.toml`：问题排查、复杂 bug 定位配置。
- `dev-cheap.toml`：低成本、轻量任务配置。
- `dev-review.toml`：代码审查、质量检查配置。
- `dev-arch.toml`：架构设计、复杂方案评审配置。
- `dev-subagent-review.toml`：PR / 分支审查和发布前质量把关配置。
- `dev-subagent-bugfix.toml`：复杂 bug 根因定位和修复配置。
- `dev-subagent-project.toml`：多步骤项目执行和并行审查配置。
- `subagents/*.toml`：只读审查角色模板，例如 `pr-explorer`、`test-impact-reviewer`、`risk-reviewer`。

使用时可根据当前任务切换对应配置，避免在速度、成本和推理强度之间频繁手动调整。

配置中的 `review_model`、`model_reasoning_effort` 等值是当前仓库模板示例，更新时间为 2026-05-27。实际使用前应按当前 Codex / CC-Switch 版本、供应商支持模型和成本策略确认。

## 使用建议

1. 先阅读 [LLM 与 Agent 原理地图](./llm-agent-principles.md)，理解 LLM、多轮对话、Agent 循环和工具调用的基本机制。
2. 再阅读 [Codex Agent 运行机制](./codex-agent-runtime.md)，理解 Codex 如何加载规则、发现上下文、调用工具和恢复任务状态。
3. 再阅读 [AI 编程工作流](./ai-workflow.md)，理解 vibe、plan、spec 和 harness engineering 的分工。
4. 继续阅读 [Agent 上下文管理](./context-management.md)，理解上下文选择、压缩、恢复和沉淀方式。
5. 长任务或跨会话任务可阅读 [planning-with-files 任务生命周期工作流](./planning-with-files-workflow.md)，明确当前任务、归档和变更记录的边界。
6. 阅读 [Agent 评估与验证](./agent-evaluation.md)，理解阶段性评估、交付评估和验收证据。
7. 如果同一仓库同时使用 Codex 和 Claude Code，阅读 [Codex 与 Claude Code 协作](./codex-claude-collaboration.md)，统一规则文件和交接流程。
8. 如果需要安装 Codex plugin 或理解 plugin 与 skill 的分工，阅读 [Codex Plugins 使用笔记](./codex-plugins.md)。
9. 如果希望在 Claude Code 中调用 Codex 做审查或救援，阅读 [codex-plugin-cc 在 Claude Code 中使用](./codex-plugin-cc.md)。
10. 新项目先参考 [AI 项目落地清单](./ai-project-checklist.md)，确认导航、验证、权限边界和变更记录。
11. 需要复盘或规划完整项目实战时，阅读 [AI Agent 完整项目实战复盘](./ai-agent-project-practice.md)，理解 Spec、垂直切片、验收清单和 token 控制方式。
12. 阅读 [Git 融入 AI 工作流](./git-workflow.md)，明确分支、worktree、commit 和 push 的授权边界。
13. 新项目或高风险任务先参考 [Spec 模式实践](./spec-workflow.md)，确认边界、方案和进度。
14. Spec 确认后，核心业务逻辑参考 [TDD 实践手册](./tdd-workflow.md)，先用测试锁定行为再实现。
15. 再阅读 [Codex CLI 使用笔记](./codex-cli.md)，完成 Codex CLI 安装与登录。
16. 继续阅读 [Codex Skills 使用笔记](./codex-skills.md)，了解 skills 的触发方式、安装方式、排查流程和 Understand Anything 的 Codex 接入方式。
17. 如果要把多个 skills、MCP 或外部工具固定成个人工作流，阅读 [个人自建 Skills](./personal-skills.md)。
18. 如需使用多 agent 审查或专项取证，阅读 [Codex Subagent 使用笔记](./subagent.md)。
19. 如需理解知识库问答和检索增强生成，阅读 [RAG 使用笔记](./rag.md)，并运行 `demos/rag-notes/`。
20. 如需接入外部工具或文档服务，阅读 [Codex MCP 使用笔记](./mcp.md)。
21. 如需让 AI 使用 CodeGraph、Next AI Draw.io 等外部工具，阅读 [外部工具使用笔记](./external-tools)。
22. 如需在 AI 重构项目中接入 n8n 编排 GitHub、CI、通知、审批和 MCP 工具网关，阅读 [AI 重构项目中 n8n 可以做什么](./external-tools/n8n-ai-refactor-workflow.md)。
23. 阅读 [CC-Switch 使用笔记](./cc-swtich.md)，安装 CC-Switch 并了解配置切换方式。
24. 根据实际任务选择 `cc-switch-configs/` 中的配置模板。
25. 修改配置前建议保留原始配置备份，方便回滚。

执行会产生仓库改动的任务时，建议遵循 [Codex 核心指令](./codex-core-commands.md) 中的流程：先规划、再让 AI 检查计划、执行后在 `docs/changes/` 生成日期开头的变更文档，最后检查 diff 并审查风险。

## 适用场景

本仓库更偏向工程开发场景，适合：

- 后端服务开发。
- API / RPC / 中间件开发。
- 系统问题排查与调试。
- 代码审查与重构评估。
- 架构设计和技术方案评审。

## 注意事项

- 配置文件包含个人使用偏好，直接复用前建议按自己的模型、供应商和成本策略调整。
- `cc-swtich.md` 文件名当前保持原样，避免影响已有链接或引用。
- 第三方工具版本和安装方式可能变化，实际使用时以对应项目官方文档为准。
