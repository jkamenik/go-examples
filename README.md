# go-examples

Just some playing around with GoLang.  This repo is meant to be used *as the GOPATH* (see `go help gopath` for details on what that means).

## Simple Setup

Use the `jkamenik/go:1.9` docker container.

```bash
$ git clone git@github.com:jkamenik/go-examples.git
$ cd go-examples
$ docker run -ti --rm -v $(pwd):/go jkamenik/go:1.9
```

## Usage

Each directory is a separate example.