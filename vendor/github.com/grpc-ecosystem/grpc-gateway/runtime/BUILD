load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "context.go",
        "convert.go",
        "doc.go",
        "errors.go",
        "handler.go",
        "marshal_json.go",
        "marshal_jsonpb.go",
        "marshaler.go",
        "marshaler_registry.go",
        "mux.go",
        "pattern.go",
        "proto2_convert.go",
        "proto_errors.go",
        "query.go",
    ],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/github.com/grpc-ecosystem/grpc-gateway/runtime",
    importpath = "github.com/grpc-ecosystem/grpc-gateway/runtime",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/jsonpb:go_default_library",
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/runtime/internal:go_default_library",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/utilities:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
        "//vendor/google.golang.org/grpc/grpclog:go_default_library",
        "//vendor/google.golang.org/grpc/metadata:go_default_library",
        "//vendor/google.golang.org/grpc/status:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/runtime/internal:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
