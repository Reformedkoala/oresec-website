package main

import (
    "html/template"
	"log"
	"net/http"
	"os"
    //"fmt"
    "path/filepath"
    "strings"
)

type TemplateData struct{
    Active string
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
    lp := filepath.Join("static", "templates", "layout.html")
	fp := filepath.Join("static", "templates", filepath.Clean(r.URL.Path))
    // Return a 404 if the template doesn't exist

    if filepath.Clean(r.URL.Path) == "/" {
        fp = filepath.Join("static", "templates", "index.html")
        tmpl, err := template.ParseFiles(lp, fp)
        if err != nil {
            log.Print(err.Error())
            http.Error(w, http.StatusText(500), 500)
            return
        }
        data := TemplateData {
            Active: "index.html",
        }
        //log.Print(data)
        err = tmpl.ExecuteTemplate(w, "layout", data)
        return
    }


	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

    tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Print(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}
    
    data := TemplateData {
        Active: strings.TrimPrefix(filepath.Clean(r.URL.Path),"/"),
    }

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

}

func main() {
    port := "4040"

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/",fs))

    http.HandleFunc("/", serveTemplate)

    log.Print("Listening on :" + port +"...")
    err := http.ListenAndServe(":4040", nil)
	if err != nil {
		log.Fatal(err)
	}
}
