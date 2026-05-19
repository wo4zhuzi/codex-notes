# 2026-05-17 TDD 工作流文档

## 任务背景

用户希望在当前知识库中补充一份 TDD 操作文档，并明确它和 Spec 模式的关系，供后续新项目参考。

## 根因定位

当前仓库已有 `spec-workflow.md` 说明如何在实现前确认边界、方案和进度，但缺少与之配套的 TDD 实践手册。`ai-workflow.md` 中虽然说明了 TDD 的位置，但没有提供可直接复制的执行流程和 prompt，后续 agent 在实现核心逻辑时仍需要临时推导 TDD 步骤。

## 执行计划

1. 新增 `tdd-workflow.md`，作为 `spec-workflow.md` 的平级实践手册。
2. 更新 `README.md`，加入 TDD 文档入口和推荐阅读顺序。
3. 更新 `ai-workflow.md`，补充 TDD 文档链接和高风险任务流程。
4. 更新 `AGENTS.md`，只增加短规则，保持“地图而非手册”的定位。
5. 执行仓库文档检查，确认没有占位符、疑似密钥或明显链接问题。

## 变更内容

- 新增 `tdd-workflow.md`：
  - 说明 TDD 的核心目标。
  - 说明 TDD 和 Spec 的区别。
  - 补充 TDD 的触发方式，包括用户显式触发、agent 自动建议、不自动升级场景和 Spec 模式中的触发点。
  - 记录 TDD 适用场景和不适用场景。
  - 提供 `Spec -> 测试清单 -> Red -> Green -> Refactor -> 验证 -> docs/changes` 标准流程。
  - 提供普通功能、bugfix、重构、夜间委托的可复制 prompt。
  - 说明无测试框架项目的替代验证方式。
- 更新 `README.md`：
  - 目录结构加入 `tdd-workflow.md`。
  - 快速入口加入 TDD 实践手册。
  - 使用建议加入 Spec 后参考 TDD 文档推进核心逻辑。
- 更新 `ai-workflow.md`：
  - 在 TDD 章节加入 `tdd-workflow.md` 链接。
  - 补充 Spec 后的两条执行路径：更通用的 Spec + Plan，以及适合核心逻辑和回归保护的 Spec + TDD。
  - 明确默认优先使用 Spec + Plan，只有当任务正确性适合用测试约束时，再升级为 Spec + TDD。
- 更新 `AGENTS.md`：
  - 增加 TDD 短规则。
  - 未把完整 TDD 手册塞入 `AGENTS.md`。

## 验证结果

已执行：

```bash
git status --short --untracked-files=all
rg -n "TODO|FIXME|your-api-key|sk-" .
rg -n "\[.*\]\(([^)]+)\)" README.md ai-workflow.md spec-workflow.md tdd-workflow.md docs/changes/2026-05-17-tdd-workflow-doc.md
```

结果：

- `git status --short --untracked-files=all` 显示本次新增 `tdd-workflow.md` 和 `docs/changes/2026-05-17-tdd-workflow-doc.md`，并修改 `README.md`、`ai-workflow.md`；仓库中仍存在此前未跟踪的 `.gitignore` 和历史变更记录。
- 敏感信息扫描仅命中文档中记录的检查命令示例，未发现真实 API Key、Token 或遗留占位符。
- Markdown 链接扫描确认 `README.md`、`ai-workflow.md`、`spec-workflow.md`、`tdd-workflow.md` 中的相关链接已出现。
- 根据用户反馈修正 `ai-workflow.md` 中高风险任务流程的表达，避免把 TDD 写成替代 Plan 的唯一默认路径。
- 根据用户反馈补充 `tdd-workflow.md` 的“如何触发 TDD”章节，说明显式触发、自动建议和不自动升级边界。

## 后续建议

- 后续在多个真实项目中实践该 TDD 流程后，再评估是否抽象为独立 skill。
- 如果新项目需要自动化测试，应在 Spec 阶段明确测试框架和验证命令。
