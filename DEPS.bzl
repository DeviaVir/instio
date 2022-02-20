load("@bazel_gazelle//:deps.bzl", "go_repository")
def go_deps():
  go_repository(
      name = "com_github_armon_consul_api",
      importpath = "github.com/armon/consul-api",
      sum = "h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
      version = "v0.0.0-20180202201655-eb2c6b5be1b6",
  )
  
  go_repository(
      name = "com_github_bits_and_blooms_bitset",
      importpath = "github.com/bits-and-blooms/bitset",
      sum = "h1:Kn4yilvwNtMACtf1eYDlG8H77R07mZSPbMjLyS07ChA=",
      version = "v1.2.0",
  )
  
  go_repository(
      name = "com_github_blevesearch_bleve_index_api",
      importpath = "github.com/blevesearch/bleve_index_api",
      sum = "h1:nx9++0hnyiGOHJwQQYfsUGzpRdEVE5LsylmmngQvaFk=",
      version = "v1.0.1",
  )
  
  go_repository(
      name = "com_github_blevesearch_bleve_v2",
      build_file_proto_mode = "disable_global",
      importpath = "github.com/blevesearch/bleve/v2",
      sum = "h1:5XKlSdpcjeJdE7n0FUEDeJRJwLuhPxq+k5n7h5UaJkg=",
      version = "v2.3.0",
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
      sum = "h1:7QkALgFNooSq3a46AE+pWeKASAZc9SiNFJhDGF1NDx4=",
      version = "v1.0.3",
  )
  
  go_repository(
      name = "com_github_blevesearch_scorch_segment_api_v2",
      importpath = "github.com/blevesearch/scorch_segment_api/v2",
      sum = "h1:NFwteOpZEvJk5Vg0H6gD0hxupsG3JYocE4DBvsA2GZI=",
      version = "v2.1.0",
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
      name = "com_github_blevesearch_upsidedown_store_api",
      importpath = "github.com/blevesearch/upsidedown_store_api",
      sum = "h1:1SYRwyoFLwG3sj0ed89RLtM15amfX2pXlYbFOnF8zNU=",
      version = "v1.0.1",
  )
  
  go_repository(
      name = "com_github_blevesearch_vellum",
      importpath = "github.com/blevesearch/vellum",
      sum = "h1:+vn8rfyCRHxKVRgDLeR0FAXej2+6mEb5Q15aQE/XESQ=",
      version = "v1.0.7",
  )
  
  go_repository(
      name = "com_github_blevesearch_zapx_v11",
      importpath = "github.com/blevesearch/zapx/v11",
      sum = "h1:TDdcbaA0Yz3Y5zpTrpvyW1AeicqWTJL3g8D5g48RiHM=",
      version = "v11.3.2",
  )
  
  go_repository(
      name = "com_github_blevesearch_zapx_v12",
      importpath = "github.com/blevesearch/zapx/v12",
      sum = "h1:XB09XMg/3ibeIJRCm2zjkaVwrtAuk6c55YRSmVlwUDk=",
      version = "v12.3.2",
  )
  
  go_repository(
      name = "com_github_blevesearch_zapx_v13",
      importpath = "github.com/blevesearch/zapx/v13",
      sum = "h1:mTvALh6oayreac07VRAv94FLvTHeSBM9sZ1gmVt0N2k=",
      version = "v13.3.2",
  )
  
  go_repository(
      name = "com_github_blevesearch_zapx_v14",
      importpath = "github.com/blevesearch/zapx/v14",
      sum = "h1:oW36JVaZDzrzmBa1X5jdTIYzdhkOQnr/ie13Cb2X7MQ=",
      version = "v14.3.2",
  )
  
  go_repository(
      name = "com_github_blevesearch_zapx_v15",
      importpath = "github.com/blevesearch/zapx/v15",
      sum = "h1:OZNE4CQ9hQhnB21ySC7x2/9Q35U3WtRXLAh5L2gdCXc=",
      version = "v15.3.2",
  )
  
  go_repository(
      name = "com_github_boltdb_bolt",
      importpath = "github.com/boltdb/bolt",
      sum = "h1:JQmyP4ZBrce+ZQu0dY660FMfatumYDLun9hBCUVIkF4=",
      version = "v1.3.1",
  )
  
  go_repository(
      name = "com_github_burntsushi_toml",
      importpath = "github.com/BurntSushi/toml",
      sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
      version = "v0.3.1",
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
      sum = "h1:VCYrMzFwEryyhRSeI+/b3tRBSeTpi/8gn5Kf6dxqn+o=",
      version = "v0.2.0",
  )
  
  go_repository(
      name = "com_github_cpuguy83_go_md2man",
      importpath = "github.com/cpuguy83/go-md2man",
      sum = "h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
      version = "v1.0.10",
  )
  
  go_repository(
      name = "com_github_cpuguy83_go_md2man_v2",
      importpath = "github.com/cpuguy83/go-md2man/v2",
      sum = "h1:U+s90UTSYgptZMwQh2aRr3LuazLJIa+Pg3Kc1ylSYVY=",
      version = "v2.0.0-20190314233015-f79a8a8ca69d",
  )
  
  go_repository(
      name = "com_github_davecgh_go_spew",
      importpath = "github.com/davecgh/go-spew",
      sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
      version = "v1.1.1",
  )
  
  go_repository(
      name = "com_github_fsnotify_fsnotify",
      importpath = "github.com/fsnotify/fsnotify",
      sum = "h1:IXs+QLmnXW2CcXuY+8Mzv/fWEsPGWxqefPtCP5CnV9I=",
      version = "v1.4.7",
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
      name = "com_github_inconshreveable_mousetrap",
      importpath = "github.com/inconshreveable/mousetrap",
      sum = "h1:Z8tu5sraLXCXIcARxBp/8cbvlwVa7Z1NHg9XEKhtSvM=",
      version = "v1.0.0",
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
      name = "com_github_nilslice_email",
      importpath = "github.com/nilslice/email",
      sum = "h1:QBb/9NM1sbrBRfDco60YiyoJplwHHU1yRocpT5i2uJU=",
      version = "v0.1.0",
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
      name = "com_github_pmezard_go_difflib",
      importpath = "github.com/pmezard/go-difflib",
      sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
      version = "v1.0.0",
  )
  
  go_repository(
      name = "com_github_rcrowley_go_metrics",
      importpath = "github.com/rcrowley/go-metrics",
      sum = "h1:dY6ETXrvDG7Sa4vE8ZQG4yqWg6UnOcbqTAahkV813vQ=",
      version = "v0.0.0-20190826022208-cac0b30c2563",
  )
  
  go_repository(
      name = "com_github_roaringbitmap_roaring",
      importpath = "github.com/RoaringBitmap/roaring",
      sum = "h1:ckvZSX5gwCRaJYBNe7syNawCU5oruY9gQmjXlp4riwo=",
      version = "v0.9.4",
  )
  
  go_repository(
      name = "com_github_russross_blackfriday",
      importpath = "github.com/russross/blackfriday",
      sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
      version = "v1.5.2",
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
      name = "com_github_ugorji_go_codec",
      importpath = "github.com/ugorji/go/codec",
      sum = "h1:3SVOIvH7Ae1KRYyQWRjXWJEA9sS/c/pjvH++55Gr648=",
      version = "v0.0.0-20181204163529-d75b2dcb6bc8",
  )
  
  go_repository(
      name = "com_github_urfave_cli_v2",
      importpath = "github.com/urfave/cli/v2",
      sum = "h1:qph92Y649prgesehzOrQjdWyxFOp/QVM+6imKHad91M=",
      version = "v2.3.0",
  )
  
  go_repository(
      name = "com_github_xordataexchange_crypt",
      importpath = "github.com/xordataexchange/crypt",
      sum = "h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
      version = "v0.0.3-0.20170626215501-b2862e3d0a77",
  )
  
  go_repository(
      name = "in_gopkg_check_v1",
      importpath = "gopkg.in/check.v1",
      sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
      version = "v0.0.0-20161208181325-20d25e280405",
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
      name = "in_gopkg_yaml_v2",
      importpath = "gopkg.in/yaml.v2",
      sum = "h1:fvjTMHxHEw/mxHbtzPi3JCcKXQRAnQTBRo6YCJSVHKI=",
      version = "v2.2.3",
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
  
  go_repository(
      name = "com_github_bazelbuild_rules_go",
      importpath = "github.com/bazelbuild/rules_go",
      sum = "h1:kX4jVcstqrsRqKPJSn2mq2o+TI21edRzEJSrEOMQtr0=",
      version = "v0.30.0",
  )