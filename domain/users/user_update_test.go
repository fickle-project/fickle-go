package users_test

import (
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"testing"
)

func TestUser_Update(t *testing.T) {
	type args struct {
		r users.IRepository
		p users.UpdateUserParam
	}
	tests := []struct {
		name              string
		initialData       []users.CreateUserParam
		toUpdateUserEmail string
		args              args
		want              users.User
		wantErr           bool
	}{
		{
			name: "ok",
			initialData: []users.CreateUserParam{{
				Name:     "fickle",
				Email:    "test@fickle.com",
				Password: "ficklePassword",
			}},
			toUpdateUserEmail: "test@fickle.com",
			args: args{
				r: inmemory.NewRepositoryUsers(),
				p: users.UpdateUserParam{
					Name:  func() *string { n := "updated"; return &n }(),
					Email: func() *string { n := "test2@fickle.com"; return &n }(),
				},
			},
			want: users.User{
				Name:  "updated",
				Email: "test2@fickle.com",
			},
			wantErr: false,
		},
		{
			name:              "not found",
			initialData:       []users.CreateUserParam{},
			toUpdateUserEmail: "test@fickle.com",
			args: args{
				r: inmemory.NewRepositoryUsers(),
				p: users.UpdateUserParam{
					Name: func() *string { n := "updated"; return &n }(),
				},
			},
			want:    users.User{},
			wantErr: true,
		},
		{
			name: "failed to validate: no update",
			initialData: []users.CreateUserParam{{
				Name:     "fickle",
				Email:    "test@fickle.com",
				Password: "ficklePassword",
			}},
			toUpdateUserEmail: "test@fickle.com",
			args: args{
				r: inmemory.NewRepositoryUsers(),
				p: users.UpdateUserParam{},
			},
			want:    users.User{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Name' empty",
			initialData: []users.CreateUserParam{{
				Name:     "fickle",
				Email:    "test@fickle.com",
				Password: "ficklePassword",
			}},
			toUpdateUserEmail: "test@fickle.com",
			args: args{
				r: inmemory.NewRepositoryUsers(),
				p: users.UpdateUserParam{
					Name: func() *string { n := ""; return &n }(),
				},
			},
			want:    users.User{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Email' empty",
			initialData: []users.CreateUserParam{{
				Name:     "fickle",
				Email:    "test@fickle.com",
				Password: "ficklePassword",
			}},
			toUpdateUserEmail: "test@fickle.com",
			args: args{
				r: inmemory.NewRepositoryUsers(),
				p: users.UpdateUserParam{
					Email: func() *string { n := ""; return &n }(),
				},
			},
			want:    users.User{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Email' already used",
			initialData: []users.CreateUserParam{{
				Name:     "fickle",
				Email:    "test@fickle.com",
				Password: "ficklePassword",
			}, {
				Name:     "fickle",
				Email:    "test2@fickle.com",
				Password: "ficklePassword",
			}},
			toUpdateUserEmail: "test@fickle.com",
			args: args{
				r: inmemory.NewRepositoryUsers(),
				p: users.UpdateUserParam{
					Email: func() *string { n := "test2@fickle.com"; return &n }(),
				},
			},
			want:    users.User{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Email' invalid",
			initialData: []users.CreateUserParam{{
				Name:     "fickle",
				Email:    "test@fickle.com",
				Password: "ficklePassword",
			}},
			toUpdateUserEmail: "test@fickle.com",
			args: args{
				r: inmemory.NewRepositoryUsers(),
				p: users.UpdateUserParam{
					Email: func() *string { n := "invalid"; return &n }(),
				},
			},
			want:    users.User{},
			wantErr: true,
		},
		{
			name: "failed to validate: no update",
			initialData: []users.CreateUserParam{{
				Name:     "fickle",
				Email:    "test@fickle.com",
				Password: "ficklePassword",
			}},
			toUpdateUserEmail: "test@fickle.com",
			args: args{
				r: inmemory.NewRepositoryUsers(),
				p: users.UpdateUserParam{},
			},
			want:    users.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := users.NewService()
			var u users.User
			for _, p := range tt.initialData {
				u2, err := s.NewUser(inmemory.NewFactoryUsers(), tt.args.r, p)
				if err != nil {
					t.Errorf("userService.NewUser() error = %v", err)
					return
				}
				if u2.Email == tt.toUpdateUserEmail {
					u = u2
				}
			}
			got, err := u.Update(tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && (got.Name != tt.want.Name || got.Email != tt.want.Email) {
				t.Errorf("User.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
