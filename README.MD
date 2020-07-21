# Gologger

![Go report card](https://goreportcard.com/badge/github.com/syrinsecurity/gologger)
[![Build Status](https://travis-ci.org/syrinsecurity/gologger.svg?branch=master)](https://travis-ci.org/syrinsecurity/gologger)
[![GoDoc](https://godoc.org/github.com/syrinsecurity/gologger?status.svg)](https://godoc.org/github.com/syrinsecurity/gologger)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/syrinsecurity/gologger/graphs/commit-activity)
[![License](https://img.shields.io/github/license/syrinsecurity/gologger.svg)](https://github.com/syrinsecurity/gologger/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/syrinsecurity/gologger.svg)](https://GitHub.com/syrinsecurity/gologger/releases/)
[![GitHub issues](https://img.shields.io/github/issues/syrinsecurity/gologger.svg)](https://GitHub.com/syrinsecurity/gologger/issues/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)




Gologger is a concurrent **thread safe** logging system which writes to the filesystem. Gologger utilises a service worker model with buffered channels/queues.

As soon as data is append to the queue the Write() function will return resulting in **none blocking** logging.

### Install

> go get github.com/syrinsecurity/gologger

### Examples

Custom loggers

```go
func main() {

	/*
	"./file.log" is where the logs will be written to. If this file does not exist it will be created.

	"200" is the buffer size of the channel
	*/
	logger, err := gologger.New("./file.log", 200)
	if err != nil {
		panic(err)
	}

	logger.WriteString("log any data you want.")
	logger.WriteString("This is a interface so you can use any type you like.", "UserID:", 9039832898, "Timestamp:", time.Now().Unix())
}
```

```go
package main

import "github.com/syrinsecurity/gologger"

var (
	logger, _ = gologger.New("./file.log", 200)
)

func main() {
	logger.Write("data", 123, 345)
}
```

Use the built in loggers for fast setup

```go
package main

import "github.com/syrinsecurity/gologger"

func main() {

	//Start the service worker
	//Leave the path blank for loggers you do not want to use
	go gologger.Service("./error.log", "", "")

	gologger.Write(gologger.LogError, "error data")

	//Quick logger syntax
	gologger.Error.Write(errors.New("example error"))
}
```

Get the size of all queues:

> gologger.QueueSize()

For more examples look at https://github.com/syrinsecurity/gologger/tree/master/examples
