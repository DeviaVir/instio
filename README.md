# instio

Instantly generate Golang CMS framework components, paired with an easy to use API.

### Local Development

Run local version:

```
$ bazel run //cmd/instio:instio -- [args]
```

Update go dependencies:

```
$ go mod tidy
```

Update dependencies from `go.mod`:

```
$ bazel run //:gazelle -- update-repos -from_file=go.mod
```

### Inspiration

Heavily inspired by [ponzu](https://github.com/ponzu-cms/ponzu).