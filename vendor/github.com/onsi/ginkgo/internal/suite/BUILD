load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["suite.go"],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/github.com/onsi/ginkgo/internal/suite",
    importpath = "github.com/onsi/ginkgo/internal/suite",
    visibility = ["//vendor/github.com/onsi/ginkgo:__subpackages__"],
    deps = [
        "//vendor/github.com/onsi/ginkgo/config:go_default_library",
        "//vendor/github.com/onsi/ginkgo/internal/containernode:go_default_library",
        "//vendor/github.com/onsi/ginkgo/internal/failer:go_default_library",
        "//vendor/github.com/onsi/ginkgo/internal/leafnodes:go_default_library",
        "//vendor/github.com/onsi/ginkgo/internal/spec:go_default_library",
        "//vendor/github.com/onsi/ginkgo/internal/spec_iterator:go_default_library",
        "//vendor/github.com/onsi/ginkgo/internal/specrunner:go_default_library",
        "//vendor/github.com/onsi/ginkgo/internal/writer:go_default_library",
        "//vendor/github.com/onsi/ginkgo/reporters:go_default_library",
        "//vendor/github.com/onsi/ginkgo/types:go_default_library",
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
