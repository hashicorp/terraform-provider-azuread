#!/bin/bash

DIR="$(cd "$(dirname "$0")"/.. && pwd)"
cd "${DIR}"

TRUNK="main"

usage() {
  echo "Usage: $0 -y [-c] [-f]" >&2
  echo >&2
  echo " -y  Proceed with release. Must be specified." >&2
  echo " -c  Only prepare the changelog; do not commit, tag or push" >&2
  echo " -f  Force release prep when \`${TRUNK}\` branch is not checked out" >&2
  echo >&2
}

while getopts ':ycfh' opt; do
  case "$opt" in
    y)
      GOTIME=1
      ;;
    c)
      NOTAG=1
      ;;
    f)
      FORCE=1
      ;;
    h)
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
  echo "Using BSD sed"
  SED="sed -E"
else
  echo "Using GNU sed"
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

set -e

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

# Replace [GH-nnnn] references with issue links
$SED -i.bak "s/\[GH-([0-9]+)\]/\(\[#\1\]\(${PROVIDER_URL}\/\1\)\)/g" CHANGELOG.md

# Set the date for the latest release
$SED -i.bak "s/^(## v?[0-9.]+) \(Unreleased\)/\1 (${DATE})/i" CHANGELOG.md

rm CHANGELOG.md.bak

if [[ "${NOTAG}" == 1 ]]; then
  echo "Skipping commit, tag and push."
  exit 0
fi

echo "Releasing v${RELEASE}..."

# Commit the changelog
git commit CHANGELOG.md -m "v${RELEASE}"

# Push
git push origin "${BRANCH}"

# Tag
git tag v"${RELEASE}"

# Push the tag
git push origin "v${RELEASE}"
