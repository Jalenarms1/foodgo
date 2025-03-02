build_dir=./bin/foodgo

run: build
	$(build_dir)

build:
	cd client && pnpm run build
	go build -o $(build_dir) ./cmd/foodgo