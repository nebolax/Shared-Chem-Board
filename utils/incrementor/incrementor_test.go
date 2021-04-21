package incrementor

import "testing"

func TestNext(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"emptykey", args{""}, 1},
		{"emptykey", args{""}, 2},
		{"k1", args{"k1"}, 1},
		{"clutch", args{"clutch"}, 1},
		{"k1", args{"k1"}, 2},
		{"k1", args{"k1"}, 3},
		{"clutch", args{"clutch"}, 2},
		{"k1", args{"k1"}, 4},
		{"clutch", args{"clutch"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Next(tt.args.key); got != tt.want {
				t.Errorf("Next(\"%v\") = %v, want %v", tt.args.key, got, tt.want)
			}
		})
	}
}
