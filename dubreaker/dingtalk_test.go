package dubreaker

import "testing"

func Test_getLocalIp(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLocalIp()
			if (err != nil) != tt.wantErr {
				t.Errorf("getLocalIp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getLocalIp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
