package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func main() {
	defaultCost := bcrypt.DefaultCost
	fmt.Printf("Default Cost: '%d'\n", defaultCost)
	for i := 0; i <= 6; i++ {
		start := time.Now()
		bcrypt.GenerateFromPassword([]byte("password"), defaultCost + i)
		fmt.Printf("Cost: %d, Duration: %v\n", defaultCost + i, time.Since(start))
	}
}
