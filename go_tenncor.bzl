load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def dependencies():
    if "com_github_mingkaic_tenncor" not in native.existing_rules():
        git_repository(
            name = "com_github_mingkaic_tenncor",
            remote = "https://github.com/mingkaic/tenncor",
            commit = "30c44662f261279e8e6b0feec5de1a13d0963acd",
        )

    if "io_bazel_rules_go" not in native.existing_rules():
        http_archive(
            name = "io_bazel_rules_go",
            urls = [ "https://github.com/bazelbuild/rules_go/releases/download/0.10.3/rules_go-0.10.3.tar.gz" ],
            sha256 = "feba3278c13cde8d67e341a837f69a029f698d7a27ddbb2a202be7a10b22142a",
        )
