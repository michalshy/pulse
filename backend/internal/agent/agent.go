package agent

import (
	"fmt"
	"os"
	"os/exec"
)

// Implemenntation of the Pulse agent.
// Serves as a documentation on how to implement communication with agent in GO
type Agent struct {
	available bool
}

// First - we have to launch Pulse agent on the service init itself
func New(path string) *Agent {
	a := Agent{}
	a.available = false
	cmd := exec.Command(path)
	// We take agents stdout / stderr to the backend process
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("start failed:", err)
		return nil
	}
	a.available = true
	return &a
}

// Then, we need to connect via websocket with the spawned agent
func (a *Agent) ConnectToAgent() {

}
