# 补充 CodeGraph 图表命令

## 任务背景

用户询问 CodeGraph 建立索引后是否可以自己查看索引图表，并要求把 `codegraph plot` 等常用命令记录到当前文档中。

## 根因定位

`external-tools/codegraph.md` 已记录 CodeGraph 的安装、建索引、MCP 接入、查询和影响面命令，但“常用命令”和“推荐日常流程”中缺少图表查看与结构化导出命令。用户需要的是建索引后的人工查看路径，而不是 MCP 查询流程。

## 执行计划

1. 在 `external-tools/codegraph.md` 的“常用命令”中补充 `stats`、`plot` 和 `export` 命令。
2. 新增“查看索引图表”小节，说明小项目、大项目、函数级图和导出场景的推荐命令。
3. 在“每天开始复杂任务前”流程中加入建索引后可选查看图表步骤。
4. 执行 Git 状态和敏感占位符检查，确认只包含本次文档变更。

## 变更内容

- 补充 `codegraph stats`，用于查看当前索引统计。
- 补充 `codegraph plot`、`codegraph plot -o codegraph.html`、`codegraph plot --functions` 等交互式 HTML 图表命令。
- 补充大项目推荐命令：`codegraph plot --cluster directory --seed top-fanin --seed-count 50 -o codegraph.html`。
- 补充 `codegraph export` 的 Mermaid、JSON、GraphML 导出示例。
- 在日常流程中加入建索引后查看图表的可选步骤。

## 验证结果

- 已执行 `git status --short`：仅包含 `external-tools/codegraph.md` 和本变更记录文件。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。
- 已手动检查 Markdown 命令块和章节位置，新增内容位于“常用命令”“查看索引图表”和“推荐日常流程”中。

## 后续建议

如果后续确认 CodeGraph 新版本调整了 `plot` 或 `export` 参数，应以本机 `codegraph plot --help` 和 `codegraph export --help` 输出为准同步更新本文档。
