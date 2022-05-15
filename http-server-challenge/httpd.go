package main

/* Key/Value database

$ curl -d'hello' http://localhost:8080/k1
$ curl http://localhost:8080/k1
hello
$ curl -i http://localhost:8080/k2
404 not found

Limit value size to 1k
*/
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Server struct {
	db DB
}

// POST /key Store request body as value
// GET /<key> Send back value, or 404 if key not found
func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {

	params := strings.Split(r.URL.Path, "/")
	key := params[len(params)-1]

	if strings.TrimSpace(key) == "" {
		log.Println("error: Invalid Request")
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	resp := &DbResponse{}

	if r.Method == http.MethodGet {
		// GET /<key> Send back value, or 404 if key not found
		result := s.db.Get(key)

		if result == nil {
			log.Printf("error: Value with key %s not found", key)

			//http.Error(w, "Value not found", http.StatusNotFound)
			resp.Error = "Value not found"
		} else {
			resp.Result = string(result)
		}

	} else if r.Method == http.MethodPost {
		// POST /key Store request body as value
		body, err1 := io.ReadAll(r.Body)
		if err1 != nil {
			resp.Error = fmt.Sprintf("Error : %v", err1)
		} else {
			s.db.Set(key, body)
			resp.Result = fmt.Sprintf("Data Saved %v => %v", key, string(body))
		}

	} else {
		log.Println("error: Invalid Method")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//Step 3: Encode result
	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)

	if err := enc.Encode(resp); err != nil {
		//Cant return error to client here
		log.Printf("can't encode %v - %s", resp, err)
	}

}

//DbResponse is a response to DbRequest
type DbResponse struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

func main() {
	db := DB{}
	var s = Server{db: db}

	http.HandleFunc("/", s.Handler)

	addr := ":8080"
	log.Printf("server readyon %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

// type DB struct {
// 	m sync.Map
// }

// func (db *DB) Get(key string) []byte {
// 	val, ok := db.m.Load(key)
// 	if !ok {
// 		return nil
// 	}

// 	return val.([]byte)
// }

// func (db *DB) Set(key string, value []byte) {
// 	db.m.Store(key, value)
// }
