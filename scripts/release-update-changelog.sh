#!/bin/bash
# Copyright IBM Corp. 2019, 2026
# SPDX-License-Identifier: MPL-2.0

set -euo pipefail

changeloggy="./.tools/bin/changeloggy"
if [[ ! -x "${changeloggy}" ]]; then
  echo "Error: ${changeloggy} not found; run \`make release-changelog\` instead" >&2
  exit 1
fi

output="$("${changeloggy}" generate)"
echo "${output}"

RELEASE="$(echo "${output}" | sed -E -n 's/^Generated v?([0-9][0-9.]*) .*/\1/p')"
if [[ -z "${RELEASE}" ]]; then
  echo "Error: could not determine release version from changeloggy output" >&2
  exit 3
fi

printf "%s" "${RELEASE}" > version/VERSION
