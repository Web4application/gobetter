Install golangci-lint
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6

# Run code generation, build, test and linters
./scripts/presubmit.sh
# Run build, test and linters but skip code generation
./scripts/presubmit.sh  --no-generate

# Or just run the linters alone:
golangci-lint run
module github.com/mobiletoly/gobetter
docker build -f ./integration/Dockerfile -t ctgo-builder .
docker run -it --mount type=bind,src="$(pwd)",target=/src ctgo-builder /bin/bash -c "cd /src; ./scripts/install_deps.sh; go generate -x ./..."
go 1.24.0

toolchain go1.24.4

require golang.org/x/tools v0.38.0

require (
	golang.org/x/mod v0.29.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
)
