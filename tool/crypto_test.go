package tool

import (
	"testing"
)

func TestMD5(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "md5",
			args: args{
				msg: "123456" + "capi91nt10f5j95umsvg",
			},
			want: "4df4eeb0879e2b89cb606bb5d6319a83",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.msg); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}
