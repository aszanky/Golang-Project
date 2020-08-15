# GOLANG WITH GRPC and Memory Cache LRU Algorithms
To run client and server, run:
   - Please ```dep ensure``` first
   - on client filepath, type ```go run client.go```
   - on server filepath, type ```go run main.go```

Structuring the code:
- Client : to demonstrate communication between client and server
- cmd : binary for server code
- delivery: grpc
- entity: business data structure
- repository: memory cache using gcache library
- usecase: currently it is empty folder, it is usually used for business logic
- generate.sh - command to create proto
- README.md - documentation the code structure