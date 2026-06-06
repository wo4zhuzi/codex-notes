# my-codebase-intel 支持 draw.io 架构图和流程图

## 任务背景

用户希望将 `next-ai-draw-io` 的 MCP 出图能力接入现有 `my-codebase-intel` 工作流，让 Understand Anything 和 CodeGraph 负责代码库理解与图谱事实来源，再由 `next-ai-draw-io` 生成 draw.io 架构图或流程图。

## 根因定位

当前 `my-codebase-intel` 已覆盖项目看板、项目介绍、模块介绍、代码分析、bug 排查、重构评估和改动审查，但没有明确的“架构图/流程图生成”任务类型，也没有规定 `next-ai-draw-io` 与 Understand Anything、CodeGraph 的职责边界。初版触发示例还把内部执行要求写进了用户提示词，表达偏长。

## 执行计划

1. 在 `my-skills/my-codebase-intel/SKILL.md` 增加架构图/流程图生成任务类型。
2. 明确事实来源与出图工具分工：Understand Anything、CodeGraph、源码和验证命令负责事实，`next-ai-draw-io` MCP 只负责 draw.io 图创建、编辑、读取或导出。
3. 补充架构图输出要求，包括文字版架构摘要、节点、边、分组、层级、标签和颜色含义。
4. 补充流程图输出要求，包括流程摘要、开始节点、结束节点、步骤、判断分支、异常路径和跨模块调用。
5. 更新 `agents/openai.yaml`，让默认提示词使用短表达。

## 变更内容

- 更新 `my-skills/my-codebase-intel/SKILL.md`：
  - 扩展 skill 描述和默认原则，加入架构图/流程图生成与可选 `next-ai-draw-io` MCP。
  - 新增“架构图/流程图生成”任务类型，包含短触发示例、执行流程、默认图谱分层和常用标记。
  - 在输出要求和边界中强调图谱事实来源与 draw.io 渲染结果的区别。
- 更新 `my-skills/my-codebase-intel/agents/openai.yaml`：
  - 调整 `short_description` 和 `default_prompt`，突出代码库分析与 draw.io 架构图/流程图生成。

## 验证结果

- 已检查 `SKILL.md` frontmatter 保留 `name` 和 `description`。
- 已检查 `agents/openai.yaml` 可被 Ruby YAML 解析。
- 已执行敏感信息扫描，未发现真实密钥形态；`sk-` 仅出现在检查说明中。

## 后续建议

- 如果后续本地安装了 `next-ai-draw-io` MCP，可补一篇外部工具笔记，记录 MCP 配置、启动方式和常用绘图提示词。
