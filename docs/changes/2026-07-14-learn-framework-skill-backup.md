# 备份 learn-framework Skill

## 任务背景

`learn-framework` 原先仅保存在用户目录 `~/.codex/skills/learn-framework/`。为支持换机迁移，需要将其纳入本仓库的个人 Skill 模板库，并复用现有安装器恢复到新设备。同时需要在仓库的 Skill 使用笔记中记录安装、首次协议生成和后续执行方式，避免备份完成后缺少可发现的使用入口。

## 根因定位

用户目录中的 Skill 不会随本仓库迁移。若旧设备不可用，仅保存在 `~/.codex/skills/` 下的文件无法通过 Git clone 恢复。

仓库已将 `my-skills/` 定义为个人 Skill 模板的维护源，并由 `my-skills/install-skill.sh` 负责安装，因此不需要新增其他备份目录或修改安装流程。

## 执行计划

1. 检查工作区状态和仓库既有个人 Skill 目录约定。
2. 核对本机 `learn-framework` 的完整文件清单和内容。
3. 将源目录完整复制到 `my-skills/learn-framework/`。
4. 比较源目录、仓库模板和临时安装结果，确认内容一致。
5. 扫描本次改动中的占位符、调试标记和疑似密钥。

## 变更内容

- 新增 `my-skills/learn-framework/SKILL.md`。
- 新增 `my-skills/learn-framework/agents/openai.yaml`。
- 新增 `my-skills/learn-framework/references/workflow.md`。
- 新增 `my-skills/learn-framework/references/protocol-template.md`。
- 保持上述文件与本机源目录一致，未调整 Skill 行为和文档内容。
- 更新 `codex-skills.md`，将 `learn-framework` 加入当前可用 Skill 和仓库模板清单。
- 在 `codex-skills.md` 的“使用方式”中补充安装、`protocol` 模式、三个决策门、`execute` 续跑、状态恢复和 L2 能力边界。

## 验证结果

执行源目录与仓库模板的递归比较：

```bash
diff -r ~/.codex/skills/learn-framework my-skills/learn-framework
```

结果：通过，无差异。

使用隔离的临时 `CODEX_HOME` 执行安装：

```bash
CODEX_HOME=<临时目录> bash my-skills/install-skill.sh learn-framework
diff -r my-skills/learn-framework <临时目录>/skills/learn-framework
```

结果：安装成功，安装后目录与仓库模板无差异。

执行安装器回归测试：

```bash
bash my-skills/tests/test-install-skill.sh
```

结果：通过，输出 `PASS: install-skill.sh`。

执行仓库级占位符和疑似密钥扫描，并对本次新增文件使用精确密钥形态正则复核。结果仅命中仓库既有的检查命令和普通示例文本；本次新增文件无命中，未发现真实 API Key、Token、私有代理地址、内部服务 URL 或遗留调试标记。

对 `codex-skills.md` 新增使用说明执行以下检查：

- `git diff --check`：通过。
- 标准调用示例与 `my-skills/learn-framework/SKILL.md` 逐项比对：一致。
- 行尾空白和精确敏感信息模式扫描：无命中。
- Markdown 代码围栏计数：数量为偶数，未发现未闭合围栏。
- `markdownlint`：本机未安装该命令，未执行；以上述本地检查替代。

## 后续建议

换机后在仓库根目录执行：

```bash
bash my-skills/install-skill.sh learn-framework
```

若目标设备已存在同名 Skill，应先确认差异，再按需使用 `--replace`；安装器会在替换前备份旧版本。
