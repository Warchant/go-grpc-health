# go-grpc-health
Simple binary, which checks if grpc server is started or not

# Build

As any GO package.

If you're not familiar with go or don't have golang installed in a system, then run:
```bash
$ ./build.sh
```

This script should produce static binary `go-grpc-health` for your host OS.

# Usage


```bash
$ ./go-grpc-health -h
Usage of ./go-grpc-health:
  -host string
        Specify grpc host (default "localhost")
  -port int
        Specify grpc port (default 50051)
  -timeout int
        Timeout, ms (default 1000)
  -verbosity string
        Verbosity level [panic,fatal,error,warn,info,debug] (default "warn")

$ ./go-grpc-health -host 127.0.0.1 -port 32864 && echo UP || echo DOWN
WARN[0000] 127.0.0.1:32864 is up
UP

$ ./go-grpc-health -host 127.0.0.1 -port 11111 && echo UP || echo DOWN
WARN[0000] 127.0.0.1:11111 is down
DOWN
```

