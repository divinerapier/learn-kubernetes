package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = ["parsers.go"],
    importpath = "github.com/divinerapier/learn-kubernetes/pkg/util/parsers",
    deps = ["//vendor/github.com/docker/distribution/reference:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["parsers_test.go"],
    embed = [":go_default_library"],
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
