package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type ApiResponse struct {
    Results []interface{} `json:"results"`
    NextURL string        `json:"next_url"`
}

func fetchAllPages(url string, headers map[string]string) ([]interface{}, error) {
    var allResults []interface{}
    for url != "" {
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            return nil, err
        }
        for key, value := range headers {
            req.Header.Set(key, value)
        }
        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            return nil, err
        }
        body, err := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        if err != nil {
            return nil, err
        }
        var apiResp ApiResponse
        if err := json.Unmarshal(body, &apiResp); err != nil {
            return nil, err
        }
        allResults = append(allResults, apiResp.Results...)
        url = apiResp.NextURL
    }
    return allResults, nil
}

func main() {
    results, err := fetchAllPages("https://api.example.com/items", nil)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Total results:", len(results))
}

