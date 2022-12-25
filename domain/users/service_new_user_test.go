package users_test

import (
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"testing"
)

func Test_userService_NewUser(t *testing.T) {
	type args struct {
		f users.IFactory
		r users.IRepository
		p users.CreateUserParam
	}
	tests := []struct {
		name        string
		initialData []users.CreateUserParam
		args        args
		want        users.User
		wantErr     bool
	}{{
		name:        "success",
		initialData: []users.CreateUserParam{},
		args: args{
			f: inmemory.NewFactoryUsers(),
			r: inmemory.NewRepositoryUsers(),
			p: users.CreateUserParam{
				Name:     "fickle",
				Email:    "test@fickle.com",
				Password: "ficklePasswd",
			},
		},
		want:    users.User{Name: "fickle", Email: "test@fickle.com"},
		wantErr: false,
	}, {
		name:        "failed to validation: 'Name' empty",
		initialData: []users.CreateUserParam{},
		args: args{
			f: inmemory.NewFactoryUsers(),
			r: inmemory.NewRepositoryUsers(),
			p: users.CreateUserParam{
				Name:     "",
				Email:    "test@fickle.com",
				Password: "ficklePasswd",
			},
		},
		want:    users.User{},
		wantErr: true,
	}, {
		name:        "failed to validation: 'Email' empty",
		initialData: []users.CreateUserParam{},
		args: args{
			f: inmemory.NewFactoryUsers(),
			r: inmemory.NewRepositoryUsers(),
			p: users.CreateUserParam{
				Name:     "fickle",
				Email:    "",
				Password: "ficklePasswd",
			},
		},
		want:    users.User{},
		wantErr: true,
	}, {
		name:        "failed to validation: 'Email' invalid",
		initialData: []users.CreateUserParam{},
		args: args{
			f: inmemory.NewFactoryUsers(),
			r: inmemory.NewRepositoryUsers(),
			p: users.CreateUserParam{
				Name:     "fickle",
				Email:    "invalid email address",
				Password: "ficklePasswd",
			},
		},
		want:    users.User{},
		wantErr: true,
	}, {
		name: "failed to validation: 'Email' already used",
		initialData: []users.CreateUserParam{{
			Name:     "elkcif",
			Email:    "test@fickle.com",
			Password: "ficklePasswd",
		}},
		args: args{
			f: inmemory.NewFactoryUsers(),
			r: inmemory.NewRepositoryUsers(),
			p: users.CreateUserParam{
				Name:     "fickle",
				Email:    "test@fickle.com",
				Password: "ficklePasswd",
			},
		},
		want:    users.User{},
		wantErr: true,
	}, {
		name:        "failed to validation: 'Password' too short",
		initialData: []users.CreateUserParam{},
		args: args{
			f: inmemory.NewFactoryUsers(),
			r: inmemory.NewRepositoryUsers(),
			p: users.CreateUserParam{
				Name:     "fickle",
				Email:    "invalid email address",
				Password: "fickle",
			},
		},
		want:    users.User{},
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := users.NewService()
			for _, p := range tt.initialData {
				_, err := u.NewUser(tt.args.f, tt.args.r, p)
				if err != nil {
					t.Errorf("userService.NewUser() error = %v", err)
					return
				}
			}
			got, err := u.NewUser(tt.args.f, tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Name != tt.want.Name || got.Email != tt.want.Email {
				t.Errorf("userService.NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
