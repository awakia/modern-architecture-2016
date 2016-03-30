package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/dimfeld/httptreemux"
	"github.com/unrolled/render"
)

func main() {
	ren := render.New()
	router := httptreemux.New()

	router.GET("/", func(w http.ResponseWriter, req *http.Request, params map[string]string) {
		ren.JSON(w, http.StatusOK, map[string]string{"message": "hello go server!"})
	})
	router.GET("/:name", func(w http.ResponseWriter, req *http.Request, params map[string]string) {
		ren.JSON(w, http.StatusOK, map[string]string{"message": "hello " + params["name"] + "!"})
	})
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":5000")
}
