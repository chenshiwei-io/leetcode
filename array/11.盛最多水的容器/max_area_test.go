package _1_盛最多水的容器

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入：[1,8,6,2,5,4,8,3,7]
输出：49 `, args: args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, want: 49},
		{name: `输入：height = [1,1]
输出：1`, args: args{height: []int{1, 1}}, want: 1},
		{name: `输入：height = [4,3,2,1,4]
输出：16`, args: args{height: []int{4, 3, 2, 1, 4}}, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
