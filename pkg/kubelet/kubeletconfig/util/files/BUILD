package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = ["files.go"],
    importpath = "github.com/divinerapier/learn-kubernetes/pkg/kubelet/kubeletconfig/util/files",
    deps = ["//pkg/util/filesystem:go_default_library"],
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

go_test(
    name = "go_default_test",
    srcs = ["files_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//pkg/kubelet/kubeletconfig/util/test:go_default_library",
        "//pkg/util/filesystem:go_default_library",
    ],
)
