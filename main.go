package main

import (
	"net/url"
	"net/http"
	"encoding/json"
)

// docker's ports
type APIPort struct {
	PrivatePort		int64
	PublicPort		int64
	Type			string
	IP				string
}

// the options of listing containers
type ListContainersOptions struct {
	All     bool
	Limit   int
	Since   string
	Before  string
	Size    bool
	Filters map[string][]string
}

// the options each container represents
type APIContainers struct {
	ID			string		//`json:"Id" yaml:"Id"`
	Names		[]string
	Image		string
	Command		string
	Created		int64
	Status		string
	Ports		[]APIPort
	Labels		map[string]string
	SizeRw		int64
	SizeRootFs	int64
}

func main() {
	var options ListContainersOptions

	options.All = true
	options.Limit = 5


	ListContainers(options)
}

func ListContainers(options ListContainersOptions) ([]APIContainers, error) {
	path := "/containers/json?" + url.QueryEscape(options)
	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var containers []APIContainers
	if err := json.NewDecoder(res.Body).Decode(&containers); err != nil {
		return nil, err
	}

	return containers, nil
}