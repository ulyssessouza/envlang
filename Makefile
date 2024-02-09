all: setup
	rm -rf gen
	antlr -Dlanguage=Go -o gen/fileparser -package fileparser EnvLangFile.g4
	antlr -Dlanguage=Go -o gen/valueparser -package valueparser EnvLangValue.g4
	@gotestsum --format testname ./...

setup:
	go install gotest.tools/gotestsum@v1.11.0
