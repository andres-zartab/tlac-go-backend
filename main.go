package main

import (
	"fmt"

	"github.com/andres-zartab/tlac-go-backend/tlac"
)

func main() {
	db, err := tlac.NewDB()
	if err != nil {
		fmt.Println("wut")
	}
	fmt.Println(db)
}
