package main

import (
	"net/http"

	"github.com/fatma49/DTS_GO_2022/controllers/taskcontroller"
)

func main() {
	http.HandleFunc("/", taskcontroller.Index)
	http.HandleFunc("/task/get_form", taskcontroller.GetForm)
	http.HandleFunc("/task/store", taskcontroller.Store)
	http.HandleFunc("/task/delete", taskcontroller.Delete)

	http.ListenAndServe(":8000", nil)

}
