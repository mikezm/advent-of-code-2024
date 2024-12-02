package day2

import "testing"

func Test_isSafe(t *testing.T) {
	type test map[string]struct {
		report report
		want   bool
	}

	tests := test{
		"one": {
			report: report{7, 6, 4, 2, 1},
			want:   true,
		},
		"two": {
			report: report{1, 2, 7, 8, 9},
			want:   false,
		},
		"three": {
			report: report{9, 7, 6, 2, 1},
			want:   false,
		},
		"four": {
			report: report{1, 3, 2, 4, 5},
			want:   false,
		},
		"five": {
			report: report{8, 6, 4, 4, 1},
			want:   false,
		},
		"six": {
			report: report{1, 3, 6, 7, 9},
			want:   true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := isSafe(tt.report); got != tt.want {
				t.Errorf("isSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSafeWhenDampened(t *testing.T) {
	type test map[string]struct {
		report report
		want   bool
	}

	tests := test{
		"one": {
			report: report{7, 6, 4, 2, 1},
			want:   true,
		},
		"two": {
			report: report{1, 2, 7, 8, 9},
			want:   false,
		},
		"three": {
			report: report{9, 7, 6, 2, 1},
			want:   false,
		},
		"four": {
			report: report{1, 3, 2, 4, 5},
			want:   true,
		},
		"five": {
			report: report{8, 6, 4, 4, 1},
			want:   true,
		},
		"six": {
			report: report{1, 3, 6, 7, 9},
			want:   true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := isSafeWhenDampened(tt.report); got != tt.want {
				t.Errorf("isSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
