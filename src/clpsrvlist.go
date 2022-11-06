package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var jsonServers [][]map[string]interface{}
	var jsonResult map[string]interface{}
	var (
		d = flag.Int("d", 0, "Debug flag")
	)

	flag.Parse()

	/* Read servers.json */
	servers, err := ioutil.ReadFile("./servers.json")
	if err != nil {
		fmt.Println("ReadFile() failed: ", err)
		os.Exit(1)
	}
	json.Unmarshal(servers, &jsonServers)
	if *d == 1 {
		fmt.Println(string(servers))
		fmt.Println(jsonServers)
	}

	/*Get server status */
	for i := 0; i < len(jsonServers); i++ {
		for j := 0; j < len(jsonServers[i]); j++ {
			ipaddress := jsonServers[i][j]["ipaddress"]
			port := jsonServers[i][j]["port"]
			user := jsonServers[i][j]["user"]
			password := jsonServers[i][j]["password"]
			if *d == 1 {
				fmt.Println(ipaddress, port, user, password)
			}

			/* Get server status */
			url := fmt.Sprintf("http://%s:%s@%s:%s/api/v1/servers", user, password, ipaddress, port)
			resp, err := http.Get(url)
			if err != nil {
				if *d == 1 {
					fmt.Println("http.Get() failed: ", err)
				}
				continue
			}
			defer resp.Body.Close()
			byteArray, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				if *d == 1 {
					fmt.Println("ioutil.ReadAll() failed: ", err)
				}
				continue
			}
			if *d == 1 {
				fmt.Println(string(byteArray))
			}
			json.Unmarshal([]byte(byteArray), &jsonResult)
			array := jsonResult["servers"].([]interface{})
			for k := 0; k < len(array); k++ {
				fmt.Println((array[k].(map[string]interface{}))["name"], ",", (array[k].(map[string]interface{}))["status"])
			}

			/* Get group status */
			url = fmt.Sprintf("http://%s:%s@%s:%s/api/v1/groups", user, password, ipaddress, port)
			resp, err = http.Get(url)
			if err != nil {
				if *d == 1 {
					fmt.Println("http.Get() failed: ", err)
				}
				continue
			}
			defer resp.Body.Close()
			byteArray, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				if *d == 1 {
					fmt.Println("ioutil.ReadAll() failed: ", err)
				}
				continue
			}
			json.Unmarshal([]byte(byteArray), &jsonResult)
			array = jsonResult["groups"].([]interface{})
			for k := 0; k < len(array); k++ {
				fmt.Println((array[k].(map[string]interface{}))["name"], ",",
					(array[k].(map[string]interface{}))["status"], ",",
					(array[k].(map[string]interface{}))["current"])
			}
			break
		}
	}
}
