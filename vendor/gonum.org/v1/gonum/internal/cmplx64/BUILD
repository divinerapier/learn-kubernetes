load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "abs.go",
        "conj.go",
        "doc.go",
        "isinf.go",
        "isnan.go",
        "sqrt.go",
    ],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/gonum.org/v1/gonum/internal/cmplx64",
    importpath = "gonum.org/v1/gonum/internal/cmplx64",
    visibility = ["//vendor/gonum.org/v1/gonum:__subpackages__"],
    deps = ["//vendor/gonum.org/v1/gonum/internal/math32:go_default_library"],
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
