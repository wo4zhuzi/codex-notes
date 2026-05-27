# 2026-05-27 修正推理强度配置值为 xhigh

## 任务背景

用户指出当前文档中有错误的最高推理强度写法，因为支持的推理模式是 `xhigh`。

## 根因定位

`cc-swtich.md`、`cc-switch-configs/` 配置模板和历史变更记录中误用了不存在的推理强度值。该错误会误导使用者复制无效配置，并影响 CC-Switch 场景模板的可用性。

## 执行计划

1. 搜索仓库中错误旧值和 `xhigh` 的出现位置。
2. 将所有错误旧值统一修正为 `xhigh`。
3. 新增本次日期变更文档。
4. 执行文本搜索、工作区状态和敏感占位符检查。

## 变更内容

- 修正 `cc-swtich.md` 中推理强度说明和示例配置。
- 修正 `cc-switch-configs/` 中相关 TOML 模板的 `model_reasoning_effort` 配置值。
- 修正 subagent 场景配置注释中的推理强度描述。
- 修正既有变更记录中的错误配置值描述。

## 验证结果

- 全仓库搜索错误旧值：无命中，确认已移除错误配置值字面量。
- `rg -n "xhigh" cc-swtich.md cc-switch-configs docs/changes`：命中 `cc-swtich.md`、相关 TOML 模板和变更记录，确认修正后的上下文合理。
- TOML 解析检查：`cc-switch-configs/*.toml` 均可被 `tomllib` 成功解析。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例、既有示例分支名 `project-task-branch` 和 `risk-reviewer` 名称，未发现真实密钥或遗留调试标记。
- `git status --short`：仅包含本次计划内的文档和配置模板改动。

## 后续建议

- 后续新增 CC-Switch 配置模板时，推理强度只使用当前工具支持的枚举值，避免从旧文档复制无效字段值。
