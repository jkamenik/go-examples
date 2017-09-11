# echo

A simple echo server.  This shows the usage of io using the `net.Listen` and `io.Copy` functions.  `net.Listen` sets up a raw TCP socket listening on a given port.  Currently that is hardcoded to `localhost:4000`.

This function returns a listener but doesn't actually listen for incoming request.  The `Accept` function listens for and then returns the active connection as a `io.Reader` like object.  A goroutine running `io.Copy` is used so that the main thread can go back to listening for incoming requests.

## Usage

Because this listens on localhost to use it in a docker container you must `exec` in a second time.

```bash
$ cd /path/to/go-examples
$ docker run -ti --rm $(pwd):/go --name echo jkamenik/go:1.9
bash-4.3# cd src/echo
bash-4.3# go build
bash-4.3# ./echo
Starting
Listening on  localhost:4000
Waiting for connection
```

In a second terminal.

```bash
$ docker exec -ti echo /bin/bash
bash-4.3# telnet localhost:4000
type in anything you want echoed
type in anything you want echoed
```
