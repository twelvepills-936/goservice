proto.gen:
	docker run --rm \
		-v $(PWD):/workspace \
		-w /workspace \
		bufbuild/buf:1.44.0 generate

proto.deps.update:
	docker run --rm \
		-v $(PWD):/workspace \
		-w /workspace \
		bufbuild/buf:1.44.0 dep update

lint:
	docker run --rm -v `pwd`:/app -w /app golangci/golangci-lint:v1.55.2 golangci-lint run --timeout 5m0s -v