package users_test

import (
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"reflect"
	"testing"
)

func TestUser_Update(t *testing.T) {
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
		p users.UpdateUserParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    users.User
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Id:    "1",
				Name:  "fickle",
				Email: "test@fickle.com",
			},
			args: args{
				r: r,
				p: users.UpdateUserParam{
					Name: func() *string { n := "updated"; return &n }(),
				},
			},
			want: users.User{
				Id:    "1",
				Name:  "updated",
				Email: "test@fickle.com",
			},
			wantErr: false,
		},
		{
			name: "failed to validate: 'Name' cannot be empty",
			fields: fields{
				Id:    "1",
				Name:  "fickle",
				Email: "test@fickle.com",
			},
			args: args{
				r: r,
				p: users.UpdateUserParam{
					Name: func() *string { n := ""; return &n }(),
				},
			},
			want:    users.User{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Email' cannot be empty",
			fields: fields{
				Id:    "1",
				Name:  "fickle",
				Email: "test@fickle.com",
			},
			args: args{
				r: r,
				p: users.UpdateUserParam{
					Email: func() *string { n := ""; return &n }(),
				},
			},
			want:    users.User{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Email' invalid",
			fields: fields{
				Id:    "1",
				Name:  "fickle",
				Email: "test@fickle.com",
			},
			args: args{
				r: r,
				p: users.UpdateUserParam{
					Email: func() *string { n := "invalid"; return &n }(),
				},
			},
			want:    users.User{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Email' already used",
			fields: fields{
				Id:    "1",
				Name:  "fickle",
				Email: "test@fickle.com",
			},
			args: args{
				r: r,
				p: users.UpdateUserParam{
					Email: func() *string { n := "test2@elkcif.com"; return &n }(),
				},
			},
			want:    users.User{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Password' too short",
			fields: fields{
				Id:    "2",
				Name:  "elkcif",
				Email: "test2@elkcif.com",
			},
			args: args{
				r: r,
				p: users.UpdateUserParam{
					Password: func() *string { n := "test"; return &n }(),
				},
			},
			want:    users.User{},
			wantErr: true,
		},
		{
			name: "ok",
			fields: fields{
				Id:    "2",
				Name:  "elkcif",
				Email: "test2@elkcif.com",
			},
			args: args{
				r: r,
				p: users.UpdateUserParam{
					Password: func() *string { n := "newPassword"; return &n }(),
				},
			},
			want: users.User{
				Id:    "2",
				Name:  "elkcif",
				Email: "test2@elkcif.com",
			},
			wantErr: false,
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
			got, err := u.Update(tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.PasswordHash = got.PasswordHash
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
