package utils

import "testing"

func TestRunCommand(t *testing.T) {
	type args struct {
		name string
		arg  []string
	}
	tests := []struct {
		name        string
		args        args
		wantToExist string
		wantErr     bool
	}{
		{
			name: "it runs Shell CMD",
			args: args{
				name: "/bin/sh",
				arg:  []string{"-c", "touch tmp.test.runcmd.out"},
			},
			wantToExist: "tmp.test.runcmd.out",
			wantErr:     false,
		},
		{
			name: "it runs another Shell CMD based on previous CMD",
			args: args{
				name: "/bin/sh",
				arg:  []string{"-c", "rm tmp.test.runcmd.out"},
			},
			wantErr: false,
		},
		{
			name: "it fails if a CLI not exist",
			args: args{
				name: "/bin/sh",
				arg:  []string{"-c", "mycliwhichnotexist"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RunCommand(tt.args.name, tt.args.arg...); (err != nil) != tt.wantErr {
				t.Errorf("RunCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
