build:
	rm -rf ./bin
	@echo "compiling"
	go build -o bin/scraper ./cmd/main.go
	@echo "program successfuly compiled"

run: build
	@echo "Running SCRAPER program"
	./bin/scraper
