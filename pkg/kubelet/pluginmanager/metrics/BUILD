load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["metrics.go"],
    importpath = "github.com/divinerapier/learn-kubernetes/pkg/kubelet/pluginmanager/metrics",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/kubelet/pluginmanager/cache:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
        "//vendor/k8s.io/klog:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["metrics_test.go"],
    embed = [":go_default_library"],
    deps = ["//pkg/kubelet/pluginmanager/cache:go_default_library"],
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
