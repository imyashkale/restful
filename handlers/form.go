package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/schema"
)

type Profile struct {
	Name string
	Age  int
}

var decoder *schema.Decoder

func init() {
	decoder = schema.NewDecoder()
}

func LoadForm(req *restful.Request, resp *restful.Response) {
	var tmp *template.Template
	var err error
	if tmp, err = template.ParseFiles("./templates/address.html"); err != nil {
		resp.WriteError(http.StatusInternalServerError, errors.New("internal server error"))
	}
	tmp.Execute(resp.ResponseWriter, tmp)
}

func SaveForm(req *restful.Request, resp *restful.Response) {
	var err error
	if err = req.Request.ParseForm(); err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	p := new(Profile)
	if err = decoder.Decode(p, req.Request.PostForm); err != nil {
		resp.WriteErrorString(http.StatusBadGateway, err.Error())
		return
	}
	resp.AddHeader("Content-Type", "text/html; charset=utf-8")
	io.WriteString(resp.ResponseWriter, fmt.Sprintf("<html><body>Name=%s, Age=%d</body></html>", p.Name, p.Age))
}

func PostHandler(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp.ResponseWriter, "Resp From PostHandler!")
}
func UserHandler(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp.ResponseWriter, "Resp From UserHandler!")
}
