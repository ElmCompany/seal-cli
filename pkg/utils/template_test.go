package utils

import (
	"testing"

	"git.elm.sa/devops/seal-cli/pkg/seal"
)

func TestProcessString(t *testing.T) {
	type args struct {
		str  string
		vars interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "it can interpolate variable",
			args: args{
				str: "-namespace {{ .Namespace }}",
				vars: &seal.Seal{
					Namespace: "hello",
				},
			},
			want: "-namespace hello",
		},
		{
			name: "it can interpolate control flow",
			args: args{
				str: "{{ if .Namespace }}namespace-wide{{ end }}",
				vars: &seal.Seal{
					Namespace: "hello",
				},
			},
			want: "namespace-wide",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessString(tt.args.str, tt.args.vars); got != tt.want {
				t.Errorf("ProcessString() = %v, want %v", got, tt.want)
			}
		})
	}
}
