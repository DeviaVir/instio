{ pkgs }: {
    deps = [
        pkgs.go
        pkgs.gopls
        pkgs.bazel_4
        pkgs.adoptopenjdk-bin
        pkgs.less
        pkgs.docker
    ];
}