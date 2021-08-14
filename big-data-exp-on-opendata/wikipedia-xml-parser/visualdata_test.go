package main

import "testing"

func Test_parsingTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"one word", args{title: "Wikipedia: Albedo"}, "A"},
		{"two words", args{title: "Wikipedia: Julie Doucet"}, "J"},
		{"one letter", args{title: "Wikipedia: A"}, "A"},
		{"sentence", args{title: "Wikipedia: An American in Paris"}, "A"},
		{"long", args{title: "Wikipedia: Academy Award for Best Production Design"}, "A"},
		{"not only A", args{title: "Wikipedia: Mitochondrion"}, "M"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsingTitle(tt.args.title); got != tt.want {
				t.Errorf("parsingTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
