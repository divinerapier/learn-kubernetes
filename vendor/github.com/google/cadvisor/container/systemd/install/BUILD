load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["install.go"],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/github.com/google/cadvisor/container/systemd/install",
    importpath = "github.com/google/cadvisor/container/systemd/install",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/cadvisor/container:go_default_library",
        "//vendor/github.com/google/cadvisor/container/systemd:go_default_library",
        "//vendor/k8s.io/klog:go_default_library",
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
