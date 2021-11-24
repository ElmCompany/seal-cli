package seal

import (
	_ "embed"
	"reflect"
	"testing"
)

func TestParseFlags(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantS      *Seal
		wantOutput string
		wantErr    bool
	}{
		{
			name:  "dry run flag",
			args:  []string{"-dry-run"},
			wantS: &Seal{DryRun: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, gotOutput, err := ParseFlags("myprogram", tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// merr := mergo.Merge(tt.wantS, Instance())
			// if merr != nil {
			// 	t.Errorf("unexpected mergo error: %v", merr)
			// }
			if !reflect.DeepEqual(gotS.DryRun, tt.wantS.DryRun) {
				t.Errorf("ParseFlags() gotS = %v, want %v", gotS.DryRun, tt.wantS.DryRun)
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("ParseFlags() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
