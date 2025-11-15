package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/ulyssessouza/envlang"
	"github.com/ulyssessouza/envlang/store"
)

type (
	options struct {
		files        string
		ignore       bool
		overwrite    bool
		template     string
		headerless   bool
		promptPrefix string
	}
)

func main() {
	var opts options

	flag.StringVar(&opts.files, "f", "", "Comma-separated list of env files to load (e.g., \".env.local,.env\")")
	flag.BoolVar(&opts.ignore, "i", false, "Ignore missing files")
	flag.BoolVar(&opts.ignore, "ignore", false, "Ignore missing files")
	flag.BoolVar(&opts.overwrite, "o", false, "Overwrite existing environment variables")
	flag.BoolVar(&opts.overwrite, "overwrite", false, "Overwrite existing environment variables")
	flag.StringVar(&opts.template, "t", "", "Generate template file from specified env file")
	flag.StringVar(&opts.template, "template", "", "Generate template file from specified env file")
	flag.BoolVar(&opts.headerless, "H", false, "Run in headerless mode (suppress initial messages)")
	flag.BoolVar(&opts.headerless, "headerless", false, "Run in headerless mode (suppress initial messages)")
	flag.StringVar(&opts.promptPrefix, "P", "=> ", "Set the prompt prefix for interactive mode")
	flag.StringVar(&opts.promptPrefix, "prompt-prefix", "=> ", "Set the prompt prefix for interactive mode")

	flag.Parse()

	if opts.template != "" {
		if err := generateTemplate(opts.template); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating template: %v\n", err)
			os.Exit(1)
		}
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		runInteractiveMode(opts)
		return
	}

	files := parseFiles(opts.files)
	if err := loadEnvFiles(files, opts.ignore, opts.overwrite); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading env files: %v\n", err)
		os.Exit(1)
	}

	if err := runCommand(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}

func parseFiles(filesFlag string) []string {
	if filesFlag == "" {
		return []string{".env"}
	}

	parts := strings.Split(filesFlag, ",")
	files := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			files = append(files, trimmed)
		}
	}
	return files
}

func loadEnvFiles(files []string, ignoreNotExist, overwrite bool) error {
	d := store.NewDefaultStoreFromEnv(os.Environ())

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			if os.IsNotExist(err) && ignoreNotExist {
				continue
			}
			return fmt.Errorf("failed to open %s: %w", file, err)
		}

		vars := envlang.GetVariablesFromInputStream(d, f)
		f.Close()

		for key, value := range vars {
			if value == nil {
				continue
			}

			if !overwrite {
				if _, exists := os.LookupEnv(key); exists {
					continue
				}
			}

			os.Setenv(key, *value)
		}
	}
	return nil
}

func runCommand(args []string) error {
	cmdPath, err := exec.LookPath(args[0])
	if err != nil {
		return fmt.Errorf("command not found: %s", args[0])
	}

	return syscall.Exec(cmdPath, args, os.Environ())
}

func generateTemplate(sourceFile string) error {
	f, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", sourceFile, err)
	}
	defer f.Close()

	d := store.NewDefaultStore()
	vars := envlang.GetVariablesFromInputStream(d, f)

	var lines []string
	for key := range vars {
		lines = append(lines, fmt.Sprintf("%s=%s", key, key))
	}

	templateFile := sourceFile + ".template"
	content := strings.Join(lines, "\n") + "\n"
	if err = os.WriteFile(templateFile, []byte(content), 0o600); err != nil { //nolint:govet
		return fmt.Errorf("failed to write template file: %w", err)
	}

	fmt.Printf("Template file generated: %s\n", templateFile)
	return nil
}

func runInteractiveMode(opts options) {
	if !opts.headerless {
		fmt.Fprintln(os.Stderr, "envlang interactive mode (Ctrl+D to finish)")
		fmt.Fprintln(os.Stderr, "Enter environment variable definitions or commands (e.g., /print $VAR, /unset $VAR, /vars):")
	}

	reader := bufio.NewReader(os.Stdin)
	d := store.NewDefaultStoreFromEnv(os.Environ()) // Initialize store with current environment

	for {
		fmt.Fprint(os.Stderr, opts.promptPrefix) // Print the prompt prefix
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "/") {
			parts := strings.Fields(line)
			command := parts[0]

			switch command {
			case "/print":
				if len(parts) != 2 { //nolint:mnd
					fmt.Fprintln(os.Stderr, "Usage: /print $VAR_NAME")
					continue
				}
				varName := strings.TrimPrefix(parts[1], "$")
				if val, ok := os.LookupEnv(varName); ok {
					fmt.Printf("%s=%q\n", varName, val)
				} else {
					fmt.Printf("%s is not set.\n", varName)
				}
			case "/unset":
				if len(parts) != 2 {
					fmt.Fprintln(os.Stderr, "Usage: /unset $VAR_NAME")
					continue
				}
				varName := strings.TrimPrefix(parts[1], "$")
				os.Unsetenv(varName)
				fmt.Printf("%s unset.\n", varName)
			case "/vars":
				envVars := os.Environ()
				for _, envVar := range envVars {
					fmt.Println(envVar)
				}
			default:
				fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
			}
		} else {
			// Process as environment variable definition
			vars := envlang.GetVariables(d, line)
			for key, value := range vars {
				if value == nil {
					os.Unsetenv(key)
					fmt.Printf("Unset %s\n", key)
				} else {
					os.Setenv(key, *value)
					fmt.Printf("Set %s=%q\n", key, *value)
				}
			}
		}
	}
}
