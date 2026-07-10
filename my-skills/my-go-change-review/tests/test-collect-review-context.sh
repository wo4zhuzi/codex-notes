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
