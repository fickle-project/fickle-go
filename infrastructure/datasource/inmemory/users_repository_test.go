package inmemory

import (
	"fickle/domain/users"
	"reflect"
	"testing"
)

func Test_rUsers_AddUser(t *testing.T) {
	type fields struct {
		data []users.User
	}
	type args struct {
		u users.User
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       users.User
		wantFields fields
		wantErr    bool
	}{
		{
			name: "ok ",
			fields: fields{
				data: []users.User{},
			},
			args: args{
				u: users.User{
					Id:           "1",
					Name:         "fickle",
					Email:        "test@fickle.com",
					PasswordHash: "fickle",
				},
			},
			want: users.User{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: "fickle",
			},
			wantFields: fields{
				data: []users.User{{
					Id:           "1",
					Name:         "fickle",
					Email:        "test@fickle.com",
					PasswordHash: "fickle",
				}},
			},
			wantErr: false,
		},
		{
			name: "ok ",
			fields: fields{
				data: []users.User{{
					Id:           "1",
					Name:         "fickle",
					Email:        "test@fickle.com",
					PasswordHash: "fickle",
				}},
			},
			args: args{
				u: users.User{
					Id:           "2",
					Name:         "fickle2",
					Email:        "test2@fickle.com",
					PasswordHash: "fickle",
				},
			},
			want: users.User{
				Id:           "2",
				Name:         "fickle2",
				Email:        "test2@fickle.com",
				PasswordHash: "fickle",
			},
			wantFields: fields{
				data: []users.User{{
					Id:           "1",
					Name:         "fickle",
					Email:        "test@fickle.com",
					PasswordHash: "fickle",
				}, {
					Id:           "2",
					Name:         "fickle2",
					Email:        "test2@fickle.com",
					PasswordHash: "fickle",
				}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rUsers{
				data: tt.fields.data,
			}
			got, err := r.AddUser(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("rUsers.AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rUsers.AddUser() = %v, want %v", got, tt.want)
				return
			}
			data, err := r.FindUsers(users.QueryUserParam{})
			if err != nil {
				t.Errorf("rUsers.FindUsers() error = %v", err)
				return
			}
			if !reflect.DeepEqual(data, tt.wantFields.data) {
				t.Errorf("rUsers.AddUser(), rUsers.FindUsers() = %v, want %v", data, tt.wantFields.data)
				return
			}
		})
	}
}

func Test_rUsers_FindUser(t *testing.T) {
	type fields struct {
		data []users.User
	}
	type args struct {
		id users.IdUser
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
				data: []users.User{{
					Id:           "1",
					Name:         "fickle",
					Email:        "test@fickle.com",
					PasswordHash: "fickle",
				}},
			},
			args: args{
				id: "1",
			},
			want: users.User{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: "fickle",
			},
			wantErr: false,
		},
		{
			name: "not found",
			fields: fields{
				data: []users.User{},
			},
			args: args{
				id: "1",
			},
			want:    users.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rUsers{
				data: tt.fields.data,
			}
			got, err := r.FindUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("rUsers.FindUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rUsers.FindUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rUsers_FindUsers(t *testing.T) {
	type fields struct {
		data []users.User
	}
	type args struct {
		q users.QueryUserParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []users.User
		wantErr bool
	}{{
		name: "find all",
		fields: fields{
			data: []users.User{{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: "fickle",
			}, {
				Id:           "2",
				Name:         "elkcif",
				Email:        "test@elkcif.com",
				PasswordHash: "elkcif",
			}},
		},
		args: args{},
		want: []users.User{{
			Id:           "1",
			Name:         "fickle",
			Email:        "test@fickle.com",
			PasswordHash: "fickle",
		}, {
			Id:           "2",
			Name:         "elkcif",
			Email:        "test@elkcif.com",
			PasswordHash: "elkcif",
		}},
		wantErr: false,
	}, {
		name: "query by Name",
		fields: fields{
			data: []users.User{{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: "fickle",
			}, {
				Id:           "2",
				Name:         "elkcif",
				Email:        "test@elkcif.com",
				PasswordHash: "elkcif",
			}},
		},
		args: args{
			q: users.QueryUserParam{
				Name: func() *string { n := "fickle"; return &n }(),
			},
		},
		want: []users.User{{
			Id:           "1",
			Name:         "fickle",
			Email:        "test@fickle.com",
			PasswordHash: "fickle",
		}},
		wantErr: false,
	}, {
		name: "query by NameContain",
		fields: fields{
			data: []users.User{{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: "fickle",
			}, {
				Id:           "2",
				Name:         "elkcif",
				Email:        "test@elkcif.com",
				PasswordHash: "elkcif",
			}},
		},
		args: args{
			q: users.QueryUserParam{
				NameContain: func() *string {
					n := "lkci"
					return &n
				}(),
			},
		},
		want: []users.User{{
			Id:           "2",
			Name:         "elkcif",
			Email:        "test@elkcif.com",
			PasswordHash: "elkcif",
		}},
		wantErr: false,
	}, {
		name: "query by Email",
		fields: fields{
			data: []users.User{{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: "fickle",
			}, {
				Id:           "2",
				Name:         "elkcif",
				Email:        "test@elkcif.com",
				PasswordHash: "elkcif",
			}},
		},
		args: args{
			q: users.QueryUserParam{
				Email: func() *string {
					n := "test@fickle.com"
					return &n
				}(),
			},
		},
		want: []users.User{{
			Id:           "1",
			Name:         "fickle",
			Email:        "test@fickle.com",
			PasswordHash: "fickle",
		}},
	}, {
		name: "query by Name & NameContain & Email 1",
		fields: fields{
			data: []users.User{{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: "fickle",
			}, {
				Id:           "2",
				Name:         "elkcif",
				Email:        "test@elkcif.com",
				PasswordHash: "elkcif",
			}},
		},
		args: args{
			q: users.QueryUserParam{
				Name:        func() *string { n := "fickle"; return &n }(),
				NameContain: func() *string { n := "ickl"; return &n }(),
				Email:       func() *string { n := "test@fickle.com"; return &n }(),
			},
		},
		want: []users.User{{
			Id:           "1",
			Name:         "fickle",
			Email:        "test@fickle.com",
			PasswordHash: "fickle",
		}},
		wantErr: false,
	}, {
		name: "query by Name & NameContain & Email 2",
		fields: fields{
			data: []users.User{{
				Id:           "1",
				Name:         "fickle",
				Email:        "test@fickle.com",
				PasswordHash: "fickle",
			}, {
				Id:           "2",
				Name:         "elkcif",
				Email:        "test@elkcif.com",
				PasswordHash: "elkcif",
			}},
		},
		args: args{
			q: users.QueryUserParam{
				Name:        func() *string { n := "fickle"; return &n }(),
				NameContain: func() *string { n := "lkci"; return &n }(),
				Email:       func() *string { n := "test@fickle.com"; return &n }(),
			},
		},
		want:    nil,
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rUsers{
				data: tt.fields.data,
			}
			got, err := r.FindUsers(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("rUsers.FindUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rUsers.FindUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rUsers_RemoveUser(t *testing.T) {
	type fields struct {
		data []users.User
	}
	type args struct {
		id users.IdUser
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantFields fields
		wantErr    bool
	}{
		{
			name: "ok",
			fields: fields{
				data: []users.User{{
					Id:           "1",
					Name:         "fickle",
					Email:        "test@fickle.com",
					PasswordHash: "fickle",
				}, {
					Id:           "2",
					Name:         "elkcif",
					Email:        "test@elkcif.com",
					PasswordHash: "elkcif",
				}},
			},
			args: args{
				id: "1",
			},
			wantFields: fields{
				data: []users.User{{
					Id:           "2",
					Name:         "elkcif",
					Email:        "test@elkcif.com",
					PasswordHash: "elkcif",
				}},
			},
			wantErr: false,
		},
		{
			name: "not found",
			fields: fields{
				data: []users.User{},
			},
			args: args{
				id: "1",
			},
			wantFields: fields{
				data: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rUsers{
				data: tt.fields.data,
			}
			if err := r.RemoveUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("rUsers.RemoveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			data, err := r.FindUsers(users.QueryUserParam{})
			if err != nil {
				t.Errorf("rUsers.FindUsers() error = %v", err)
				return
			}
			if !reflect.DeepEqual(data, tt.wantFields.data) {
				t.Errorf("rUsers.RemoveUser(), rUsers.FindUsers() = %v, want %v", data, tt.wantFields.data)
				return
			}
		})
	}
}

func Test_rUsers_UpdateUser(t *testing.T) {
	type fields struct {
		data []users.User
	}
	type args struct {
		id users.IdUser
		p  users.UpdateUserParam
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       users.User
		wantFields fields
		wantErr    bool
	}{
		{
			name: "ok",
			fields: fields{
				data: []users.User{{
					Id:           "1",
					Name:         "fickle",
					Email:        "test@fickle.com",
					PasswordHash: "fickle",
				}, {
					Id:           "2",
					Name:         "elkcif",
					Email:        "test@elkcif.com",
					PasswordHash: "elkcif",
				}},
			},
			args: args{
				id: "1",
				p: users.UpdateUserParam{
					Name: func() *string { n := "updated"; return &n }(),
				},
			},
			want: users.User{
				Id:           "1",
				Name:         "updated",
				Email:        "test@fickle.com",
				PasswordHash: "fickle",
			},
			wantFields: fields{
				data: []users.User{{
					Id:           "2",
					Name:         "elkcif",
					Email:        "test@elkcif.com",
					PasswordHash: "elkcif",
				}, {
					Id:           "1",
					Name:         "updated",
					Email:        "test@fickle.com",
					PasswordHash: "fickle",
				}},
			},
			wantErr: false,
		},
		{
			name: "not found",
			fields: fields{
				data: []users.User{},
			},
			args: args{
				id: "1",
				p: users.UpdateUserParam{
					Name: func() *string { n := "updated"; return &n }(),
				},
			},
			want: users.User{},
			wantFields: fields{
				data: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rUsers{
				data: tt.fields.data,
			}
			got, err := r.UpdateUser(tt.args.id, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("rUsers.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rUsers.UpdateUser() = %v, want %v", got, tt.want)
			}
			data, err := r.FindUsers(users.QueryUserParam{})
			if err != nil {
				t.Errorf("rUsers.FindUsers() error = %v", err)
				return
			}
			if !reflect.DeepEqual(data, tt.wantFields.data) {
				t.Errorf("rUsers.UpdateUser(), rUsers.FindUsers() = %v, want %v", data, tt.wantFields.data)
				return
			}
		})
	}
}
