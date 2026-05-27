# 2026-05-27 修正 CC-Switch 配置 profile 加载错误

## 任务背景

用户反馈在 CC-Switch 配置中加入 `profile = "dev-subagent-review"` 后，启动时报错：

```text
Error loading configuration: config profile dev-subagent-review not found
```

## 根因定位

`profile` 在 Codex 配置中不是当前模板的场景声明字段，而是用于选择已有配置 profile 的引用。Codex CLI 的 `--profile <CONFIG_PROFILE>` 会从 `config.toml` 中查找已定义的 `[profiles.<name>]`。

当前 `cc-switch-configs/*.toml` 是用于覆盖整份配置的模板，并没有定义 `[profiles.dev-subagent-review]`。因此在顶层添加 `profile = "dev-subagent-review"` 后，加载器会尝试查找不存在的 profile，最终报错。

## 执行计划

1. 移除 `cc-switch-configs/dev-*.toml` 中错误的顶层 `profile` 字段。
2. 在 `cc-swtich.md` 的多配置说明中补充注意事项，说明 `profile` 与整份配置模板的区别。
3. 验证所有 TOML 文件仍可解析，并确认不再残留顶层 `profile`。

## 变更内容

- 移除以下配置模板中的顶层 `profile = "..."`：
  - `cc-switch-configs/dev-main.toml`
  - `cc-switch-configs/dev-cheap.toml`
  - `cc-switch-configs/dev-arch.toml`
  - `cc-switch-configs/dev-review.toml`
  - `cc-switch-configs/dev-debug.toml`
  - `cc-switch-configs/dev-subagent-review.toml`
  - `cc-switch-configs/dev-subagent-bugfix.toml`
  - `cc-switch-configs/dev-subagent-project.toml`
- 更新 `cc-swtich.md`，明确不要在整份配置模板顶层添加 `profile = "..."`。

## 验证结果

- `python3 -c 'import pathlib,tomllib; [tomllib.load(open(p,"rb")) for p in pathlib.Path("cc-switch-configs").glob("**/*.toml")]; print("ok")'`：输出 `ok`，确认 `cc-switch-configs/` 下所有 TOML 文件语法可解析。
- `rg -n "^profile\s*=" cc-switch-configs cc-swtich.md README.md`：无输出，确认配置模板中不再残留顶层 `profile` 字段。
- `rg -n "config profile|profile-v2|不要在模板顶层" cc-swtich.md docs/changes/2026-05-27-cc-switch-profile-error.md`：确认错误说明和规避方式已写入文档。

## 后续建议

如果后续要使用 Codex 原生 `--profile`，应在主 `config.toml` 中增加 `[profiles.<name>]` 配置块，而不是在独立模板顶层写 `profile`。
