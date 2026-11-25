#!/usr/bin/env bash
# Copyright IBM Corp. 2019, 2026
# SPDX-License-Identifier: MPL-2.0


echo "==> Checking acceptance test terraform blocks are formatted..."

if ! terrafmt diff -c -q -f -p "*_test.go" ./internal; then
  echo "------------------------------------------------"
  echo ""
  echo "The preceding files contain terraform blocks that are not correctly formatted or contain errors."
  echo "You can fix this by running 'make terrafmt', or the pinned .tools/bin/terrafmt on them directly:"
  echo ""
  echo "format only acceptance test config blocks:"
  echo "$ terrafmt fmt -f -p \"*_test.go\" ./internal"
  echo ""
  echo "format a single test file:"
  echo "$ terrafmt fmt -f ./internal/services/service/resource_test.go"
  echo ""
  exit 1
fi

exit 0
