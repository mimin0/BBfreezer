package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	config "github.com/BBfreezer/helper/config"
)

func main() {
	// reading YAML config and validation congig path
	cfg, err := config.NewConfig("./config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	url := fmt.Sprintf("%s/rest/api/1.0/projects/%s/repos/%s/pull-requests", cfg.Host, cfg.Project, cfg.Repo)
	jsonData := map[string]string{"order": "OLDEST"}
	jsonValue, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(os.Getenv("BB_USER"), os.Getenv("BB_PASS"))
	fmt.Println("++", req)

	/* declared map of string with empty interface which will hold the value of the parsed json. */
	var result map[string]interface{}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	/* Unmarshal the json string string by converting it to byte into map */
	json.Unmarshal([]byte(body), &result)

	for _, item := range result["values"].([]interface{}) {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Printf("%v\n", item.(map[string]interface{})["title"])
	}
	// fmt.Println("response Body:", string(body))
}
