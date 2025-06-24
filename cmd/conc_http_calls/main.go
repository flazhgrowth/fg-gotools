package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/flazhgrowth/fg-gotools/printer"
	"github.com/go-resty/resty/v2"
)

func main() {
	numberOfWorker := 2
	numberOfIterationsPerWorker := 3
	cl := resty.New()

	mapBearerToken := map[int]string{
		1: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDY5MzI5MTcsImlhdCI6MTc0Njg0NjUxNywiaWQiOiIwMUpUSzcxWEtNVlRSUE02WFdUSjlTRjU5WCIsInVzZXJuYW1lIjoiZmUuZXZhbG9zIiwiZW1haWwiOiJmZS5ldmFsb3NAZ21haWwuY29tIn0.0vIWnAOwnB4t10qozWmkyTs2a-ByZNv9I0cYFnkf_iI",    // fe.evalos
		2: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDY5MzI5NDUsImlhdCI6MTc0Njg0NjU0NSwiaWQiOiIwMUpUSzBSVFBIU0ZOOVI0NzRDOTI0MDQ2OSIsInVzZXJuYW1lIjoiYmFhc2thcmFhYSIsImVtYWlsIjoiYmFhc2thcmFhYUBnbWFpbC5jb20ifQ.FpJpRoti88Y-NV6wuXexm45BlCq47yku_zq3lY9qpMw", // baaskaraaa
	}

	wg := sync.WaitGroup{}
	for i := range numberOfWorker {
		wg.Add(1)
		go func(workerNumber int) {
			defer wg.Done()
			bearerToken := mapBearerToken[workerNumber]
			fmt.Println("Running worker number", workerNumber)
			for iter := range numberOfIterationsPerWorker {
				resp, err := cl.
					R().
					SetAuthToken(bearerToken).
					Get("http://localhost:11011/api/v1/me")
				if err != nil {
					panic(err)
				}

				decodedBody := map[string]any{}
				json.Unmarshal(resp.Body(), &decodedBody)
				decodedBody["status_code"] = resp.StatusCode()
				decodedBody["iter"] = iter + 1
				decodedBody["worker"] = workerNumber
				printer.PrintInJSONFormat(decodedBody)
			}
		}(i + 1)
	}

	wg.Wait()
}
