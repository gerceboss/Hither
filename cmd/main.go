package main

import (
	"fmt"
	"os"
)

var containers = map[string]string{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid command. Try `hither help`.")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "help":
		showHelp()
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Usage: hither run <executable_file>")
			os.Exit(1)
		}
		executable := os.Args[2]
		runExecutable(executable)
	case "list":
		listContainers()
	case "stop":
		if len(os.Args) < 3 {
			fmt.Println("Usage: hither stop <container_name>")
			os.Exit(1)
		}
		containerName := os.Args[2]
		stopContainer(containerName)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
	}
}

func showHelp() {
	fmt.Println("Hither CLI - Command List:")
	fmt.Println("  hither help                 - Show this help message")
	fmt.Println("  hither run <executable>     - Run an executable in a container")
	fmt.Println("  hither list                 - List all running containers")
	fmt.Println("  hither stop <container>     - Stop a running container")
}

func runExecutable(executable string) {
	// you have the name and the content of the file ab iska kr lenge kuch to
	// filename := os.Args[1]
	// content, err := ioutil.ReadAll(filename) 
	// if err != nil {
	// 	log.Fatalf("Failed to read file content: %v", err)
	// }
	// contentStr := string(content)
		fmt.Println("running")

}

func listContainers() {
	if len(containers) == 0 {
		fmt.Println("No running containers.")
	} else {
		fmt.Println("Currently running containers:")
		for container, executable := range containers {
			fmt.Printf("  %s - Running %s\n", container, executable)
		}
	}
}

func stopContainer(containerName string) {
	executable, exists := containers[containerName]
	if !exists {
		fmt.Printf("Container %s does not exist.\n", containerName)
		return
	}
	fmt.Printf("Stopping container %s (running %s)...\n", containerName, executable)
	delete(containers, containerName)
	fmt.Printf("Container %s stopped and removed.\n", containerName)
}
