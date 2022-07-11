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

SHA1 : 3419c8da7c400bb1ceeb00767b9a84dfc315e031

MD5 : 503457e4670121141fe0c181b7fd3f4e

### How to get checksum with my CMD ?

`Certutil -hashfile gotodo.exe MD5`

`Certutil -hashfile gotodo.exe SHA1`
