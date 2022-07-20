package rottenOrange

import "testing"

func Test_rottenOrange(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				grid: [][]int{
					{2, 1, 1},
					{1, 1, 0},
					{0, 1, 1},
				},
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				grid: [][]int{
					{2, 1, 1},
					{0, 1, 1},
					{1, 0, 1},
				},
			},
			want: -1,
		},
		{
			name: "test3",
			args: args{
				grid: [][]int{
					{0, 2},
					{1, 1},
				},
			},
			want: 2,
		},
		{
			name: "test4",
			args: args{
				grid: [][]int{
					{0, 1, 2, 2},
					{1, 1, 1, 2},
					{0, 1, 1, 2},
					{2, 1, 0, 1},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rottenOrange(tt.args.grid); got != tt.want {
				t.Errorf("rottenOrange() = %v, want %v", got, tt.want)
			}
		})
	}
}
