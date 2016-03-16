default:
	go build --ldflags '-extldflags "-static"' monte_carlo.go
