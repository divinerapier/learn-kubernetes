load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "clientset_generated.go",
        "doc.go",
        "register.go",
    ],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/k8s.io/sample-apiserver/pkg/generated/clientset/versioned/fake",
    importpath = "k8s.io/sample-apiserver/pkg/generated/clientset/versioned/fake",
    visibility = ["//visibility:public"],
    deps = [
        "//staging/src/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/runtime/serializer:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/watch:go_default_library",
        "//staging/src/k8s.io/client-go/discovery:go_default_library",
        "//staging/src/k8s.io/client-go/discovery/fake:go_default_library",
        "//staging/src/k8s.io/client-go/testing:go_default_library",
        "//staging/src/k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1:go_default_library",
        "//staging/src/k8s.io/sample-apiserver/pkg/apis/wardle/v1beta1:go_default_library",
        "//staging/src/k8s.io/sample-apiserver/pkg/generated/clientset/versioned:go_default_library",
        "//staging/src/k8s.io/sample-apiserver/pkg/generated/clientset/versioned/typed/wardle/v1alpha1:go_default_library",
        "//staging/src/k8s.io/sample-apiserver/pkg/generated/clientset/versioned/typed/wardle/v1alpha1/fake:go_default_library",
        "//staging/src/k8s.io/sample-apiserver/pkg/generated/clientset/versioned/typed/wardle/v1beta1:go_default_library",
        "//staging/src/k8s.io/sample-apiserver/pkg/generated/clientset/versioned/typed/wardle/v1beta1/fake:go_default_library",
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
