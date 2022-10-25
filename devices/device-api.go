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

type Sensor struct {
	Id   uint32
	Name string
}

func (s *Sensor) sendAlert(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	loc, _ := time.LoadLocation("America/New_York")
	myTime := t.In(loc)

	switch s.Name {
	case "test":
		log.Println(myTime.Hour(), myTime.Minute(), " [+] TEST SUCCESSFUL [+]")
		w.Write([]byte("[+] TEST SUCCESSFUL [+]\n"))

	case "back":
		go runScript("alert.py", "Back Door", "", "")
		w.Write([]byte("[!] Alert Email Sent -> BACK\n"))
		log.Println(myTime.Hour(), myTime.Minute(), " [!] Alert Email Sent -> BACK")

	default:
		go runScript("alert.py", "Front Door", "", "")
		log.Println(myTime.Hour(), myTime.Minute(), " [!] Alert Email Sent -> FRONT")
		w.Write([]byte("[!] Alert Email Sent -> FRONT\n"))
	}
}

func runScript(args ...string) {
	comm := exec.Command("/usr/bin/python3", args...)
	_ = comm.Run()
}

func redirection(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.awesomeinc.com/index.html", http.StatusMovedPermanently)
}

func init() {

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("device-log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println("* Device Server Started *")
}

func main() {
	front := Sensor{1, "front"}
	back := Sensor{2, "back"}
	test := Sensor{3, "test"}

	cert, err := tls.LoadX509KeyPair("./fullchain.pem", "./privkey.pem")

	if err != nil {
		log.Fatal("[!] Error loading Server Certificates [!]")
	}

	r := mux.NewRouter()

	r.HandleFunc("/device_alert/front", front.sendAlert).Methods("GET")
	r.HandleFunc("/device_alert/back", back.sendAlert).Methods("GET")
	r.HandleFunc("/test", test.sendAlert).Methods("GET")
	r.HandleFunc("/", redirection).Methods("GET")

	tlsConf := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	serverProxy := &http.Server{
		Addr:      ":8080",
		Handler:   r,
		TLSConfig: tlsConf,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(serverProxy.ListenAndServeTLS("", ""))
}
