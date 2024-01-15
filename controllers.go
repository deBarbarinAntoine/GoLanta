package GoLanta

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// rootHandler redirects to index handler.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

// Index page handler.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/index" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := struct {
		Base BaseData
	}{
		Base: BaseData{
			Title:      "Index - GoLanta",
			StaticPath: "static/",
		},
	}
	err := tmpl["index"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	homeHandler
//
// fetch and show a list of all Character present in `characters.json`.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/home" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	chars, err := retrieveChars()
	if err != nil {
		log.Fatal("log: retrieveChars() error!\n", err)
	}
	data := struct {
		Base  BaseData
		Chars []Character
	}{
		Base: BaseData{
			Title:      "Home - GoLanta",
			StaticPath: "static/",
		},
		Chars: chars,
	}
	//fmt.Printf("log: data: %#v\n", data) // testing
	err = tmpl["home"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	articleHandler
//
// fetch and show a specific Article which id number is indicated in the query params.
//
// Query params: ?article=<article-id>
func characterHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/character" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if !r.URL.Query().Has("char") {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("char"))
	if err != nil {
		log.Println("log: characterHandler() strconv.Atoi error!\n", err)
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	char, ok := selectChar(id)
	if !ok {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := struct {
		Base BaseData
		Char Character
	}{
		Base: BaseData{
			Title:      char.Name + " - GoLanta",
			StaticPath: "static/",
		},
		Char: char,
	}
	//fmt.Printf("log: data: %#v\n", data) // testing
	err = tmpl["character"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	searchHandler
//
// fetch and show all Character which Character.Name matches the search indicated in the query params.
//
// Query params: ?q=<search>
func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/search" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if !r.URL.Query().Has("q") {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	search := r.URL.Query().Get("q")
	chars := searchChar(search)
	var message template.HTML
	if len(chars) == 0 {
		message = "<div class=\"message\">There is no character matching your research!</div>"
	}
	data := struct {
		Base    BaseData
		Chars   []Character
		Search  string
		Message template.HTML
	}{
		Base: BaseData{
			Title:      "Research",
			StaticPath: "static/",
		},
		Chars:   chars,
		Search:  search,
		Message: message,
	}
	//fmt.Printf("log: data: %#v\n", data) // testing
	err := tmpl["search"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	searchHandler
//
// fetch and show all Article which title matches the search indicated in the query params.
//
// Query params: ?q=<search>
func createHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/create" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := struct {
		Base BaseData
	}{
		Base: BaseData{
			Title:      "Research - GoLanta",
			StaticPath: "static/",
		},
	}
	//fmt.Printf("log: data: %#v\n", data) // testing
	err := tmpl["create"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	createTreatmentHandler
//
// checks the form values sent by createHandler and calls createChar() to create the new Character.
//
// In case of invalid values, it redirects to createUserHandler with ?pass=error or ?user=error query params.
func createTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/create/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	strength, _ := strconv.Atoi(r.FormValue("strength"))
	agility, _ := strconv.Atoi(r.FormValue("agility"))
	stamina, _ := strconv.Atoi(r.FormValue("stamina"))
	vitality, _ := strconv.Atoi(r.FormValue("vitality"))
	initiative, _ := strconv.Atoi(r.FormValue("initiative"))
	intelligence, _ := strconv.Atoi(r.FormValue("intelligence"))
	knowledge, _ := strconv.Atoi(r.FormValue("knowledge"))
	fame, _ := strconv.Atoi(r.FormValue("fame"))
	var newChar = Character{
		Id:           getIdNewChar(),
		Name:         r.FormValue("name"),
		Nationality:  r.FormValue("nationality"),
		Strength:     strength,
		Agility:      agility,
		Stamina:      stamina,
		Vitality:     vitality,
		Initiative:   initiative,
		Intelligence: intelligence,
		Knowledge:    knowledge,
		Fame:         fame,
	}
	createChar(newChar)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

//	loginHandler
//
// takes the User info to send it to loginTreatmentHandler via Post Method.
//
// Optional query params: ?status=<error>	(error: "error" or "restricted")
func updateHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/update" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if !r.URL.Query().Has("char") {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("char"))
	if err != nil {
		log.Println("log: characterHandler() strconv.Atoi error!\n", err)
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	char, ok := selectChar(id)
	if !ok {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := struct {
		Base BaseData
		Char Character
	}{
		Base: BaseData{
			Title:      "Update Character - GoLanta",
			StaticPath: "static/",
		},
		Char: char,
	}
	//fmt.Printf("log: data: %#v\n", data) // testing
	err = tmpl["update"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	loginTreatmentHandler
//
// checks the form values sent by loginHandler to open the session and redirect to adminHandler
// or redirect to loginHandler with query params: ?status=error
func updateTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/update/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println("log: characterHandler() strconv.Atoi error!\n", err)
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	strength, _ := strconv.Atoi(r.FormValue("strength"))
	agility, _ := strconv.Atoi(r.FormValue("agility"))
	stamina, _ := strconv.Atoi(r.FormValue("stamina"))
	vitality, _ := strconv.Atoi(r.FormValue("vitality"))
	initiative, _ := strconv.Atoi(r.FormValue("initiative"))
	intelligence, _ := strconv.Atoi(r.FormValue("intelligence"))
	knowledge, _ := strconv.Atoi(r.FormValue("knowledge"))
	fame, _ := strconv.Atoi(r.FormValue("fame"))
	var updatedChar = Character{
		Id:           id,
		Name:         r.FormValue("name"),
		Nationality:  r.FormValue("nationality"),
		Strength:     strength,
		Agility:      agility,
		Stamina:      stamina,
		Vitality:     vitality,
		Initiative:   initiative,
		Intelligence: intelligence,
		Knowledge:    knowledge,
		Fame:         fame,
	}
	updateChar(updatedChar)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

//	logoutHandler
//
// close and clear the Session opened.
// It also clears the cache so that the Session can be closed.
func removeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/remove" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if !r.URL.Query().Has("char") {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("char"))
	if err != nil {
		log.Println("log: characterHandler() strconv.Atoi error!\n", err)
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	char, ok := selectChar(id)
	if !ok {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := struct {
		Base BaseData
		Char Character
	}{
		Base: BaseData{
			Title:      "Remove Character - GoLanta",
			StaticPath: "static/",
		},
		Char: char,
	}
	err = tmpl["remove"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	modifyUserTreatmentHandler
//
// checks the User new info and runs User.modifyUser with the new info.
//
// If new info is invalid, it redirects to modifyUserHandler with ?status=error query params.
func removeTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/remove/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("char"))
	if err != nil {
		log.Println("log: characterHandler() strconv.Atoi error!\n", err)
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	removeChar(id)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

//	errorHandler
//
// shows the custom error 404 page.
// It is called whenever the url is unknown or incorrect.
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	log.Printf("log: status: %#v\n", status) // testing
	if status == http.StatusNotFound {
		data := struct {
			Base BaseData
		}{
			Base: BaseData{
				Title:      "GoLanta - 404 Not Found",
				StaticPath: "static/",
			},
		}
		err := tmpl["error404"].ExecuteTemplate(w, "base", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
