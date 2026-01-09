package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 输入明文密码
	plainPassword := "123456" // 可以修改为你想要的密码

	// 使用 bcrypt 加密密码，cost factor = 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("加密失败: %v\n", err)
		return
	}

	// 输出加密后的密码
	fmt.Printf("明文密码: %s\n", plainPassword)
	fmt.Printf("加密后的密码: %s\n", string(hashedPassword))
	fmt.Println("\n将加密后的密码复制到数据库中即可")
}

