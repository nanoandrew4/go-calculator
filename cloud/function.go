package cloud

import (
	"calculator/pkg/calculator"
	"encoding/json"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"io"
	"net/http"
)

var (
	errEmptyBody      = mustMarshal(CalculateError{Error: "Body cannot be empty"})
	errEmptyOperation = mustMarshal(CalculateError{Error: "Body must contain an operation"})
	errGeneric        = mustMarshal(CalculateError{Error: "Unexpected error"})
)

func init() {
	functions.HTTP("calculate", calculate)
}

func mustMarshal(entityToMarshal any) string {
	marshalledEntity, err := json.Marshal(entityToMarshal)
	if err != nil {
		panic(err)
	}
	return string(marshalledEntity)
}

func calculate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, errGeneric, http.StatusInternalServerError)
		return
	} else if len(body) == 0 {
		http.Error(w, errEmptyBody, http.StatusBadRequest)
		return
	}

	var operationRequest CalculateRequest
	err = json.Unmarshal(body, &operationRequest)
	if err != nil {
		http.Error(w, errGeneric, http.StatusInternalServerError)
		return
	} else if len(operationRequest.Operation) == 0 {
		http.Error(w, errEmptyOperation, http.StatusBadRequest)
		return
	}

	response, err := calculator.Calculate(operationRequest.Operation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marshalledResponse, err := json.Marshal(CalculateResponse{Result: response})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(marshalledResponse)
}
