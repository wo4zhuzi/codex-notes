---
name: my-go-change-review
description: Go 双分支深度改动审查工作流。用于用户指定原始 Git ref 和开发 Git ref 后，审查新功能、模块重构、Service、Repository 或 SQL 改动是否逻辑自洽，重点排查事务边界、tx 逃逸、数据库锁范围、死锁、隔离级别、幂等性、历史逻辑回归、兼容性和测试缺口；默认只读，审查结束后询问是否生成报告。
---

# My Go Change Review

对两个固定 Git ref 之间的 Go 改动执行只读深度审查。先判断业务逻辑是否自洽，再重点检查事务与数据库锁，最后补充兼容性、并发、资源、安全和测试风险。

## 输入

获取：

- `base_ref`：原始分支、tag 或 commit SHA，必填。
- `head_ref`：开发分支、tag 或 commit SHA，必填。
- `change_intent`：功能目标、验收条件和必须保持不变的行为，可选。

缺少必填 ref 时只追问缺失值。不要自行把 `main`、`master` 或当前分支当作默认值。

调用示例：

```text
使用 my-go-change-review 审查 main 和 feature/order-tx，本次目标是调整订单扣减事务。
```

## 硬性边界

- 只审查并输出 findings，不修改业务代码、测试、配置或文档。
- 不在用户仓库执行 `git checkout`、`git switch`、`git reset` 或 `git clean`。
- 不执行 `git fetch`，不安装项目依赖，不自动访问外部服务。
- 不让未提交改动混入固定 SHA 的 diff 或验证结果。
- 不把命名、格式或个人风格偏好作为缺陷。
- 不根据 CodeGraph、LLM 总结或函数名替代源码和测试证据。
- 不在审查完成前询问是否生成报告。
- 不在用户明确确认前写入报告文件。

允许在系统临时目录创建本地 clone 以验证固定 head；只操作该临时目录并在结束时清理。

## 加载资源

从本 `SKILL.md` 所在目录解析 `<skill-dir>`，不要假设用户当前目录就是 Skill 目录。

1. 始终读取 `references/logic-consistency.md`。
2. 触达事务、SQL、Repository 或数据库锁时读取 `references/transaction-locking.md`。
3. 输出 findings 前读取 `references/severity-guide.md`。
4. 不一次性加载与当前阶段无关的 reference。

## 审查流程

### 1. 固定审查基线

在用户仓库根目录运行：

```bash
bash "<skill-dir>/scripts/collect-review-context.sh" "<base-ref>" "<head-ref>"
```

记录 `BASE_SHA`、`HEAD_SHA` 和 `MERGE_BASE_SHA`。后续命令只使用这些固定 SHA，审查范围固定为 `MERGE_BASE_SHA..HEAD_SHA`。

遇到以下情况时停止或降级：

- ref 无法解析、两个 ref 相同或无 merge-base：停止并说明错误。
- `SHALLOW=true`：继续可见范围审查，但标记历史不完整；不要自动 fetch。
- submodule、Git LFS 指针或生成代码：标记未覆盖范围；不要自动初始化或生成。
- `WORKTREE_DIRTY=true`：仍可读取固定 SHA；需要运行测试时使用临时 head clone。

### 2. 恢复改动意图

优先采用 `change_intent`。没有说明时读取：

```bash
git log --reverse --format='%H%x09%s%n%b' "<merge-base-sha>..<head-sha>"
git diff --stat "<merge-base-sha>" "<head-sha>"
git diff --name-status --find-renames "<merge-base-sha>" "<head-sha>"
```

结合新增测试、注释和相关文档，把改动分类为新功能、重构或混合改动。列出目标、必须保持的行为和假设；无法确认的业务规则放入开放问题。

### 3. 建立变更地图

读取 Go diff 的函数上下文：

```bash
git diff --find-renames --function-context "<merge-base-sha>" "<head-sha>" -- '*.go'
```

识别：

- 新增、删除、重命名的文件。
- 变化的函数、方法、类型、接口和错误值。
- 配置、序列化字段、数据库模型和 migration。
- SQL、事务封装、Repository、状态机和并发路径。
- 测试新增、删除和断言变化。

不要只复述 diff。为每个变更对象记录它改变的行为或契约。

### 4. 定向读取新旧代码和历史

使用固定 SHA 读取文件：

```bash
git show "<merge-base-sha>:<path>"
git show "<head-sha>:<path>"
```

使用固定字符串搜索符号，避免把符号名解释成正则：

```bash
git grep -n -F -e '<symbol>' "<merge-base-sha>" -- '*.go'
git grep -n -F -e '<symbol>' "<head-sha>" -- '*.go'
```

默认读取直接调用方、直接依赖、直接实现方、对应测试、mock、注册入口和配置引用。接口、共享模型、持久化、事务、权限、并发或核心规则变化时继续扩大，直到可以说明最终状态和副作用边界。

对被删除的校验、特殊分支、错误处理和锁逻辑定向查询：

```bash
git log -S '<关键表达式>' -- <path>
git log -G '<相关正则>' -- <path>
git blame "<merge-base-sha>" -- <path>
git log --follow -- <path>
```

读取相关历史提交和测试后再判断是否恢复了历史问题。

### 5. 审查逻辑自洽性

按 `references/logic-consistency.md` 提取业务不变量，检查成功路径、错误路径、部分失败、状态转换、幂等、调用方兼容和测试有效性。

必须回答：

- 输入和前置状态是否在所有入口一致。
- 状态变化和副作用是否在成功、失败和重试下闭合。
- 新路径是否遗漏旧校验或绕过已有规则。
- 调用方是否仍按旧错误、返回值和时序工作。
- 历史修复保护的行为是否仍然存在。

### 6. 审查事务与数据库锁

出现 `Begin`、`Transaction`、`Commit`、`Rollback`、`FOR UPDATE`、Repository 写操作、状态抢占、余额、库存、配额、批量更新或删除时，读取 `references/transaction-locking.md`。

追踪整个事务入口、每个数据库句柄、所有退出点和 commit 后副作用。重点回答：

- 所有必须原子完成的操作是否使用同一个 tx。
- 是否存在误用外层 db 的 tx 逃逸。
- SQL 谓词和索引是否支持预期锁粒度。
- 多条路径的加锁顺序是否可能形成等待环。
- 隔离级别、affected rows 和重试是否保护业务不变量。
- 外部调用是否延长事务，提交与消息副作用是否一致。

数据库类型、索引、隔离级别或执行计划不明确时，只报告条件性风险，不声称一定锁表或一定死锁。

### 7. 执行分层验证

先确认当前工作树是否可代表固定 head：

```bash
git rev-parse HEAD
git status --porcelain=v1 --untracked-files=all
```

仅当当前 HEAD 等于 `HEAD_SHA` 且工作树干净时在当前目录验证。否则用 `mktemp -d` 创建临时目录，从本地仓库执行 `git clone --no-checkout --local`，只在临时 clone 中 checkout detached `HEAD_SHA`；不要切换用户仓库。

从受影响文件所在目录向上定位最近的 `go.mod`。按以下顺序验证：

1. 受影响 package。
2. 直接依赖 package。
3. public interface、共享模型或核心组件变化时验证对应 module。

阻止 Go 命令隐式联网：

```bash
GOTOOLCHAIN=local GOPROXY=off go test <affected-package>
GOTOOLCHAIN=local GOPROXY=off go vet <affected-package>
```

依赖未缓存、需要外部服务、测试耗时明显或可能改变外部状态时停止该项并记录原因。head 出现疑似回归时，可在临时 merge-base clone 复跑同一命令：head 失败而 base 通过才作为强回归证据；两边都失败时记录为既有失败。

### 8. 输出审查结果

按 `references/severity-guide.md` 输出：

1. 审查结论。
2. 按严重程度排序的 findings。
3. 逻辑自洽性结论。
4. 事务与锁专项结论。
5. 实际验证命令和结果。
6. 假设、开放问题和未覆盖范围。

每条确定 finding 都必须包含文件和行号、触发条件、问题、影响、证据和建议方向。没有发现问题时明确说明，同时列出验证缺口和残余风险。

## 报告确认门

完成全部审查输出后，最后单独询问：

```text
是否生成审查报告？
```

用户未确认时停止，不写文件。用户确认后：

1. 先读取目标仓库关于文档位置和文件改动的规则。
2. 将 ref 中的 `/`、空格和文件名非法字符转换为短横线。
3. 默认写入 `docs/reviews/YYYY-MM-DD-<base>-to-<head>.md`。
4. 记录原始 ref、完整 SHA、merge-base、改动意图、findings、两个专项结论、验证结果和未覆盖范围。
5. 只生成报告，不借机修改业务代码。
