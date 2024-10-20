package handler

import "fmt"

func ErrorHnadler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}