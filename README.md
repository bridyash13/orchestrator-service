# Mage Take home exercise - Orchestrator Service

## Run Locally

Clone the project in required folder

```bash
  git clone https://github.com/bridyash13/orchestrator-service.git
```

Change to project folder

```bash
  cd orchestrator-service
```

Start the datamock server

```bash
  go run datamock/server/server.go
```

Start the logic server in a new terminal

```bash
  go run logic/server/server.go
```

Start the orchestrator server in a new terminal

```bash
  go run server/server.go
```

Start the client in a new terminal

```bash
  go run client/client.go
```

Doc Link:
https://docs.google.com/document/d/1WxmRHVPzUszk85ZiawcGXF7pzDAUNGqf5QphpScWPXI/edit?usp=sharing