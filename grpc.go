Greeting: Hello world

service Greeter {
  
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}


message HelloRequest {
  string name = 1;
}


message HelloReply {
  string message = 1;
}

service Greeter {
  
  rpc SayHello (HelloRequest) returns (HelloReply) {}
  
  rpc SayHelloAgain (HelloRequest) returns (HelloReply) {}
}


message HelloRequest {
  string name = 1;
}


message HelloReply {
  string message = 1;
}

$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

 func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
        return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
if err != nil {
        log.Fatalf("could not greet: %v", err)
}
log.Printf("Greeting: %s", r.GetMessage())

$ go run greeter_server/main.go

$ go run greeter_client/main.go Alice

Greeting: Hello Alice
Greeting: Hello again Alice
