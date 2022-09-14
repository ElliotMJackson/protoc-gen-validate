load("@rules_java//java:defs.bzl", "java_library")

java_library(
    name = "java_harness",
    srcs = glob(["*.java"]),
    visibility = ["//tests/harness:__subpackages__"],
    deps = [
        "//java/pgv-java-stub/src/main/java/io/envoyproxy/pgv",
        "//java/pgv-java-validation/src/main/java/io/envoyproxy/pgv",
        "//tests/harness:harness_java_proto",
        "//tests/harness/proto/base:java",
        "//tests/harness/proto/base:proto_java_proto",
        "//tests/harness/proto/other_package:embed_java_proto",
        "//tests/harness/proto/other_package:java",
        "//tests/harness/proto/yet_another_package:embed_java_proto",
        "//tests/harness/proto/yet_another_package:java",
        "//validate:validate_java",
        "@com_google_guava//jar",
        "@com_google_protobuf//:protobuf_java",
    ],
)
