# codex-notes

Codex / Claude Code 相关使用笔记与配置模板。

## 内容说明

本仓库主要用于沉淀 AI 编程工具的本地使用经验，包括：

- Codex CLI 安装、登录和基础配置。
- CC-Switch 安装与服务商切换说明。
- 面向不同开发场景的 CC-Switch 配置模板。

## 目录结构

```text
.
├── README.md
├── codex-cli.md
├── codex-core-commands.md
├── codex-skills.md
├── cc-swtich.md
├── docs/
│   └── changes/
│       └── YYYY-MM-DD-<topic>.md
└── cc-switch-configs/
    ├── dev-main.toml
    ├── dev-debug.toml
    ├── dev-cheap.toml
    ├── dev-review.toml
    └── dev-arch.toml
```

## 快速入口

- [Codex CLI 使用笔记](./codex-cli.md)
- [Codex 核心指令](./codex-core-commands.md)
- [Codex Skills 使用笔记](./codex-skills.md)
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

使用时可根据当前任务切换对应配置，避免在速度、成本和推理强度之间频繁手动调整。

## 使用建议

1. 先阅读 [Codex CLI 使用笔记](./codex-cli.md)，完成 Codex CLI 安装与登录。
2. 再阅读 [Codex Skills 使用笔记](./codex-skills.md)，了解 skills 的触发方式和排查流程。
3. 继续阅读 [CC-Switch 使用笔记](./cc-swtich.md)，安装 CC-Switch 并了解配置切换方式。
4. 根据实际任务选择 `cc-switch-configs/` 中的配置模板。
5. 修改配置前建议保留原始配置备份，方便回滚。

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
