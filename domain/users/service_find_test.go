package users_test

import (
	"fickle/domain/users"
	"fickle/infrastructure/datasource/inmemory"
	"reflect"
	"testing"
)

func Test_userService_FindUser(t *testing.T) {
	r := inmemory.NewRepositoryUsers()
	r.AddUser(users.User{
		Id:           "1",
		Name:         "fickle",
		Email:        "test@fickle.com",
		PasswordHash: "hash",
	})
	r.AddUser(users.User{
		Id:           "2",
		Name:         "elkcif",
		Email:        "test2@elkcif.com",
		PasswordHash: "hash",
	})

	type args struct {
		r  users.IRepository
		id users.IdUser
	}
	tests := []struct {
		name    string
		args    args
		want    users.User
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				r:  r,
				id: "1",
			},
			want: users.User{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: "hash",
			},
			wantErr: false,
		},
		{
			name: "ok",
			args: args{
				r:  r,
				id: "2",
			},
			want: users.User{
				Id:           "2",
				Name:         "elkcif",
				Email:        "test2@elkcif.com",
				PasswordHash: "hash",
			},
			wantErr: false,
		},
		{
			name: "fail: not found",
			args: args{
				r:  r,
				id: "3",
			},
			want:    users.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := users.NewService()
			got, err := u.FindUser(tt.args.r, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.FindUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.FindUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_FindUsers(t *testing.T) {
	r := inmemory.NewRepositoryUsers()
	r.AddUser(users.User{
		Id:           "1",
		Name:         "fickle",
		Email:        "test@fickle.com",
		PasswordHash: "hash",
	})
	r.AddUser(users.User{
		Id:           "2",
		Name:         "elkcif",
		Email:        "test2@elkcif.com",
		PasswordHash: "hash",
	})
	r.AddUser(users.User{
		Id:           "3",
		Name:         "fickle2",
		Email:        "test2@fickle.com",
		PasswordHash: "hash",
	})

	type args struct {
		r users.IRepository
		p users.QueryUserParam
	}
	tests := []struct {
		name    string
		args    args
		want    []users.User
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				r: r,
				p: users.QueryUserParam{},
			},
			want: []users.User{
				{
					Id:           "1",
					Name:         "fickle",
					Email:        "test@fickle.com",
					PasswordHash: "hash",
				}, {
					Id:           "2",
					Name:         "elkcif",
					Email:        "test2@elkcif.com",
					PasswordHash: "hash",
				}, {
					Id:           "3",
					Name:         "fickle2",
					Email:        "test2@fickle.com",
					PasswordHash: "hash",
				},
			},
			wantErr: false,
		},
		{
			name: "ok: query 'Name'",
			args: args{
				r: r,
				p: users.QueryUserParam{
					Name: func() *string { s := "fickle2"; return &s }(),
				},
			},
			want: []users.User{
				{
					Id:           "3",
					Name:         "fickle2",
					Email:        "test2@fickle.com",
					PasswordHash: "hash",
				},
			},
			wantErr: false,
		},
		{
			name: "ok: query 'NameContain'",
			args: args{
				r: r,
				p: users.QueryUserParam{
					NameContain: func() *string { s := "fickle"; return &s }(),
				},
			},
			want: []users.User{
				{
					Id:           "1",
					Name:         "fickle",
					Email:        "test@fickle.com",
					PasswordHash: "hash",
				}, {
					Id:           "3",
					Name:         "fickle2",
					Email:        "test2@fickle.com",
					PasswordHash: "hash",
				},
			},
			wantErr: false,
		},
		{
			name: "ok: query 'Email'",
			args: args{
				r: r,
				p: users.QueryUserParam{
					Email: func() *string { s := "test2@elkcif.com"; return &s }(),
				},
			},
			want: []users.User{
				{
					Id:           "2",
					Name:         "elkcif",
					Email:        "test2@elkcif.com",
					PasswordHash: "hash",
				},
			},
			wantErr: false,
		},
		{
			name: "ok: query 'Name', 'NameContain' and 'Email'",
			args: args{
				r: r,
				p: users.QueryUserParam{
					Name:        func() *string { n := "fickle"; return &n }(),
					NameContain: func() *string { n := "lkci"; return &n }(),
					Email:       func() *string { n := "test@fickle.com"; return &n }(),
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "ok: query 'Name', 'NameContain' and 'Email'",
			args: args{
				r: r,
				p: users.QueryUserParam{
					Name:        func() *string { n := "fickle"; return &n }(),
					NameContain: func() *string { n := "ickl"; return &n }(),
					Email:       func() *string { n := "test@fickle.com"; return &n }(),
				},
			},
			want: []users.User{{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: "hash",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := users.NewService()
			got, err := u.FindUsers(tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.FindUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.FindUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
