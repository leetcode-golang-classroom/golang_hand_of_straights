package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	groupSize := 3
	for idx := 0; idx < b.N; idx++ {
		isNStraightHand(hand, groupSize)
	}

}
func Test_isNStraightHand(t *testing.T) {
	type args struct {
		hand      []int
		groupSize int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "hand = [1,2,3,6,2,3,4,7,8], groupSize = 3",
			args: args{hand: []int{1, 2, 3, 6, 2, 3, 4, 7, 8}, groupSize: 3},
			want: true,
		},
		{
			name: "hand = [1,2,3,4,5], groupSize = 4",
			args: args{hand: []int{1, 2, 3, 4, 5}, groupSize: 4},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNStraightHand(tt.args.hand, tt.args.groupSize); got != tt.want {
				t.Errorf("isNStraightHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
