package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() { //单向加密存储
	password := "jack"
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return
	}
	check(bytes, []byte("jack"))
	fmt.Println(string(bytes))
}

func check(bytes []byte, password []byte) {
	err := bcrypt.CompareHashAndPassword(bytes, password)
	if err != nil {
		fmt.Println("与密文不匹配")
	} else {
		fmt.Println("匹配,验证成功")
	}
}
