#!/bin/bash
# Copyright IBM Corp. 2023, 2026
# SPDX-License-Identifier: MPL-2.0


REPO_DIR="$(cd "$(dirname "$0")"/.. && pwd)"
cd "${REPO_DIR}"

TRUNK="main"

# set debug="echo " to print the release commands instead of running them
debug="${debug:-}"

usage() {
  echo "Usage: $0 -y [-C] [-f]" >&2
  echo >&2
  echo " -y  Proceed with release. Must be specified." >&2
  echo " -C  Only prepare the changelog; do not commit, tag or push" >&2
  echo " -t  Override trunk branch (default: ${TRUNK}), useful for patch releases"
  echo " -T  Skip tests before preparing release" >&2
  echo " -f  Force release prep when \`${TRUNK}\` branch is not checked out" >&2
  echo >&2
}

while getopts ':yCt:Tfh' opt; do
  case "$opt" in
    y)
      GOTIME=1
      ;;
    C)
      NOTAG=1
      ;;
    t)
      TRUNK="${OPTARG}"
      ;;
    T)
      NOTEST=1
      ;;
    f)
      FORCE=1
      ;;
    *)
      usage
      exit 1
      ;;
  esac
done

if [[ "${GOTIME}" != "1" ]]; then
  echo "Specify \`-y\` to initiate release!" >&2
  usage
  exit 1
fi

if [[ "$(uname)" == "Darwin" ]]; then
  echo "(Using BSD sed)"
  SED="sed -E"
else
  echo "(Using GNU sed)"
  SED="sed -r"
fi


DATE="$(date '+%B %d, %Y')"
PROVIDER_URL="https:\/\/github.com\/hashicorp\/terraform-provider-azuread\/issues"

BRANCH="$(git rev-parse --abbrev-ref HEAD)"
if [[ "${BRANCH}" != "${TRUNK}" ]]; then
  if [[ "${FORCE}" == "1" ]]; then
    echo "Caution: Proceeding with release prep on branch: ${BRANCH}"
  else
    echo "Release must be prepped on \`${TRUNK}\` branch. Specify \`-f\` to override." >&2
    exit 1
  fi
fi

if [[ "$(git status --short)" != "" ]]; then
  echo "Error: working tree is dirty" >&2
  exit 4
fi

set -e

if [[ "${NOTEST}" == "1" ]]; then
  echo "Warning: Skipping tests"
else
  echo "Running tests..."
  ( set -x; TF_ACC='' scripts/checks/test.sh )
fi

echo "Preparing changelog for release..."

if [[ ! -f CHANGELOG.md ]]; then
  echo "Error: CHANGELOG.md not found."
  exit 2
fi

# Get the next release
RELEASE="$($SED -n 's/^## v?([0-9.]+) \(Unreleased\)/\1/p' CHANGELOG.md)"
if [[ "${RELEASE}" == "" ]]; then
  echo "Error: could not determine next release in CHANGELOG.md" >&2
  exit 3
fi

# Ensure latest changes are checked out
# shellcheck disable=SC2086 # debug is intentionally unquoted for command prefix pattern
( set -x; ${debug}git pull --rebase origin "${TRUNK}" )

# Replace [GH-nnnn] references with issue links
# shellcheck disable=SC2086 # debug is intentionally unquoted for command prefix pattern
( set -x; ${debug}$SED -i.bak "s/\[GH-([0-9]+)\]/\(\[#\1\]\(${PROVIDER_URL}\/\1\)\)/g" CHANGELOG.md )

# Set the date for the latest release
# shellcheck disable=SC2086 # debug is intentionally unquoted for command prefix pattern
( set -x; ${debug}$SED -i.bak "s/^(## v?[0-9.]+) \(Unreleased\)/\1 (${DATE})/i" CHANGELOG.md )

# shellcheck disable=SC2086 # debug is intentionally unquoted for command prefix pattern
${debug}rm CHANGELOG.md.bak

if [[ "${NOTAG}" == "1" ]]; then
  echo "Warning: Skipping commit, tag and push."
  exit 0
fi

echo "Committing changelog..."
(
  set -x
  # shellcheck disable=SC2086 # debug is intentionally unquoted for command prefix pattern
  ${debug}git commit CHANGELOG.md -m v"${RELEASE}"
  # shellcheck disable=SC2086 # debug is intentionally unquoted for command prefix pattern
  ${debug}git push origin "${BRANCH}"
)


echo "Releasing v${RELEASE}..."

(
  set -x
  # shellcheck disable=SC2086 # debug is intentionally unquoted for command prefix pattern
  ${debug}git tag v"${RELEASE}"
  # shellcheck disable=SC2086 # debug is intentionally unquoted for command prefix pattern
  ${debug}git push origin v"${RELEASE}"
)
