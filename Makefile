build:
	@go build ./braingockcli
	@mv braingockcli/braingockcli braingock

test:
	@go test ./...
