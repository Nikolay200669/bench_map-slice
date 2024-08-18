.PHONY: bench

bench:
	@echo "Running benchmarks..."
	go test -bench=. ./...
