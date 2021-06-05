package main

import "awesomeProject/httpServer"

func main() {
h:= httpServer.Hello{}
httpServer.StartHttpServer(h)
//readFromWeb.ReadTextUrlFromWeb()
}
