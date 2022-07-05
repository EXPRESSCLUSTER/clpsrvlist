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
	var jsonServers map[string]interface{}
	var jsonResult map[string]interface{}
	var (
		m = flag.Int("m", 0, "Debug flag")
	)

	flag.Parse()

	/* Read servers.json */
	servers, err := ioutil.ReadFile("./servers.json")
	if err != nil {
		fmt.Println("ReadFile() failed: ", err)
		os.Exit(1)
	}
	json.Unmarshal([]byte(servers), &jsonServers)

	/*Get server status */
	arrayServers := jsonServers["servers"].([]interface{})
	for i := 0; i < len(arrayServers); i++ {
		user := arrayServers[i].(map[string]interface{})["user"]
		password := arrayServers[i].(map[string]interface{})["password"]
		ipaddress := arrayServers[i].(map[string]interface{})["ipaddress"]
		port := arrayServers[i].(map[string]interface{})["port"]
		hostname := arrayServers[i].(map[string]interface{})["hostname"]
		url := fmt.Sprintf("http://%s:%s@%s:%s/api/v1/servers/%s", user, password, ipaddress, port, hostname)

		resp, err := http.Get(url)
		if err != nil {
			if *m == 1 {
				fmt.Println("http.Get() failed: ", err)
			}
			fmt.Println(arrayServers[i].(map[string]interface{})["hostname"], "Unknown")
			continue
		}
		defer resp.Body.Close()

		byteArray, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			if *m == 1 {
				fmt.Println("ioutil.ReadAll() failed: ", err)
			}
			fmt.Println(arrayServers[i].(map[string]interface{})["hostname"], "Unknown")
			continue
		}

		json.Unmarshal([]byte(byteArray), &jsonResult)
		array := jsonResult["servers"].([]interface{})
		fmt.Println((array[0].(map[string]interface{}))["name"], (array[0].(map[string]interface{}))["status"])
	}
}
