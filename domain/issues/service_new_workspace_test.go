package issues_test

import (
	"fickle/domain/issues"
	"fickle/infrastructure/datasource/inmemory"
	"testing"
)

func Test_boardService_NewWorkspace(t *testing.T) {
	type args struct {
		f issues.IFactory
		r issues.IRepository
		p issues.CreateWorkspaceParam
	}
	tests := []struct {
		name    string
		args    args
		want    issues.Workspace
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				f: inmemory.NewFactoryIssues(),
				r: inmemory.NewRepositoryIssues(),
				p: issues.CreateWorkspaceParam{
					Name: "General",
					Columns: []issues.AddBoardColumnParam{{
						Name:    "Backlog",
						Color:   "#2c3e50",
						Hidden:  false,
						Order:   "1",
						Default: true,
						UserId:  "1",
					}, {
						Name:    "Next up",
						Color:   "#f39c12",
						Hidden:  false,
						Order:   "11",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "In Progress",
						Color:   "#3498db",
						Hidden:  false,
						Order:   "111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Pending",
						Color:   "#e67e22",
						Hidden:  false,
						Order:   "1111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Done",
						Color:   "#2ecc71",
						Hidden:  false,
						Order:   "11111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Archived",
						Color:   "#95a5a6",
						Hidden:  true,
						Order:   "111111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Canceled",
						Color:   "#e74c3c",
						Hidden:  true,
						Order:   "1111111",
						Default: false,
						UserId:  "1",
					}},
					UserId: "1",
				},
			},
			want: issues.Workspace{
				Name: "General",
				Columns: []issues.Column{{
					Name:    "Backlog",
					Color:   "#2c3e50",
					Hidden:  false,
					Order:   "1",
					Default: true,
					UserId:  "1",
				}, {
					Name:    "Next up",
					Color:   "#f39c12",
					Hidden:  false,
					Order:   "11",
					Default: false,
					UserId:  "1",
				}, {
					Name:    "In Progress",
					Color:   "#3498db",
					Hidden:  false,
					Order:   "111",
					Default: false,
					UserId:  "1",
				}, {
					Name:    "Pending",
					Color:   "#e67e22",
					Hidden:  false,
					Order:   "1111",
					Default: false,
					UserId:  "1",
				}, {
					Name:    "Done",
					Color:   "#2ecc71",
					Hidden:  false,
					Order:   "11111",
					Default: false,
					UserId:  "1",
				}, {
					Name:    "Archived",
					Color:   "#95a5a6",
					Hidden:  true,
					Order:   "111111",
					Default: false,
					UserId:  "1",
				}, {
					Name:    "Canceled",
					Color:   "#e74c3c",
					Hidden:  true,
					Order:   "1111111",
					Default: false,
					UserId:  "1",
				}},
				Archived: false,
				UserId:   "1",
			},
			wantErr: false,
		},
		{
			name: "ok",
			args: args{
				f: inmemory.NewFactoryIssues(),
				r: inmemory.NewRepositoryIssues(),
				p: issues.CreateWorkspaceParam{
					Name: "General",
					Columns: []issues.AddBoardColumnParam{{
						Name:    "Backlog",
						Color:   "#2c3e50",
						Hidden:  false,
						Order:   "1",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Next up",
						Color:   "#f39c12",
						Hidden:  false,
						Order:   "11",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "In Progress",
						Color:   "#3498db",
						Hidden:  false,
						Order:   "111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Pending",
						Color:   "#e67e22",
						Hidden:  false,
						Order:   "1111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Done",
						Color:   "#2ecc71",
						Hidden:  false,
						Order:   "11111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Archived",
						Color:   "#95a5a6",
						Hidden:  true,
						Order:   "111111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Canceled",
						Color:   "#e74c3c",
						Hidden:  true,
						Order:   "1111111",
						Default: false,
						UserId:  "1",
					}},
					UserId: "1",
				},
			},
			want: issues.Workspace{
				Name: "General",
				Columns: []issues.Column{{
					Name:    "Backlog",
					Color:   "#2c3e50",
					Hidden:  false,
					Order:   "1",
					Default: true,
					UserId:  "1",
				}, {
					Name:    "Next up",
					Color:   "#f39c12",
					Hidden:  false,
					Order:   "11",
					Default: false,
					UserId:  "1",
				}, {
					Name:    "In Progress",
					Color:   "#3498db",
					Hidden:  false,
					Order:   "111",
					Default: false,
					UserId:  "1",
				}, {
					Name:    "Pending",
					Color:   "#e67e22",
					Hidden:  false,
					Order:   "1111",
					Default: false,
					UserId:  "1",
				}, {
					Name:    "Done",
					Color:   "#2ecc71",
					Hidden:  false,
					Order:   "11111",
					Default: false,
					UserId:  "1",
				}, {
					Name:    "Archived",
					Color:   "#95a5a6",
					Hidden:  true,
					Order:   "111111",
					Default: false,
					UserId:  "1",
				}, {
					Name:    "Canceled",
					Color:   "#e74c3c",
					Hidden:  true,
					Order:   "1111111",
					Default: false,
					UserId:  "1",
				}},
				Archived: false,
				UserId:   "1",
			},
			wantErr: false,
		},
		{
			name: "failed to validate: 'Default' duplicated",
			args: args{
				f: inmemory.NewFactoryIssues(),
				r: inmemory.NewRepositoryIssues(),
				p: issues.CreateWorkspaceParam{
					Name: "General",
					Columns: []issues.AddBoardColumnParam{{
						Name:    "Backlog",
						Color:   "#2c3e50",
						Hidden:  false,
						Order:   "1",
						Default: true,
						UserId:  "1",
					}, {
						Name:    "Next up",
						Color:   "#f39c12",
						Hidden:  false,
						Order:   "11",
						Default: true,
						UserId:  "1",
					}, {
						Name:    "In Progress",
						Color:   "#3498db",
						Hidden:  false,
						Order:   "111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Pending",
						Color:   "#e67e22",
						Hidden:  false,
						Order:   "1111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Done",
						Color:   "#2ecc71",
						Hidden:  false,
						Order:   "11111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Archived",
						Color:   "#95a5a6",
						Hidden:  true,
						Order:   "111111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Canceled",
						Color:   "#e74c3c",
						Hidden:  true,
						Order:   "1111111",
						Default: false,
						UserId:  "1",
					}},
					UserId: "1",
				},
			},
			want:    issues.Workspace{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Name' empty",
			args: args{
				f: inmemory.NewFactoryIssues(),
				r: inmemory.NewRepositoryIssues(),
				p: issues.CreateWorkspaceParam{
					Name: "",
					Columns: []issues.AddBoardColumnParam{{
						Name:    "Backlog",
						Color:   "#2c3e50",
						Hidden:  false,
						Order:   "1",
						Default: true,
						UserId:  "1",
					}, {
						Name:    "Next up",
						Color:   "#f39c12",
						Hidden:  false,
						Order:   "11",
						Default: true,
						UserId:  "1",
					}, {
						Name:    "In Progress",
						Color:   "#3498db",
						Hidden:  false,
						Order:   "111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Pending",
						Color:   "#e67e22",
						Hidden:  false,
						Order:   "1111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Done",
						Color:   "#2ecc71",
						Hidden:  false,
						Order:   "11111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Archived",
						Color:   "#95a5a6",
						Hidden:  true,
						Order:   "111111",
						Default: false,
						UserId:  "1",
					}, {
						Name:    "Canceled",
						Color:   "#e74c3c",
						Hidden:  true,
						Order:   "1111111",
						Default: false,
						UserId:  "1",
					}},
					UserId: "1",
				},
			},
			want:    issues.Workspace{},
			wantErr: true,
		},
		{
			name: "failed to validate: 'Column[].Name' empty",
			args: args{
				f: inmemory.NewFactoryIssues(),
				r: inmemory.NewRepositoryIssues(),
				p: issues.CreateWorkspaceParam{
					Name: "",
					Columns: []issues.AddBoardColumnParam{{
						Name:    "",
						Color:   "#2c3e50",
						Hidden:  false,
						Order:   "1",
						Default: true,
						UserId:  "1",
					}},
					UserId: "1",
				},
			},
			want:    issues.Workspace{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := issues.NewService()
			got, err := s.NewWorkspace(tt.args.f, tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("boardService.NewWorkspace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Name != tt.want.Name || got.Archived != tt.want.Archived || got.UserId != tt.want.UserId || len(got.Columns) != len(tt.want.Columns) {
				t.Errorf("boardService.NewWorkspace() = %v, want %v", got, tt.want)
				return
			}
			for i := 0; i < len(got.Columns); i++ {
				g := got.Columns[i]
				w := tt.want.Columns[i]
				if g.Name != w.Name || g.Color != w.Color || g.Hidden != w.Hidden || g.Order != w.Order || g.Default != w.Default || g.UserId != w.UserId {
					t.Errorf("boardService.NewWorkspace() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
