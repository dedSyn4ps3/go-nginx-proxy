package main

import (
	"crypto/tls"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func runScript(args ...string) {
	comm := exec.Command("/usr/bin/python3", args...)
	_ = comm.Run()
}

func redirection(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.awesomeinc.com/index.html", http.StatusMovedPermanently)
}

func newSignup(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	loc, _ := time.LoadLocation("America/New_York")
	myTime := t.In(loc)

	log.WithFields(log.Fields{
		"Time":       myTime.String(),
		"Email":      r.FormValue("email"),
		"User-Agent": r.UserAgent(),
	}).Info("New Signup")

	email := r.FormValue("email")
	go runScript("signup.py", email, "", "")
	http.Redirect(w, r, "https://www.awesomeinc.com/index.html", http.StatusFound)
}

func newContact(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	loc, _ := time.LoadLocation("America/New_York")
	myTime := t.In(loc)

	log.WithFields(log.Fields{
		"Time":  myTime.String(),
		"Name":  r.FormValue("name"),
		"Email": r.FormValue("email"),
		"Phone": r.FormValue("phone"),
	}).Info("New Contact")

	name, email, phone := r.FormValue("name"), r.FormValue("email"), r.FormValue("phone")
	go runScript("contact.py", name, email, phone)
	http.Redirect(w, r, "https://www.awesomeinc.com/index.html", http.StatusFound)
}

func init() {

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("contact-log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println("* Server Started *")
}

func main() {

	cert, err := tls.LoadX509KeyPair("./fullchain.pem", "./privkey.pem")

	if err != nil {
		log.Fatal("[!] Error loading Server Certificates [!]")
	}

	r := mux.NewRouter()

	r.HandleFunc("/new_signup", newSignup).Methods("POST")
	r.HandleFunc("/contact", newContact).Methods("POST")
	r.HandleFunc("/", redirection).Methods("GET")

	tlsConf := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	serverProxy := &http.Server{
		Addr:      ":8081",
		Handler:   r,
		TLSConfig: tlsConf,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(serverProxy.ListenAndServeTLS("", ""))
}
