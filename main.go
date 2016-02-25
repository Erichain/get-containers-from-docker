package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

// docker's ports
type APIPort struct {
	PrivatePort			int64
	PublicPort			int64
	Type				string
	IP					string
}

// the options of listing containers
type ListContainersOptions struct {
	All     			bool
	Limit   			int
	Since   			string
	Before  			string
	Size    			bool
	Filters 			map[string][]string
}

// the options each container represents
type APIContainers struct {
	ID					string		//`json:"Id" yaml:"Id"`
	Names				[]string
	Image				string
	Command				string
	Created				int64
	Status				string
	Ports				[]APIPort
	Labels				map[string]string
	SizeRw				int64
	SizeRootFs			int64
}

func main() {
	var options ListContainersOptions

	options.All = true
	options.Limit = 5

	ListContainers(options)
}

func ListContainers(options ListContainersOptions) ([]APIContainers, error) {
	host := "unix:///var/run/docker.sock"
	path := host + "/containers/json?" + formatConvertion(options)

	fmt.Println("%s", path)

	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var containers []APIContainers
	if err := json.NewDecoder(res.Body).Decode(&containers); err != nil {
		return nil, err
	}

	fmt.Println(containers)

	return containers, nil
}

// convert the options to json
func formatConvertion(options ListContainersOptions) string {
	config := &options

	url, err := json.Marshal(config)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(string(url))

	return string(url)
}