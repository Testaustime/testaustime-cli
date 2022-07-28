# testaustime-cli
Command line utility for [Testaustime](https://testaustime.fi).

## Features
List of currently implemented features

### Account
Manage your account

- Account information (username, registration time, friendcode)
- Registration
- Login with username and password
- Show authentication token
- Generate a new authentication token
- Generate a new friend code
- Change password

### Statistics
Show coding statistics

- Past 24 hours
- Past week
- Pasth month
- All time
- Top languages, projects, editors and hosts

### Friends
Show friends' coding statistics

- Past week
- Past month
- Add a new friend
- Remove a friend

### Find user
Show specific friend's coding statistics

- Past 24 hours
- Past week
- Pasth month
- All time
- Top languages, projects, editors and hosts

## Installation

```sh
git clone https://github.com/romeq/testaustime-cli
cd testaustime-cli

# install config
mkdir -p ~/.config/testaustime-cli
cp config.toml.example ~/.config/testaustime-cli/config.toml

# install dependencies and compile binary
go get -u
go build main.go

# link binary to path
ln -s $PWD/main ~/.local/bin/testaustime-cli
```


