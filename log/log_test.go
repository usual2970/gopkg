package log

import (
	"testing"
)

func TestAddGlobalFields(t *testing.T) {
	type args struct {
		fields map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddGlobalFields(tt.args.fields)
		})
	}
}

func Test_dlogger_Debug(t *testing.T) {

	type args struct {
		v []interface{}
	}
	tests := []struct {
		name   string
		args   args
		fields map[string]interface{}
	}{
		{
			name: "debug1",
			args: args{
				v: []interface{}{"1234"},
			},
			fields: map[string]interface{}{
				"a":  1,
				"aa": 11,
			},
		},
		{
			name: "debug2",
			args: args{
				v: []interface{}{"5678"},
			},
			fields: map[string]interface{}{
				"b":  2,
				"bb": 22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := WithFields(tt.fields)
			l.Debug(tt.args.v...)
		})
	}
}
