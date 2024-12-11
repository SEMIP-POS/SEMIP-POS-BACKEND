# Default target
.PHONY: run
run:
	go run main.go serveHttp --config config.yaml --secret secret.yaml
