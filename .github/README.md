# testaustime-cli
Fast and secure command line utility for [Testaustime](https://testaustime.fi). 

## Privacy and Security
More of the security practices this program implies

### How passwords are processed

**NOTE**: These practices are not done with authentication token, as it already lives in a file accessible by any program you run.

- Passwords are never shown in terminal when typing
- Passwords are overwritten in memory with random data (from `/dev/urandom`) 
  when they are not needed anymore
- Passwords shall not be copied in memory, but rather passed as pointers to other functions

### Usage of third-party libraries
I am actively trying to restrict the usage of third-party libraries especially if core libraries implement the same things 
or if they are not exactly essential for the program.

### Your privacy

**Your Testaustime data**

Your testaustime data is never sent anywhere except the [open-source backend](https://github.com/Testaustime/testaustime-backend). If you don't want
to use the shared instance, you can setup your own backend instance. 
All coding data is fetched from the backend server you have configured in this program's configuration file.

**Application telemetries**
This application doesn't collect any telemetries as of now


## Installation

```sh
git clone https://github.com/Testaustime/testaustime-cli
cd testaustime-cli

# install config
mkdir -p ~/.config/testaustime-cli
cp configtemplate.toml ~/.config/testaustime-cli/config.toml

# install dependencies and compile binary
make clean install
```

## Features
List of currently implemented features

### Account
Manage your account

- Account information (username, registration time, friend code)
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

### Data for a specific user
Show specific friend's coding statistics

- Past 24 hours
- Past week
- Past month
- All time
- Top languages, projects, editors and hosts



## Contributing

I am always pleased with more contributors in this project.
I'd appreciate if you would make sure following things are fine before making a new pull request:

- `go.mod` shall not contain any libraries that are not used
  - `go mod tidy`
- code builds :) 
  - `make build`
- all tests shall pass 
  - `make test`
- linter has no issues
  - `golangci-lint run`
- `.github/README.md` is up to date with your pull request
