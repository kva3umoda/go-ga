package helper

import (
	"log"
	"os"
)

func GetCurDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
	//fmt.Println(dir)
}
