all: clean gen test lint
	@echo "All done!"

test:
	@gotestsum --format testname -- -failfast -race -covermode=atomic -coverprofile=coverage.out ./...

gen:
	@antlr4 -Dlanguage=Go -o gen/fileparser -package fileparser EnvLangFile.g4
	@antlr4 -Dlanguage=Go -o gen/valueparser -package valueparser EnvLangValue.g4

clean:
	@rm -rf gen

setup:
	go install gotest.tools/gotestsum@v1.11.0

grun:
	antlr4 EnvLangValue.g4
	javac -cp "/usr/share/java/antlr-4.13.1-complete.jar:$CLASSPATH" EnvLangValue*.java
	java -Xmx500M -cp "/usr/share/java/antlr-4.13.1-complete.jar:$CLASSPATH" org.antlr.v4.gui.TestRig EnvLangValue dqstrings -tokens

lint:
	golangci-lint run ./... --timeout 2m
