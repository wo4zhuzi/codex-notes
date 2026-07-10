# Go 双分支深度审查 Skill 实现计划

> **For agentic workers:** REQUIRED SUB-SKILL: 使用 `skill-creator` 初始化并验证 Skill。当前环境未提供 `superpowers:executing-plans`，因此在当前会话按本计划内联执行，每个步骤完成后立即验证。

**目标：** 在 `my-skills/` 中创建可复用的 `my-go-change-review` 模板，根据两个 Git ref 深度审查 Go 改动，优先判断逻辑自洽性以及事务和数据库锁风险。

**架构：** `SKILL.md` 只编排只读审查流程；确定性的 Git 基线收集放入 shell 脚本；逻辑自洽、事务锁和 finding 分级分别放入按需加载的 reference。脚本只读取指定 commit，测试通过临时 Git fixture 覆盖 merge-base、ref 异常和工作树隔离。

**技术栈：** Markdown、YAML、Bash、Git、`rg`、Go CLI、Codex `skill-creator`。

---

## 文件映射

- 创建 `my-skills/my-go-change-review/SKILL.md`：触发说明、输入契约、只读边界、审查编排和报告确认门。
- 创建 `my-skills/my-go-change-review/agents/openai.yaml`：UI 名称、短描述和默认提示词。
- 创建 `my-skills/my-go-change-review/scripts/collect-review-context.sh`：校验两个 ref，输出固定 SHA、merge-base、提交、文件和 Go module 清单。
- 创建 `my-skills/my-go-change-review/tests/test-collect-review-context.sh`：在临时仓库验证脚本行为和只读性。
- 创建 `my-skills/my-go-change-review/references/logic-consistency.md`：逻辑自洽专项检查。
- 创建 `my-skills/my-go-change-review/references/transaction-locking.md`：事务与数据库锁专项检查。
- 创建 `my-skills/my-go-change-review/references/severity-guide.md`：finding 分级、证据状态和输出模板。
- 创建 `my-skills/install-skill.sh`：把仓库模板安装到用户级 Codex Skills 目录，支持覆盖前备份。
- 创建 `my-skills/tests/test-install-skill.sh`：在临时 `CODEX_HOME` 验证安装和替换流程。
- 修改 `README.md`：在模板目录树中加入新 Skill。
- 修改 `personal-skills.md`：增加新模板定位、触发示例和启用命令。
- 修改 `codex-skills.md`：补充模板清单中的新 Skill。
- 更新 `docs/changes/2026-07-10-go-change-review-skill-design.md`：记录实际实现和验证结果。

## 根因与风险边界自检

- 根因：普通 diff review 缺少业务不变量、历史修复、事务边界和锁风险分析。
- 最小修改：只新增一个独立 Skill 模板及三个现有入口文档的短说明，不重构已有 Skill。
- 验证方式：shell fixture 测试、`bash -n`、`skill-creator` 校验、YAML 解析、占位符扫描、Markdown 空白检查和 Git 范围检查。
- 风险边界：只在全部验证通过后通过已测试安装器写入用户级 Skills 目录；不修改业务代码，不自动 commit，不使用 CodeGraph，不引入仓库运行时依赖。

### Task 1：初始化 Skill 骨架

**Files:**

- Create: `my-skills/my-go-change-review/SKILL.md`
- Create: `my-skills/my-go-change-review/agents/openai.yaml`
- Create: `my-skills/my-go-change-review/scripts/`
- Create: `my-skills/my-go-change-review/references/`
- Create: `my-skills/my-go-change-review/tests/`

- [x] **Step 1：使用官方初始化脚本创建 Skill**

Run:

```bash
python3 /Users/david/.codex/skills/.system/skill-creator/scripts/init_skill.py \
  my-go-change-review \
  --path my-skills \
  --resources scripts,references \
  --interface 'display_name=个人 · Go Change Review' \
  --interface 'short_description=深度审查 Go 双分支改动，重点检查逻辑、事务与数据库锁风险' \
  --interface 'default_prompt=使用 $my-go-change-review 审查 main 与 feature/order-tx 之间的 Go 改动，重点检查逻辑自洽性、事务边界和数据库锁风险。'
```

Expected: 创建 `my-skills/my-go-change-review/`、`SKILL.md`、`agents/openai.yaml`、`scripts/` 和 `references/`，退出码为 0。

- [x] **Step 2：创建测试目录**

Run:

```bash
mkdir -p my-skills/my-go-change-review/tests
```

Expected: `tests/` 存在，仓库其他目录不变。

- [x] **Step 3：检查初始化产物**

Run:

```bash
find my-skills/my-go-change-review -maxdepth 2 -type f -print
```

Expected: 仅显示 `SKILL.md` 和 `agents/openai.yaml`；resource 目录为空且没有示例占位文件。

### Task 2：用 TDD 实现 Git 上下文收集脚本

**Files:**

- Create: `my-skills/my-go-change-review/tests/test-collect-review-context.sh`
- Create: `my-skills/my-go-change-review/scripts/collect-review-context.sh`

- [x] **Step 1：写入失败测试**

Create `my-skills/my-go-change-review/tests/test-collect-review-context.sh`：

```bash
#!/usr/bin/env bash
set -euo pipefail

TEST_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SKILL_DIR="$(cd "${TEST_DIR}/.." && pwd)"
SCRIPT="${SKILL_DIR}/scripts/collect-review-context.sh"

fail() {
  printf 'FAIL: %s\n' "$1" >&2
  exit 1
}

assert_contains() {
  local haystack="$1"
  local needle="$2"
  [[ "${haystack}" == *"${needle}"* ]] || fail "missing output: ${needle}"
}

assert_not_contains() {
  local haystack="$1"
  local needle="$2"
  [[ "${haystack}" != *"${needle}"* ]] || fail "unexpected output: ${needle}"
}

expect_failure() {
  local expected_message="$1"
  shift
  local output
  local status
  set +e
  output="$("$@" 2>&1)"
  status=$?
  set -e
  [[ ${status} -ne 0 ]] || fail "command unexpectedly succeeded: $*"
  assert_contains "${output}" "${expected_message}"
}

[[ -x "${SCRIPT}" ]] || fail "missing executable script: ${SCRIPT}"

TMP_DIR="$(mktemp -d)"
trap 'rm -rf "${TMP_DIR}"' EXIT
REPO="${TMP_DIR}/repo"

git init -q -b main "${REPO}"
git -C "${REPO}" config user.name "Skill Test"
git -C "${REPO}" config user.email "skill-test@example.com"
printf 'module example.com/review\n\ngo 1.22\n' >"${REPO}/go.mod"
printf 'package review\n\nfunc Existing() {}\n' >"${REPO}/old.go"
git -C "${REPO}" add go.mod old.go
git -C "${REPO}" commit -q -m "initial"
INITIAL_SHA="$(git -C "${REPO}" rev-parse HEAD)"

git -C "${REPO}" checkout -q -b feature/order
git -C "${REPO}" mv old.go new.go
printf 'package review\n\nfunc Feature() {}\n' >"${REPO}/feature.go"
mkdir -p "${REPO}/services/billing"
printf 'module example.com/review/services/billing\n\ngo 1.22\n' >"${REPO}/services/billing/go.mod"
git -C "${REPO}" add feature.go new.go services/billing/go.mod
git -C "${REPO}" commit -q -m "add order feature"

git -C "${REPO}" checkout -q main
printf 'package review\n\nfunc BaselineOnly() {}\n' >"${REPO}/baseline.go"
git -C "${REPO}" add baseline.go
git -C "${REPO}" commit -q -m "advance baseline"
printf 'local-only\n' >"${REPO}/dirty.txt"

STATUS_BEFORE="$(git -C "${REPO}" status --porcelain=v1 --untracked-files=all)"
OUTPUT="$(cd "${REPO}" && "${SCRIPT}" main feature/order)"
STATUS_AFTER="$(git -C "${REPO}" status --porcelain=v1 --untracked-files=all)"

assert_contains "${OUTPUT}" "MERGE_BASE_SHA=${INITIAL_SHA}"
assert_contains "${OUTPUT}" "WORKTREE_DIRTY=true"
assert_contains "${OUTPUT}" $'R100\told.go\tnew.go'
assert_contains "${OUTPUT}" $'A\tfeature.go'
assert_contains "${OUTPUT}" $'A\tservices/billing/go.mod'
assert_contains "${OUTPUT}" "services/billing/go.mod"
assert_not_contains "${OUTPUT}" "baseline.go"
[[ "${STATUS_BEFORE}" == "${STATUS_AFTER}" ]] || fail "script changed the worktree"

expect_failure "cannot resolve head ref: missing-ref" \
  bash -c "cd '${REPO}' && '${SCRIPT}' main missing-ref"
expect_failure "base and head resolve to the same commit" \
  bash -c "cd '${REPO}' && '${SCRIPT}' main main"

SHALLOW="${TMP_DIR}/shallow"
git clone -q --branch feature/order --depth 2 "file://${REPO}" "${SHALLOW}"
SHALLOW_OUTPUT="$(cd "${SHALLOW}" && "${SCRIPT}" HEAD~1 HEAD)"
assert_contains "${SHALLOW_OUTPUT}" "SHALLOW=true"

git -C "${REPO}" checkout -q --orphan unrelated
git -C "${REPO}" rm -q -rf .
printf 'unrelated\n' >"${REPO}/unrelated.txt"
git -C "${REPO}" add unrelated.txt
git -C "${REPO}" commit -q -m "unrelated root"
expect_failure "no merge base between main and unrelated" \
  bash -c "cd '${REPO}' && '${SCRIPT}' main unrelated"

printf 'PASS: collect-review-context.sh\n'
```

- [x] **Step 2：运行测试并确认 Red**

Run:

```bash
bash my-skills/my-go-change-review/tests/test-collect-review-context.sh
```

Expected: 非零退出，错误包含 `missing executable script`。

- [x] **Step 3：写入最小实现**

Create `my-skills/my-go-change-review/scripts/collect-review-context.sh`：

```bash
#!/usr/bin/env bash
set -euo pipefail

fail() {
  local status="$1"
  shift
  printf 'error: %s\n' "$*" >&2
  exit "${status}"
}

if [[ $# -ne 2 ]]; then
  fail 2 "usage: collect-review-context.sh <base-ref> <head-ref>"
fi

BASE_REF="$1"
HEAD_REF="$2"

REPOSITORY_ROOT="$(git rev-parse --show-toplevel 2>/dev/null)" || \
  fail 2 "current directory is not inside a Git repository"
cd "${REPOSITORY_ROOT}"

resolve_ref() {
  local label="$1"
  local ref="$2"
  local sha
  if ! sha="$(git rev-parse --verify "${ref}^{commit}" 2>/dev/null)"; then
    fail 3 "cannot resolve ${label} ref: ${ref}"
  fi
  printf '%s\n' "${sha}"
}

BASE_SHA="$(resolve_ref base "${BASE_REF}")"
HEAD_SHA="$(resolve_ref head "${HEAD_REF}")"

if [[ "${BASE_SHA}" == "${HEAD_SHA}" ]]; then
  fail 4 "base and head resolve to the same commit"
fi

if ! MERGE_BASE_SHA="$(git merge-base "${BASE_SHA}" "${HEAD_SHA}" 2>/dev/null)"; then
  fail 5 "no merge base between ${BASE_REF} and ${HEAD_REF}"
fi

SHALLOW="$(git rev-parse --is-shallow-repository)"
WORKTREE_DIRTY=false
if [[ -n "$(git status --porcelain=v1 --untracked-files=all)" ]]; then
  WORKTREE_DIRTY=true
fi

COMMIT_COUNT="$(git rev-list --count "${MERGE_BASE_SHA}..${HEAD_SHA}")"
GO_MODULE_COUNT="$(git ls-tree -r --name-only "${HEAD_SHA}" | awk '$0 == "go.mod" || /\/go\.mod$/ { count++ } END { print count + 0 }')"

printf 'REPOSITORY_ROOT=%s\n' "${REPOSITORY_ROOT}"
printf 'BASE_REF=%s\n' "${BASE_REF}"
printf 'BASE_SHA=%s\n' "${BASE_SHA}"
printf 'HEAD_REF=%s\n' "${HEAD_REF}"
printf 'HEAD_SHA=%s\n' "${HEAD_SHA}"
printf 'MERGE_BASE_SHA=%s\n' "${MERGE_BASE_SHA}"
printf 'SHALLOW=%s\n' "${SHALLOW}"
printf 'WORKTREE_DIRTY=%s\n' "${WORKTREE_DIRTY}"
printf 'COMMIT_COUNT=%s\n' "${COMMIT_COUNT}"
printf 'GO_MODULE_COUNT=%s\n' "${GO_MODULE_COUNT}"

printf '\n[COMMITS]\n'
git log --reverse --format='%H%x09%s' "${MERGE_BASE_SHA}..${HEAD_SHA}"

printf '\n[CHANGED_FILES]\n'
git diff --name-status --find-renames "${MERGE_BASE_SHA}" "${HEAD_SHA}"

printf '\n[GO_MODULES_AT_HEAD]\n'
git ls-tree -r --name-only "${HEAD_SHA}" | awk '$0 == "go.mod" || /\/go\.mod$/ { print }'
```

- [x] **Step 4：赋予执行权限并确认 Green**

Run:

```bash
chmod +x my-skills/my-go-change-review/scripts/collect-review-context.sh
chmod +x my-skills/my-go-change-review/tests/test-collect-review-context.sh
bash my-skills/my-go-change-review/tests/test-collect-review-context.sh
```

Expected: 输出 `PASS: collect-review-context.sh`，退出码为 0。

- [x] **Step 5：执行 shell 语法检查**

Run:

```bash
bash -n my-skills/my-go-change-review/scripts/collect-review-context.sh
bash -n my-skills/my-go-change-review/tests/test-collect-review-context.sh
```

Expected: 两条命令均无输出且退出码为 0。

### Task 3：编写三个专项 Reference

**Files:**

- Create: `my-skills/my-go-change-review/references/logic-consistency.md`
- Create: `my-skills/my-go-change-review/references/transaction-locking.md`
- Create: `my-skills/my-go-change-review/references/severity-guide.md`

- [x] **Step 1：编写逻辑自洽检查表**

Create `my-skills/my-go-change-review/references/logic-consistency.md`：

````markdown
# 逻辑自洽审查

## 目标

判断 head 引入的业务行为能否在所有入口、状态和失败路径下保持一致。始终以源码、测试、历史和命令结果为证据，不根据函数名或注释直接推断行为正确。

## 证据优先级

```text
head/base 源码、测试和运行结果
> 直接调用关系和配置
> Git 历史与提交说明
> 改动意图和代码注释
> 推断
```

把推断单独标为假设。无法确认的业务规则进入开放问题，不作为确定 finding。

## 建立业务不变量

对每条变更路径记录：

| 维度 | 必须回答的问题 |
| --- | --- |
| 输入 | 哪些输入合法，空值、零值和重复请求如何处理？ |
| 前置状态 | 操作允许从哪些状态开始？ |
| 状态变化 | 哪些数据必须一起变化，顺序是否重要？ |
| 成功结果 | 返回值、持久化、副作用和可观测信号是什么？ |
| 失败结果 | 哪些变化必须撤销，错误如何传递？ |
| 重试 | 重复执行是否安全，幂等键或去重条件是什么？ |
| 兼容性 | 调用方是否依赖旧错误、旧默认值或旧时序？ |

如果无法从 `change_intent` 获得不变量，从 base 实现、base 测试、head 新增测试和历史提交中恢复，并在结论中标明来源。

## 检查成功路径

- 从入口沿直接调用关系追到持久化或最终副作用。
- 确认前置校验在所有入口一致执行。
- 确认新旧路径使用相同的状态定义、默认值和计算口径。
- 确认返回成功前所有必须完成的状态变化已经完成。
- 确认新增分支不会绕过鉴权、幂等、状态机、审计、日志或指标。

## 检查失败与部分失败

- 枚举每个提前返回、错误包装、超时、取消和 panic 路径。
- 确认失败后不存在已经提交但未返回、已经返回成功但未提交的矛盾状态。
- 确认循环或批处理的中途失败策略与业务要求一致。
- 确认补偿、重试或回滚不会重复产生副作用。
- 对“记录失败但继续成功”的代码，确认调用方能观察到降级结果。

## 检查状态机和跨入口一致性

- 比较 Handler、Service、Repository、定时任务和消息消费者中的相同规则。
- 检查新增状态是否被所有 switch、查询过滤、序列化和指标维度处理。
- 检查删除或合并状态后是否仍有旧入口、旧数据或异步消息引用。
- 检查并发请求下的状态前置条件是否在写入时再次验证。

## 检查调用方兼容性

- 比较 public 函数、方法、接口、结构体字段和错误值的 base/head 定义。
- 检查直接调用方是否仍按旧返回值、错误类型、nil 语义或执行顺序工作。
- 检查 mock、fixture、生成代码和测试辅助函数是否同步。
- 接口或共享类型变化时，扩大到所有实现方和跨 module 使用方。

## 追踪历史原因

对被删除的校验、特殊分支、错误处理和锁逻辑执行定向历史查询：

```bash
git log -S '<关键表达式>' -- <path>
git log -G '<相关正则>' -- <path>
git blame <base-sha> -- <path>
git log --follow -- <path>
```

阅读引入相关逻辑的提交及相邻测试。只有历史明确说明修复目的时，才能声称 head 恢复了历史问题。

## 控制审查范围

默认读取变更对象的直接调用方、直接依赖、直接实现方、测试、mock、注册入口和配置引用。出现以下情况时继续扩大：

- public interface 或共享类型变化。
- 持久化、序列化或配置契约变化。
- 权限、事务、并发或核心业务规则变化。
- 当前范围不能解释最终状态或副作用。

停止扩展时要能明确说明影响边界；不能仅以“文件未改动”为停止理由。

## 检查测试有效性

- 确认测试断言业务结果和状态，不只断言 mock 被调用。
- 检查正常、边界、失败、重复请求和并发场景。
- 对重构，确认 base 的行为保护测试在 head 仍然成立。
- 对新功能，确认测试覆盖 `change_intent` 中的验收条件。
- 缺少可执行证据时，把结论降级并记录残余风险。
````

- [x] **Step 2：编写事务与锁检查表**

Create `my-skills/my-go-change-review/references/transaction-locking.md`：

````markdown
# 事务与数据库锁审查

## 触发条件

变更文件或直接调用路径出现以下任一信号时，读取本文件并追踪完整事务边界：

- `Begin`、`BeginTx`、`Commit`、`Rollback`、`Transaction`、`RunInTx`。
- `FOR UPDATE`、`FOR SHARE`、`LOCK`、`NOWAIT`、`SKIP LOCKED`。
- Repository 写操作、余额或库存扣减、状态抢占、批量更新或删除。
- 数据库唯一约束、隔离级别、重试或幂等逻辑变化。

不要因为事务入口不在 diff 中而跳过。只要新路径进入既有事务，就从入口追到 commit/rollback 和事务后的副作用。

## 先确认数据库事实

从 `go.mod`、驱动 import、配置、migration、SQL 方言和部署文件识别：

- 使用 `database/sql`、GORM、`sqlx`、Ent、Bun 还是自定义封装。
- 使用 MySQL、PostgreSQL 或其他数据库。
- 显式隔离级别、读写分离、连接池和超时配置。
- 表结构、主键、唯一约束和与查询条件相关的索引。

数据库类型、索引或执行计划不明确时，只能报告条件性风险。不要把“可能扩大锁范围”写成“一定锁表”。

## 绘制事务边界

按执行顺序记录：

| 步骤 | 操作 | 使用的句柄 | 必须原子完成 | 可能持有的锁 |
| --- | --- | --- | --- | --- |
| 1 | 读取或校验 | `db` 或 `tx` | 是/否 | 行锁、间隙锁或无 |
| 2 | 写入 | `db` 或 `tx` | 是/否 | 依据 SQL 和索引判断 |
| 3 | 外部副作用 | HTTP、RPC、MQ | 是/否 | 事务是否仍开启 |

确认事务入口、所有退出点、commit、rollback 和 commit 后动作。事务边界无法闭合时，先补充调用链证据再分级。

## 检查 tx 逃逸

- `database/sql`：事务内读写必须使用 `*sql.Tx`，不能误用外层 `*sql.DB`。
- GORM：`db.Transaction(func(tx *gorm.DB) ...)` 内必须把 `tx` 传到所有相关 Repository，不能继续使用捕获的 `db`。
- `sqlx`：事务内使用 `*sqlx.Tx`，检查 helper 是否把它降回 `*sqlx.DB`。
- Ent：检查 `tx.Client()` 或事务 client 是否贯穿相关写操作。
- Bun：检查 `RunInTx` 回调内是否使用回调提供的 `bun.Tx`。
- 自定义封装：追踪接口的实际动态类型和 context 中保存的事务句柄。

只要一个必须原子完成的操作绕过事务，就检查失败时是否可能产生部分提交，并优先按数据一致性风险分级。

## 检查 commit、rollback 和错误语义

- 检查所有提前返回和 panic 是否触发 rollback。
- 检查 commit 错误是否返回给调用方，不能在 commit 失败后返回成功。
- 检查 rollback 错误是否被记录，同时不覆盖原始业务错误。
- 检查 defer 中的变量捕获是否能区分成功、失败和 panic。
- 检查 context 取消或超时后事务是否可靠结束。
- 检查嵌套事务是新事务、复用外层事务还是 savepoint，并确认行为符合预期。

## 检查锁范围和索引

- 提取 `SELECT ... FOR UPDATE`、UPDATE 和 DELETE 的完整谓词。
- 将谓词字段与主键、唯一索引和普通索引顺序对照。
- 检查缺少 WHERE、范围过宽、隐式类型转换、函数包裹索引列和低选择性条件。
- 检查批量操作是否一次锁定过多记录，是否有稳定分批键和批次上限。
- 检查访问不存在记录时，在当前数据库和隔离级别下是否涉及间隙锁或谓词锁。
- 缺少 `EXPLAIN` 或生产索引证据时，用 `高风险` 或 `待确认` 表达，不下确定结论。

## 检查并发一致性

- 检查“先读后写”是否可能发生丢失更新。
- 检查状态条件是否包含在 UPDATE 谓词中，是否验证 affected rows。
- 检查余额、库存、配额和任务抢占是否使用原子条件或正确锁定。
- 检查唯一约束冲突是否按幂等成功、可重试错误或业务冲突处理。
- 检查锁等待超时和死锁重试是否有次数、退避和完整事务重放。
- 确认重试不会重复发送消息、扣款或调用外部服务。

## 检查锁顺序和死锁

- 列出每条相关业务路径的表和记录加锁顺序。
- 比较批量 ID 是否使用稳定排序。
- 检查路径 A 先锁资源 X 再锁 Y，而路径 B 先锁 Y 再锁 X 的环路。
- 检查索引变化是否改变扫描和加锁顺序。
- 只有形成可解释的等待环或有运行证据时，才标为确定死锁；否则说明并发前提并标为高风险。

## 检查事务时长与外部副作用

- 标出事务内 HTTP、RPC、文件 I/O、消息发送、sleep 和大循环。
- 检查这些操作是否延长连接和锁持有时间。
- 检查事务提交前发送消息导致回滚后外部状态已变化的问题。
- 检查事务提交后发送消息失败导致数据库成功但事件丢失的问题。
- 根据现有架构判断是否已有 outbox、幂等消费或补偿机制；不要无条件要求引入新架构。

## 数据库差异边界

- MySQL：结合存储引擎、隔离级别、访问索引判断 record lock、gap lock 和 next-key lock。
- PostgreSQL：区分 row-level lock、table-level lock mode、谓词锁和 MVCC 可见性。
- 不确认数据库版本和隔离级别时，不引用特定实现作为确定事实。
- 不用代码静态阅读代替 `EXPLAIN`、锁监控或并发复现；缺少运行证据时明确残余风险。
````

- [x] **Step 3：编写 finding 分级规范**

Create `my-skills/my-go-change-review/references/severity-guide.md`：

````markdown
# Finding 分级与输出规范

## 严重程度

- `P0`：可能造成大范围数据损坏、不可逆丢失、严重安全事故或无法控制的生产影响。
- `P1`：明确行为回归、事务逃逸、可形成的死锁、高概率大范围锁等待或生产故障。
- `P2`：特定边界下的逻辑错误、兼容性问题或具有明确触发条件的风险。
- `P3`：非阻断改进建议。不要把纯风格偏好列为 finding。

严重程度按实际影响和触发概率判断，不按改动行数、文件数量或代码复杂度判断。

## 证据状态

- `已确认`：可由 diff、base/head 源码、测试、历史或验证命令直接证明。
- `高风险`：触发路径明确，但缺少数据库执行计划、并发复现或部署环境证据。
- `待确认`：依赖未提供的业务规则或部署配置。放入开放问题，不作为确定缺陷。

## Finding 模板

### [P1][已确认] <问题标题>

- 位置：`path/to/file.go:line`
- 触发条件：<可以复现问题的输入、状态或并发顺序>
- 问题：<违反的业务不变量或技术契约>
- 影响：<用户、数据或系统后果>
- 证据：<base/head 源码、历史或验证命令>
- 建议方向：<只给修复方向，不修改代码>

## 输出顺序

1. 审查结论：不通过、有条件通过或未发现阻断问题。
2. Findings：按 `P0`、`P1`、`P2`、`P3` 排序。
3. 逻辑自洽性：说明已验证不变量、新旧路径差异和假设。
4. 事务与锁专项：说明事务边界、tx 传递、锁范围、锁顺序、隔离和重试。
5. 验证结果：列出实际命令、结果和未执行原因。
6. 开放问题与未覆盖范围。
7. 固定询问：`是否生成审查报告？`

## 输出约束

- 每条 finding 必须有文件和行号；删除代码引用 base SHA 和旧路径。
- 每条 finding 必须说明触发条件和实际影响。
- 相同根因合并为一条 finding，并列出所有受影响位置。
- 没有发现问题时，明确说明未发现阻断问题，同时列出验证缺口和残余风险。
- 不输出大段原始 diff、日志或命令结果，只保留支持结论的证据。
- 不在用户确认前写审查报告文件。
````

- [x] **Step 4：验证 Reference 可被主 Skill 直接引用**

Run:

```bash
rg -n "业务不变量|部分失败|历史修复|事务边界|tx.*逃逸|锁顺序|隔离级别|P0|已确认|建议方向" \
  my-skills/my-go-change-review/references
```

Expected: 三个文件分别命中其负责的术语，没有空文件或占位段落。

### Task 4：实现主 Skill 与 UI 元数据

**Files:**

- Modify: `my-skills/my-go-change-review/SKILL.md`
- Modify: `my-skills/my-go-change-review/agents/openai.yaml`

- [x] **Step 1：替换初始化占位内容**

Replace `my-skills/my-go-change-review/SKILL.md` with：

````markdown
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
````

确认正文少于 500 行，且三个 reference 都从主 Skill 直接链接。

- [x] **Step 2：重新生成 `agents/openai.yaml`**

Run:

```bash
python3 /Users/david/.codex/skills/.system/skill-creator/scripts/generate_openai_yaml.py \
  my-skills/my-go-change-review \
  --name my-go-change-review \
  --interface 'display_name=个人 · Go Change Review' \
  --interface 'short_description=深度审查 Go 双分支改动，重点检查逻辑、事务与数据库锁风险' \
  --interface 'default_prompt=使用 $my-go-change-review 审查 main 与 feature/order-tx 之间的 Go 改动，重点检查逻辑自洽性、事务边界和数据库锁风险。'
```

Expected: `agents/openai.yaml` 只包含 `interface` 下三个已确认字段，所有字符串带引号，默认提示显式包含 `$my-go-change-review`。

完整预期内容：

```yaml
interface:
  display_name: "个人 · Go Change Review"
  short_description: "深度审查 Go 双分支改动，重点检查逻辑、事务与数据库锁风险"
  default_prompt: "使用 $my-go-change-review 审查 main 与 feature/order-tx 之间的 Go 改动，重点检查逻辑自洽性、事务边界和数据库锁风险。"
```

- [x] **Step 3：扫描初始化占位符**

Run:

```bash
rg -n "TODO|TBD|FIXME|Structuring This Skill|Example helper|Reference Documentation" \
  my-skills/my-go-change-review
```

Expected: 无命中，退出码为 1。

### Task 5：用 TDD 实现可迁移安装器

**Files:**

- Create: `my-skills/tests/test-install-skill.sh`
- Create: `my-skills/install-skill.sh`

- [x] **Step 1：写入失败测试**

Create `my-skills/tests/test-install-skill.sh`：

```bash
#!/usr/bin/env bash
set -euo pipefail

TEST_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MY_SKILLS_DIR="$(cd "${TEST_DIR}/.." && pwd)"
INSTALLER="${MY_SKILLS_DIR}/install-skill.sh"

fail() {
  printf 'FAIL: %s\n' "$1" >&2
  exit 1
}

assert_contains() {
  local haystack="$1"
  local needle="$2"
  [[ "${haystack}" == *"${needle}"* ]] || fail "missing output: ${needle}"
}

expect_failure() {
  local expected_message="$1"
  shift
  local output
  local status
  set +e
  output="$("$@" 2>&1)"
  status=$?
  set -e
  [[ ${status} -ne 0 ]] || fail "command unexpectedly succeeded: $*"
  assert_contains "${output}" "${expected_message}"
}

[[ -x "${INSTALLER}" ]] || fail "missing executable installer: ${INSTALLER}"

TMP_DIR="$(mktemp -d)"
trap 'rm -rf "${TMP_DIR}"' EXIT
TEST_CODEX_HOME="${TMP_DIR}/codex-home"
TARGET="${TEST_CODEX_HOME}/skills/my-go-change-review"

OUTPUT="$(CODEX_HOME="${TEST_CODEX_HOME}" "${INSTALLER}" my-go-change-review)"
assert_contains "${OUTPUT}" "installed: my-go-change-review"
[[ -f "${TARGET}/SKILL.md" ]] || fail "installed SKILL.md is missing"
[[ -f "${TARGET}/agents/openai.yaml" ]] || fail "installed openai.yaml is missing"

expect_failure "target already exists" \
  env CODEX_HOME="${TEST_CODEX_HOME}" "${INSTALLER}" my-go-change-review

printf 'old-version\n' >"${TARGET}/installed-marker.txt"
REPLACE_OUTPUT="$(CODEX_HOME="${TEST_CODEX_HOME}" "${INSTALLER}" my-go-change-review --replace)"
assert_contains "${REPLACE_OUTPUT}" "backup:"
[[ ! -e "${TARGET}/installed-marker.txt" ]] || fail "replacement kept stale files"

BACKUP_MARKER="$(find "${TEST_CODEX_HOME}/skills/.backups" -type f -name installed-marker.txt -print -quit)"
[[ -n "${BACKUP_MARKER}" ]] || fail "backup does not contain previous installation"
[[ "$(<"${BACKUP_MARKER}")" == "old-version" ]] || fail "backup content is incorrect"

expect_failure "invalid skill name" \
  env CODEX_HOME="${TEST_CODEX_HOME}" "${INSTALLER}" '../invalid'
expect_failure "skill template not found" \
  env CODEX_HOME="${TEST_CODEX_HOME}" "${INSTALLER}" missing-skill

printf 'PASS: install-skill.sh\n'
```

- [x] **Step 2：运行测试并确认 Red**

Run:

```bash
bash my-skills/tests/test-install-skill.sh
```

Expected: 非零退出，错误包含 `missing executable installer`。

- [x] **Step 3：写入最小安装器**

Create `my-skills/install-skill.sh`：

```bash
#!/usr/bin/env bash
set -euo pipefail

usage() {
  printf 'usage: install-skill.sh <skill-name> [--replace]\n' >&2
}

fail() {
  printf 'error: %s\n' "$*" >&2
  exit 1
}

if [[ $# -lt 1 || $# -gt 2 ]]; then
  usage
  exit 2
fi

SKILL_NAME="$1"
REPLACE=false
if [[ $# -eq 2 ]]; then
  [[ "$2" == "--replace" ]] || fail "unknown option: $2"
  REPLACE=true
fi

[[ "${SKILL_NAME}" =~ ^[a-z0-9][a-z0-9-]*$ ]] || \
  fail "invalid skill name: ${SKILL_NAME}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SOURCE="${SCRIPT_DIR}/${SKILL_NAME}"
[[ -f "${SOURCE}/SKILL.md" ]] || fail "skill template not found: ${SOURCE}"

CODEX_HOME_DIR="${CODEX_HOME:-${HOME}/.codex}"
TARGET_ROOT="${CODEX_HOME_DIR}/skills"
TARGET="${TARGET_ROOT}/${SKILL_NAME}"
STAGED="${TARGET_ROOT}/.${SKILL_NAME}.install.$$"
BACKUP=""

cleanup() {
  if [[ -e "${STAGED}" || -L "${STAGED}" ]]; then
    rm -rf -- "${STAGED}"
  fi
}
trap cleanup EXIT

mkdir -p "${TARGET_ROOT}"

if [[ -e "${TARGET}" || -L "${TARGET}" ]]; then
  [[ "${REPLACE}" == true ]] || fail "target already exists: ${TARGET}; use --replace"
fi

cp -R "${SOURCE}" "${STAGED}"
[[ -f "${STAGED}/SKILL.md" ]] || fail "staged installation is missing SKILL.md"

if [[ -e "${TARGET}" || -L "${TARGET}" ]]; then
  BACKUP_ROOT="${TARGET_ROOT}/.backups"
  mkdir -p "${BACKUP_ROOT}"
  BACKUP="${BACKUP_ROOT}/${SKILL_NAME}-$(date '+%Y%m%d-%H%M%S')-$$"
  mv "${TARGET}" "${BACKUP}"
fi

if ! mv "${STAGED}" "${TARGET}"; then
  if [[ -n "${BACKUP}" && ! -e "${TARGET}" ]]; then
    mv "${BACKUP}" "${TARGET}"
  fi
  fail "failed to install skill: ${SKILL_NAME}"
fi

trap - EXIT
printf 'installed: %s -> %s\n' "${SKILL_NAME}" "${TARGET}"
if [[ -n "${BACKUP}" ]]; then
  printf 'backup: %s\n' "${BACKUP}"
fi
printf 'restart Codex to load the installed skill\n'
```

- [x] **Step 4：赋予执行权限并确认 Green**

Run:

```bash
chmod +x my-skills/install-skill.sh
chmod +x my-skills/tests/test-install-skill.sh
bash my-skills/tests/test-install-skill.sh
bash -n my-skills/install-skill.sh
bash -n my-skills/tests/test-install-skill.sh
```

Expected: 测试输出 `PASS: install-skill.sh`；两条语法检查无输出；全部退出码为 0。

### Task 6：更新仓库入口文档

**Files:**

- Modify: `README.md`
- Modify: `personal-skills.md`
- Modify: `codex-skills.md`

- [x] **Step 1：更新 README 目录树**

将：

```text
├── my-skills/
│   └── my-codebase-intel/
```

改为：

```text
├── my-skills/
│   ├── my-codebase-intel/
│   └── my-go-change-review/
```

- [x] **Step 2：更新个人 Skill 模板说明**

在 `personal-skills.md` 中：

- 把 `my-go-change-review` 加入命名示例和推荐表。
- 把模板目录树更新为同时包含 `my-codebase-intel` 和 `my-go-change-review`。
- 新增一段简短说明：该 Skill 接受两个 Git ref，使用 merge-base 审查逻辑自洽性、事务和数据库锁，不自动修改代码。
- 增加触发示例：

```text
使用 my-go-change-review 审查 main 和 feature/order-tx，本次目标是调整订单扣减事务。
```

- 增加用户级启用命令：

```bash
cp -R my-skills/my-go-change-review ~/.codex/skills/my-go-change-review
```

- [x] **Step 3：更新 Codex Skills 模板入口**

在 `codex-skills.md` 的模板库说明中同时列出：

```text
my-skills/my-codebase-intel/SKILL.md
my-skills/my-go-change-review/SKILL.md
```

- [x] **Step 4：检查文档引用**

Run:

```bash
rg -n "my-go-change-review|Go Change Review|逻辑自洽|事务.*数据库锁" \
  README.md personal-skills.md codex-skills.md
```

Expected: README 目录树、个人 Skill 说明和 Codex Skills 模板入口均有命中。

### Task 7：完整验证、用户级安装与变更记录

**Files:**

- Update: `docs/changes/2026-07-10-go-change-review-skill-design.md`

- [x] **Step 1：运行脚本测试和语法检查**

Run:

```bash
bash my-skills/my-go-change-review/tests/test-collect-review-context.sh
bash my-skills/tests/test-install-skill.sh
bash -n my-skills/my-go-change-review/scripts/collect-review-context.sh
bash -n my-skills/my-go-change-review/tests/test-collect-review-context.sh
bash -n my-skills/install-skill.sh
bash -n my-skills/tests/test-install-skill.sh
```

Expected: 测试分别输出 `PASS: collect-review-context.sh` 和 `PASS: install-skill.sh`，语法检查无输出，全部退出码为 0。

- [x] **Step 2：运行 Skill 官方校验**

当前系统 Python 缺少 `PyYAML`。在获得执行环境安装许可后，只在临时目录创建校验环境，不修改仓库依赖：

```bash
python3 -m venv /tmp/my-go-change-review-validator
/tmp/my-go-change-review-validator/bin/pip install PyYAML
/tmp/my-go-change-review-validator/bin/python \
  /Users/david/.codex/skills/.system/skill-creator/scripts/quick_validate.py \
  my-skills/my-go-change-review
```

Expected: 输出 `Skill is valid!`。若网络或安装权限不可用，必须改用 YAML 解析和 frontmatter 手动校验，并在变更记录中明确官方校验未运行，不能伪造通过结果。

- [x] **Step 3：验证 UI 元数据和默认提示**

Run:

```bash
/tmp/my-go-change-review-validator/bin/python -c \
  'import pathlib, yaml; data=yaml.safe_load(pathlib.Path("my-skills/my-go-change-review/agents/openai.yaml").read_text()); assert "$my-go-change-review" in data["interface"]["default_prompt"]; print("openai.yaml valid")'
```

Expected: 输出 `openai.yaml valid`。

- [x] **Step 4：执行仓库级检查**

Run:

```bash
git diff --check
rg -n "TODO|FIXME|your-api-key|sk-" .
git status --short --untracked-files=all
```

Expected: `git diff --check` 无输出；敏感扫描只命中既有示例和检查命令；Git 状态只包含本次 Spec、Skill、入口文档、`.gitignore` 和变更记录。

- [x] **Step 5：用已测试安装器执行用户级安装**

Run:

```bash
bash my-skills/install-skill.sh my-go-change-review
```

若目标已存在，先向用户说明并使用：

```bash
bash my-skills/install-skill.sh my-go-change-review --replace
```

Expected: 安装到 `${CODEX_HOME:-$HOME/.codex}/skills/my-go-change-review`；替换时输出备份路径。此步骤写入工作区外，执行时使用权限审批。

- [x] **Step 6：更新变更记录**

将实际实现文件、Red/Green 结果、官方 Skill 校验、YAML 校验、文档检查和残余风险写入 `docs/changes/2026-07-10-go-change-review-skill-design.md`，不得保留计划态描述。

- [x] **Step 7：准备提交但不自动提交**

汇报待提交文件、变更摘要、验证结果和建议提交信息：

```text
新增 Go 双分支审查 Skill
```

只有用户明确回复 `确认提交`、`yes` 或 `commit` 后，才执行 `git add` 和 `git commit`；默认不执行 `git push`。
