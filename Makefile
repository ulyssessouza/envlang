all:
	rm -rf fileparser
	rm -rf valueparser
	antlr -Dlanguage=Go -o parsers/fileparser -package fileparser EnvLangFile.g4
	antlr -Dlanguage=Go -o parsers/valueparser -package valueparser EnvLangValue.g4
	@go run main.go
