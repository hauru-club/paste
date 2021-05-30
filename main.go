package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/bmizerany/pat"
)

const keyLen = 10

//go:embed readme.txt
var readme []byte

func run() error {
	store := newDataStore()

	port := os.Getenv("PASTE_PORT")
	host := os.Getenv("PASTE_HOST")
	target := os.Getenv("PASTE_TARGET")

	if target == "" {
		target = host + port
	}

	m := pat.New()
	m.Get("/:key", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get(":key")

		res := store.get(key)
		if res == nil {
			http.NotFound(w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}))
	m.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(readme)
	}))
	m.Post("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		key := genKey(keyLen)

		delKey := genKey(keyLen)

		// store file
		store.set(key, b)

		// store pointer to delete posted file
		store.set(delKey, []byte(key))

		buff := bytes.NewBuffer([]byte{})

		if err := newFileTmpl.Execute(buff, newFileResponse{
			Host:      target,
			Key:       key,
			DeleteKey: delKey,
		}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(buff.Bytes())
	}))
	m.Del("/:key", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get(":key")

		fileKey := store.get(key)
		if fileKey == nil {
			http.NotFound(w, r)
			return
		}

		store.delete(string(fileKey))
		store.delete(key)

		w.WriteHeader(http.StatusOK)
	}))

	http.Handle("/", m)
	log.Printf("listening at %s:%s...", host, port)
	return http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type newFileResponse struct {
	Host      string
	Key       string
	DeleteKey string
}

const newFileTmplString = `
To acces your file execute below command:

  $ curl {{ .Host }}/{{ .Key }}

If you wish to delete file, use this:

  $ curl -XDELETE {{ .Host }}/{{ .DeleteKey }}

`

var newFileTmpl = template.Must(template.New("newFile").Parse(newFileTmplString))
