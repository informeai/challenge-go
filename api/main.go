package main

import (
	"fmt"
	delivery "github.com/informeai/challenge-go/api/delivery"
	"log"
)

func main() {
	//server running...
	d := delivery.UserDeliveryMemory{}
	fmt.Println("server running...")
	log.Fatalln(d.Run(":4000"))

}
