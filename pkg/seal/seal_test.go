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
			wantS: Seal{DryRun: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, gotOutput, err := ParseFlags("myprogram", tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("ParseFlags() gotS = %v, want %v", gotS, tt.wantS)
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("ParseFlags() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
