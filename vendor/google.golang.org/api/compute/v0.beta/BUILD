load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["compute-gen.go"],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/google.golang.org/api/compute/v0.beta",
    importpath = "google.golang.org/api/compute/v0.beta",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/google.golang.org/api/gensupport:go_default_library",
        "//vendor/google.golang.org/api/googleapi:go_default_library",
        "//vendor/google.golang.org/api/option:go_default_library",
        "//vendor/google.golang.org/api/transport/http:go_default_library",
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
