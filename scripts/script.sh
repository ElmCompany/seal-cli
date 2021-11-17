#!/bin/sh
echo -n {{ .Secret }} | kubeseal --cert {{ .Host }}/v1/cert.pem --raw --from-file=/dev/stdin --scope cluster-wide