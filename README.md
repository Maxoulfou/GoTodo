# Simple TODO CLI application written in Go

## Basic

To start, you will have to put the executable in an isolated folder, then the application create a file named
".TodoFile.json".
Don't edit them

## Detail

```bash
GoTodo is a simple CLI application written in Golang
to manage your tasks only with a command line tool.

Usage:
  gotodo [command]

Available Commands:
  add         Add a new todo
  complete    Mark todo as completed
  completion  Generate the autocompletion script for the specified shell
  del         Delete todo
  help        Help about any command
  list        List all todos

Flags:
  -h, --help   help for gotodo

Use "gotodo [command] --help" for more information about a command.

```

## Compilation

`go build main.go -o bin/`

## Other

- Terminal color should fix in future
- MySql storage / SQLite3 / PostgreSQL ???

## Verifications checksum

> Theses checksum are valid only for my executable in Releases

SHA1 : e632fe821723432c8a9cdc401c1678ead3b0ddb2

MD5 : 2e219204eab5c131c5125f6b1a49003c

### How to get checksum with my CMD ?

`Certutil -hashfile gotodo.exe MD5`

`Certutil -hashfile gotodo.exe SHA1`