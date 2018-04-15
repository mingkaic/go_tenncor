all: test

proto: serial graphmgr

serial:
	bazel build //serial:go_default_library

graphmgr:
	bazel build //graphmgr:go_default_library

test: proto
	bazel test //:test
