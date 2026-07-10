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
