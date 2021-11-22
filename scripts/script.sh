#!/bin/sh
{{- if and .Namespace .Name }}
echo -n {{ .Secret }} | kubeseal --scope strict --namespace {{ .Namespace }} --name {{ .Name }} --cert {{ .Host }}/v1/cert.pem --raw --from-file=/dev/stdin
{{- else if .Namespace }}
echo -n {{ .Secret }} | kubeseal --scope namespace-wide --namespace {{ .Namespace }} --cert {{ .Host }}/v1/cert.pem --raw --from-file=/dev/stdin
{{- else }}
echo -n {{ .Secret }} | kubeseal --scope cluster-wide --cert {{ .Host }}/v1/cert.pem --raw --from-file=/dev/stdin
{{- end }}