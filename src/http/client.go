package http

import (
	"github.com/cortiz/fakeTrafic/src/yaml"
	"io"
	"log"
	"net/http"
)

func DoRequest(request *yaml.Request) {
	client := &http.Client{}
	req, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	for key, value := range request.Headers {
		req.Header.Add(key,value)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	println(string(body))
	defer resp.Body.Close()
}
