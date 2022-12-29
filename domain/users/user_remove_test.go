package users_test

import (
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"testing"
)

func TestUser_Remove(t *testing.T) {
	r := inmemory.NewRepositoryUsers()
	r.AddUser(users.User{
		Id:           "1",
		Name:         "fickle",
		Email:        "test@fickle.com",
		PasswordHash: "",
	})
	r.AddUser(users.User{
		Id:           "2",
		Name:         "elkcif",
		Email:        "test2@elkcif.com",
		PasswordHash: "",
	})

	type fields struct {
		Id           users.IdUser
		Name         string
		Email        string
		PasswordHash string
	}
	type args struct {
		r users.IRepository
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: "",
			},
			args: args{
				r: r,
			},
			wantErr: false,
		},
		{
			name: "ok",
			fields: fields{
				Id:           "2",
				Name:         "elkcif",
				Email:        "test2@elkcif.com",
				PasswordHash: "",
			},
			args: args{
				r: r,
			},
			wantErr: false,
		},
		{
			name: "not found",
			fields: fields{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: "",
			},
			args: args{
				r: r,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &users.User{
				Id:           tt.fields.Id,
				Name:         tt.fields.Name,
				Email:        tt.fields.Email,
				PasswordHash: tt.fields.PasswordHash,
			}
			if err := u.Remove(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("User.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
