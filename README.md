go-examples
===========

Just some playing around with GoLang

Setup
-----

This requires "go".  For Macs install it from homebrew.

```
$ brew install go
```

Install
-------

```
$ git clone git@github.com:jkamenik/go-examples.git
```

Usage
-----

Each dir is a self-contained executable.  You will need to setup the Go ENV variables and build the executables.

```
$ export GOPATH=`pwd`
$ export GOBIN=`pwd`/bin
$ go install src/new/*.go
$ ./bin/new
# output of the "new" program
```
