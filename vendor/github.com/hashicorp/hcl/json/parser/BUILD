load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "flatten.go",
        "parser.go",
    ],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/github.com/hashicorp/hcl/json/parser",
    importpath = "github.com/hashicorp/hcl/json/parser",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/hcl/hcl/ast:go_default_library",
        "//vendor/github.com/hashicorp/hcl/hcl/token:go_default_library",
        "//vendor/github.com/hashicorp/hcl/json/scanner:go_default_library",
        "//vendor/github.com/hashicorp/hcl/json/token:go_default_library",
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
