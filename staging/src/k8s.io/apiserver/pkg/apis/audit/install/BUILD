package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = ["install.go"],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/install",
    importpath = "k8s.io/apiserver/pkg/apis/audit/install",
    deps = [
        "//staging/src/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
        "//staging/src/k8s.io/apiserver/pkg/apis/audit:go_default_library",
        "//staging/src/k8s.io/apiserver/pkg/apis/audit/v1:go_default_library",
        "//staging/src/k8s.io/apiserver/pkg/apis/audit/v1alpha1:go_default_library",
        "//staging/src/k8s.io/apiserver/pkg/apis/audit/v1beta1:go_default_library",
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

go_test(
    name = "go_default_test",
    srcs = ["roundtrip_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//staging/src/k8s.io/apimachinery/pkg/api/apitesting/roundtrip:go_default_library",
        "//staging/src/k8s.io/apiserver/pkg/apis/audit/fuzzer:go_default_library",
    ],
)
