package handlers

import (
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/emicklei/go-restful/v3"
)

type Message struct {
	Text string
}

func RootHandler(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "Yash Kale")
}

func TemplateHandler(req *restful.Request, resp *restful.Response) {
	msg := &Message{Text: "Yash Kale"}
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(resp.ResponseWriter, msg)
}

func JsonHandler(req *restful.Request, resp *restful.Response) {
	resp.WriteJson(struct {
		Name string `json:"name"`
	}{
		Name: "Yash",
	}, "application/json")

}

func WriteHeaderAndJsonHandler(req *restful.Request, resp *restful.Response) {
	content := struct {
		Name string `json:"name"`
	}{
		Name: "Yash",
	}
	resp.WriteHeaderAndJson(http.StatusOK, content, "applicaion/json")
}
