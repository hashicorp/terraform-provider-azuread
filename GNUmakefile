TEST?=$$(go list ./... |grep -v 'vendor')
WEBSITE_REPO=github.com/hashicorp/terraform-website
PKG_NAME=internal
PROVIDER=azuread


.EXPORT_ALL_VARIABLES:
  TF_SCHEMA_PANIC_ON_ERROR=1
  GO111MODULE=on
  GOFLAGS=-mod=vendor

default: build

tools:
	@echo "==> installing required tooling..."
	@sh "$(CURDIR)/scripts/gogetcookie.sh"
	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell
	GO111MODULE=off go get -u github.com/bflad/tfproviderlint/cmd/tfproviderlintx
	GO111MODULE=off go get -u github.com/bflad/tfproviderdocs
	GO111MODULE=off go get -u github.com/katbyte/terrafmt
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$GOPATH/bin v1.27.0

build: fmtcheck
	go install

fmt:
	@echo "==> Fixing source code with gofmt..."
	# This logic should match the search logic in scripts/gofmtcheck.sh
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

# Currently required by tf-deploy compile
fmtcheck:
	@sh "$(CURDIR)/scripts/gofmtcheck.sh"

goimports:
	@echo "==> Fixing imports code with goimports..."
	goimports -local "github.com/terraform-providers/terraform-provider-azuread" -w $(PKG_NAME)/

lint:
	@echo "==> Checking source code against linters..."
	golangci-lint run ./... -v

tflint:
	@echo "==> Checking source code against terraform provider linters..."
	@tfproviderlintx \
        -AT005 -AT006 -AT007 -AT007\
        -R001 -R002 -R003 -R004 -R006 -R007 -R008 -R010 -R012 -R013 -R014\
        -S001 -S002 -S003 -S004 -S005 -S006 -S007 -S008 -S009 -S010 -S011 -S012 -S013 -S014 -S015 -S016 -S017 -S018 -S019 -S020\
        -S021 -S022 -S023 -S024 -S025 -S026 -S027 -S028 -S029 -S030 -S031 -S032 -S033 -S034\
        -V002 -V003 -V004 -V005 -V006 -V007\
        -XR002\
        ./$(PKG_NAME)/...
	@sh -c "'$(CURDIR)/scripts/terrafmt-acctests.sh'"

whitespace:
	@echo "==> Fixing source code with whitespace linter..."
	golangci-lint run ./... --no-config --disable-all --enable=whitespace --fix

depscheck:
	@echo "==> Checking source code with go mod tidy..."
	@go mod tidy
	@git diff --exit-code -- go.mod go.sum || \
		(echo; echo "Unexpected difference in go.mod/go.sum files. Run 'go mod tidy' command or revert any go.mod/go.sum changes and commit."; exit 1)
	@echo "==> Checking source code with go mod vendor..."
	@go mod vendor
	@git diff --compact-summary --exit-code -- vendor || \
		(echo; echo "Unexpected difference in vendor/ directory. Run 'go mod vendor' command or revert any go.mod/go.sum/vendor changes and commit."; exit 1)


test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 180m -ldflags="-X=github.com/terraform-providers/terraform-provider-azuread/version.ProviderVersion=acc"

acctests: fmtcheck
	TF_ACC=1 go test -v ./internal/services/$(SERVICE)/tests/ $(TESTARGS) -timeout $(TESTTIMEOUT) -ldflags="-X=github.com/terraform-providers/terraform-provider-azuread/version.ProviderVersion=acc"

debugacc: fmtcheck
	TF_ACC=1 dlv test $(TEST) --headless --listen=:2345 --api-version=2 -- -test.v $(TESTARGS)

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

website-lint:
	@echo "==> Checking documentation spelling..."
	@misspell -error -source=text -i hdinsight -locale UK website/
	@echo "==> Checking documentation for errors..."
	@tfproviderdocs check -provider-name=azuread -require-resource-subcategory \
		-allowed-resource-subcategories-file website/allowed-subcategories
	@sh -c "'$(CURDIR)/scripts/terrafmt-website.sh'"

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PROVIDER)

.PHONY: build test testacc vet fmt fmtcheck errcheck vendor-status test-compile website website-test
