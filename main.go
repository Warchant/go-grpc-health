package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"time"
)

var (
	host      string
	port      int
	timeout   int
	verbosity string
)

const (
	checkEndpoint string = "933fdc21e012e7a28a1cb4ff49a9b749218b8568d7f6663ce99ebc57db9e3922da18994c358064029112c656cc9342d185688cd0fe33209d3ae140436dc733b5"
)

func main() {
	flag.StringVar(&host, "host", "localhost", "Specify grpc host")
	flag.IntVar(&port, "port", 50051, "Specify grpc port")
	flag.IntVar(&timeout, "timeout", 1000, "Timeout, ms")
	flag.StringVar(&verbosity, "verbosity", "warn", "Verbosity level [panic,fatal,error,warn,info,debug]")
	flag.Parse()

	level, err := logrus.ParseLevel(verbosity)
	if err != nil {
		logrus.Errorf("%v\n", err)
		os.Exit(2)
	}
	logrus.SetLevel(level)

	logrus.Debugf("Starting go-grpc-healthcheck...")

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	logrus.Debugf("Checking availability of %s...", conn.Target())
	if err != nil {
		down(conn.Target())
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Millisecond)
	defer cancel()

	err = conn.Invoke(ctx, checkEndpoint, nil, nil)
	if err != nil {
		code := status.Code(err);
		logrus.Infof("Returned status code: %d, message: %v", code, err)

		switch code {
		case codes.Unimplemented:
			fallthrough
		case codes.OK:
			up(conn.Target())
		}
	}

	down(conn.Target())
}

func down(address string) {
	logrus.Warnf("GRPC server at %s is down", address)
	os.Exit(1)
}

func up(address string) {
	logrus.Warnf("GRPC server at %s is up", address)
	os.Exit(0)
}
