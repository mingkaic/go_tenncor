workspace(name = "com_github_mingkaic_go_tenncor")

load("//:go_tenncor.bzl", "dependencies")
dependencies()

load("@com_github_mingkaic_tenncor//:tenncor.bzl", "dependencies", "test_dependencies")
dependencies()

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
