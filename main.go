package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/emicklei/go-restful/v3"
)

type User struct {
	UserId   string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// container contains group of webservices
	wc := restful.NewContainer()

	// creating new webservices
	wsUser := new(restful.WebService)
	wsPost := new(restful.WebService)
	wsAuth := new(restful.WebService)

	// root path
	rAuth := wsAuth.Path("").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	rUser := wsUser.Path("/user")
	rPost := wsPost.Path("/post")

	rUser.Route(rUser.GET("").To(UserHandler))
	rPost.Route(rPost.GET("{userid}").To(PostHandler))

	rAuth.Route(rAuth.POST("/login").To(LoginHandler))

	wc.Add(rUser)
	wc.Add(rPost)
	wc.Add(rAuth)

	fmt.Println("Running on :8080")
	log.Fatal(http.ListenAndServe(":8080", wc))
}

func LoginHandler(req *restful.Request, resp *restful.Response) {
	user := new(User)
	var err error
	if err = req.ReadEntity(user); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, "Required fields are not found")
	}

	if err = resp.WriteEntity(user); err != nil {
		log.Fatal(err)
		resp.WriteErrorString(http.StatusInternalServerError, "contact admin")
	}

	
}

func LogoutHandler(req *restful.Request, resp *restful.Response) {

}

// PostHandler
func PostHandler(req *restful.Request, resp *restful.Response) {
	userId := req.PathParameter("userid")
	q := req.QueryParameter("limit")
	sb := strings.Builder{}
	sb.WriteString("Resp From PostHandler! ")
	sb.WriteString(" :")
	sb.WriteString(userId)
	sb.WriteString(" :")
	sb.WriteString(q)

	io.WriteString(resp.ResponseWriter, sb.String())
}

// UserHandler
func UserHandler(req *restful.Request, resp *restful.Response) {
	q := req.QueryParameter("limit")
	sb := strings.Builder{}
	sb.WriteString("Resp From UserHandler!")
	sb.WriteString(" :")
	sb.WriteString(q)
	io.WriteString(resp.ResponseWriter, sb.String())
}
