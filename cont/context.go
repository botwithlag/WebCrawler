package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Post struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

var (
	posts   = make(map[int]Post)
	nextID  = 1
	postsMu sync.Mutex
)

type Logger struct {
	RespHandler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	l.RespHandler.ServeHTTP(w, r)
	log.Printf("The duration taken is %v to complete the request ", time.Since(start))

}

func NewMux(handlerTobeWrapper http.Handler) *Logger {
	return &Logger{handlerTobeWrapper}
}

func main() {
	mux := http.NewServeMux()
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/posts/", postHandler)
	fmt.Println("The server is up and running on Port 8000")
	wrappedMux := NewMux(mux)
	log.Fatal(http.ListenAndServe(":8000", wrappedMux))

}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	// used to get al requests
	switch r.Method {
	case "GET":
		handleGetPosts(w, r)
	default:
		http.Error(w, "Method hot allowed", http.StatusMethodNotAllowed)

	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
	if err != nil {
		http.Error(w, "Invalud post Id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		handleGetPost(w, r, id)
	case "POST":
		handlePostPost(w, r)
	case "DELETE":
		handleDeletePost(w, r, id)
	}

}

func handleGetPost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()
	p, ok := posts[id]
	if !ok {
		http.Error(w, "Unable to find id", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func handlePostPost(w http.ResponseWriter, r *http.Request) {

	var p Post
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unexpected Error", int(http.StatusBadGateway))
	}
	json.Unmarshal(body, &p)
	defer r.Body.Close()

	postsMu.Lock()
	defer postsMu.Unlock()

	p.ID = nextID
	nextID++
	posts[p.ID] = p

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)

}

func handleDeletePost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()
	_, ok := posts[id]
	if !ok {
		http.Error(w, "UnabletofindRequest", http.StatusBadRequest)
	}
	delete(posts, id)
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(posts)
	w.WriteHeader(http.StatusAccepted)

}

func handleGetPosts(w http.ResponseWriter, r *http.Request) {
	postsMu.Lock()
	defer postsMu.Unlock()

	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(posts)
	w.WriteHeader(http.StatusAccepted)

}
