package mo

import "testing"

func TestMo_Info(t *testing.T) {
	type fields struct {
	}
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"t1", fields{}, args{[]interface{}{"Log1"}}},
		{"t1", fields{}, args{[]interface{}{"Log1", 123}}},
		{"t1", fields{}, args{[]interface{}{"Log1 ", 123, " de"}}},
		{"t1", fields{}, args{[]interface{}{"Log1 ", 123, TestMo_Info}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Mo{}
			d.Info(tt.args.args...)
		})
	}
}
