package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "mounter",
    embed = [":go_default_library"],
)

go_library(
    name = "go_default_library",
    srcs = ["mounter.go"],
    importpath = "github.com/divinerapier/learn-kubernetes/cluster/gce/gci/mounter",
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
)
