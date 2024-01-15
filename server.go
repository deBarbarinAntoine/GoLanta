package GoLanta

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

var tmpl = make(map[string]*template.Template)

var (
	_, b, _, _ = runtime.Caller(0)
	path       = filepath.Dir(b) + "/"
)

// Establishing the fileserver to make assets directory available for the client.
func fileServer() {
	fs := http.FileServer(http.Dir(path + "assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

// runServer runs the server in a goroutine and open the browser at the correct URL and then loops waiting for the "stop" input to end all goroutines.
func runServer() {
	port := "localhost:8080"
	url := "http://" + port + "/"
	go http.ListenAndServe(port, nil)
	fmt.Println("Server is running...")
	time.Sleep(time.Second * 5)
	cmd := exec.Command("explorer", url)
	cmd.Run()
	fmt.Println("If the navigator didn't open on its own, just go to ", url, " on your browser.")
	isRunning := true
	for isRunning {
		fmt.Println("If you want to end the server, type 'stop' here :")
		var command string
		fmt.Scanln(&command)
		if command == "stop" {
			isRunning = false
		}
	}
}

// Run is the public function that executes all necessary functions to run the server and website.
func Run() {
	tmplPath := path + "templates/"
	tmpl["index"] = template.Must(template.ParseFiles(tmplPath+"index.gohtml", tmplPath+"base.gohtml"))
	tmpl["home"] = template.Must(template.ParseFiles(tmplPath+"home.gohtml", tmplPath+"base.gohtml"))
	tmpl["character"] = template.Must(template.ParseFiles(tmplPath+"character.gohtml", tmplPath+"base.gohtml"))
	tmpl["create"] = template.Must(template.ParseFiles(tmplPath+"create.gohtml", tmplPath+"base.gohtml"))
	tmpl["update"] = template.Must(template.ParseFiles(tmplPath+"update.gohtml", tmplPath+"base.gohtml"))
	tmpl["remove"] = template.Must(template.ParseFiles(tmplPath+"remove.gohtml", tmplPath+"base.gohtml"))
	tmpl["search"] = template.Must(template.ParseFiles(tmplPath+"search.gohtml", tmplPath+"base.gohtml"))
	tmpl["error404"] = template.Must(template.ParseFiles(tmplPath+"error404.gohtml", tmplPath+"base.gohtml"))
	routes()
	fileServer()
	runServer()
}
