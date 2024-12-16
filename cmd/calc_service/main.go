package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
    "calculator-service/pkg/calculator" 
)

type request struct {
	Expression string `json:"expression"`
}

type response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}
    defer r.Body.Close()


	var req request
	if err := json.Unmarshal(reqBody, &req); err != nil {
		log.Printf("Error unmarshaling JSON request: %v", err)
		sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON request")
		return
	}

    result, err := calculator.Calculate(req.Expression) // Use your calculator logic

	if err != nil {
		log.Printf("Error evaluating expression: %v", err)
        var statusCode int
        var message string
        switch err.(type){
            case *calculator.InvalidExpressionError:
                statusCode = http.StatusUnprocessableEntity
                message = "Expression is not valid"
            default:
                statusCode = http.StatusInternalServerError
                message = "Internal server error"
        }
        sendErrorResponse(w, statusCode, message)

        return
	}


	resp := response{Result: result}
	sendJsonResponse(w, http.StatusOK, resp)
}

func sendJsonResponse(w http.ResponseWriter, status int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
    if err := json.NewEncoder(w).Encode(payload); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}

func sendErrorResponse(w http.ResponseWriter, status int, message string) {
    resp := response{Error: message}
    sendJsonResponse(w, status, resp)
}

func main() {
	http.HandleFunc("/api/v1/calculate", calculateHandler)

	fmt.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Error starting server: ", err)
	}
}
