package gologger

import (
	"fmt"
	"os"
	"time"
)

//CustomLogger is heavily customisable logger which allows for complex file naming conventions and backups
type CustomLogger struct {
	Path string
	//Extention is the file extention; .txt, .json, .log
	Extention string
	//NameingConvention returns a string which will be used as the filename
	NameingConvention func() string

	//LineTerminator usually is "\n"
	LineTerminator string

	//ConventionUpdate is how long the logger should wait until it checks the current naming convention and if its time to change the file handle will be changed
	ConventionUpdate time.Duration

	//Callbacks

	//ConventionUpdated can be used to backup the old log file after the convention has change
	ConventionUpdated func(oldFile string, newFile string)

	queue chan ([]byte)
	close chan (struct{})
}

//NewCustomLogger creates a new custom logger
func NewCustomLogger(path string, extention string, bufferSize int) *CustomLogger {
	logger := &CustomLogger{
		Extention: extention,
		queue:     make(chan []byte, bufferSize),
		close:     make(chan struct{}),

		NameingConvention: func() string {
			return ""
		},

		//Never bother updating the convention
		ConventionUpdate: ((time.Hour * 24) * 7) * 2000,

		LineTerminator: "\n",
		Path:           path,
	}

	return logger
}

func (l *CustomLogger) Write(data interface{}) {
	l.queue <- []byte(fmt.Sprint(data))
}

func (l *CustomLogger) getFileName() string {
	return l.Path + l.NameingConvention() + l.Extention
}

func (l *CustomLogger) getFileHandle() (*os.File, string, error) {
	fileName := l.getFileName()

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, "", err
	}

	return file, fileName, nil
}

//Service will start the logging service
func (l *CustomLogger) Service() error {

	handle, fileName, err := l.getFileHandle()
	if err != nil {
		return err
	}

	var newFileName string
	for {
		select {

		case <-time.NewTicker(l.ConventionUpdate).C:

			handle, newFileName, err = l.getFileHandle()
			if err != nil {
				return err
			}

			fileName = newFileName

			if l.ConventionUpdated != nil {
				//Run on another goroutine to prevent the service worker from hanging
				go l.ConventionUpdated(fileName, newFileName)
			}

			break
		case data := <-l.queue:
			handle.Write(append(data, l.LineTerminator...))
			break
		case <-l.close:
			return handle.Close()
		}
	}
}

//Close will shutdown the service worker
func (l *CustomLogger) Close() {
	l.close <- struct{}{}
}
