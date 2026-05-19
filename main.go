package main

import (
	"log"

	"sigap-backend/bootstrap"
)

func main() {
	err := bootstrap.Start()
	if err != nil {
		log.Fatalf("Gagal memuat aplikasi: %v", err)
	}

}
