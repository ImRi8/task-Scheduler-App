package main

import (
	"bytes"
	"gofr.dev/pkg/gofr/request"
	"net/http"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	go main()
	time.Sleep(3 * time.Second)

	test := []struct {
		desc       string
		method     string
		endpoint   string
		statusCode int
		body       []byte
	}{
		{"get Health", http.MethodGet, "task/health", http.StatusOK, nil},
		{"get TaskById", http.MethodGet, "task/getTaskById?id=1", http.StatusOK, nil},
		{"post DeleteById", http.MethodPost, "task/deleteTaskById?id=15", http.StatusCreated, nil},
		{"post CreateTask", http.MethodPost, "task/createTask", http.StatusCreated, []byte(`{
		
			"title": "newtask1",
			"description": "sleep1 at",
			"priority": 2,
			"dueDate": "2023-12-19T15:04:08Z"
			
		}`)},
		{"post Updatetask", http.MethodPost, "task/updateTask", http.StatusCreated, []byte(`{
		
			"id" : 1,
			"description" : "new-update"
		
		}`),
		}}

	for i, tc := range test {
		req, _ := request.NewMock(tc.method, "http://localhost:8000/"+tc.endpoint, bytes.NewBuffer(tc.body))

		c := http.Client{}

		resp, err := c.Do(req)
		if err != nil {
			t.Errorf("TEST[%v] Failed.\tHTTP request encountered Err: %v\n%s", i, err, tc.desc)
			continue
		}

		if resp.StatusCode != tc.statusCode {
			t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.statusCode, resp.StatusCode, tc.desc)
		}

		_ = resp.Body.Close()

	}

}
