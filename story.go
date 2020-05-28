package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"strings"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func ParseJsonFileToStoryType(r io.Reader) (Story, error){
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	layout := path.Join("", "story_temp.html")
	tmpl, err := template.ParseFiles(layout)

	if err != nil {
		log.Fatal(err)
		return
	}

	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]

	if story, ok := h.s[path]; ok {
		err = tmpl.Execute(w, story)
		if err != nil {
			log.Fatal(err)
			return
		}
	}else {
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

