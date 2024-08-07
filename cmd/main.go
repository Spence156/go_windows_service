// go_windows_service/main.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Spence156/go_windows_service/internal/handlers"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
)

type myservice struct{}

func (m *myservice) Execute(args []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (svcSpecificEC bool, exitCode uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	s <- svc.Status{State: svc.StartPending}
	go handlers.StartServer()
	s <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	for c := range r {
		switch c.Cmd {
		case svc.Interrogate:
			s <- c.CurrentStatus
			time.Sleep(100 * time.Millisecond)
			s <- c.CurrentStatus
		case svc.Stop, svc.Shutdown:
			s <- svc.Status{State: svc.StopPending}
			return false, 0
		default:
			// unexpected control request
		}
	}
	return false, 0
}

func main() {
	serviceName := "example_service"
	fmt.Printf("Service Name: %v", serviceName)
	isWindowsService, err := svc.IsWindowsService()
	if err != nil {
		log.Fatalf("failed to determine if we are running in a Windows service: %v", err)
	}

	if isWindowsService {
		runService(&serviceName, false)
	} else {
		runService(&serviceName, true)
	}
}

func runService(name *string, isDebug bool) {
	var err error
	if isDebug {
		err = debug.Run(*name, &myservice{})
	} else {
		err = svc.Run(*name, &myservice{})
	}
	if err != nil {
		log.Fatalf("%s service failed: %v", *name, err)
	}
}
