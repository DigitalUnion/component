package duinflux

import "testing"

func Test_createGrafanaDashboard(t *testing.T) {
	type args struct {
		grafanaUrl string
		token      string
		bucket     string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"", args{"http://172.17.147.39:3000/", "Bearer eyJrIjoiNHBYSE8yNWRKVTc5OHZ4ZDA1bk4ySWM4R1lEd01pekwiLCJuIjoiY2xvdWQiLCJpZCI6MX0=", "testc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createGrafanaDashboard(tt.args.grafanaUrl, tt.args.token, tt.args.bucket)
		})
	}
}
