package main

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
)

type HtmlData struct {
	Title string
	Imports []string
	Layout string
	JavaScript template.HTML
}

func handleRoute(htmlData HtmlData) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		writePage(w, htmlData)
	}
}

func getPagesPaths(paths []string, template string) []string {
	result := make([]string, 0)

	for _, v := range paths {
		result = append(result, "pages/" + v)
	}

	if len(template) > 0 {
		result = append(result, "pages/" + template)
	}

	return result
}

func writePage(w http.ResponseWriter, htmlData HtmlData) {
	pagesPaths := getPagesPaths(htmlData.Imports, htmlData.Layout)

	templates := template.Must(template.ParseFiles(pagesPaths...))
	var err error

	if len(htmlData.Layout) > 0 {
		err = templates.ExecuteTemplate(w, htmlData.Layout, htmlData)
	} else {
		err = templates.Execute(w, htmlData)
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func routeHandler() *http.ServeMux {
	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./pages/static/"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("/", handleRoute(GetIndex()))
	router.HandleFunc("/index/partial", handleRoute(GetIndex_Partial()))
	
	return router
}

func main() {
	router := routeHandler()

	fmt.Println("Now listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
