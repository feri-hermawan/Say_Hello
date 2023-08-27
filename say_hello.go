package sayhello

import (
	"strconv"
)

func SayHello(name string, age int) string {
	return "Hello my name is " + name + "and my age is " + strconv.Itoa(age)
}
