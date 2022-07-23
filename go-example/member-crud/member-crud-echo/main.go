package main

import "github.com/MichaelYcCho/go-practice/go-example/member-crud/member-crud-echo/member"

func main() {
	server := member.NewDefaultServer()
	server.Run()
}
