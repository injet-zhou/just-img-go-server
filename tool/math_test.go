package tool

import "testing"

func TestTwo2TheNthPower(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2^0",
			args: args{
				number: 1,
			},
			want: 0,
		},
		{
			name: "2^1",
			args: args{
				number: 2,
			},
			want: 1,
		},
		{
			name: "0",
			args: args{
				number: 0,
			},
			want: -1,
		},
		{
			name: "-1",
			args: args{
				number: -1,
			},
			want: -1,
		},
		{
			name: "2^10",
			args: args{
				number: 1024,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Two2TheNthPower(tt.args.number); got != tt.want {
				t.Errorf("Two2TheNthPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
