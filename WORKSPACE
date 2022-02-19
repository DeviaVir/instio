load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
    ],
)

git_repository(
    name = "com_google_protobuf",
    commit = "09745575a923640154bcf307fba8aedff47f240a",
    remote = "https://github.com/protocolbuffers/protobuf",
    shallow_since = "1558721209 -0700",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

go_rules_dependencies()

go_register_toolchains(version = "1.16.5")

gazelle_dependencies()

protobuf_deps()

go_repository(
    name = "com_github_nilslice_email",
    importpath = "github.com/nilslice/email",
    sum = "h1:QBb/9NM1sbrBRfDco60YiyoJplwHHU1yRocpT5i2uJU=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man_v2",
    importpath = "github.com/cpuguy83/go-md2man/v2",
    sum = "h1:U+s90UTSYgptZMwQh2aRr3LuazLJIa+Pg3Kc1ylSYVY=",
    version = "v2.0.0-20190314233015-f79a8a8ca69d",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_russross_blackfriday_v2",
    importpath = "github.com/russross/blackfriday/v2",
    sum = "h1:lPqVAte+HuHNfhJ/0LC98ESWRz8afy9tM/0RK8m9o+Q=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_shurcool_sanitized_anchor_name",
    importpath = "github.com/shurcooL/sanitized_anchor_name",
    sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_urfave_cli_v2",
    importpath = "github.com/urfave/cli/v2",
    sum = "h1:qph92Y649prgesehzOrQjdWyxFOp/QVM+6imKHad91M=",
    version = "v2.3.0",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
    version = "v0.0.0-20161208181325-20d25e280405",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:fvjTMHxHEw/mxHbtzPi3JCcKXQRAnQTBRo6YCJSVHKI=",
    version = "v2.2.3",
)

go_repository(
    name = "com_github_armon_consul_api",
    importpath = "github.com/armon/consul-api",
    sum = "h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
    version = "v0.0.0-20180202201655-eb2c6b5be1b6",
)

go_repository(
    name = "com_github_blevesearch_bleve",
    importpath = "github.com/blevesearch/bleve",
    sum = "h1:Q8r+fHTt35jtGXJUM0ULwM3Tzg+MRfyai4ZkWDy2xO4=",
    version = "v1.0.14",
)

go_repository(
    name = "com_github_blevesearch_blevex",
    importpath = "github.com/blevesearch/blevex",
    sum = "h1:pnilj2Qi3YSEGdWgLj1Pn9Io7ukfXPoQcpAI1Bv8n/o=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_blevesearch_cld2",
    importpath = "github.com/blevesearch/cld2",
    sum = "h1:/4ikScMMYMqsRFWJjCyzd3CNWB0lxvqDkqa5nEv6NMc=",
    version = "v0.0.0-20200327141045-8b5f551d37f5",
)

go_repository(
    name = "com_github_blevesearch_go_porterstemmer",
    importpath = "github.com/blevesearch/go-porterstemmer",
    sum = "h1:GtmsqID0aZdCSNiY8SkuPJ12pD4jI+DdXTAn4YRcHCo=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_blevesearch_mmap_go",
    importpath = "github.com/blevesearch/mmap-go",
    sum = "h1:JtMHb+FgQCTTYIhtMvimw15dJwu1Y5lrZDMOFXVWPk0=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_blevesearch_segment",
    importpath = "github.com/blevesearch/segment",
    sum = "h1:5lG7yBCx98or7gK2cHMKPukPZ/31Kag7nONpoBt22Ac=",
    version = "v0.9.0",
)

go_repository(
    name = "com_github_blevesearch_snowballstem",
    importpath = "github.com/blevesearch/snowballstem",
    sum = "h1:lMQ189YspGP6sXvZQ4WZ+MLawfV8wOmPoD/iWeNXm8s=",
    version = "v0.9.0",
)

go_repository(
    name = "com_github_blevesearch_zap_v11",
    importpath = "github.com/blevesearch/zap/v11",
    sum = "h1:IrDAvtlzDylh6H2QCmS0OGcN9Hpf6mISJlfKjcwJs7k=",
    version = "v11.0.14",
)

go_repository(
    name = "com_github_blevesearch_zap_v12",
    importpath = "github.com/blevesearch/zap/v12",
    sum = "h1:2o9iRtl1xaRjsJ1xcqTyLX414qPAwykHNV7wNVmbp3w=",
    version = "v12.0.14",
)

go_repository(
    name = "com_github_blevesearch_zap_v13",
    importpath = "github.com/blevesearch/zap/v13",
    sum = "h1:r+VNSVImi9cBhTNNR+Kfl5uiGy8kIbb0JMz/h8r6+O4=",
    version = "v13.0.6",
)

go_repository(
    name = "com_github_blevesearch_zap_v14",
    importpath = "github.com/blevesearch/zap/v14",
    sum = "h1:NdcT+81Nvmp2zL+NhwSvGSLh7xNgGL8QRVZ67njR0NU=",
    version = "v14.0.5",
)

go_repository(
    name = "com_github_blevesearch_zap_v15",
    importpath = "github.com/blevesearch/zap/v15",
    sum = "h1:Ylj8Oe+mo0P25tr9iLPp33lN6d4qcztGjaIsP51UxaY=",
    version = "v15.0.3",
)

go_repository(
    name = "com_github_boltdb_bolt",
    importpath = "github.com/boltdb/bolt",
    sum = "h1:JQmyP4ZBrce+ZQu0dY660FMfatumYDLun9hBCUVIkF4=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_coreos_etcd",
    importpath = "github.com/coreos/etcd",
    sum = "h1:jFneRYjIvLMLhDLCzuTuU4rSJUjRplcJQ7pD7MnhC04=",
    version = "v3.3.10+incompatible",
)

go_repository(
    name = "com_github_coreos_go_etcd",
    importpath = "github.com/coreos/go-etcd",
    sum = "h1:bXhRBIXoTm9BYHS3gE0TtQuyNZyeEMux2sDi4oo5YOo=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_coreos_go_semver",
    importpath = "github.com/coreos/go-semver",
    sum = "h1:3Jm3tLmsgAYcjC+4Up7hJrFBPr+n7rAqYeSw/SZazuY=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_couchbase_ghistogram",
    importpath = "github.com/couchbase/ghistogram",
    sum = "h1:b95QcQTCzjTUocDXp/uMgSNQi8oj1tGwnJ4bODWZnps=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_couchbase_moss",
    importpath = "github.com/couchbase/moss",
    sum = "h1:HCL+xxHUwmOaL44kMM/gU08OW6QGCui1WVFO58bjhNI=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_couchbase_vellum",
    importpath = "github.com/couchbase/vellum",
    sum = "h1:BrbP0NKiyDdndMPec8Jjhy0U47CZ0Lgx3xUC2r9rZqw=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man",
    importpath = "github.com/cpuguy83/go-md2man",
    sum = "h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
    version = "v1.0.10",
)

go_repository(
    name = "com_github_cznic_b",
    importpath = "github.com/cznic/b",
    sum = "h1:SwD98825d6bdB+pEuTxWOXiSjBrHdOl/UVp75eI7JT8=",
    version = "v0.0.0-20181122101859-a26611c4d92d",
)

go_repository(
    name = "com_github_cznic_mathutil",
    importpath = "github.com/cznic/mathutil",
    sum = "h1:iwZdTE0PVqJCos1vaoKsclOGD3ADKpshg3SRtYBbwso=",
    version = "v0.0.0-20181122101859-297441e03548",
)

go_repository(
    name = "com_github_cznic_strutil",
    importpath = "github.com/cznic/strutil",
    sum = "h1:MZRmHqDBd0vxNwenEbKSQqRVT24d3C05ft8kduSwlqM=",
    version = "v0.0.0-20181122101858-275e90344537",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_facebookgo_ensure",
    importpath = "github.com/facebookgo/ensure",
    sum = "h1:8ISkoahWXwZR41ois5lSJBSVw4D0OV19Ht/JSTzvSv0=",
    version = "v0.0.0-20200202191622-63f1cf65ac4c",
)

go_repository(
    name = "com_github_facebookgo_stack",
    importpath = "github.com/facebookgo/stack",
    sum = "h1:JWuenKqqX8nojtoVVWjGfOF9635RETekkoH6Cc9SX0A=",
    version = "v0.0.0-20160209184415-751773369052",
)

go_repository(
    name = "com_github_facebookgo_subset",
    importpath = "github.com/facebookgo/subset",
    sum = "h1:7HZCaLC5+BZpmbhCOZJ293Lz68O7PYrF2EzeiFMwCLk=",
    version = "v0.0.0-20200203212716-c811ad88dec4",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:IXs+QLmnXW2CcXuY+8Mzv/fWEsPGWxqefPtCP5CnV9I=",
    version = "v1.4.7",
)

go_repository(
    name = "com_github_glycerine_go_unsnap_stream",
    importpath = "github.com/glycerine/go-unsnap-stream",
    sum = "h1:Ujru1hufTHVb++eG6OuNDKMxZnGIvF6o/u8q/8h2+I4=",
    version = "v0.0.0-20181221182339-f9677308dec2",
)

go_repository(
    name = "com_github_glycerine_goconvey",
    importpath = "github.com/glycerine/goconvey",
    sum = "h1:gclg6gY70GLy3PbkQ1AERPfmLMMagS60DKF78eWwLn8=",
    version = "v0.0.0-20190410193231-58a59202ab31",
)

go_repository(
    name = "com_github_gofrs_uuid",
    importpath = "github.com/gofrs/uuid",
    sum = "h1:yyYWMnhkhrKwwr8gAOcOCYxOOscHgDS9yZgBrnJfGa0=",
    version = "v4.2.0+incompatible",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:6nsPYzhq5kReh6QImI3k5qWzO4PEbvbIW2cwSfR/6xs=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_golang_snappy",
    importpath = "github.com/golang/snappy",
    sum = "h1:Qgr9rKW7uDUkrbSmQeiDsGa8SjGyCOGtuasMWwvp2P4=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    importpath = "github.com/gopherjs/gopherjs",
    sum = "h1:twflg0XRTjwKpxb/jFExr4HGq6on2dEOmnL6FV+fgPw=",
    version = "v0.0.0-20190910122728-9d188e94fb99",
)

go_repository(
    name = "com_github_gorilla_schema",
    importpath = "github.com/gorilla/schema",
    sum = "h1:YufUaxZYCKGFuAq3c96BOhjgd5nmXiOY9NGzF247Tsc=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    importpath = "github.com/hashicorp/hcl",
    sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hpcloud_tail",
    importpath = "github.com/hpcloud/tail",
    sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_ikawaha_kagome_ipadic",
    importpath = "github.com/ikawaha/kagome.ipadic",
    sum = "h1:pFxZ1PpMpc6ZoBK712YN5cVK0u/ju2DZ+gRIOriJFFs=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    importpath = "github.com/inconshreveable/mousetrap",
    sum = "h1:Z8tu5sraLXCXIcARxBp/8cbvlwVa7Z1NHg9XEKhtSvM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jmhodges_levigo",
    importpath = "github.com/jmhodges/levigo",
    sum = "h1:q5EC36kV79HWeTBWsod3mG11EgStG3qArTKcvlksN1U=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jtolds_gls",
    importpath = "github.com/jtolds/gls",
    sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
    version = "v4.20.0+incompatible",
)

go_repository(
    name = "com_github_kljensen_snowball",
    importpath = "github.com/kljensen/snowball",
    sum = "h1:6DZLCcZeL0cLfodx+Md4/OLC6b/bfurWUOUGs1ydfOU=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_magiconair_properties",
    importpath = "github.com/magiconair/properties",
    sum = "h1:LLgXmsheXeRoUOBOjtwPQCWIYqM/LU1ayDtDePerRcY=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    importpath = "github.com/mitchellh/go-homedir",
    sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    sum = "h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_mschoch_smat",
    importpath = "github.com/mschoch/smat",
    sum = "h1:8imxQsjDm8yFEAVBe7azKmKSgzSkZXDuKkSq9374khM=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_nilslice_jwt",
    importpath = "github.com/nilslice/jwt",
    sum = "h1:T03TBmHTQJFjeAXgX9kGsfWoGTH14lwHQkrPXxvQnT0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    sum = "h1:WSHQ+IS43OoUrWtD1/bbclrwK8TTH5hzp+umCiuxHgs=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    sum = "h1:RE1xgDvH7imwFD45h+u2SgIfERHlS2yNG4DObb5BSKU=",
    version = "v1.4.3",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    importpath = "github.com/pelletier/go-toml",
    sum = "h1:T5zMGML61Wp+FlcbWjRDT7yAxhJNAiPPLOFECq181zc=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_philhofer_fwd",
    importpath = "github.com/philhofer/fwd",
    sum = "h1:UbZqGr5Y38ApvM/V/jEljVxwocdweyH+vmYvRPBnbqQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_rcrowley_go_metrics",
    importpath = "github.com/rcrowley/go-metrics",
    sum = "h1:dY6ETXrvDG7Sa4vE8ZQG4yqWg6UnOcbqTAahkV813vQ=",
    version = "v0.0.0-20190826022208-cac0b30c2563",
)

go_repository(
    name = "com_github_remyoudompheng_bigfft",
    importpath = "github.com/remyoudompheng/bigfft",
    sum = "h1:OdAsTTz6OkFY5QxjkYwrChwuRruF69c169dPK26NUlk=",
    version = "v0.0.0-20200410134404-eec4a21b6bb0",
)

go_repository(
    name = "com_github_roaringbitmap_roaring",
    importpath = "github.com/RoaringBitmap/roaring",
    sum = "h1:gpyfd12QohbqhFO4NVDUdoPOCXsyahYRQhINmlHxKeo=",
    version = "v0.4.23",
)

go_repository(
    name = "com_github_russross_blackfriday",
    importpath = "github.com/russross/blackfriday",
    sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    sum = "h1:m8/z1t7/fwjysjQRYbP0RD+bUIF/8tJwPdEZsI83ACI=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_spf13_cast",
    importpath = "github.com/spf13/cast",
    sum = "h1:oget//CVOEoFewqQxwr0Ej5yjygnqGkvggSE/gB35Q8=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_spf13_cobra",
    importpath = "github.com/spf13/cobra",
    sum = "h1:f0B+LkLX6DtmRH1isoNA9VTtNUK9K8xYd28JNNfOv/s=",
    version = "v0.0.5",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    importpath = "github.com/spf13/jwalterweatherman",
    sum = "h1:XHEdyB+EcvlqZamSM4ZOMGlc93t6AcsBEu9Gc1vn7yk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    sum = "h1:zPAT6CGy6wXeQ7NtTnaTerfKOsV6V6F8agHXFiazDkg=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_spf13_viper",
    importpath = "github.com/spf13/viper",
    sum = "h1:VUFqw5KcqRf7i70GOzW7N+Q7+gxVBkSSqiXB12+JQ4M=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_steveyen_gtreap",
    importpath = "github.com/steveyen/gtreap",
    sum = "h1:CjhzTa274PyJLJuMZwIzCO1PfC00oRa8d1Kc78bFXJM=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:2E4SXV/wtOkTonXsotYi4li6zVWxYlZuYNCXe9XRJyk=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_syndtr_goleveldb",
    importpath = "github.com/syndtr/goleveldb",
    sum = "h1:fBdIW9lB4Iz0n9khmH8w27SJ3QEJ7+IgjPEwGSZiFdE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_tebeka_snowball",
    importpath = "github.com/tebeka/snowball",
    sum = "h1:ujvgLOr6IHbsvB2Vgz27IcxWqDrNu9/oPhhe74lN/Kc=",
    version = "v0.4.2",
)

go_repository(
    name = "com_github_tecbot_gorocksdb",
    importpath = "github.com/tecbot/gorocksdb",
    sum = "h1:g+WoO5jjkqGAzHWCjJB1zZfXPIAaDpzXIEJ0eS6B5Ok=",
    version = "v0.0.0-20191217155057-f0fad39f321c",
)

go_repository(
    name = "com_github_tidwall_gjson",
    importpath = "github.com/tidwall/gjson",
    sum = "h1:6aeJ0bzojgWLa82gDQHcx3S0Lr/O51I9bJ5nv6JFx5w=",
    version = "v1.14.0",
)

go_repository(
    name = "com_github_tidwall_match",
    importpath = "github.com/tidwall/match",
    sum = "h1:+Ho715JplO36QYgwN9PGYNhgZvoUSc9X2c80KVTi+GA=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_tidwall_pretty",
    importpath = "github.com/tidwall/pretty",
    sum = "h1:RWIZEg2iJ8/g6fDDYzMpobmaoGh5OLl4AXtGUGPcqCs=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_tidwall_sjson",
    importpath = "github.com/tidwall/sjson",
    sum = "h1:cuiLzLnaMeBhRmEv00Lpk3tkYrcxpmbU81tAY4Dw0tc=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_tinylib_msgp",
    importpath = "github.com/tinylib/msgp",
    sum = "h1:9fQd+ICuRIu/ue4vxJZu6/LzxN0HwMds2nq/0cFvxHU=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_ugorji_go_codec",
    importpath = "github.com/ugorji/go/codec",
    sum = "h1:3SVOIvH7Ae1KRYyQWRjXWJEA9sS/c/pjvH++55Gr648=",
    version = "v0.0.0-20181204163529-d75b2dcb6bc8",
)

go_repository(
    name = "com_github_willf_bitset",
    importpath = "github.com/willf/bitset",
    sum = "h1:NotGKqX0KwQ72NUzqrjZq5ipPNDQex9lo3WpaS8L2sc=",
    version = "v1.1.10",
)

go_repository(
    name = "com_github_xordataexchange_crypt",
    importpath = "github.com/xordataexchange/crypt",
    sum = "h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
    version = "v0.0.3-0.20170626215501-b2862e3d0a77",
)

go_repository(
    name = "in_gopkg_fsnotify_v1",
    importpath = "gopkg.in/fsnotify.v1",
    sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
    version = "v1.4.7",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    importpath = "gopkg.in/tomb.v1",
    sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
    version = "v1.0.0-20141024135613-dd632973f1e7",
)

go_repository(
    name = "io_etcd_go_bbolt",
    importpath = "go.etcd.io/bbolt",
    sum = "h1:XAzx9gjCb0Rxj7EoqcClPD1d5ZBxZJk0jbuoPHenBt0=",
    version = "v1.3.5",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:f+lwQ+GtmgoY+A2YaQxlSOnDjXcQ7ZRLWOHbC6HtRqE=",
    version = "v0.0.0-20220214200702-86341886e292",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:CIJ76btIcR3eFI5EgSo6k1qKw9KJexJuRLI9G7Hp5wE=",
    version = "v0.0.0-20211112202133-69e39bad7dc2",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:wMNYb4v58l5UBM7MYRLPG6ZhfOqbKu7X5eyFl8ZhKvA=",
    version = "v0.0.0-20180314180146-1d60e4601c6f",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:SrN+KX8Art/Sf4HNj6Zcz06G7VEz+7w9tdXTPOZ7+l4=",
    version = "v0.0.0-20210615035016-665e8c7367d1",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:v+OssWQX+hTHEmOBgwxdZxK4zHq3yOs8F9J7mk0PY8E=",
    version = "v0.0.0-20201126162022-7de9c90e9dd1",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:olpwvP2KacW1ZWvsR7uQhoyTYvKAupfQrRGBFM352Gk=",
    version = "v0.3.7",
)
