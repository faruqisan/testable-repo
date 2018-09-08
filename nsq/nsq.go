package nsq

import (
	"fmt"
)

func Publish(data string) error {
	fmt.Println("Real Publisher Called ", data)
	return nil
}