/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/12/06 17:43
 */

package dugindb

import (
	"reflect"
	"testing"
)

func TestGetSliceLMR(t *testing.T) {
	type args struct {
		sliceL []string
		sliceR []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
		want2 []string
	}{
		// TODO: Add test cases.
		{"", args{[]string{"a", "a"}, []string{"a", "a", "a", "a"}}, nil, []string{"a"}, nil},
		{"", args{[]string{"a"}, []string{"c"}}, []string{"a"}, nil, []string{"c"}},
		{"", args{[]string{"a", "b"}, []string{"a", "c"}}, []string{"b"}, []string{"a"}, []string{"c"}},
		{"", args{[]string{"a", "b"}, []string{"a", "c"}}, []string{"b"}, []string{"a"}, []string{"c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := GetSliceLMR(tt.args.sliceL, tt.args.sliceR)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSliceLMR() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetSliceLMR() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetSliceLMR() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
