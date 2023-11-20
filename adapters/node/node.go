package node

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

type Node struct {
	Cmd *exec.Cmd
	Stdin io.WriteCloser
	Stdout io.ReadCloser
}

func(n *Node) Init() *Node  {
	// Create New Node Process
	n.Cmd = exec.Command("node", "runner.js")
	
	// Create pipes for the input/output of the script
	stdin, err := n.Cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error creating stdin pipe:", err)
        return nil
    }
	n.Stdin = stdin

	stdout, _ := n.Cmd.StdoutPipe()
	if err != nil {
        fmt.Println("Error creating stdin pipe:", err)
        return nil
    }
	n.Stdout = stdout

	return n
}

func(n *Node) SafelyRunJSCode(code string) (*string, error) {
    // Start the Node process
    if err := n.Cmd.Start(); err != nil {
        fmt.Println("Error starting Node process:", err)
        return nil, err
    }

	// Write code to Node.js process
    _, err := n.Stdin.Write([]byte(code))
    if err != nil {
        fmt.Println("Error writing to stdin:", err)
        return nil, err
    }

    // Close stdin to signal that no more input will be sent
    if err := n.Stdin.Close(); err != nil {
        fmt.Println("Error closing stdin:", err)
        return nil, err
    }

    // Read output from Node.js process
    scanner := bufio.NewScanner(n.Stdout)
    output := ""
    for scanner.Scan() {
        output += scanner.Text() + "\n"
    }

    // Wait for the Node process to finish
    if err := n.Cmd.Wait(); err != nil {
        fmt.Println("Error waiting for Node process:", err)
        return nil, err
    }

    return &output, nil
}