package main

import (
	"reflect"
	"testing"
)


func Test_fetchLongestStablePrices(t *testing.T) {
	type args struct {
		data []int
		x int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test data 1",
			args: args{
				[]int{2,4,3,6,6,3},
				0,
			},
			want: []int{6,6},
		},
		{
			name: "Test data 1",
			args: args{
				[]int{3,1,2,1,2,2,1,3,1,1,2,2,2,2},
				1,
			},
			want: []int{1,2,1,2,2,1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fetchLongestStablePrices(tt.args.data, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchLongestStablePrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
