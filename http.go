package main

import (
	"html"
	"log"
	"net/http"
	"strings"

	"github.com/alecthomas/template"
)

type Info struct {
	Prefix string
	Author string
}

func welcomePage(w http.ResponseWriter, r *http.Request) {

	htmltxt := `<h1> Welcome to my page </h1> You will find something interesting apps. Browse local files  <a href={{.Prefix}}> {{.Prefix}} </a>. <tiny>{{.Author}}</tiny> </br>
	  The <a href="/api/apps">apps</a>  and
		<a href="/api/apps/2"/> /api/apps/2 </a>are sample API handlers`
	tmpl := template.New("welcome") //  ParseFiles("template.txt")
	tmpl.Parse(htmltxt)

	// tmpl, err := template.New("template.txt") //  ParseFiles("template.txt")
	// if err != nil {
	// 	log.Print(err)
	// }
	info := Info{"/www", "@iamssk"}
	// tmpl.Execute(w, nil)
	tmpl.ExecuteTemplate(w, tmpl.Name(), info)

}
func Index(w http.ResponseWriter, r *http.Request) {

	log.Print("Root is here ", r.URL.Path, " escpaed = ", html.EscapeString(r.URL.Path))
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fname := "./" + strings.TrimPrefix(r.URL.Path, "/www")
	// if fname != "" {
	http.ServeFile(w, r, fname)
	log.Print("Will server this file ", fname)
	// } else {
	// 	http.ServeFile(w, r, "./")
	// }

}
