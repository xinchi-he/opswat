package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
)

//DataID struct
type DataID struct {
	DataID string `json:"data_id"`
}

//Engine struct
type Engine struct {
	Time   string `json:"def_time"`
	Threat string `json:"threat_found"`
	Scan   int    `json:"scan_result_i"`
}

//your API KEY goes here!!!
const APIKEY = "YOUR API KEY"

func getHash() string {
	//read file from command line argument
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//calculate sha256 for the given file
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func queryByHash(hash string) bool {
	//start a new HTTP client for GET request
	url := "https://api.metadefender.com/v4/hash/"

	req, err := http.NewRequest("GET", url+hash, nil)
	if err != nil {
		log.Fatal("GET ERROR ", err)
	}

	//set API header
	req.Header.Set("apikey", APIKEY)

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("RESPONSE ERROR ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("READ RESPONSE ERROR ", err)
	}

	results := fmt.Sprintf("%s\n", body)

	//retrive body content, if no error is found, the parser will start to parse the result
	if strings.Contains(results, "error") {
		fmt.Println("no result found, will upload the file")
	} else {
		fmt.Println("record found!")
		parseReport(body)
		return false
	}
	return true
}

func uploadFile() string {
	//start a new HTTP client for POST request
	url := "https://api.metadefender.com/v4/file"

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	req, err := http.NewRequest("POST", url, f)

	if err != nil {
		log.Fatal("POST ERROR ", err)
	}

	//set headers
	req.Header.Set("apikey", APIKEY)
	req.Header.Set("content-type", "application/octet-stream")
	req.Header.Set("filename", os.Args[1])

	//client := &http.Client{Timeout: time.Second * 10}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("SEND ERROR, ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	//parse the HTTP response and retrieve data_id
	var ID DataID
	err = json.Unmarshal(body, &ID)
	if err != nil {
		log.Fatal("ERROR ON UNMARSHAL JSON ", err)
	}

	return ID.DataID
}

func getReport(id string) {
	//start a new HTTP client for GET request
	url := "https://api.metadefender.com/v4/file/"

	req, err := http.NewRequest("GET", url+id, nil)
	if err != nil {
		log.Fatal("GET ERROR ", err)
	}

	req.Header.Set("apikey", APIKEY)

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("RESPONSE ERROR ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("READ RESPONSE ERROR ", err)
	}

	results := fmt.Sprintf("%s\n", body)

	//if no error is found from the response body, generate the report
	if strings.Contains(results, "error") {
		fmt.Println("no result found")
	} else {
		fmt.Println("record found, will generate report...")
		parseReport(body)
	}
}

func parseReport(report []byte) {
	//utilize a 3rd party JSON parser library to parse response to generate report as requested
	filename, _, _, _ := jsonparser.Get(report, "file_info", "display_name")
	fmt.Printf("filename: %s\n", string(filename))

	overallStatus, _, _, _ := jsonparser.Get(report, "scan_results", "scan_all_result_a")
	fmt.Printf("overall_status: %s\n\n", string(overallStatus))

	jsonparser.ObjectEach(report, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		fmt.Printf("engine: %s\n", string(key))

		var entries Engine
		err := json.Unmarshal(value, &entries)
		if err != nil {
			log.Fatal("ERROR UNMARSHAL JSON ", err)
		}
		if entries.Threat != "" {
			fmt.Printf("threat_found: %s\n", entries.Threat)

		} else {
			fmt.Printf("threat_found: CLEAN \n")

		}
		fmt.Printf("scan_result: %s\n", strconv.Itoa(entries.Scan))
		fmt.Printf("def_time: %s\n\n", entries.Time)

		return nil
	}, "scan_results", "scan_details")

}

func main() {

	//step (1) get the hash of the input file
	h := getHash()

	//step (2) query the API with hash
	if queryByHash(h) {
		// step (4), no hash found, will upload file
		id := uploadFile()
		// step (5) & (6) pull report and generate report
		getReport(id)

	}

	//in queryByHash(), if false is returned, execute step (3) (6) and directly generate report.

}
