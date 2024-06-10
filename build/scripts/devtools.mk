#!/usr/bin/make -f
## DevTools:

bet: build format test-unit ## yo bet
	@git add .
	@git commit -m 'bet'
	@git push

checkpoint: format  ## checkpoint and push to remote
	@git add -A
	@git commit -m "yo bet"
	@git push
	@echo "checkpointed and pushed to remote"

repo-rinse: | ## dangerous!!! make sure you know what you are doing
	git clean -xfd
	git submodule foreach --recursive git clean -xfd
	git submodule foreach --recursive git reset --hard
	git submodule update --init --recursive

sync: 
	@echo "Running go mod download && go work sync"
	@go mod download
	@go work sync

tidy: ## run go mod tidy in all modules
	@echo "Running go mod tidy in all modules"
	@go env -w GOPRIVATE=github.com/berachain
	@dirs=$$(find . -name 'go.mod' -exec dirname {} \;); \
	count=0; \
	total=$$(echo "$$dirs" | wc -l); \
	for dir in $$dirs; do \
		printf "[%d/%d modules complete] Running in %s\n" $$count $$total $$dir && \
		(cd $$dir && go mod tidy) || exit 1; \
		count=$$((count + 1)); \
	done
	@printf "go mod tidy complete for all modules\n"

yap: ## the yap cave
	@go run ./mod/node-core/pkg/utils/yap/yap.go

tidy-sync-check:
	@$(MAKE) repo-rinse tidy sync 
	@if [ -n "$$(git status --porcelain)" ]; then \
		echo "Tidy and sync operations resulted in changes"; \
		git status -s; \
		git diff --exit-code; \
	fi


.PHONY: format build test-unit bet

