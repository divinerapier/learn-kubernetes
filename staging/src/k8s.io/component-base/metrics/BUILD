load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "counter.go",
        "gauge.go",
        "histogram.go",
        "labels.go",
        "metric.go",
        "opts.go",
        "processstarttime.go",
        "registry.go",
        "summary.go",
        "version.go",
        "version_parser.go",
        "wrappers.go",
    ],
    importmap = "github.com/divinerapier/learn-kubernetes/vendor/k8s.io/component-base/metrics",
    importpath = "k8s.io/component-base/metrics",
    visibility = ["//visibility:public"],
    deps = [
        "//staging/src/k8s.io/apimachinery/pkg/version:go_default_library",
        "//staging/src/k8s.io/component-base/version:go_default_library",
        "//vendor/github.com/blang/semver:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
        "//vendor/github.com/prometheus/client_model/go:go_default_library",
        "//vendor/github.com/prometheus/procfs:go_default_library",
        "//vendor/k8s.io/klog:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "counter_test.go",
        "gauge_test.go",
        "histogram_test.go",
        "opts_test.go",
        "registry_test.go",
        "summary_test.go",
        "version_parser_test.go",
    ],
    embed = [":go_default_library"],
    deps = [
        "//staging/src/k8s.io/apimachinery/pkg/version:go_default_library",
        "//vendor/github.com/blang/semver:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
        "//vendor/github.com/prometheus/common/expfmt:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//staging/src/k8s.io/component-base/metrics/legacyregistry:all-srcs",
        "//staging/src/k8s.io/component-base/metrics/prometheus/clientgo:all-srcs",
        "//staging/src/k8s.io/component-base/metrics/prometheus/ratelimiter:all-srcs",
        "//staging/src/k8s.io/component-base/metrics/prometheus/restclient:all-srcs",
        "//staging/src/k8s.io/component-base/metrics/prometheus/version:all-srcs",
        "//staging/src/k8s.io/component-base/metrics/prometheus/workqueue:all-srcs",
        "//staging/src/k8s.io/component-base/metrics/testutil:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
