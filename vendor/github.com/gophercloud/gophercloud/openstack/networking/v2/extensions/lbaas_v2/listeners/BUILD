load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "requests.go",
        "results.go",
        "urls.go",
    ],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/listeners",
    importpath = "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/listeners",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gophercloud/gophercloud:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/l7policies:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/pools:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/pagination:go_default_library",
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
