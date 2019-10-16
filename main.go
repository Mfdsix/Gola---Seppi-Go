package main

import(
	"net/http"
	"html/template"
	"fmt"
)

type M map[string]interface{}

func main(){
	var tmpl, err = template.ParseGlob("views/*")
	if(err != nil){
		fmt.Println(err.Error())
		return
	}

	http.Handle("/static/", 
		http.StripPrefix("/static/", 
			http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		var data = M{
			"title" : "Go - Index Page",
			"name" : "Mahfudz Ainur Rif'an",
		}

		err = tmpl.ExecuteTemplate(w, "index", data)
		if(err != nil){
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
		var data = M{
			"title" : "Go - About Page",
		}

		err = tmpl.ExecuteTemplate(w, "about", data)
		if(err != nil){
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server Started at :6511")
	http.ListenAndServe(":6511", nil)
}