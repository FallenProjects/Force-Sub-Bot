package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token       string
	OwnerId     int64
	LoggerId    int64
	DatabaseURI string
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func init() {
	_ = godotenv.Load()

	DatabaseURI = os.Getenv("DB_URI")
	Token = os.Getenv("8653203719:AAFXFxCaFYJuRyJ9sBaDtJ3dc7XSs7zBUos")
	OwnerId = toInt64(os.Getenv("@AMIT_OF"))
	LoggerId = toInt64(os.Getenv("LOGGER_ID"))

	// Log to stdout instead of a file because vercel doesn't support writing to files
	InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	setDefaults()
}
