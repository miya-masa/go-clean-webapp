package web

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "web", log.LstdFlags)
