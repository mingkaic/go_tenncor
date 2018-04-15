package(
    default_visibility = [ "//visibility:public" ],
)

licenses(["notice"])

load("@io_bazel_rules_go//go:def.bzl", "go_test")

GRPC_COMPILE_DEPS = [
    "@com_github_golang_protobuf//proto:go_default_library",
    "@com_github_golang_glog//:go_default_library",
    "@org_golang_google_grpc//:go_default_library",
    "@org_golang_x_net//context:go_default_library",
]

go_test(
    name = "test",
    size = "small",
    srcs = ["test.go"],
    deps = [
        "//serial:go_default_library",
        "//graphmgr:go_default_library",
    ] + GRPC_COMPILE_DEPS,
)
