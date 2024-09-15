//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/optiop/next_century_sdk/core/client"
	"github.com/optiop/next_century_sdk/core/schema"
)

// simple example
func main() {

	// sample user
	if err := os.Setenv("NCM_EMAIL", "test@test.com"); err != nil {
		log.Fatal(err)
	}

	if err := os.Setenv("NCM_PASS", "test"); err != nil {
		log.Fatal(err)
	}

	client := client.New("https://api.test.com/")

	data, err := client.GetDailyReads("test", schema.TimeRequest{Date: time.Now()})
	if err != nil {
		log.Fatal(err)
	}

	// save data json in one file
	file, err := os.Create("sample_data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := file.Write(jsonData); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully request & save data!")
}
