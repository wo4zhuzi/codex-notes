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
