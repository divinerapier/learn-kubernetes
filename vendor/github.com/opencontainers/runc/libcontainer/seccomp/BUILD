load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "seccomp_unsupported.go",
    ],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/github.com/opencontainers/runc/libcontainer/seccomp",
    importpath = "github.com/opencontainers/runc/libcontainer/seccomp",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/opencontainers/runc/libcontainer/configs:go_default_library"],
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
