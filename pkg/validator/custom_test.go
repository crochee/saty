package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type structCustomValidation struct {
	Order string `binding:"order"`
}

func TestOrderWithDBSort(t *testing.T) {
	engine, err := New()
	assert.Nil(t, err)
	err = RegisterValidation(engine, "order", OrderWithDBSort)
	assert.Nil(t, err)

	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "nil",
			args: args{
				value: "",
			},
			wantErr: true,
		},
		{
			name: "one_ok",
			args: args{
				value: "update",
			},
			wantErr: false,
		},
		{
			name: "one_single_char",
			args: args{
				value: "u",
			},
			wantErr: true,
		},
		{
			name: "one_single_invalid_char",
			args: args{
				value: "%",
			},
			wantErr: true,
		},
		{
			name: "one_two_char_contain_",
			args: args{
				value: "u_",
			},
			wantErr: true,
		},
		{
			name: "one_many_char",
			args: args{
				value: "ucfd 98he'	s'[",
			},
			wantErr: true,
		},
		{
			name: "mult_ok",
			args: args{
				value: "io asc,created_at,updated_at ASC,name desc,deleted DESC",
			},
			wantErr: false,
		},
		{
			name: "mult_failed",
			args: args{
				value: "upDate DESC,created",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err = engine.ValidateStruct(structCustomValidation{Order: tt.args.value})
			if tt.wantErr {
				assert.NotNil(t, err, "err:%v", err)
				return
			}
			assert.Nil(t, err, "err:%v", err)
		})
	}
}
