all:
	@antlr -Dlanguage=Go -o parser EnvLang.g4
	@go run main.go
