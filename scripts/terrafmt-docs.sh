#!/usr/bin/env bash

DIR="$(cd "$(dirname "$0")"/.. && pwd)"

echo "==> Checking documentation terraform blocks are formatted..."

files=$(find "${DIR}"/docs -type f -name "*.md")
error=false

for f in $files; do
  terrafmt diff -c -q "$f" || error=true
done

if ${error}; then
  echo "------------------------------------------------"
  echo ""
  echo "The preceding files contain terraform blocks that are not correctly formatted or contain errors."
  echo "You can fix this by running make tools and then terrafmt on them."
  echo ""
  echo "format a single file:"
  echo "$ terrafmt fmt ./docs/resources/name.md"
  echo ""
  echo "format all documentation files:"
  echo "$ find docs | egrep '\.md$' | sort | while read f; do terrafmt fmt \$f; done"
  echo ""
  echo "on windows:"
  echo "$ Get-ChildItem -Path docs -Recurse -Filter \"*.md\" | foreach {terrafmt fmt $_.fullName}"
  echo ""
  exit 1
fi

exit 0
