package httpServer

import (
	"fmt"
	"net/http"
)

var hostServer = "localhost"
var port = "4000"

type Hello struct{}

func (h Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<H1> Hello Server is started </H1>")
	fmt.Fprint(writer, "<H4>Remote Cliente Address is :</H4>"+request.Host)
}

func StartHttpServer(hello Hello) {
	err := http.ListenAndServe(hostServer+":"+port, hello)
	CheckError(err)
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
