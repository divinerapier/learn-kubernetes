load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "requests.go",
        "results.go",
    ],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/github.com/gophercloud/gophercloud/openstack/identity/v3/extensions/trusts",
    importpath = "github.com/gophercloud/gophercloud/openstack/identity/v3/extensions/trusts",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/gophercloud/gophercloud/openstack/identity/v3/tokens:go_default_library"],
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
