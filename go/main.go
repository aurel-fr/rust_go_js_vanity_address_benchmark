package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/algorand/go-algorand-sdk/crypto"
)

var mu sync.Mutex

func main() {
	count := 0
	t0 := time.Now()
	var list [5]string
	for count < 5 {
		acc := crypto.GenerateAccount()
		if strings.HasPrefix(acc.Address.String(), "TEST") {
			list[count] = acc.Address.String()
			count++

		}
	}
	fmt.Printf("Singlethreaded Go found 5 adresses in %v secs\n%v\n", time.Since(t0).Seconds(), list)

	count = 0
	t0 = time.Now()
	for count < 5 {
		go func() {
			acc := crypto.GenerateAccount()
			if strings.HasPrefix(acc.Address.String(), "TEST") {
				mu.Lock()
				list[count] = acc.Address.String()
				count++
				mu.Unlock()
			}
		}()
	}
	fmt.Printf("Multithreaded Go found 5 adresses in %v secs\n%v\n", time.Since(t0).Seconds(), list)

	os.Exit(0)
}
