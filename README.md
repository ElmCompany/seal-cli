# Overview

This is a wrapper CLI of[kubeseal CLI](https://github.com/bitnami-labs/sealed-secrets/releases).

# Prerequisies

- Go 1.16 +

# Getting started

- built it `make build`
- use it `./seal-cli -s MyPassWord -h https://seal.company.lan`

where `https://seal.company.lan` is [the base URL of your seale-secret controller](https://github.com/bitnami-labs/sealed-secrets/blob/main/helm/sealed-secrets/values.yaml#L50)

# Usage

TODO

# License

[LGPL v3](LICENSE)
