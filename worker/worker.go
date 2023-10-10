package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hjunior29/go-tasker/models"
	"github.com/hjunior29/go-tasker/store"
)

// Initializes the worker to process tasks.
// It continuously dequeues tasks and processes them.
func StartWorker() {
	for {
		TasksRequest, err := store.DequeueTask()
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		processTask(TasksRequest)
	}
}

// processTask handles the processing of a dequeued task.
// It sends an HTTP request based on the task details and logs the process.
func processTask(Tasks *models.Tasks) {
	var TasksRequest map[string]interface{}
	var status, logMessage string
	logType := "Send"

	err := json.Unmarshal([]byte(Tasks.Payload), &TasksRequest)
	if err != nil {
		status := "Sent Failed"
		logMessage = fmt.Sprintf("Error unmarshalling TasksRequest: %v", err)
		store.LogTask(store.DB, Tasks.TaskID, Tasks.Payload, status, logMessage, logType, Tasks.Method, Tasks.URL, time.Now())
		return
	}

	if Tasks.URL == "" {
		status = "Sent Failed"
		logMessage = "URL is empty"
		store.LogTask(store.DB, Tasks.TaskID, Tasks.Payload, status, logMessage, logType, Tasks.Method, Tasks.URL, time.Now())
		return
	}

	if Tasks.Method == "" {
		status = "Sent Failed"
		logMessage = "Method is empty"
		store.LogTask(store.DB, Tasks.TaskID, Tasks.Payload, status, logMessage, logType, Tasks.Method, Tasks.URL, time.Now())
		return
	}

	payloadBytes, err := json.Marshal(TasksRequest)
	if err != nil {
		status = "Sent Failed"
		logMessage = fmt.Sprintf("Error marshalling payload: %v", err)
		store.LogTask(store.DB, Tasks.TaskID, Tasks.Payload, status, logMessage, logType, Tasks.Method, Tasks.URL, time.Now())
		return
	}

	req, err := http.NewRequest(Tasks.Method, Tasks.URL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		status = "Sent Failed"
		logMessage = fmt.Sprintf("Error creating HTTP request: %v", err)
		store.LogTask(store.DB, Tasks.TaskID, Tasks.Payload, status, logMessage, logType, Tasks.Method, Tasks.URL, time.Now())
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		status = "Sent Failed"
		logMessage = fmt.Sprintf("Error sending HTTP request: %v", err)
		store.LogTask(store.DB, Tasks.TaskID, Tasks.Payload, status, logMessage, logType, Tasks.Method, Tasks.URL, time.Now())
		return
	}
	defer resp.Body.Close()

	if resp.Status == "200 OK" {
		status = "Sent Success"
		logMessage = fmt.Sprintf("HTTP Response Status: %v", resp.Status)
	} else {
		status = "Sent Failed"
		logMessage = fmt.Sprintf("HTTP Response Status: %v", resp.Status)
	}

	store.LogTask(store.DB, Tasks.TaskID, Tasks.Payload, status, logMessage, logType, Tasks.Method, Tasks.URL, time.Now())
}
