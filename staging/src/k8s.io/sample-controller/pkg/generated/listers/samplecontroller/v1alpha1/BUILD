load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "expansion_generated.go",
        "foo.go",
    ],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/k8s.io/sample-controller/pkg/generated/listers/samplecontroller/v1alpha1",
    importpath = "k8s.io/sample-controller/pkg/generated/listers/samplecontroller/v1alpha1",
    visibility = ["//visibility:public"],
    deps = [
        "//staging/src/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/labels:go_default_library",
        "//staging/src/k8s.io/client-go/tools/cache:go_default_library",
        "//staging/src/k8s.io/sample-controller/pkg/apis/samplecontroller/v1alpha1:go_default_library",
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
