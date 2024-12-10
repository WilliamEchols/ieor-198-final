package utils

import (
	"io"
	"log"
	"os"
	"time"
)

func SetupLogging() (*os.File, error) {
	filename := "./logs/" + time.Now().Format("20060102_150405") + ".log"
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	startLog()

	return logFile, nil
}

func startLog() {
	log.Printf("")
	log.Printf("  ███╗   █████╗  █████╗ ")
	log.Printf(" ████║  ██╔══██╗██╔══██╗")
	log.Printf("██╔██║  ╚██████║╚█████╔╝")
	log.Printf("╚═╝██║   ╚═══██║██╔══██╗")
	log.Printf("███████╗ █████╔╝╚█████╔╝")
	log.Printf("╚══════╝ ╚════╝  ╚════╝ ")
	log.Printf("")
	log.Printf("        ieor-198        ")
	log.Printf("")
}
