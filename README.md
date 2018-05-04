# go-video

Video API written in go gRPC and gRPC Gateway.


# Installation

```bash
$ make setup
```

# Proto

Create a `.proto` file in `/protos` folder:

```proto
syntax = "proto3";
package video;

message GetVideosRequest {
  string query = 1;
}

message GetVideosResponse {
  string query = 1;
}

service VideoService {
  rpc Echo(GetVideosRequest) returns (GetVideosResponse) {
    option (google.api.http) {
      post: "/v1/videos",
      body: "*"
    }
  }
}
```

## Create gRPC Stub
