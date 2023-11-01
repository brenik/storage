package main

import (
	"fmt"
	"github/brenik/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("It works", st)
}
