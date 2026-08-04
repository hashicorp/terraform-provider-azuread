TEST?=$$(go list ./... |grep -v 'vendor')
TESTTIMEOUT=180m
TF_SCHEMA_PANIC_ON_ERROR=1

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
	go install github.com/client9/misspell/cmd/misspell@latest
	go install github.com/bflad/tfproviderlint/cmd/tfproviderlintx@latest
	go install github.com/bflad/tfproviderdocs@latest
	go install github.com/katbyte/terrafmt@latest
	go install mvdan.cc/gofumpt@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $$(go env GOPATH || $$GOPATH)/bin v2.12.2

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
goimports: ## Fix Go import ordering (slower than fmt, so kept separate)
	@echo "==> Fixing imports with goimports..."
	@golangci-lint fmt -E goimports

quick-checks: ## Run the quick CI checks (formatting + provider policies)
	@echo "==> Running the set of quick CI checks (formatting + provider policies)..."
	@sh "$(CURDIR)/scripts/checks/fmt-check.sh"
	@sh "$(CURDIR)/scripts/checks/test-package-check.sh"

terrafmt: ## Fix terraform blocks in acceptance tests and docs
	@echo "==> Fixing acceptance test terraform blocks code with terrafmt..."
	@terrafmt fmt -f -p "*_test.go" ./internal
	@echo "==> Fixing documentation terraform blocks code with terrafmt..."
	@terrafmt fmt -p "*.md" ./docs

##@ Linting & Dependencies
lint: ## Check source code with the golangci linters
	@echo "==> Checking source code with golangci-lint..."
	@golangci-lint run -v ./...

lint-fix: ## Fix source code with all golangci linters
	@echo "==> Fixing source code with all golangci linters..."
	@golangci-lint run ./... --fix

tfproviderlint: ## Check terraform schema definitions with tfproviderlint
	@echo "==> Checking terraform schemas with tfproviderlint..."
	@tfproviderlintx \
        -AT005 -AT006 -AT007 -AT007\
        -R001 -R002 -R003 -R004 -R006 -R007 -R008 -R010 -R012 -R013 -R014\
        -S001 -S002 -S003 -S004 -S005 -S006 -S007 -S008 -S009 -S010 -S011 -S012 -S013 -S014 -S015 -S016 -S017 -S018 -S019 -S020\
        -S021 -S022 -S023 -S024 -S025 -S026 -S027 -S028 -S029 -S030 -S031 -S032 -S033 -S034\
        -V002 -V003 -V004 -V005 -V006 -V007\
        -XR002\
        ./internal/...
	@sh -c "'$(CURDIR)/scripts/checks/terrafmt-acctests.sh'"

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

pr-check: generate build test lint tfproviderlint docs-lint ## Run the same set of checks CI runs against a PR

.PHONY: default help tools build debug fmt goimports quick-checks fmtcheck terrafmt generate lint lint-fix shellcheck depscheck gencheck tfproviderlint tflint test testacc acctests debugacc docs-lint validate-examples teamcity-test todo pr-check
