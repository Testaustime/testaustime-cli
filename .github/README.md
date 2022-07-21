# testaustime-cli
Command line utility for [Testaustime](https://testaustime.fi).

**NOTE**: Highly under developement. Please open an issue if you find any bugs.

## Features
- Colored output (optionally disabled)
- Profile information
- Simple coding statistics

## Installation
```sh
git clone https://github.com/romeq/testaustime-cli
cd testaustime-cli

# install dependencies and compile binary
go get -u
go build main.go

# link binary to path
ln -s $PWD/main ~/.local/bin/testaustime-cli
```

