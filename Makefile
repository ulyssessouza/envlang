all: clean gen test
	@echo "All done!"

test:
	@gotestsum --format testname ./...

gen:
	@antlr -Dlanguage=Go -o gen/fileparser -package fileparser EnvLangFile.g4
	@antlr -Dlanguage=Go -o gen/valueparser -package valueparser EnvLangValue.g4

clean:
	@rm -rf gen

setup:
	go install gotest.tools/gotestsum@v1.11.0
