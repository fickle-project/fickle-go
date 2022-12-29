package users

import "testing"

func TestUser_Verify(t *testing.T) {
	type fields struct {
		Id           IdUser
		Name         string
		Email        string
		PasswordHash string
	}
	type args struct {
		password []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: func() string { h, _ := generatePasswordHash("ficklePassword"); return string(h) }(),
			},
			args: args{
				password: []byte("ficklePassword"),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "fail: mismatch",
			fields: fields{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: func() string { h, _ := generatePasswordHash("ficklePassword"); return string(h) }(),
			},
			args: args{
				password: []byte("mismatch"),
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Id:           tt.fields.Id,
				Name:         tt.fields.Name,
				Email:        tt.fields.Email,
				PasswordHash: tt.fields.PasswordHash,
			}
			got, err := u.Verify(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}
