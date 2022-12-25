package errors

import (
	"errors"
	"testing"
)

func TestErrValidation_As(t *testing.T) {
	type fields struct {
		Property    string
		Given       *string
		Description string
	}
	type args struct {
		t interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "true",
			fields: fields{
				Property:    "Name",
				Given:       new(string),
				Description: "cannot be empty",
			},
			args: args{
				t: &ErrValidation{},
			},
			want: true,
		},
		{
			name: "false",
			fields: fields{
				Property:    "Name",
				Given:       new(string),
				Description: "cannot be empty",
			},
			args: args{
				t: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrValidation{
				Property:    tt.fields.Property,
				Given:       tt.fields.Given,
				Description: tt.fields.Description,
			}
			if got := e.As(tt.args.t); got != tt.want {
				t.Errorf("ErrValidation.As() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrValidation_Error(t *testing.T) {
	type fields struct {
		Property    string
		Given       *string
		Description string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "nil",
			fields: fields{
				Property:    "Name",
				Given:       nil,
				Description: "given nil",
			},
			want: "the property \"Name\" given 'nil' given nil.",
		},
		{
			name: "non-nil pointer",
			fields: fields{
				Property:    "Name",
				Given:       new(string),
				Description: "given non-nil pointer",
			},
			want: "the property \"Name\" given \"\" given non-nil pointer.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrValidation{
				Property:    tt.fields.Property,
				Given:       tt.fields.Given,
				Description: tt.fields.Description,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrValidation.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrNotFound_As(t *testing.T) {
	type fields struct {
		Object string
		Id     string
	}
	type args struct {
		t interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "true",
			fields: fields{
				Object: "User",
				Id:     "1",
			},
			args: args{
				t: &ErrNotFound{},
			},
			want: true,
		},
		{
			name: "false",
			fields: fields{
				Object: "User",
				Id:     "1",
			},
			args: args{
				t: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrNotFound{
				Object: tt.fields.Object,
				Id:     tt.fields.Id,
			}
			if got := e.As(tt.args.t); got != tt.want {
				t.Errorf("ErrNotFound.As() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrNotFound_Error(t *testing.T) {
	type fields struct {
		Object string
		Id     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "not found 1",
			fields: fields{
				Object: "Name",
				Id:     "1",
			},
			want: "\"Name\" not found (id: \"1\")",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrNotFound{
				Object: tt.fields.Object,
				Id:     tt.fields.Id,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrNotFound.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
