load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "AppendSliceValue.go",
        "main.go",
        "version.go",
    ],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/github.com/go-bindata/go-bindata/go-bindata",
    importpath = "github.com/go-bindata/go-bindata/go-bindata",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/go-bindata/go-bindata:go_default_library"],
)

go_binary(
    name = "go-bindata",
    embed = [":go_default_library"],
    tags = ["manual"],
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
