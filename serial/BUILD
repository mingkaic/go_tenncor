package(
    default_visibility = [ "//visibility:public" ],
)

licenses(["notice"])

load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_proto_library(
    name = "go_default_library",
    importpath = "github.com/mingkaic/go_tenncor/serial",
    proto = "@com_github_mingkaic_tenncor//proto:tenncor_serial_proto",
)
