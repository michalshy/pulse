package agent

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"pulse/internal/config"
	pb "pulse/internal/gen/pulse"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Implemenntation of the Pulse agent.
// Serves as a documentation on how to implement communication with agent in GO
type Agent struct {
	client    pb.PulseClient
	available bool
}

// launches agent exe
func create(path string) error {
	cmd := exec.Command(path)
	// We take agents stdout / stderr to the backend process
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("start failed:", err)
		return err
	}
	time.Sleep(2 * time.Second)
	return nil
}

func connect(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return conn, err
	}
	return conn, nil
}

// Gets us launched and configured agent
func New(config config.PulseAgent) (*Agent, error) {
	var err error
	var conn *grpc.ClientConn
	// launch
	err = create(config.Binary)
	if err != nil {
		return nil, err
	}

	// connect
	conn, err = connect(
		fmt.Sprintf("%s:%d", config.Host, config.Port),
	)
	if err != nil {
		return nil, err
	}

	return &Agent{
		pb.NewPulseClient(conn),
		true,
	}, nil
}

func (agent *Agent) heartbeat() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := agent.client.Heartbeat(ctx, &pb.HeartbeatRequest{
		Timestamp: time.Now().UnixMilli(),
	})
	if err != nil {
		return err
	}

	slog.Info("heartbeat response:", "status", resp.Ok, "timestamp", resp.Timestamp)
	return nil
}

func (agent *Agent) StartHeartbeat(ctx context.Context, interval uint) {
	ticker := time.NewTicker(time.Duration(interval))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := agent.heartbeat(); err != nil {
				slog.Error("Heartbeat failed", "error", err)
			}
		case <-ctx.Done():
			slog.Info("Heartbeat stopped")
			return
		}
	}
}
