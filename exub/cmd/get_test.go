package cmd

import "testing"

func Test_getUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test", args: args{url: "http://localhost:8080"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getUrl(tt.args.url)
		})
	}
}
