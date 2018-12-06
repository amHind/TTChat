package main

import (
	"fmt"
	"net"
	"strings"
)


func main() {
	ln, _ := net.Listen("tcp", ":8080") //le serveur écoute sur toutes les interfaces
	for {
		conn, _:= ln.Accept()
		cmdLine := make([]byte,(1024 * 4))
		n,_= conn.Read(cmdLine)
		}
	fmt.Println("WELCOME TO TCCHAT")
	//fmt.Printf("%q\n", strings.Split("a,b,c", ","))










	var names []string
	names = strings.Split("Ta,b,c", ",")
	fmt.Printf("%q\n", names)

	for _, name := range names {
		fmt.Println(name)

	}
	for name := range names {
		fmt.Println(name)

	}
}

