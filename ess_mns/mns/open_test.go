/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/09/14 12:15 PM
 */

package mns

import "testing"

func TestGetInstanceIdInfo(t *testing.T) {
	type args struct {
		instanceIds []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{[]string{"i-2zecdh933ojb21ym1ctf"}}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := GetInstanceIdInfo(tt.args.instanceIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInstanceIdInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("GetInstanceIdInfo() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
