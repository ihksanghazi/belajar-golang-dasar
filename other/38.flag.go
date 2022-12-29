package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "put Your Database Host")
	var user *string = flag.String("user", "root", "put Your Database User")
	var pass *string = flag.String("password", "root", "put Your Database Password")

	flag.Parse()

	fmt.Println("Host = ", *host)
	fmt.Println("User = ", *user)
	fmt.Println("Password = ", *pass)
}
