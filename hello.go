package guestbook

import (
	"net/http"

	"io"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	hc := urlfetch.Client(c)
	res, err := hc.Get("https://gcpug.jp" + r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	w.Header().Add("Content-Type", res.Header.Get("Content-Type"))
	io.Copy(w, res.Body)
}
