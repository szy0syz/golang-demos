package server

import (
	"bookstore/server/middleware"
	"bookstore/store"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type BookStoreServer struct {
	s    store.Store
	serv *http.Server
}

func NewBookStoreServer(addr string, s store.Store) *BookStoreServer {
	serv := &BookStoreServer{
		s: s,
		serv: &http.Server{
			Addr: addr,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/book", serv.createBookHandler).Methods("POST")

	// 加载中间件，感觉没有nodejs优雅
	serv.serv.Handler = middleware.Logging(middleware.Validating(router))

	return serv
}

func (bs *BookStoreServer) createBookHandler(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := bs.s.Create(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(201)
	response(w, book)
}

func (bs *BookStoreServer) getBookHandler(w http.ResponseWriter, req *http.Request) {
	id, ok := mux.Vars(req)["id"]
	if !ok {
		http.Error(w, "no id found in request", http.StatusBadRequest)
		return
	}
	book, err := bs.s.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response(w, book)
}

func response(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
