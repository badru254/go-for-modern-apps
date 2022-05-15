//HTTP server in go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//healthHandler returns a server health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ok")
}

//MathRequest is a request of math operation
type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

//MathResponse is a response to MathRequest
type MathResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

//mathHandler returns result of calculation
func mathHandler(w http.ResponseWriter, r *http.Request) {

	//Step 1: Decode & Validate
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	req := &MathRequest{}

	if err := dec.Decode(req); err != nil {
		log.Printf("error: bad JSON: %s", err)
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	if !strings.Contains("+-*/", req.Op) {
		log.Printf("error: bad operator: %q", req.Op)
		http.Error(w, "unknown operator", http.StatusBadRequest)
		return
	}

	//Step 2: work
	resp := &MathResponse{}
	switch req.Op {
	case "+":
		resp.Result = req.Left + req.Right
	case "-":
		resp.Result = req.Left - req.Right
	case "*":
		resp.Result = req.Left * req.Right
	case "/":
		if req.Right == 0.0 {
			resp.Error = "division by 0"
		} else {
			resp.Result = req.Left / req.Right
		}
	default:
		resp.Error = fmt.Sprintf("unknown operation: %s", req.Op)

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

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/math", mathHandler)

	addr := ":8080"
	log.Printf("server readyon %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

//curl -d "{ \"op\": \"-\",\"left\": 2.4,\"right\": 4.5 }" -H "Content-Type: application/json" -X POST  http://localhost:8080/math
