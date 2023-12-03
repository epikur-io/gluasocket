.PHONY: check
check:
	go vet ./...
	go test ./...

.PHONY: clean
clean:
	go clean -i ./...
	go clean -testcache

.PHONY: outdated
outdated:
	go install github.com/psampaz/go-mod-outdated@latest
	go list -u -m -json -mod=mod all | go-mod-outdated -direct
