load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "constants.go",
        "fileinfo.go",
        "fix.go",
        "generate.go",
        "kinds.go",
        "known_go_imports.go",
        "known_proto_imports.go",
        "lang.go",
        "package.go",
        "resolve.go",
        "std_package_list.go",
    ],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/github.com/bazelbuild/bazel-gazelle/internal/language/go",
    importpath = "github.com/bazelbuild/bazel-gazelle/internal/language/go",
    visibility = ["//vendor/github.com/bazelbuild/bazel-gazelle:__subpackages__"],
    deps = [
        "//vendor/github.com/bazelbuild/bazel-gazelle/internal/config:go_default_library",
        "//vendor/github.com/bazelbuild/bazel-gazelle/internal/flag:go_default_library",
        "//vendor/github.com/bazelbuild/bazel-gazelle/internal/label:go_default_library",
        "//vendor/github.com/bazelbuild/bazel-gazelle/internal/language:go_default_library",
        "//vendor/github.com/bazelbuild/bazel-gazelle/internal/language/proto:go_default_library",
        "//vendor/github.com/bazelbuild/bazel-gazelle/internal/pathtools:go_default_library",
        "//vendor/github.com/bazelbuild/bazel-gazelle/internal/repos:go_default_library",
        "//vendor/github.com/bazelbuild/bazel-gazelle/internal/resolve:go_default_library",
        "//vendor/github.com/bazelbuild/bazel-gazelle/internal/rule:go_default_library",
        "//vendor/github.com/bazelbuild/buildtools/build:go_default_library",
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
