# Password Generator CLI

## Purpose
A simple command line random string generator. Its purpose is only to serve as a milestone in my studies on Go.

## Usage
**Clone this repo:**
```bash
git clone git@github.com:7Cass/pwd-gen-cli.git
```
**Install the packages**
```bash
go get -u
go mod tidy
```
**Build and run**
```bash
go build main.go
# Without args default length is 9
./main
# Max length is 100
./main rnd 20
```
