package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

type RequestBody struct {
	LogEntrySizeBytes     int `json:"logEntrySizeBytes"`
	NumOfEntries          int `json:"numOfEntries"`
	SleepMsBetweenEntries int `json:"sleepMsBetweenEntries"`
}

func GenerateLogsOfSize(resWriter http.ResponseWriter, request *http.Request) {
	var reqBody RequestBody
	if err := json.NewDecoder(request.Body).Decode(&reqBody); err != nil {
		http.Error(resWriter, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if reqBody.LogEntrySizeBytes <= 10 {
		http.Error(resWriter, "Invalid logEntryMessageBytes value in request body. Value must be greater than 10 because a single log entry of '{\"msg\":\"\"}' uses 10 bytes", http.StatusBadRequest)
		return
	}

	if reqBody.NumOfEntries <= 0 {
		http.Error(resWriter, "Invalid numOfEntries value in request body. Value must be greater than 0", http.StatusBadRequest)
		return
	}

	if reqBody.SleepMsBetweenEntries <= 0 {
		http.Error(resWriter, "Invalid sleepMsBetweenEntries value in request body. Value must be greater than 0", http.StatusBadRequest)
		return
	}

	resWriter.WriteHeader(http.StatusAccepted)
	resWriter.Write([]byte("Log generation started"))

	go func() {
		for i := 0; i < reqBody.NumOfEntries; i++ {
			logEntryMsg := make([]byte, reqBody.LogEntrySizeBytes)
			for i := range logEntryMsg {
				logEntryMsg[i] = letters[rand.Intn(len(letters))]
			}

			jsonOutput := map[string]string{"msg": string(logEntryMsg)}
			jsonBytes, err := json.Marshal(jsonOutput)
			if err != nil {
				os.Stderr.WriteString("Error marshaling JSON: " + err.Error() + "\n")
				continue
			}
			os.Stdout.Write(jsonBytes)
			os.Stdout.Write([]byte("\n"))

			time.Sleep(time.Duration(reqBody.SleepMsBetweenEntries) * time.Millisecond)
		}
	}()

}
