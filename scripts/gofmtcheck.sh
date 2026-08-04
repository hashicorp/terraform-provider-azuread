#!/usr/bin/env bash
# Copyright (c) HashiCorp, Inc.
# Check gofmt
echo "==> Checking that code complies with gofmt requirements..."
gofmt_files=$(gofmt -l $(find . -type d -name vendor -prune -o -name '*.go' -print))
if [ -n "${gofmt_files}" ]; then
    echo 'gofmt needs running on the following files:'
    echo "${gofmt_files}"
    echo "You can use the command: \`make fmt\` to reformat code."
    exit 1
fi

exit 0
