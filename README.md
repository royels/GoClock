# GoClock

I'll be using one of my old systems programming assignments as a base with which to create this (won't be translated code, but definitely...inspired.)

This project is essentially a binary decimal clock implemented in Go(lang). 

To create: Build executable with `go build ~/path/to/main/dir/pa2.go` or `go install github.com/royels/main` with your $GOPATH set to `~/project/`, then run.

To run:
```shell
Usage: bin/main startTime [intervalSeconds [numTicks]]
    (required startTime = HH:MM:SS)
    (HH = [0-23]; MM = [0-59]; SS = [0-59])
    (optional intervalSeconds = [1-86399], default = 1)
    (optional numTicks = [0-61], default = 7)
```
