package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/k8s.io/apiserver/plugin/pkg/audit",
    importpath = "k8s.io/apiserver/plugin/pkg/audit",
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//staging/src/k8s.io/apiserver/plugin/pkg/audit/buffered:all-srcs",
        "//staging/src/k8s.io/apiserver/plugin/pkg/audit/dynamic:all-srcs",
        "//staging/src/k8s.io/apiserver/plugin/pkg/audit/fake:all-srcs",
        "//staging/src/k8s.io/apiserver/plugin/pkg/audit/log:all-srcs",
        "//staging/src/k8s.io/apiserver/plugin/pkg/audit/truncate:all-srcs",
        "//staging/src/k8s.io/apiserver/plugin/pkg/audit/webhook:all-srcs",
    ],
    tags = ["automanaged"],
)
