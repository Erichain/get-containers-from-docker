/**
 * implemented by go dockerClient
 */
package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)
func main() {
	var endpoint = "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	containers, _ := client.ListContainers(docker.ListContainersOptions{All:true})

	for _, container := range containers {
		fmt.Println("ID: ", container.ID)
		fmt.Println("Created: ", container.Created)
	}
}