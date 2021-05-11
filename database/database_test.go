package database

import (
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestUint8To64Ar(t *testing.T) {
	type args struct {
		inp []uint8
	}
	tests := []struct {
		args    args
		wantRes []uint64
	}{
		{
			args:    args{[]uint8{0, 0, 0, 0, 0, 2, 1, 1, 0, 0, 0, 0, 0, 0, 1, 2}},
			wantRes: []uint64{131329, 258},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if gotRes := uint8To64Ar(tt.args.inp); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Uint8To64Ar() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestUint64To8Ar(t *testing.T) {
	type args struct {
		inp []uint64
	}
	tests := []struct {
		args    args
		wantRes []uint8
	}{
		{
			args:    args{[]uint64{131329, 258}},
			wantRes: []uint8{0, 0, 0, 0, 0, 2, 1, 1, 0, 0, 0, 0, 0, 0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if gotRes := uint64To8Ar(tt.args.inp); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Uint64To8Ar() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
