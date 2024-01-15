package GoLanta

import "net/http"

//	routes
//
// initialises all the routes.
func routes() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/character", characterHandler) // use query params: ?char=<char-id>
	http.HandleFunc("/search", searchHandler)       // use query params: ?q=<search>
	http.HandleFunc("/create", createHandler)       // use query params: ?error=<error>
	http.HandleFunc("/create/treatment", createTreatmentHandler)
	http.HandleFunc("/update", updateHandler) // use query params: ?char=<char-id>
	http.HandleFunc("/update/treatment", updateTreatmentHandler)
	http.HandleFunc("/remove", removeHandler)                    // use query params: ?char=<char-id>
	http.HandleFunc("/remove/treatment", removeTreatmentHandler) // use query params: ?char=<char-id>
}
