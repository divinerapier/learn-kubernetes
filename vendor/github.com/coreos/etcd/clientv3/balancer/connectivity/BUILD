load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["connectivity.go"],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/github.com/coreos/etcd/clientv3/balancer/connectivity",
    importpath = "github.com/coreos/etcd/clientv3/balancer/connectivity",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.uber.org/zap:go_default_library",
        "//vendor/google.golang.org/grpc/connectivity:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
