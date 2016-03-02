/**
 * implemented with plain go
 */
package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
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
	All     			int						`url:"all"`
	Limit   			int						`url:"limit"`
	Before  			string					`url:"before"`
	Size    			int						`url:"size"`
}

// the options each container represents
type APIContainers struct {
	ID					string
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
	options := ListContainersOptions{1, 5, "8dfafdbc3a40", 1}

	ListContainers(options)
}

func ListContainers(options ListContainersOptions) ([]APIContainers, error) {
	var containers []APIContainers

	v, _ := query.Values(options)
	host := "unix:///var/run/docker.sock"
	path := host + "/containers/json?" + v.Encode()

	fmt.Println("%s", path)

	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&containers); err != nil {
		return nil, err
	}

	for _, container := range containers {
		fmt.Println("ID: ", container.ID)
		fmt.Println("Created: ", container.Created)
	}

	return containers, nil
}
