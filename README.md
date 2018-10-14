# go-grpc-health
Simple binary, which checks if grpc server is started or not

# Build

As any GO package.

If you're not familiar with go or don't have golang installed in a system, then run:
```bash
$ ./build.sh
```

This script should produce static binary `go-grpc-health` for your host OS.

### Crosscompilation

You also can produce binary from any os for any other OS, thanks to GO:
```bash
$ uname
Darwin

$ GOOS=linux ./build.sh
$ file go-grpc-health
go-grpc-health: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, not stripped

$ GOOS=windows ./build.sh
$ file go-grpc-health.exe                                                                      
go-grpc-health.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows

$ GOOS=darwin ./build.sh
$ file go-grpc-health                                                                          
go-grpc-health: Mach-O 64-bit executable x86_64
```

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

