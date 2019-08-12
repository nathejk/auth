package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"local/token"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	activePrivateKey  string
	activeFingerprint string
	kid               string
	publicKeys        map[string]string
)

func main() {
	fmt.Println("Starting keygen")

	var pub string
	publicKeys = make(map[string]string, 0)
	activePrivateKey, pub, kid = token.GenerateRSAKeys(4096)
	publicKeys[kid] = pub

	http.HandleFunc("/basicauth", basicauthHandler)
	http.HandleFunc("/pubkey", pubkeyHandler)
	http.Handle("/", http.FileServer(http.Dir("/www")))

	fmt.Println("Running webserver")
	log.Fatal(http.ListenAndServe(":80", nil))
}

type Credentials struct {
	Username string
	Password string
}

func basicauthHandle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data Credentials
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://basicauth", nil)
	req.SetBasicAuth(data.Username, data.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	log.Println(s)
}
func basicauthHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	//v := req.Form
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	log.Println("Testing " + username + ":" + password)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://basicauth", nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	referer := r.Header.Get("Referer")
	if resp.StatusCode == http.StatusOK {
		u, err := url.Parse(referer)
		if err != nil {
			panic(err)
		}
		cookieDsn, _ := url.Parse(os.Getenv("COOKIE_DSN"))
		loginCookie := &http.Cookie{
			Name:     cookieDsn.User.Username(),
			Value:    token.Sign(map[string]string{"sub": username}, activePrivateKey, kid),
			HttpOnly: true,
			Domain:   cookieDsn.Host,
			Path:     "/" + cookieDsn.Path,
		}
		http.SetCookie(w, loginCookie)
		log.Printf("Setting cookie %#v\n", loginCookie)
		query, _ := url.ParseQuery(u.RawQuery)
		if _, found := query["goto"]; !found {
			http.Redirect(w, r, referer, http.StatusFound)
			return
		}
		gotoUrl := query["goto"][0]
		if _, err = url.Parse(gotoUrl); err != nil {
			http.Redirect(w, r, referer, http.StatusFound)
			return
		}
		http.Redirect(w, r, gotoUrl, 302)

	} else {
		http.Redirect(w, r, referer, http.StatusFound)
		return
	}
}

func pubkeyHandler(w http.ResponseWriter, r *http.Request) {
	kid := strings.Replace(r.URL.Path, "/", "", 1)

	fmt.Fprintf(w, publicKeys[kid])
}
