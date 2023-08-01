# Go
GO = go

GOFMT = gofmt
GOFMT_FILES = $(shell find . -name '*.go' | grep -v vendor)

GOMOD   = $(GO) mod
GOTEST  = $(GO) test
GOCLEAN = $(GO) clean

# OS
RM    = rm -f
ECHO  = /bin/echo

# first rule is default rule
default: test

# run tests
test: fmtcheck deps
	@$(ECHO) "==> $@"
	$(GOTEST)

# format go code
fmt:
	@$(ECHO) "==> $@"
	$(GOFMT) -w $(GOFMT_FILES)

# check if code has been formatted
fmtcheck:
	@$(ECHO) "==> $@"
	@files=$$($(GOFMT) -l $(GOFMT_FILES)) ; \
	if [[ -n $${files} ]]; then \
		$(ECHO) "error: file(s) not gofmt formatted:" ; \
		$(ECHO) "$${files}" ; \
		$(ECHO) "run: 'make fmt' to reformat file(s)" ; \
		exit 1 ; \
	fi

# tidy go modules
tidy: deps
	@$(ECHO) "==> $@"
	$(GOMOD) tidy

# verify modules
verify:
	@$(ECHO) "==> $@"
	$(GOMOD) verify

# download modules to pre-fill the cache 
deps:
	@$(ECHO) "==> $@"
	$(GOMOD) download

# clean
clean:
	@$(ECHO) "==> $@"
	$(GOCLEAN) -i

# remove build cache
cleaner: clean
	@$(ECHO) "==> $@"
	$(GOCLEAN) -cache

# remove module cache
cleanest: cleaner test-clean
	@$(ECHO) "==> $@"
	$(GOCLEAN) --modcache

# remove test cache
test-clean:
	@$(ECHO) "==> $@"
	$(GOCLEAN) -testcache

# store local copies of external dependenciew
vendor: deps
	@$(ECHO) "==> $@"
	$(GOMOD) vendor

.PHONY: test fmt fmtcheck tidy verify deps clean cleaner cleanest test-clean vendor