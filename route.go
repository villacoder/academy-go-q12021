package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func getInfoCSV(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	news := csvToMap()
	result, err := json.Marshal(news)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error generating JSON string"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func getInfoCSVByID(resp http.ResponseWriter, req *http.Request) {
	var result []byte
	var err error

	ids, ok := req.URL.Query()["id"]

	if !ok || len(ids[0]) < 1 {
		log.Println("Url Param 'ID' is missing")
		return
	}

	resp.Header().Set("Content-type", "application/json")
	news := csvToMap()
	idStr := ids[0]
	id, _ := strconv.Atoi(idStr)

	for i := range news {
		if news[i].Id == id {
			result, err = json.Marshal(news[i])
		}
	}

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error generating JSON string"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}
