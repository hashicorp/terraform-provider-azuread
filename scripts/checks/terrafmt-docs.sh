#!/usr/bin/env bash
# Copyright IBM Corp. 2019, 2026
# SPDX-License-Identifier: MPL-2.0


DIR="$(cd "$(dirname "$0")"/../.. && pwd)"

echo "==> Checking documentation terraform blocks are formatted..."

if ! terrafmt diff -c -q -p "*.md" "${DIR}/docs"; then
  echo "------------------------------------------------"
  echo ""
  echo "The preceding files contain terraform blocks that are not correctly formatted or contain errors."
  echo "You can fix this by running 'make terrafmt', or the pinned .tools/bin/terrafmt on them directly:"
  echo ""
  echo "format only documentation config blocks:"
  echo "$ terrafmt fmt -p \"*.md\" ./docs"
  echo ""
  echo "format a single documentation file:"
  echo "$ terrafmt fmt ./docs/resources/name.md"
  echo ""
  exit 1
fi

exit 0
