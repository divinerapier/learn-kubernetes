load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "npb_ep.go",
        "npb_is.go",
        "tf_wide_deep.go",
        "utils.go",
        "workloads.go",
    ],
    importpath = "github.com/divinerapier/learn-kubernetes/test/e2e_node/perf/workloads",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/kubelet/apis/config:go_default_library",
        "//pkg/kubelet/cm/cpumanager:go_default_library",
        "//staging/src/k8s.io/api/core/v1:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/api/resource:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
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
