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
