TEST?=$$(go list ./... |grep -v 'vendor')
TESTTIMEOUT=180m
TF_SCHEMA_PANIC_ON_ERROR=1

# The single source of truth for the golangci-lint version is the 'version:' field in
# scripts/.custom-gcl.yml (it is required to live there for the plugin build); everything
# else, including the CI workflows, derives it from that file.
GOLANGCI_LINT_VERSION := $(shell sed -n 's/^version: *//p' scripts/.custom-gcl.yml)

TYPOS_VERSION := v1.50.1

# Go tools installed by 'make tools'. terrafmt is also installed by the quick-checks CI
# jobs, which sed its version out of this file.
MISSPELL_VERSION := v0.3.4
TFPROVIDERDOCS_VERSION := v0.12.1
TERRAFMT_VERSION := v1.0.1
GOFUMPT_VERSION := v0.12.0

# The single source of truth for the actionlint version is the go install pin
# in .github/workflows/workflow-actionlint.yml.
ACTIONLINT_VERSION := $(shell sed -n 's/.*actionlint\/cmd\/actionlint@//p' .github/workflows/workflow-actionlint.yml)

.EXPORT_ALL_VARIABLES:

default: build

help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-18s\033[0m%s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)

##@ Deprecated (remove at the end of 2026)
fmtcheck: ## renamed to quick-checks
	@echo "NOTE: 'make fmtcheck' has been renamed to 'make quick-checks' to reflect what it actually runs and will be removed in the future."
	@$(MAKE) quick-checks

tflint: ## renamed to tfproviderlint
	@echo "NOTE: 'make tflint' has been renamed to 'make tfproviderlint' to reflect what it actually runs and will be removed in the future."
	@$(MAKE) tfproviderlint

##@ Build & Generate
tools: ## Install the tools required to develop the provider
	@echo "==> installing required tooling..."
	go install github.com/client9/misspell/cmd/misspell@$(MISSPELL_VERSION)
	go install github.com/bflad/tfproviderdocs@$(TFPROVIDERDOCS_VERSION)
	go install github.com/katbyte/terrafmt@$(TERRAFMT_VERSION)
	go install mvdan.cc/gofumpt@$(GOFUMPT_VERSION)
	go install github.com/rhysd/actionlint/cmd/actionlint@$(ACTIONLINT_VERSION)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $$(go env GOPATH || $$GOPATH)/bin $(GOLANGCI_LINT_VERSION)
	@$(MAKE) golangci-with-modules

build: quick-checks generate ## Run the quick checks, generate code, and compile the provider
	go install

debug: ## Build and launch the provider under the delve debugger
	go build -gcflags="all=-N -l" -trimpath -o terraform-provider-azuread
	dlv exec --listen=:51000 --headless=true --api-version=2 --accept-multiclient --continue terraform-provider-azuread -- -debug

generate: ## Regenerate auto-generated code
	go generate ./internal/services/...
	go generate ./internal/provider/

gencheck: generate ## Check that generated code matches what is committed
	@echo "==> Comparing generated code to committed code..."
	@git diff --compact-summary --exit-code -- ./ || \
		(echo; echo "Unexpected difference in generated code. Run 'make generate' to update the generated code and commit."; echo "If you added or modified a resource, ensure 'go generate' directives are up to date."; exit 1)

##@ Formatting & Quick Checks
# All top-level locations containing Go source, excluding vendor.
GOPATHS=main.go internal version

# The fixers here (plus goimports below) should match the checks in scripts/checks/fmt-check.sh
fmt: ## Fix Go formatting (gofmt, gofumpt, whitespace)
	@echo "==> Fixing source code with gofmt..."
	@gofmt -s -w $(GOPATHS)
	@echo "==> Fixing source code with gofumpt..."
	@gofumpt -w $(GOPATHS)
	@echo "==> Fixing source code with whitespace linter..."
	@golangci-lint run ./... --no-config --enable-only=whitespace --fix

# goimports runs via `golangci-lint fmt` as the standalone binary is single-threaded and far slower
goimports: ## Fix Go import ordering/grouping (slower than fmt, so kept separate)
	@echo "==> Fixing imports with goimports and gci..."
	@golangci-lint fmt -E goimports,gci

quick-checks: ## Run the quick CI checks (formatting + provider policies)
	@echo "==> Running the set of quick CI checks (formatting + provider policies)..."
	@sh "$(CURDIR)/scripts/checks/fmt-check.sh"
	@sh "$(CURDIR)/scripts/checks/test-package-check.sh"
	@sh "$(CURDIR)/scripts/checks/terrafmt-acctests.sh"

terrafmt: ## Fix terraform blocks in acceptance tests and docs
	@echo "==> Fixing acceptance test terraform blocks code with terrafmt..."
	@terrafmt fmt -f -p "*_test.go" ./internal
	@echo "==> Fixing documentation terraform blocks code with terrafmt..."
	@terrafmt fmt -p "*.md" ./docs

##@ Linting & Dependencies
# golangci-lint module plugins (tfproviderlint) only exist in a custom-built binary, so lint
# targets use scripts/golangci-with-modules, rebuilt automatically whenever the config (which
# pins the golangci-lint version) changes. The pinned version is installed before building so
# the host binary running 'golangci-lint custom' always matches the pin. The config filename
# and its living in the build cwd are both fixed by golangci-lint, hence the cd into scripts/.
scripts/golangci-with-modules: scripts/.custom-gcl.yml
	@echo "==> Building golangci-lint with plugins (scripts/golangci-with-modules)..."
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $$(go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)
	@cd scripts && $$(go env GOPATH)/bin/golangci-lint custom

golangci-with-modules: ## Build golangci-lint with plugins (automatic when the config or pinned version changes)
	@if [ -x scripts/golangci-with-modules ] && ! ./scripts/golangci-with-modules version 2>/dev/null | grep -qF -- "$(GOLANGCI_LINT_VERSION:v%=%)"; then \
		echo "==> scripts/golangci-with-modules is not $(GOLANGCI_LINT_VERSION), rebuilding..."; \
		rm -f scripts/golangci-with-modules; \
	fi
	@$(MAKE) scripts/golangci-with-modules

lint: golangci-with-modules ## Check source code with the golangci linters
	@echo "==> Checking source code with golangci-lint..."
	@./scripts/golangci-with-modules run -v ./...

lint-fix: golangci-with-modules ## Fix source code with all golangci linters
	@echo "==> Fixing source code with all golangci linters..."
	@./scripts/golangci-with-modules run ./... --fix

# tfproviderlint runs as part of lint; this target runs just its checks
tfproviderlint: golangci-with-modules ## Check terraform schema definitions with only the tfproviderlint checks
	@echo "==> Checking terraform schemas with tfproviderlint (via golangci-lint)..."
	@./scripts/golangci-with-modules run -v --enable-only tfproviderlint ./...

typos: ## Check spelling in code, docs and examples with typos (config in .typos.toml)
	@command -v typos >/dev/null || (echo "typos not installed. Install via: brew install typos-cli (macOS) or see https://github.com/crate-ci/typos/releases/tag/$(TYPOS_VERSION)" && exit 1)
	@echo "==> Checking spelling with typos..."
	@typos || \
		(echo; echo "Spelling errors found. Fix them with 'make typos-fix', or add false positives (Azure names, enum values) to .typos.toml."; exit 1)

typos-fix: ## Fix spelling errors found by typos
	@command -v typos >/dev/null || (echo "typos not installed. Install via: brew install typos-cli (macOS) or see https://github.com/crate-ci/typos/releases/tag/$(TYPOS_VERSION)" && exit 1)
	@typos --write-changes

yamllint: ## Check YAML files with yamllint (config in .yamllint.yml)
	@command -v yamllint >/dev/null || (echo "yamllint not installed. Install via: brew install yamllint (macOS) or pip install yamllint" && exit 1)
	@echo "==> Checking YAML files with yamllint..."
	@yamllint -s .

actionlint: ## Check GitHub workflows with actionlint (incl. shellcheck on run blocks)
	@command -v actionlint >/dev/null || (echo "actionlint not installed. Install via 'make tools' or: go install github.com/rhysd/actionlint/cmd/actionlint@$(ACTIONLINT_VERSION)" && exit 1)
	@echo "==> Checking workflows with actionlint..."
	@actionlint

shellcheck: ## Check shell scripts with shellcheck
	@command -v shellcheck >/dev/null || (echo "shellcheck not installed. Install via: brew install shellcheck (macOS) or apt install shellcheck (Linux)" && exit 1)
	@echo "==> Checking shell scripts with shellcheck..."
	@shellcheck scripts/*.sh scripts/checks/*.sh scripts/automation/*.sh || \
		(echo; echo "ShellCheck found issues in shell scripts."; echo "Review the errors above and fix them. See https://www.shellcheck.net/ for detailed explanations of each rule."; exit 1)

depscheck: ## Check that go.mod/go.sum and vendor/ are in sync
	@echo "==> Checking source code with go mod tidy..."
	@go mod tidy
	@git diff --exit-code -- go.mod go.sum || \
		(echo; echo "Unexpected difference in go.mod/go.sum files. Run 'go mod tidy' command or revert any go.mod/go.sum changes and commit."; exit 1)
	@echo "==> Checking source code with go mod vendor..."
	@go mod vendor
	@git diff --compact-summary --exit-code -- vendor || \
		(echo; echo "Unexpected difference in vendor/ directory. Run 'go mod vendor' command or revert any go.mod/go.sum/vendor changes and commit."; echo "Do not modify files in the vendor/ directory directly."; exit 1)

##@ Testing
test: ## Run the unit tests
	@TEST=$(TEST) ./scripts/checks/test.sh

testacc: ## Run acceptance tests for a package (TEST=./internal/services/<service>)
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout $(TESTTIMEOUT) -ldflags="-X=github.com/hashicorp/terraform-provider-azuread/version.ProviderVersion=acc"

acctests: ## Run acceptance tests for a service (SERVICE=<service>)
	TF_ACC=1 go test -v ./internal/services/$(SERVICE)/ $(TESTARGS) -timeout $(TESTTIMEOUT) -ldflags="-X=github.com/hashicorp/terraform-provider-azuread/version.ProviderVersion=acc"

debugacc: ## Run acceptance tests under the delve debugger (TEST=./internal/services/<service>)
	TF_ACC=1 dlv test $(TEST) --headless --listen=:2345 --api-version=2 -- -test.v $(TESTARGS)

##@ Documentation
# markdown checked by markdownlint: docs guides and index, README, and the .github markdown
# (PR/issue templates etc). The resource/data-source docs are exempt for now (leading \# ignore
# glob, \# escapes the hash from make) - they carry a large backlog of emphasis-style findings.
MARKDOWN_INPUTS='docs/**/*.md' README.md '.github/**/*.md' '\#docs/resources' '\#docs/data-sources'

markdownlint: ## Check repo markdown with markdownlint (config in .markdownlint.yml)
	@command -v markdownlint-cli2 >/dev/null || (echo "markdownlint-cli2 not installed. Install via: brew install markdownlint-cli2 (macOS) or npm install -g markdownlint-cli2" && exit 1)
	@echo "==> Checking markdown with markdownlint..."
	@markdownlint-cli2 $(MARKDOWN_INPUTS)

docs-lint: ## Check the documentation for issues
	@echo "==> Checking documentation spelling..."
	@misspell -error -source=text -i hdinsight docs/
	@echo "==> Checking documentation for errors..."
	@tfproviderdocs check -provider-name=azuread -allowed-guide-subcategories="Authentication,Upgrade Guides" -enable-contents-check -require-schema-ordering -require-guide-subcategory -require-resource-subcategory
	@sh -c "'$(CURDIR)/scripts/checks/terrafmt-docs.sh'"

validate-examples: ## Check that the terraform examples are valid
	@echo "==> Validating examples..."
	@./scripts/checks/examples-validate.sh

##@ Other
teamcity-test: ## Test the TeamCity configuration
	@$(MAKE) -C .teamcity tools
	@$(MAKE) -C .teamcity test

todo: ## List all TODOs in the codebase
	@grep --color=always --exclude=GNUmakefile --exclude-dir=.git --exclude-dir=vendor --line-number --recursive TODO "$(CURDIR)"

pr-check: generate build test lint docs-lint ## Run the same set of checks CI runs against a PR

.PHONY: default help tools build debug fmt goimports quick-checks fmtcheck terrafmt generate lint lint-fix golangci-with-modules actionlint yamllint markdownlint typos typos-fix shellcheck depscheck gencheck tfproviderlint tflint test testacc acctests debugacc docs-lint validate-examples teamcity-test todo pr-check
