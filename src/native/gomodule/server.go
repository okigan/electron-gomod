package main

import (
	"log"
	"time"

	hardwaremonitoring "golangmodule/gen"
)

// Server is our struct that will handle the Hardware monitoring Logic
// It will fulfill the gRPC interface generated
type Server struct {
	hardwaremonitoring.UnimplementedHardwareMonitorServer
}

// Monitor is used to start a stream of HardwareStats
func (s *Server) Monitor(req *hardwaremonitoring.EmptyRequest,
	stream hardwaremonitoring.HardwareMonitor_MonitorServer) error {
	// Start a ticker that executes each 2 seconds
	timer := time.NewTicker(2 * time.Second)

	for {
		select {
		// Exit on stream context done
		case <-stream.Context().Done():
			log.Println("ending stream")
			return nil
		case <-timer.C:
			log.Println("sending stats")
			// Grab stats and output
			hwStats, err := s.GetStats()
			if err != nil {
				log.Println(err.Error())
			}

			// Send the Hardware stats on the stream
			err = stream.Send(hwStats)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}

var counter int32 = 0

// GetStats will extract system stats and output a Hardware Object, or an error
// if extraction fails
func (s *Server) GetStats() (*hardwaremonitoring.HardwareStats, error) {

	hwStats := &hardwaremonitoring.HardwareStats{
		Cpu:        counter,
		MemoryFree: counter,
		MemoryUsed: counter,
	}

	counter++

	return hwStats, nil
}
