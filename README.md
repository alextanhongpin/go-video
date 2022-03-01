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

The command below basically has the following meaning - find all files with the `.proto` extension in the `/proto` folder and generate the stub in go.

```bash
$ find ./proto/**.proto -exec \
	protoc -I/usr/local/include -I. \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:. "{}" \; 
```

## Create reverse-proxy

The command below basically has the following meaning - find all files with the `.proto` extension in the `/proto` folder and generate the reverse-proxy (gateway) in go.

```bash
$ find ./proto/**.proto -exec \
	protoc -I/usr/local/include -I. \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:. "{}" \;
```

## Notes

The above commands can be inlined in the code through `go generate`. However, for readability a Makefile is preferred.


## Installing ffmpeg

```
# Install the `exodus_bundler` package, if you haven't already.
pip install --user exodus_bundler
export PATH="${HOME}/.local/bin/:${PATH}"

# Create an `ffmpeg` bundle and extract it in the current directory.
exodus --tarball ffmpeg | tar -zx
```

- https://intoli.com/blog/transcoding-on-aws-lambda/
