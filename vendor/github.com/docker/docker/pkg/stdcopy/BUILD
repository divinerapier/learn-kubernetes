load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["stdcopy.go"],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/github.com/docker/docker/pkg/stdcopy",
    importpath = "github.com/docker/docker/pkg/stdcopy",
    visibility = ["//visibility:public"],
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
