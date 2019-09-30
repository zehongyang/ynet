package utils

import (
	"fmt"
	"io/ioutil"
)

func ReadConfig()  {
	bytes, err := ioutil.ReadFile("./conf/config.json")
	if err != nil {
		fmt.Println("")
	}
}
