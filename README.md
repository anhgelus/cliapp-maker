# CliApp Maker

A GoLang library aiming to help to make cli application.

Lightweight and simple to use.

## Install

To install the library, just run 
```bash
$ go get github.com/anhgelus/cliapp-maker@latest
```

## How to use

To start using this library, you must understand the vocabulary used:

- `CliApp` - Your application
- `Cmd` - Subcommand of your application (e.g.: `install` for `myapp install .`)
- `Option` - The flags (e.g.: -v, --foo bar)
- `Param` - The parameters of the subcommand (e.g.: `.` for `myapp install .`)
    :warning: The param are not handled by our system, you must do it by yourself!

**This part is not updated!** I'll do this later.

Look at the CliApp type to understand how it works (a future refactor will change this, so I'll not finish this part before)

## Features

- [x] Basic cli features (options, subcommand and params support)
- [x] Automatic help handler (with the option -h or without subcommand)
- [x] Customisable global options
- [ ] Beautiful display

### In coming refactor

- [x] Better management of basic cli features (look at #1)

## Technologies

- GoLang 1.19
- More soon
