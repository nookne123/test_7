package logic

import "testing"

func TestFindTheMostValuablePath(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "test_1",
			args: args{
				input: "[[59], [73, 41], [52, 40, 53], [26, 53, 6, 34]]",
			},
			want:    237,
			wantErr: false,
		},
		{
			name: "test_2",
			args: args{
				input: "[\n  [59],\n  [73, 41],\n  [52, 40, 9],\n  [26, 53, 6, 34],\n  [10, 51, 87, 86, 81],\n  [61, 95, 66, 57, 25, 68],\n  [90, 81, 80, 38, 92, 67, 73]]",
			},
			want:    470,
			wantErr: false,
		},
		{
			name: "test_3",
			args: args{
				input: "[[59], [73, 41], [52, 40, 53], [26, 53, 6, 34]],",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindTheMostValuablePath(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindTheMostValuablePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindTheMostValuablePath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
