COVERAGE_FILE ?= coverage.out

TARGET ?= maze
VERSION ?= v1.0.0

MAZE_WIDTH ?= 10
MAZE_HEIGHT ?= 10
END_X = $(shell expr $(MAZE_WIDTH) - 1)
END_Y = $(shell expr $(MAZE_HEIGHT) - 1)

## build: Build the maze binary executable
.PHONY: build
build:
	@echo "Building '${TARGET}' binary..."
	@mkdir -p ./bin
	@go build -ldflags="-X 'main.version=${VERSION}'" -o ./bin/${TARGET} ./cmd/${TARGET}
	@echo "✓ Binary built successfully: './bin/${TARGET}'"

## test: Run all tests with race detector and coverage report
.PHONY: test
test:
	@echo "Running tests with coverage..."
	@go test -coverpkg='gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/...' \
		--race -count=1 -coverprofile='$(COVERAGE_FILE)' ./...
	@echo ""
	@echo "Coverage summary:"
	@go tool cover -func='$(COVERAGE_FILE)' | grep ^total | tr -s '\t'

## help: Display this help message
.PHONY: help
help:
	@echo "Available targets:"
	@echo ""
	@grep -E '^##' $(MAKEFILE_LIST) | sed 's/##//g' | column -t -s ':'

## clean: Remove generated files and binaries
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf ./bin ./output $(COVERAGE_FILE)
	@echo "✓ Cleanup complete"

## generate: Generate a maze and save it to './output/maze.txt'
.PHONY: generate
generate: build
	@echo "Generating ${MAZE_WIDTH}x${MAZE_HEIGHT} maze using DFS algorithm..."
	@mkdir -p ./output
	@./bin/${TARGET} generate \
		--algorithm=dfs \
		--width=${MAZE_WIDTH} \
		--height=${MAZE_HEIGHT} \
		--output=./output/maze.txt
	@echo "✓ Maze generated and saved to './output/maze.txt'"

## solve: Solve a maze from './output/maze.txt'
.PHONY: solve
solve: build
	@echo "Solving maze from './output/maze.txt' using A* algorithm..."
	@./bin/${TARGET} solve \
		--algorithm=astar \
		--file=./output/maze.txt \
		--start=0,0 \
		--end=$(END_X),$(END_Y) \
		--output=./output/maze_solution.txt
	@echo "✓ Solution saved to './output/maze_solution.txt'"

## generate-and-solve: Generate a maze and immediately solve it
.PHONY: generate-and-solve
generate-and-solve: build
	@echo "Generating and solving ${MAZE_WIDTH}x${MAZE_HEIGHT} maze..."
	@mkdir -p ./output
	@./bin/${TARGET} generate \
		--algorithm=dfs \
		--width=${MAZE_WIDTH} \
		--height=${MAZE_HEIGHT} \
		--output=./output/maze.txt
	@./bin/${TARGET} solve \
		--algorithm=astar \
		--file=./output/maze.txt \
		--start=0,0 \
		--end=$(END_X),$(END_Y) \
		--output=./output/maze_solution.txt
	@echo "✓ Maze generated and solved. Files in './output/'"

## demo: Run a complete demonstration (generate, display, solve, display solution)
.PHONY: demo
demo: build
	@echo "=== Maze Generation Demo ==="
	@mkdir -p ./output
	@echo "Step 1: Generating maze..."
	@./bin/${TARGET} generate \
		--algorithm=dfs \
		--width=${MAZE_WIDTH} \
		--height=${MAZE_HEIGHT} \
		--output=./output/maze.txt
	@echo ""
	@echo "Generated maze:"
	@cat ./output/maze.txt
	@echo ""
	@echo "Step 2: Solving maze..."
	@./bin/${TARGET} solve \
		--algorithm=astar \
		--file=./output/maze.txt \
		--start=0,0 \
		--end=$(END_X),$(END_Y) \
		--output=./output/maze_solution.txt
	@echo ""
	@echo "Solution:"
	@cat ./output/maze_solution.txt
	@echo ""
	@echo "✓ Demo complete! Files saved in './output/'"

## generate-unicode: Run maze generation in Unicode mode and display result
.PHONY: generate-unicode
generate-unicode: build
	@echo "Generating ${MAZE_WIDTH}x${MAZE_HEIGHT} maze in Unicode mode..."
	@mkdir -p ./output
	@./bin/${TARGET} generate \
		--algorithm=dfs \
		--width=${MAZE_WIDTH} \
		--height=${MAZE_HEIGHT} \
		--output=./output/maze_unicode.txt \
		--unicode
	@echo ""
	@echo "Generated Unicode maze:"
	@cat ./output/maze_unicode.txt
	@echo ""
	@echo "✓ Unicode maze test complete! File saved in './output/maze_unicode.txt'"
