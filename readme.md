# Basic Go gRPC Server and Client


Implement a simple gRPC server and client with the following functionality:
- simple RPC
- server-side streaming RPC
- client-side streaming RPC
- bidirectional streaming RPC


# Running the application

1. Install the dependencies

```bash
go mod tidy
```

2. Run the server

```bash
go run server/main.go
```

3. Run the client

```bash
go run client/main.go
```