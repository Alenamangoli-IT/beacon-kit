#!/usr/bin/make -f
## DevTools:
tidy: ## run go mod tidy in all modules
	@echo "Running go mod tidy in all modules"
	@find . -name 'go.mod' -execdir go mod tidy \;

sync: 
	@echo "Running go mod download && go work sync"
	@go mod download
	@go work sync

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
