package logic

import "testing"

func TestCatchMeLeftRightEqual(t *testing.T) {
	type args struct {
		input      string
		startDigit int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test_1",
			args: args{
				input:      "LLRR=",
				startDigit: 0,
			},
			want:    "210122",
			wantErr: false,
		},
		{
			name: "test_2",
			args: args{
				input:      "==RLL",
				startDigit: 0,
			},
			want:    "000210",
			wantErr: false,
		},
		{
			name: "test_3",
			args: args{
				input:      "=LLRR",
				startDigit: 0,
			},
			want:    "221012",
			wantErr: false,
		},
		{
			name: "test_4",
			args: args{
				input:      "RRL=R",
				startDigit: 0,
			},
			want:    "012001",
			wantErr: false,
		},
		{
			name: "test_5",
			args: args{
				input:      "R=LLR",
				startDigit: 0,
			},
			want:    "022101",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CatchMeLeftRightEqual(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CatchMeLeftRightEqual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CatchMeLeftRightEqual() input [%v] got [%v], want [%v]", tt.args.input, got, tt.want)
			}
		})
	}
}
