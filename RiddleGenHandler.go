package mathch

import (
	"fmt"
	"net/http"
	"strconv"
)

func NewRiddle(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprint(writer, RandGenerate(-1))
		return
	}
	_, ok := request.PostForm["diff"]
	if !ok {
		fmt.Fprint(writer, RandGenerate(-1))
		return
	}
	val, err := strconv.Atoi(request.PostFormValue("diff"))
	if err != nil {
		fmt.Fprint(writer, RandGenerate(-1))
		return
	}
	fmt.Fprint(writer, RandGenerate(val))
}
func CheckSolution(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprint(writer, "0")
		return
	}
	_, ok1 := request.PostForm["eq"]
	_, ok2 := request.PostForm["ans"]
	if !ok1 || !ok2 {
		fmt.Fprint(writer, "0")
		return
	}
	if Solve(request.PostFormValue("eq"), request.PostFormValue("ans")) {
		fmt.Fprint(writer, "1")
	} else {
		fmt.Fprint(writer, "0")
	}
}
