#!/usr/bin/env bash
set -euo pipefail

pass() { printf '✅ %s\n' "$1"; }
fail() { printf '❌ %s\n' "$1"; }
warn() { printf '⚠️  %s\n' "$1"; }

FAILED=0

check_cmd() {
  local cmd="$1"
  if command -v "$cmd" >/dev/null 2>&1; then
    pass "$cmd installed"
  else
    fail "$cmd missing"
    FAILED=1
  fi
}

version_ge() {
  # usage: version_ge current required
  [ "$(printf '%s\n%s\n' "$2" "$1" | sort -V | head -n1)" = "$2" ]
}

check_go() {
  if ! command -v go >/dev/null 2>&1; then
    fail "go missing (required >= 1.21)"
    FAILED=1
    return
  fi

  local raw current
  raw="$(go version)"
  current="$(printf '%s' "$raw" | sed -E 's/.*go([0-9]+\.[0-9]+(\.[0-9]+)?).*/\1/')"
  if version_ge "$current" "1.21"; then
    pass "go version $current (>= 1.21)"
  else
    fail "go version $current (< 1.21)"
    FAILED=1
  fi
}

check_rust() {
  if ! command -v rustc >/dev/null 2>&1; then
    fail "rustc missing (required >= 1.75)"
    FAILED=1
    return
  fi

  local raw current
  raw="$(rustc --version)"
  current="$(printf '%s' "$raw" | awk '{print $2}')"
  if version_ge "$current" "1.75"; then
    pass "rustc version $current (>= 1.75)"
  else
    fail "rustc version $current (< 1.75)"
    FAILED=1
  fi
}

check_node() {
  if ! command -v node >/dev/null 2>&1; then
    fail "node missing (required >= 18)"
    FAILED=1
    return
  fi

  local raw current
  raw="$(node --version)"
  current="${raw#v}"
  if version_ge "$current" "18"; then
    pass "node version $current (>= 18)"
  else
    fail "node version $current (< 18)"
    FAILED=1
  fi
}

check_git_identity() {
  local name email
  name="$(git config user.name || true)"
  email="$(git config user.email || true)"

  if [[ -n "$name" ]]; then
    pass "git user.name configured ($name)"
  else
    fail "git user.name not configured"
    FAILED=1
  fi

  if [[ -n "$email" ]]; then
    pass "git user.email configured ($email)"
  else
    fail "git user.email not configured"
    FAILED=1
  fi
}

echo "== Task 0.1.2 Validation Checklist =="
check_go
check_rust
check_node
check_cmd docker
check_cmd bpftool
check_git_identity

echo
if [[ "$FAILED" -eq 0 ]]; then
  pass "All prerequisites satisfied"
  exit 0
else
  warn "One or more prerequisites are missing"
  exit 1
fi
