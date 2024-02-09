all:
	rm -rf gen
	antlr -Dlanguage=Go -o gen/fileparser -package fileparser EnvLangFile.g4
	antlr -Dlanguage=Go -o gen/valueparser -package valueparser EnvLangValue.g4
	@gotestsum ./...

setup:
	go install gotest.tools/gotestsum@latest
