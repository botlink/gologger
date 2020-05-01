# Gologger

Gologger is a concurrent **thread safe** logging system which writes to the filesystem. Gologger utilises a service worker model with buffered channels/queues.

As soon as data is append to the queue the Write() function will return resulting in **none blocking** logging.

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