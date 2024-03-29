package node

import (
	"context"
	"fmt"
	"io"
	"os/exec"
)

type Node struct {
    cmd *exec.Cmd
    stdin          io.WriteCloser
    stdout         io.ReadCloser
    runCodeChannel chan string
    resultChannel chan  string
}

func (n *Node) Init(ctx context.Context) *Node {
    // Create New Node Process
    n.cmd = exec.Command("node", "adapters/node/dist/index.js")

    // Set up stdin and stdout pipes
    var err error
    n.stdin, err = n.cmd.StdinPipe()
    if err != nil {
        fmt.Println("Error creating stdin pipe:", err)
        return n
    }

    n.stdout, err = n.cmd.StdoutPipe()
    if err != nil {
        fmt.Println("Error creating stdout pipe:", err)
        return n
    }

    // Start the Node.js process
    if err := n.cmd.Start(); err != nil {
        fmt.Println("Error starting Node.js process:", err)
        return n
    }

    n.runCodeChannel = make(chan string)
    n.resultChannel = make(chan string)

    go n.handleJSCode(ctx)

    return n
}

func (n *Node) SafelyRunJSCode(code string) (string, error) {
    n.runCodeChannel <- code
    result := <-n.resultChannel
    return result, nil
}

func (n *Node) handleJSCode(ctx context.Context) {
    for {
        select {
        case <- ctx.Done():
            n.stdin.Close()
            n.stdout.Close()
            return
        case c := <- n.runCodeChannel:
            n.resultChannel <- c;
        }
    }
}
