package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "bgvyzdsv"
	hashStartingWith := []string{"00000", "000000"}
	for _, hsw := range hashStartingWith {
		fmt.Printf("The secret number with hash starting with %s is: %d\n", hsw, findSecretNumber(input, hsw))
	}
}

func findSecretNumber(input, hashStartsWith string) int {
	var secretNumber int
	for {
		secretKey := []byte(input + strconv.Itoa(secretNumber))
		md5Hash := fmt.Sprintf("%x", md5.Sum(secretKey))
		if strings.HasPrefix(md5Hash, hashStartsWith) {
			break
		}
		secretNumber++
	}
	return secretNumber
}
