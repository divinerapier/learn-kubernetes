package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["storage_batch.go"],
    importpath = "github.com/divinerapier/learn-kubernetes/pkg/registry/batch/rest",
    deps = [
        "//pkg/api/legacyscheme:go_default_library",
        "//pkg/apis/batch:go_default_library",
        "//pkg/registry/batch/cronjob/storage:go_default_library",
        "//pkg/registry/batch/job/storage:go_default_library",
        "//staging/src/k8s.io/api/batch/v1:go_default_library",
        "//staging/src/k8s.io/api/batch/v1beta1:go_default_library",
        "//staging/src/k8s.io/api/batch/v2alpha1:go_default_library",
        "//staging/src/k8s.io/apiserver/pkg/registry/generic:go_default_library",
        "//staging/src/k8s.io/apiserver/pkg/registry/rest:go_default_library",
        "//staging/src/k8s.io/apiserver/pkg/server:go_default_library",
        "//staging/src/k8s.io/apiserver/pkg/server/storage:go_default_library",
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
)
