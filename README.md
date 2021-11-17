# Overview

This is a wrapper CLI of[kubeseal CLI](https://github.com/bitnami-labs/sealed-secrets/releases), specifically the [raw mode](https://github.com/bitnami-labs/sealed-secrets/tree/7ce55636a2e264239593bc197feaa76c82f35026#raw-mode-experimental).

If you just need to encrypt your secret on **RAW** mode, this CLI will be the easy way.

# Prerequisies

- Go 1.16 +
- [kubeseal CLI](https://github.com/bitnami-labs/sealed-secrets/releases) v0.17.0 or above to be installed on PATH
- [sealed-secret controller](https://github.com/bitnami-labs/sealed-secrets/tree/main/helm/sealed-secrets) v0.17.0 or above to be up & running on a k8s cluster & accessible, for example, by https://seal.company.lan

# Getting started

- built it `make build`
- use it `./seal-cli -s MyPassWord -h https://seal.company.lan`

where `https://seal.company.lan` is [the base URL of your seale-secret controller](https://github.com/bitnami-labs/sealed-secrets/blob/main/helm/sealed-secrets/values.yaml#L50)

# Usage

**cluster wide - sealed secret can be used in any namespace**
```sh
seal -s MyPassWord123 -h https://seal.company.lan
```

**namespace wide - sealed secret can be used only in a specific namespace**
```sh
seal -s MyPassWord123 -h https://seal.company.lan -n mynamespace
```

**strict  - sealed secret can be used only in a specific namespace & for specific secret resource (mysecret)**
```sh
seal -s MyPassWord123 -h https://seal.company.lan -n mynamespace -name mysecret
```

**Dry run - preview the kubeseal command behind the scenes**

```sh
# append it to any command, and it show the underline command instead of the real execution
seal [ ... ] -dry-run
```

# License

[LGPL v3](LICENSE)
