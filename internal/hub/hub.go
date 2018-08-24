package hub

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dtoebe/collab-md/internal/middleware"
	"github.com/dtoebe/collab-md/internal/utils"

	"encoding/json"

	"github.com/dtoebe/collab-md/internal/document"
)

type Hub struct {
	Instances map[string]*instance
	HostURL   string
}

type instance struct {
	ID       string
	Document *document.Document
}

func NewHub(proto, hostPort string) *Hub {
	if hostPort[0] == ':' {
		hostPort = "localhost" + hostPort
	}
	return &Hub{
		Instances: make(map[string]*instance),
		HostURL:   fmt.Sprintf("%s://%s", proto, hostPort),
	}
}

func (h Hub) genUniqueID(l int) string {
	var id string
	for {
		id = utils.GenRandStr(l)
		if _, ok := h.Instances[id]; !ok {
			break
		}
	}
	return id
}

func (h *Hub) NewInstance(t, b, u string) *instance {
	id := h.genUniqueID(6)
	return &instance{
		ID:       id,
		Document: document.NewDocument(t, b, u, id),
	}
}

func (h *Hub) NewInstanceMiddleware() middleware.Middleware {
	log.Println("Document Create requested")
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := r.FormValue("title")
			b := r.FormValue("body")
			u := "user-" + utils.GenRandStr(5)

			if t == "" || b == "" {
				log.Println("Missing title or body: title=", t, "body=", b)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			inst := h.NewInstance(t, b, u)
			resp := struct {
				Document document.Document `json:"document,omitempty"`
				ID       string            `json:"id,omitempty"`
				Metadata json.RawMessage   `json:"metadata,omitempty"`
				URL      string            `json:"url,omitempty"`
				User     string            `json:"user,omitempty"`
			}{
				Document: *inst.Document,
				ID:       inst.ID,
				URL:      fmt.Sprintf("%s/documents/edit/%s", h.HostURL, inst.ID),
				User:     u,
			}

			bb, err := json.Marshal(&resp)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(bb)
			handler.ServeHTTP(w, r)
		})
	}
}

func (h *Hub) EditInstanceMiddleware() middleware.Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			paths := strings.Split(r.URL.Path, "/")
			id := paths[len(paths)-1]

			if _, ok := h.Instances[id]; !ok {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			if r.Header.Get("Connetion-Type") != "direct" {
				// TODO: Serve Http site
			}

			user := r.FormValue("user")
			if user == "" {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

		})
	}
}
