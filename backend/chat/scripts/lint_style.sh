#!/usr/bin/env bash

set -euo pipefail

status=0

while IFS= read -r file; do
  if rg -n '[A-Za-z0-9_]+\{[^{}\n]*[A-Za-z0-9_]+\s*:' "$file" >/dev/null; then
    echo "single-line keyed struct literal: $file"
    rg -n '[A-Za-z0-9_]+\{[^{}\n]*[A-Za-z0-9_]+\s*:' "$file" || true
    status=1
  fi

  if rg -n '\b(var\s+)?[A-Za-z0-9_]+\s*:?=\s*[A-Za-z0-9_.]+\{\s*\}' "$file" >/dev/null; then
    echo "empty struct literal should prefer 'var x X': $file"
    rg -n '\b(var\s+)?[A-Za-z0-9_]+\s*:?=\s*[A-Za-z0-9_.]+\{\s*\}' "$file" || true
    status=1
  fi

  if rg -n '&[A-Za-z0-9_.]+\{\s*\}' "$file" >/dev/null; then
    echo "empty pointer-to-struct should prefer 'new(X)': $file"
    rg -n '&[A-Za-z0-9_.]+\{\s*\}' "$file" || true
    status=1
  fi

  if rg -n '\b[A-Za-z0-9_]+\s*:?=\s*\[\][A-Za-z0-9_.\*]+\{\s*\}' "$file" >/dev/null; then
    echo "empty slice should prefer 'var xs []T': $file"
    rg -n '\b[A-Za-z0-9_]+\s*:?=\s*\[\][A-Za-z0-9_.\*]+\{\s*\}' "$file" || true
    status=1
  fi
done < <(find . -name '*.go' -type f | sort)

exit "$status"
