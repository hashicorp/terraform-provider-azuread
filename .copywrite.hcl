# Copyright IBM Corp. 2019, 2026
# SPDX-License-Identifier: MPL-2.0

schema_version = 1

project {
  license        = "MPL-2.0"
  copyright_year = 2019

  header_ignore = [
    ".ci/**",
    ".github/**",
    ".teamcity/**",
    ".tools/bin/**", # gitignored tool installs (make tools)
    ".tools/npm/**",
    ".tools/venv/**",
    ".release/**",
    "vendor/**",
    "internal/**/**_gen.go", # Pandora generated files
    ".goreleaser.yml",
  ]
}
