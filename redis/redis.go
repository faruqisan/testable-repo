package redis

import (
	"fmt"
)

func Set(key string, data interface{}) {
	fmt.Println("Real redis callled", key, data)
}

func Get(key string) {
	fmt.Println("Real redis called to return : ", key)
}
