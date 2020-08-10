# [KPT](https://googlecontainertools.github.io/kpt/)

## For building declarative manifest

### Consumer Workflow:

1. Upstearm Remote

`fetch the package` (human)

```bash
kpt pkg get https://github.com/kubernetes/examples/staging/cockroachdb cockroachdb
```

`display package content`

```bash
kpt cfg count cockroachdb
```

2. Local Instance

`list the setter`

```bash
export SRC_REPO=https://github.com/GoogleContainerTools/kpt.git
kpt pkg get $SRC_REPO/package-examples/helloworld-set@v0.6.0 helloworld

kpt cfg list-setters cockroachdb
```

`modify the package` (human)

```bash
kpt cfg set helloworld/ replicas 3
```

