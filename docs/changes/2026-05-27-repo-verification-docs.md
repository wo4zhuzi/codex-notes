# 2026-05-27 仓库验证说明修正

## 任务背景

本次会话要求先使用只读 subagent 定位问题根因，再由主 agent 汇总根因判断、最小修复方案和验证计划后开始修复。

## 根因定位

三类只读审查结论一致显示：当前 `main` 分支工作树干净，未发现可审查的 PR diff 或具体失败日志，因此不能定位某个行为 bug 的根因。

已确认的文档问题是：仓库说明仍按“纯文档、无自动化测试”描述，但实际已有 `demos/function-calling-orders` 和 `demos/rag-notes` 两个 Go demo，并包含 `*_test.go`。同时，`README.md` 中的 `cc-switch-configs/` 目录结构和配置模板说明落后于实际文件。

## 执行计划

1. 更新 `AGENTS.md`，说明仓库包含文档、配置模板和 Go demo，并补充 demo 测试命令。
2. 更新 `README.md`，同步 `cc-switch-configs/` 的场景模板和 `subagents/` 角色模板。
3. 为 `demos/function-calling-orders/README.md` 补充 `go test ./...` 验证入口。
4. 运行仓库建议的本地检查和两个 Go demo 的测试。

## 变更内容

- `AGENTS.md`：修正项目结构、构建测试说明和测试指南，避免遗漏 `demos/` 下的 Go 测试。
- `README.md`：补齐 subagent 场景模板和只读审查角色模板索引，并标注配置值更新时间。
- `demos/function-calling-orders/README.md`：新增验证小节。

## 验证结果

- `git status --short`：确认本次仅修改 `AGENTS.md`、`README.md`、`demos/function-calling-orders/README.md`，并新增本变更记录。
- `cd demos/function-calling-orders && go test ./...`：沙箱内因无法写入 `~/Library/Caches/go-build` 失败；提升权限重跑后通过，输出 `ok function-calling-orders (cached)`。
- `cd demos/rag-notes && go test ./...`：沙箱内因无法写入 Go build cache trim 文件失败；提升权限重跑后通过，输出 `ok rag-notes (cached)`。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例、既有示例分支名 `project-task-branch` 和 `risk-reviewer` 名称，未发现真实密钥或遗留调试标记。
- `command -v markdownlint`：本机未安装，未运行可选 Markdown lint。

## 后续建议

- 若后续需要定位具体行为问题，应补充失败命令、错误日志、复现步骤或目标分支 diff。
- 涉及模型名称或配置枚举时，继续按当前 Codex / CC-Switch 版本和供应商文档确认。
