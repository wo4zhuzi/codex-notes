---
name: my-codebase-intel
description: 个人代码库情报工作流。用于项目看板、项目介绍、模块介绍、架构图/流程图生成、代码分析、bug 排查、重构评估和改动审查；组合使用 Understand Anything、CodeGraph、必要源码读取、验证命令、Markdown 摘要和可选 next-ai-draw-io MCP，帮助 Codex 用短指令完成代码库理解、影响面分析与 draw.io 图谱生成。
---

# My Codebase Intel

使用这个 skill 减少大范围读文件，让项目理解、代码分析、架构出图和改动审查基于可验证证据推进。

## 默认原则

1. 先判断任务类型：项目看板、项目介绍、模块介绍、架构图/流程图生成、代码分析、bug 排查、重构评估或改动审查。
2. 优先检查 `.understand-anything/knowledge-graph.json` 是否存在。若存在，用它获取架构层级、关键文件、业务域和 dashboard 上下文；若不存在且全局结构很重要，建议用户先运行 `/understand --language zh`。
3. 优先使用 CodeGraph 查询符号定义、调用方、被调用方、依赖路径、复杂度、入口点、热点和 diff impact。
4. 只读取和结论直接相关的源码、测试、配置或文档。
5. 输出时标注关键依据来自 Understand Anything、CodeGraph、源码、测试、日志还是命令结果。
6. 工具不可用时要明确说明，并降级为源码搜索、文件阅读和手动验证路径。
7. 需要画图时，Understand Anything、CodeGraph、源码和验证命令负责事实来源，next-ai-draw-io MCP 只负责 draw.io 图的创建、编辑、读取或导出。

事实优先级：

```text
源码 / 测试 / 日志 / 运行结果
> CodeGraph 精确结构查询
> Understand Anything 全局知识图谱
> LLM 总结
```

## 任务类型

### 项目看板

触发示例：

```text
使用 my-codebase-intel 帮我准备项目看板
```

执行流程：

1. 检查 `.understand-anything/knowledge-graph.json` 是否存在。
2. 如果图谱存在，按 `understand-dashboard` 的流程启动 dashboard，并给出带 token 的访问地址；如果无法启动，说明原因并继续生成 Markdown 看板。
3. 用 Understand Anything 图谱整理项目层级、关键模块、导览、关键文件和业务域。
4. 用 CodeGraph 查询入口点、热点函数、高复杂度区域、主要依赖和关键调用链。
5. 输出 Markdown 看板，包含：项目概览、模块分区、关键入口、主要依赖、复杂度热点、建议优先阅读文件、后续分析建议。

### 项目介绍

触发示例：

```text
使用 my-codebase-intel 生成项目介绍
```

执行流程：

1. 使用 Understand Anything 获取项目全局结构、层级、tour、关键文件和文档节点。
2. 使用 CodeGraph 验证入口点、核心模块、主要调用关系和复杂度热点。
3. 必要时读取 README、manifest、入口文件和核心模块源码。
4. 输出：项目用途、技术栈、目录结构、核心模块、主要数据流、运行入口、测试入口、适合新手优先阅读的 5 个文件。

### 模块介绍

触发示例：

```text
使用 my-codebase-intel 介绍 <模块路径或模块名>
```

执行流程：

1. 用 CodeGraph 查询模块文件依赖、导出符号、调用方、被调用方和影响范围。
2. 用 Understand Anything 补充该模块在项目层级或业务域中的位置。
3. 读取直接相关源码，确认模块职责和关键路径。
4. 输出：模块职责、主要文件、核心函数或类型、输入输出、上游调用方、下游依赖、常见修改风险和建议验证方式。

### 架构图/流程图生成

触发示例：

```text
使用 my-codebase-intel 把这个项目整理成 draw.io 生成架构图
使用 my-codebase-intel 把这个项目整理成 draw.io 生成流程图
```

执行流程：

1. 检查 `.understand-anything/knowledge-graph.json` 是否存在。若存在，用 Understand Anything 获取项目层级、组件、业务域、关键文件和导览；若不存在且全局结构很重要，建议用户先运行 `/understand --language zh`。
2. 用 CodeGraph 查询入口点、模块依赖、主要调用链、复杂度热点、影响半径或 diff impact，按用户目标选择项目级、模块级或改动级视角。
3. 读取 README、manifest、入口文件、核心模块源码或测试，校验图中关键节点和边是否来自真实代码。
4. 如果用户要求架构图，先输出文字版架构摘要，再输出图谱规格。图谱规格至少包含：节点、边、分组、层级、标签和颜色含义。
5. 如果用户要求流程图，先输出流程摘要，再输出流程规格。流程规格至少包含：开始节点、结束节点、步骤、判断分支、异常路径和跨模块调用。
6. 使用 next-ai-draw-io MCP 出图时，先调用 `start_session` 打开浏览器实时预览会话，确认预览页已建立 MCP 会话（通常 URL 带 `?mcp=`），再调用 `create_new_diagram` 或 `edit_diagram` 创建或编辑 draw.io 图；需要保存时调用 `get_diagram` 或 `export_diagram`。如果出现 `No active session`，先重新调用 `start_session`。
7. MCP 模式下不需要给 next-ai-draw-io 单独配置模型 API Key；LLM 推理由当前 Codex / Claude 会话负责，next-ai-draw-io MCP 只负责图的预览、创建、编辑、读取和导出。
8. 如果 MCP 不可用，输出可复制给 next-ai-draw-io 的自然语言绘图指令，并说明未能直接生成图的原因。

默认图谱分层：

```text
入口 / 用户接口
> 核心模块 / 业务域
> 数据 / 外部服务 / 工具链
```

常用标记：

- 高复杂度节点：标注复杂度或维护风险。
- 强依赖边：标注主要调用、导入或数据流关系。
- 变更影响面：标注受 diff impact 或调用方分析影响的区域。

### 代码分析与修改方案

触发示例：

```text
使用 my-codebase-intel 分析 <目标函数或模块>
```

执行流程：

1. 用 CodeGraph 查询目标定义位置、调用方、被调用方、依赖路径、复杂度和影响半径。
2. 用 Understand Anything 判断目标属于哪个层级、模块或业务域。
3. 读取直接相关源码和测试。
4. 在提出修改前先说明根因或当前判断。
5. 输出最小修改方案、涉及文件、验证命令和剩余风险；用户未要求实现时不修改代码。

### Bug 排查

触发示例：

```text
使用 my-codebase-intel 排查 <现象或错误>
```

执行流程：

1. 先复现或收集现象、失败命令、日志和错误信息。
2. 用 CodeGraph 从相关入口或错误符号追踪调用链和依赖。
3. 读取最小相关源码集合定位根因。
4. 输出根因、证据、最小修复方案、验证命令和风险边界。

### 重构评估

触发示例：

```text
使用 my-codebase-intel 评估重构 <函数、类型或模块>
```

执行流程：

1. 用 CodeGraph 查询调用方、跨模块依赖、复杂度、路径和影响半径。
2. 用 Understand Anything 判断涉及层级和业务域。
3. 输出重构风险、分阶段方案、兼容边界、测试范围和回滚建议。
4. 用户未确认前不执行重构。

### 改动审查

触发示例：

```text
使用 my-codebase-intel 审查当前改动
```

执行流程：

1. 查看当前 git diff 或 staged diff 的变更文件。
2. 如果 Understand Anything 图谱存在，分析变更文件所在层级、组件和相关节点。
3. 使用 CodeGraph diff impact 查询受影响函数、调用方和跨模块依赖。
4. 按严重程度输出 findings，重点看行为回归、遗漏测试、边界条件、文档同步和验证缺口。

## 输出要求

- 结论优先，先说当前判断和最重要风险。
- 引用具体文件、函数、模块或命令结果作为依据。
- 区分“已验证事实”和“基于图谱/调用关系的推断”。
- 需要用户决策时，给出推荐方案和取舍依据。
- 不输出大段工具原始结果，保留高信号摘要。
- 架构图任务先给文字版架构摘要，再给图谱规格；流程图任务先给流程摘要，再给流程规格；使用 next-ai-draw-io MCP 时说明是否已打开实时预览、图名、导出结果或不可用原因。

## 边界

- 不直接读取或修改 `.codegraph/graph.db`。
- 不把 Understand Anything、CodeGraph 或 LLM 总结当作最终事实。
- 不把 next-ai-draw-io 生成的图当作事实来源；图只表达已验证或已标注来源的结构判断。
- 涉及行为判断时，不跳过源码、测试、日志或运行结果验证。
- 声称完成前，不跳过测试或等价手动验证。
- 不为了理解背景而一次性读取大量无关文件。
