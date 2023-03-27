package utils

import (
	"fmt"
	"log"
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func CheckError(msg string, err error) {
	if err != nil {
		log.Printf("❗ ERROR DETECTED ❗ %s => %s", msg, err)
	}
}

func FailOnError(msg string, err error) {
	if err != nil {
		log.Fatalf(fmt.Sprintf("❗❗ PANIC ❗❗ %s => %s", msg, err))
	}
}
