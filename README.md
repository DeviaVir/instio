# instio

Instantly generate Golang CMS framework components, paired with an easy to use API.

### Usage



### Local Development

Run local version:

```
$ bazel run --run_under "cd $PWD && " //cmd/instio:instio -- --hostname 0.0.0.0 run
```

Update go dependencies:

```
$ go mod tidy
```

Update dependencies from `go.mod`:

```
$ bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=DEPS.bzl%go_deps
```

### Inspiration

Heavily inspired by [ponzu](https://github.com/ponzu-cms/ponzu).