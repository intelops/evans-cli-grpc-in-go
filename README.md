## Prerequisites

Before running this code, please ensure that you have the following prerequisites installed:
- Go programming language (version 1.16 or later)
- Evans CLI (for interacting with the gRPC service)

## Installation
To install the necessary dependencies, you can use the following command:
```go
go mod tidy
```
## Running the Code
To run the code, follow these steps:
1. Open a terminal or command prompt.
2. Navigate to the directory containing the code.
3. Run the code by running the following command:
    ```go
    go run main.go
    ```
4. The gRPC server will start running and listen for incoming connections on localhost:50051. You should see a log message indicating that the server has started successfully.

## Additional Information
- The code uses the google.golang.org/grpc package to create and run the gRPC server.
- The server listens on `localhost:50051` by default. You can modify the host and port variables in the code to change the server address.
- The code imports the `github.com/azar-intelops/grpc-for-evans/gen/api/v1` package, which presumably contains the gRPC API definitions and message types used by the server. Make sure you have the necessary proto files and dependencies installed for this code to work correctly.

## How to test the server?

Evans CLI is a powerful tool for testing and interacting with gRPC services. Here are some reasons why it is commonly used:

- `Interactive Shell:` Evans provides an interactive shell that allows you to explore and test gRPC services interactively. You can navigate through services, view available methods, and make requests right from the command line.
- `Automatic Code Generation:` Evans can generate client stubs in various programming languages based on the gRPC service definition. This can be useful for integrating the generated client code into your own applications.
- `Support for gRPC Reflection:` Evans supports gRPC reflection, which allows clients to discover the available services and methods on a gRPC server dynamically. This is particularly helpful when the server does not expose a formal service discovery mechanism.
- `Rich Feature Set:` Evans offers many features to assist in gRPC testing, such as support for server-streaming and bidirectional-streaming methods, handling of custom headers and metadata, support for TLS encryption, and more.
- `Easy-to-Use Syntax:` Evans provides a simple and intuitive syntax for making requests and working with gRPC services, making it accessible even for those new to gRPC.

Using Evans CLI, you can conveniently test your gRPC server, verify its functionality, and ensure that it is behaving as expected.

Checkout the blog to see how to work with [Evans-cli](https://intelops.ai/blog/evans-cli-a-go-grpc-client/)

