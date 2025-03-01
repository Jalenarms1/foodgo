build_dir=./bin/foodgo

run: build
	$(build_dir)

build:
	go build -o $(build_dir) ./cmd/foodgo