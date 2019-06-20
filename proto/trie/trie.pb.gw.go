// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/trie/trie.proto

/*
Package trie is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package trie

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_Trie_KeySearch_0(ctx context.Context, marshaler runtime.Marshaler, client TrieClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq KeySearchTrieRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.KeySearch(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterTrieHandlerFromEndpoint is same as RegisterTrieHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterTrieHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterTrieHandler(ctx, mux, conn)
}

// RegisterTrieHandler registers the http handlers for service Trie to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterTrieHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterTrieHandlerClient(ctx, mux, NewTrieClient(conn))
}

// RegisterTrieHandlerClient registers the http handlers for service Trie
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "TrieClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "TrieClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "TrieClient" to call the correct interceptors.
func RegisterTrieHandlerClient(ctx context.Context, mux *runtime.ServeMux, client TrieClient) error {

	mux.Handle("POST", pattern_Trie_KeySearch_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Trie_KeySearch_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Trie_KeySearch_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Trie_KeySearch_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"zldz", "trie", "keysearch"}, ""))
)

var (
	forward_Trie_KeySearch_0 = runtime.ForwardResponseMessage
)
