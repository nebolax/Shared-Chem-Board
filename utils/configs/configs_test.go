package configs

import (
	"reflect"
	"testing"
)

type testStruct struct {
	key1 string
	var2 map[int]string
	x    int
}

func TestGet(t *testing.T) {
	TestSet(t)

	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"1", args{"k1"}, 2},
		{"1", args{"k2"}, testStruct{"ts", map[int]string{5: "gq", 9: "fdsfsf"}, 27}},
		{"1", args{"k3"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{"k1", 2}},
		{"1", args{"k2", testStruct{"ts", map[int]string{5: "gq", 9: "fdsfsf"}, 27}}},
		{"1", args{"k3", true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.args.key, tt.args.value)
		})
	}
}
