package issues_test

import (
	"fickle/domain/issues"
	"fickle/infrastructure/datasource/inmemory"
	"reflect"
	"testing"
)

func newRepository() issues.IRepository {
	r := inmemory.NewRepositoryIssues()
	r.CreateWorkspace(issues.Workspace{
		Id:   "1",
		Name: "general",
		Columns: []issues.Column{
			{
				Id:          "1",
				Name:        "Backlog",
				Color:       "#2c3e50",
				Hidden:      false,
				Order:       "1",
				Default:     true,
				WorkspaceId: "1",
				UserId:      "1",
			}, {
				Id:          "2",
				Name:        "Next up",
				Color:       "#f39c12",
				Hidden:      false,
				Order:       "11",
				Default:     false,
				WorkspaceId: "1",
				UserId:      "1",
			},
		},
		Archived: false,
		UserId:   "1",
	})
	r.CreateWorkspace(issues.Workspace{
		Id:   "2",
		Name: "ws 2",
		Columns: []issues.Column{
			{
				Id:          "3",
				Name:        "Backlog",
				Color:       "#2c3e50",
				Hidden:      false,
				Order:       "1",
				Default:     true,
				WorkspaceId: "2",
				UserId:      "1",
			}, {
				Id:          "4",
				Name:        "Next up",
				Color:       "#f39c12",
				Hidden:      false,
				Order:       "11",
				Default:     false,
				WorkspaceId: "2",
				UserId:      "1",
			},
		},
		Archived: false,
		UserId:   "1",
	})
	r.CreateBoard(issues.Board{
		Id:          "1",
		Name:        "fickle",
		WorkspaceId: "1",
		Archived:    false,
		UserId:      "1",
	})
	r.CreateBoard(issues.Board{
		Id:          "2",
		Name:        "fickle-client",
		WorkspaceId: "1",
		Archived:    true,
		UserId:      "1",
	})
	r.CreateIssue(issues.Issue{
		Id:          "1",
		Name:        "issue 1",
		BoardId:     "1",
		ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
		Order:       "1",
		WorkspaceId: "1",
		UserId:      "1",
	})
	r.CreateIssue(issues.Issue{
		Id:          "2",
		Name:        "issue 2",
		BoardId:     "1",
		ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
		Order:       "1",
		WorkspaceId: "1",
		UserId:      "1",
	})

	return r
}

func Test_boardService_FindBoard(t *testing.T) {
	r := newRepository()

	type args struct {
		r  issues.IRepository
		id issues.IdBoard
		p  issues.QueryBoardParam
	}
	tests := []struct {
		name    string
		args    args
		want    issues.BoardWithEmbedDatas
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				r:  r,
				id: "1",
				p:  issues.QueryBoardParam{},
			},
			want: issues.BoardWithEmbedDatas{
				Board: issues.Board{
					Id:          "1",
					Name:        "fickle",
					WorkspaceId: "1",
					Archived:    false,
					UserId:      "1",
				},
			},
			wantErr: false,
		},
		{
			name: "ok: embed workspace",
			args: args{
				r:  r,
				id: "1",
				p: issues.QueryBoardParam{
					Embed: []string{"workspace"},
				},
			},
			want: issues.BoardWithEmbedDatas{
				Board: issues.Board{Id: "1", Name: "fickle", WorkspaceId: "1", Archived: false, UserId: "1"},
				Workspace: issues.Workspace{
					Id:   "1",
					Name: "general",
					Columns: []issues.Column{
						{
							Id:          "1",
							Name:        "Backlog",
							Color:       "#2c3e50",
							Hidden:      false,
							Order:       "1",
							Default:     true,
							WorkspaceId: "1",
							UserId:      "1",
						}, {
							Id:          "2",
							Name:        "Next up",
							Color:       "#f39c12",
							Hidden:      false,
							Order:       "11",
							Default:     false,
							WorkspaceId: "1",
							UserId:      "1",
						},
					},
					Archived: false,
					UserId:   "1",
				},
			},
			wantErr: false,
		},
		{
			name: "ok: embed issues and workspace",
			args: args{
				r:  r,
				id: "1",
				p: issues.QueryBoardParam{
					Embed: []string{"issues", "workspace"},
				},
			},
			want: issues.BoardWithEmbedDatas{
				Board: issues.Board{Id: "1", Name: "fickle", WorkspaceId: "1", Archived: false, UserId: "1"},
				Columns: []issues.ColumnWithEmbedDatas{{
					Column: issues.Column{
						Id:          "1",
						Name:        "Backlog",
						Color:       "#2c3e50",
						Hidden:      false,
						Order:       "1",
						Default:     true,
						WorkspaceId: "1",
						UserId:      "1",
					},
					Issues: []issues.IssueWithEmbedDatas{
						{
							Issue: issues.Issue{
								Id:          "1",
								Name:        "issue 1",
								BoardId:     "1",
								ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
								Order:       "1",
								WorkspaceId: "1",
								UserId:      "1",
							},
						}, {
							Issue: issues.Issue{
								Id:          "2",
								Name:        "issue 2",
								BoardId:     "1",
								ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
								Order:       "1",
								WorkspaceId: "1",
								UserId:      "1",
							},
						},
					},
				}, {
					Column: issues.Column{
						Id:          "2",
						Name:        "Next up",
						Color:       "#f39c12",
						Hidden:      false,
						Order:       "11",
						Default:     false,
						WorkspaceId: "1",
						UserId:      "1",
					},
					Issues: []issues.IssueWithEmbedDatas{},
				}},
				Workspace: issues.Workspace{
					Id:   "1",
					Name: "general",
					Columns: []issues.Column{
						{
							Id:          "1",
							Name:        "Backlog",
							Color:       "#2c3e50",
							Hidden:      false,
							Order:       "1",
							Default:     true,
							WorkspaceId: "1",
							UserId:      "1",
						}, {
							Id:          "2",
							Name:        "Next up",
							Color:       "#f39c12",
							Hidden:      false,
							Order:       "11",
							Default:     false,
							WorkspaceId: "1",
							UserId:      "1",
						},
					},
					Archived: false,
					UserId:   "1",
				},
			},
			wantErr: false,
		},
		{
			name: "fail: not found",
			args: args{
				r:  r,
				id: "0",
				p:  issues.QueryBoardParam{},
			},
			want:    issues.BoardWithEmbedDatas{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := issues.NewService()
			got, err := b.FindBoard(tt.args.r, tt.args.id, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("boardService.FindBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boardService.FindBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardService_FindBoards(t *testing.T) {
	r := newRepository()

	type args struct {
		r issues.IRepository
		p issues.QueryBoardParam
	}
	tests := []struct {
		name    string
		args    args
		want    []issues.BoardWithEmbedDatas
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				r: r,
				p: issues.QueryBoardParam{},
			},
			want: []issues.BoardWithEmbedDatas{
				{
					Board: issues.Board{
						Id:          "1",
						Name:        "fickle",
						WorkspaceId: "1",
						Archived:    false,
						UserId:      "1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ok: embed issues and workspace",
			args: args{
				r: r,
				p: issues.QueryBoardParam{
					Embed: []string{"issues", "workspace"},
				},
			},
			want: []issues.BoardWithEmbedDatas{
				{
					Board: issues.Board{Id: "1", Name: "fickle", WorkspaceId: "1", Archived: false, UserId: "1"},
					Columns: []issues.ColumnWithEmbedDatas{{
						Column: issues.Column{
							Id:          "1",
							Name:        "Backlog",
							Color:       "#2c3e50",
							Hidden:      false,
							Order:       "1",
							Default:     true,
							WorkspaceId: "1",
							UserId:      "1",
						},
						Issues: []issues.IssueWithEmbedDatas{
							{
								Issue: issues.Issue{
									Id:          "1",
									Name:        "issue 1",
									BoardId:     "1",
									ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
									Order:       "1",
									WorkspaceId: "1",
									UserId:      "1",
								},
							}, {
								Issue: issues.Issue{
									Id:          "2",
									Name:        "issue 2",
									BoardId:     "1",
									ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
									Order:       "1",
									WorkspaceId: "1",
									UserId:      "1",
								},
							},
						},
					}, {
						Column: issues.Column{
							Id:          "2",
							Name:        "Next up",
							Color:       "#f39c12",
							Hidden:      false,
							Order:       "11",
							Default:     false,
							WorkspaceId: "1",
							UserId:      "1",
						},
						Issues: []issues.IssueWithEmbedDatas{},
					}},
					Workspace: issues.Workspace{
						Id:   "1",
						Name: "general",
						Columns: []issues.Column{
							{
								Id:          "1",
								Name:        "Backlog",
								Color:       "#2c3e50",
								Hidden:      false,
								Order:       "1",
								Default:     true,
								WorkspaceId: "1",
								UserId:      "1",
							}, {
								Id:          "2",
								Name:        "Next up",
								Color:       "#f39c12",
								Hidden:      false,
								Order:       "11",
								Default:     false,
								WorkspaceId: "1",
								UserId:      "1",
							},
						},
						Archived: false,
						UserId:   "1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ok: include archived",
			args: args{
				r: r,
				p: issues.QueryBoardParam{
					IncludeArchived: true,
				},
			},
			want: []issues.BoardWithEmbedDatas{
				{
					Board: issues.Board{
						Id:          "1",
						Name:        "fickle",
						WorkspaceId: "1",
						Archived:    false,
						UserId:      "1",
					},
				},
				{
					Board: issues.Board{
						Id:          "2",
						Name:        "fickle-client",
						WorkspaceId: "1",
						Archived:    true,
						UserId:      "1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ok: name equal",
			args: args{
				r: r,
				p: issues.QueryBoardParam{
					Name: func() *string { s := "fickle"; return &s }(),
				},
			},
			want: []issues.BoardWithEmbedDatas{
				{
					Board: issues.Board{
						Id:          "1",
						Name:        "fickle",
						WorkspaceId: "1",
						Archived:    false,
						UserId:      "1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ok: name contain",
			args: args{
				r: r,
				p: issues.QueryBoardParam{
					NameContain:     func() *string { s := "ickl"; return &s }(),
					IncludeArchived: true,
				},
			},
			want: []issues.BoardWithEmbedDatas{
				{
					Board: issues.Board{
						Id:          "1",
						Name:        "fickle",
						WorkspaceId: "1",
						Archived:    false,
						UserId:      "1",
					},
				},
				{
					Board: issues.Board{
						Id:          "2",
						Name:        "fickle-client",
						WorkspaceId: "1",
						Archived:    true,
						UserId:      "1",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := issues.NewService()
			got, err := b.FindBoards(tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("boardService.FindBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boardService.FindBoards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardService_FindIssue(t *testing.T) {
	r := newRepository()

	type args struct {
		r  issues.IRepository
		id issues.IdIssue
		p  issues.QueryIssueParam
	}
	tests := []struct {
		name    string
		args    args
		want    issues.IssueWithEmbedDatas
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				r:  r,
				id: "1",
				p:  issues.QueryIssueParam{},
			},
			want: issues.IssueWithEmbedDatas{
				Issue: issues.Issue{
					Id:          "1",
					Name:        "issue 1",
					BoardId:     "1",
					ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
					Order:       "1",
					WorkspaceId: "1",
					UserId:      "1",
				},
			},
			wantErr: false,
		},
		{
			name: "ok: embed column, baord and workspace",
			args: args{
				r:  r,
				id: "1",
				p: issues.QueryIssueParam{
					Embed: []string{"column", "board", "workspace"},
				},
			},
			want: issues.IssueWithEmbedDatas{
				Issue: issues.Issue{
					Id:          "1",
					Name:        "issue 1",
					BoardId:     "1",
					ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
					Order:       "1",
					WorkspaceId: "1",
					UserId:      "1",
				},
				Column: issues.Column{
					Id:          "1",
					Name:        "Backlog",
					Color:       "#2c3e50",
					Hidden:      false,
					Order:       "1",
					Default:     true,
					WorkspaceId: "1",
					UserId:      "1",
				},
				Board: issues.Board{
					Id:          "1",
					Name:        "fickle",
					WorkspaceId: "1",
					Archived:    false,
					UserId:      "1",
				},
				Workspace: issues.Workspace{
					Id:   "1",
					Name: "general",
					Columns: []issues.Column{
						{
							Id:          "1",
							Name:        "Backlog",
							Color:       "#2c3e50",
							Hidden:      false,
							Order:       "1",
							Default:     true,
							WorkspaceId: "1",
							UserId:      "1",
						}, {
							Id:          "2",
							Name:        "Next up",
							Color:       "#f39c12",
							Hidden:      false,
							Order:       "11",
							Default:     false,
							WorkspaceId: "1",
							UserId:      "1",
						},
					},
					Archived: false,
					UserId:   "1",
				},
			},
			wantErr: false,
		},
		{
			name: "fail: not found",
			args: args{
				r:  r,
				id: "0",
				p:  issues.QueryIssueParam{},
			},
			want:    issues.IssueWithEmbedDatas{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := issues.NewService()
			got, err := b.FindIssue(tt.args.r, tt.args.id, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("boardService.FindIssue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boardService.FindIssue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardService_FindIssues(t *testing.T) {
	r := newRepository()

	type args struct {
		r issues.IRepository
		p issues.QueryIssueParam
	}
	tests := []struct {
		name    string
		args    args
		want    []issues.IssueWithEmbedDatas
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				r: r,
				p: issues.QueryIssueParam{
					BoardId: func() *issues.IdBoard { var id issues.IdBoard = "1"; return &id }(),
				},
			},
			want: []issues.IssueWithEmbedDatas{
				{
					Issue: issues.Issue{
						Id:          "1",
						Name:        "issue 1",
						BoardId:     "1",
						ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
						Order:       "1",
						WorkspaceId: "1",
						UserId:      "1",
					},
				}, {
					Issue: issues.Issue{
						Id:          "2",
						Name:        "issue 2",
						BoardId:     "1",
						ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
						Order:       "1",
						WorkspaceId: "1",
						UserId:      "1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ok",
			args: args{
				r: r,
				p: issues.QueryIssueParam{
					BoardId: func() *issues.IdBoard { var id issues.IdBoard = "2"; return &id }(),
				},
			},
			want:    []issues.IssueWithEmbedDatas{},
			wantErr: false,
		},
		{
			name: "ok: embed column, baord and workspace",
			args: args{
				r: r,
				p: issues.QueryIssueParam{
					Embed: []string{"column", "board", "workspace"},
				},
			},
			want: []issues.IssueWithEmbedDatas{
				{
					Issue: issues.Issue{
						Id:          "1",
						Name:        "issue 1",
						BoardId:     "1",
						ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
						Order:       "1",
						WorkspaceId: "1",
						UserId:      "1",
					},

					Column: issues.Column{
						Id:          "1",
						Name:        "Backlog",
						Color:       "#2c3e50",
						Hidden:      false,
						Order:       "1",
						Default:     true,
						WorkspaceId: "1",
						UserId:      "1",
					},
					Board: issues.Board{
						Id:          "1",
						Name:        "fickle",
						WorkspaceId: "1",
						Archived:    false,
						UserId:      "1",
					},
					Workspace: issues.Workspace{
						Id:   "1",
						Name: "general",
						Columns: []issues.Column{
							{
								Id:          "1",
								Name:        "Backlog",
								Color:       "#2c3e50",
								Hidden:      false,
								Order:       "1",
								Default:     true,
								WorkspaceId: "1",
								UserId:      "1",
							}, {
								Id:          "2",
								Name:        "Next up",
								Color:       "#f39c12",
								Hidden:      false,
								Order:       "11",
								Default:     false,
								WorkspaceId: "1",
								UserId:      "1",
							},
						},
						Archived: false,
						UserId:   "1",
					},
				}, {
					Issue: issues.Issue{
						Id:          "2",
						Name:        "issue 2",
						BoardId:     "1",
						ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
						Order:       "1",
						WorkspaceId: "1",
						UserId:      "1",
					},

					Column: issues.Column{
						Id:          "1",
						Name:        "Backlog",
						Color:       "#2c3e50",
						Hidden:      false,
						Order:       "1",
						Default:     true,
						WorkspaceId: "1",
						UserId:      "1",
					},
					Board: issues.Board{
						Id:          "1",
						Name:        "fickle",
						WorkspaceId: "1",
						Archived:    false,
						UserId:      "1",
					},
					Workspace: issues.Workspace{
						Id:   "1",
						Name: "general",
						Columns: []issues.Column{
							{
								Id:          "1",
								Name:        "Backlog",
								Color:       "#2c3e50",
								Hidden:      false,
								Order:       "1",
								Default:     true,
								WorkspaceId: "1",
								UserId:      "1",
							}, {
								Id:          "2",
								Name:        "Next up",
								Color:       "#f39c12",
								Hidden:      false,
								Order:       "11",
								Default:     false,
								WorkspaceId: "1",
								UserId:      "1",
							},
						},
						Archived: false,
						UserId:   "1",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := issues.NewService()
			got, err := b.FindIssues(tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("boardService.FindIssues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boardService.FindIssues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardService_FindWorkspace(t *testing.T) {
	r := newRepository()

	type args struct {
		r  issues.IRepository
		id issues.IdWorkspace
		p  issues.QueryWorkspaceParam
	}
	tests := []struct {
		name    string
		args    args
		want    issues.WorkspaceWithEmbedDatas
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				r:  r,
				id: "1",
				p:  issues.QueryWorkspaceParam{},
			},
			want: issues.WorkspaceWithEmbedDatas{
				Workspace: issues.Workspace{
					Id:   "1",
					Name: "general",
					Columns: []issues.Column{
						{
							Id:          "1",
							Name:        "Backlog",
							Color:       "#2c3e50",
							Hidden:      false,
							Order:       "1",
							Default:     true,
							WorkspaceId: "1",
							UserId:      "1",
						}, {
							Id:          "2",
							Name:        "Next up",
							Color:       "#f39c12",
							Hidden:      false,
							Order:       "11",
							Default:     false,
							WorkspaceId: "1",
							UserId:      "1",
						},
					},
					Archived: false,
					UserId:   "1",
				},
			},
			wantErr: false,
		},
		{
			name: "ok: embed boards and issues",
			args: args{
				r:  r,
				id: "1",
				p: issues.QueryWorkspaceParam{
					Embed:           []string{"boards", "issues"},
					IncludeArchived: true,
				},
			},
			want: issues.WorkspaceWithEmbedDatas{
				Workspace: issues.Workspace{Id: "1", Name: "general", Columns: []issues.Column{{Id: "1", Name: "Backlog", Color: "#2c3e50", Hidden: false, Order: "1", Default: true, WorkspaceId: "1", UserId: "1"}, {Id: "2", Name: "Next up", Color: "#f39c12", Hidden: false, Order: "11", Default: false, WorkspaceId: "1", UserId: "1"}}, Archived: false, UserId: "1"},
				Boards: []issues.BoardWithEmbedDatas{{
					Board: issues.Board{
						Id:          "1",
						Name:        "fickle",
						WorkspaceId: "1",
						Archived:    false,
						UserId:      "1",
					},
					Columns: []issues.ColumnWithEmbedDatas{
						{
							Column: issues.Column{
								Id:          "1",
								Name:        "Backlog",
								Color:       "#2c3e50",
								Hidden:      false,
								Order:       "1",
								Default:     true,
								WorkspaceId: "1",
								UserId:      "1",
							},
							Issues: []issues.IssueWithEmbedDatas{
								{
									Issue: issues.Issue{
										Id:          "1",
										Name:        "issue 1",
										BoardId:     "1",
										ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
										Order:       "1",
										WorkspaceId: "1",
										UserId:      "1",
									},
								}, {
									Issue: issues.Issue{
										Id:          "2",
										Name:        "issue 2",
										BoardId:     "1",
										ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
										Order:       "1",
										WorkspaceId: "1",
										UserId:      "1",
									},
								},
							},
						}, {
							Column: issues.Column{
								Id:          "2",
								Name:        "Next up",
								Color:       "#f39c12",
								Hidden:      false,
								Order:       "11",
								Default:     false,
								WorkspaceId: "1",
								UserId:      "1",
							},
							Issues: []issues.IssueWithEmbedDatas{},
						},
					},
				}, {
					Board: issues.Board{
						Id:          "2",
						Name:        "fickle-client",
						WorkspaceId: "1",
						Archived:    true,
						UserId:      "1",
					},
					Columns: []issues.ColumnWithEmbedDatas{
						{
							Column: issues.Column{
								Id:          "1",
								Name:        "Backlog",
								Color:       "#2c3e50",
								Hidden:      false,
								Order:       "1",
								Default:     true,
								WorkspaceId: "1",
								UserId:      "1",
							},
						}, {
							Column: issues.Column{
								Id:          "2",
								Name:        "Next up",
								Color:       "#f39c12",
								Hidden:      false,
								Order:       "11",
								Default:     false,
								WorkspaceId: "1",
								UserId:      "1",
							},
						},
					},
					Workspace: issues.Workspace{},
				}},
			},
			wantErr: false,
		},
		{
			name: "fail: not found",
			args: args{
				r:  r,
				id: "0",
				p:  issues.QueryWorkspaceParam{},
			},
			want:    issues.WorkspaceWithEmbedDatas{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := issues.NewService()
			got, err := b.FindWorkspace(tt.args.r, tt.args.id, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("boardService.FindWorkspace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got.Boards) != len(tt.want.Boards) {
				t.Errorf("boardService.FindWorkspace().Boards = %v, want %v", got, tt.want)
				return
			}
			for i := 0; i < len(got.Boards); i++ {
				g := got.Boards[i]
				w := tt.want.Boards[i]
				if len(g.Columns) != len(w.Columns) {
					t.Errorf("boardService.FindWorkspace().Columns = %v, want %v", g.Columns, w.Columns)
					return
				}
				for j := 0; j < len(g.Columns); j++ {
					if !reflect.DeepEqual(g.Columns[j].Column, w.Columns[j].Column) {
						t.Errorf("boardService.FindWorkspace().Columns[].Column = %v, want %v", g.Columns[j].Column, w.Columns[j].Column)
						return
					}
					if !reflect.DeepEqual(g.Columns[j].Workspace, w.Columns[j].Workspace) {
						t.Errorf("boardService.FindWorkspace().Column[].Workspace = %v, want %v", g.Columns[j].Workspace, w.Columns[j].Workspace)
						return
					}
					if len(g.Columns[j].Issues) != len(w.Columns[j].Issues) {
						for k := 0; k < len(g.Columns[j].Issues); k++ {
							if !reflect.DeepEqual(g.Columns[j].Issues[k], w.Columns[j].Issues[k]) {
								t.Errorf("boardService.FindWorkspace().Column[].Issues = %v, want %v", g.Columns[j].Issues[k], w.Columns[j].Issues[k])
								return
							}
						}
					}
				}
				if !reflect.DeepEqual(g.Board, w.Board) {
					t.Errorf("boardService.FindWorkspace().Boards[].Board = %v, want %v", g.Board, w.Board)
					return
				}
				if !reflect.DeepEqual(g.Workspace, w.Workspace) {
					t.Errorf("boardService.FindWorkspace().Workspace = %v, want %v", g.Workspace, w.Workspace)
					return
				}
			}
			if !reflect.DeepEqual(got.Workspace, tt.want.Workspace) {
				t.Errorf("boardService.FindWorkspace().Workspace = %v, want %v", got.Workspace, tt.want.Workspace)
			}
		})
	}
}

func Test_boardService_FindWorkspaces(t *testing.T) {
	r := newRepository()

	type args struct {
		r issues.IRepository
		p issues.QueryWorkspaceParam
	}
	tests := []struct {
		name    string
		args    args
		want    []issues.WorkspaceWithEmbedDatas
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				r: r,
				p: issues.QueryWorkspaceParam{},
			},
			want: []issues.WorkspaceWithEmbedDatas{
				{
					Workspace: issues.Workspace{
						Id:   "1",
						Name: "general",
						Columns: []issues.Column{
							{
								Id:          "1",
								Name:        "Backlog",
								Color:       "#2c3e50",
								Hidden:      false,
								Order:       "1",
								Default:     true,
								WorkspaceId: "1",
								UserId:      "1",
							}, {
								Id:          "2",
								Name:        "Next up",
								Color:       "#f39c12",
								Hidden:      false,
								Order:       "11",
								Default:     false,
								WorkspaceId: "1",
								UserId:      "1",
							},
						},
						Archived: false,
						UserId:   "1",
					},
				}, {
					Workspace: issues.Workspace{
						Id:   "2",
						Name: "ws 2",
						Columns: []issues.Column{
							{
								Id:          "3",
								Name:        "Backlog",
								Color:       "#2c3e50",
								Hidden:      false,
								Order:       "1",
								Default:     true,
								WorkspaceId: "2",
								UserId:      "1",
							}, {
								Id:          "4",
								Name:        "Next up",
								Color:       "#f39c12",
								Hidden:      false,
								Order:       "11",
								Default:     false,
								WorkspaceId: "2",
								UserId:      "1",
							},
						},
						Archived: false,
						UserId:   "1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ok: embed boards and issues",
			args: args{
				r: r,
				p: issues.QueryWorkspaceParam{
					Embed:           []string{"boards", "issues"},
					IncludeArchived: true,
				},
			},
			want: []issues.WorkspaceWithEmbedDatas{
				{
					Workspace: issues.Workspace{Id: "1", Name: "general", Columns: []issues.Column{{Id: "1", Name: "Backlog", Color: "#2c3e50", Hidden: false, Order: "1", Default: true, WorkspaceId: "1", UserId: "1"}, {Id: "2", Name: "Next up", Color: "#f39c12", Hidden: false, Order: "11", Default: false, WorkspaceId: "1", UserId: "1"}}, Archived: false, UserId: "1"},
					Boards: []issues.BoardWithEmbedDatas{
						{
							Board: issues.Board{
								Id:          "1",
								Name:        "fickle",
								WorkspaceId: "1",
								Archived:    false,
								UserId:      "1",
							},
							Columns: []issues.ColumnWithEmbedDatas{
								{
									Column: issues.Column{
										Id:          "1",
										Name:        "Backlog",
										Color:       "#2c3e50",
										Hidden:      false,
										Order:       "1",
										Default:     true,
										WorkspaceId: "1",
										UserId:      "1",
									},
									Issues: []issues.IssueWithEmbedDatas{
										{
											Issue: issues.Issue{
												Id:          "1",
												Name:        "issue 1",
												BoardId:     "1",
												ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
												Order:       "1",
												WorkspaceId: "1",
												UserId:      "1",
											},
										}, {
											Issue: issues.Issue{
												Id:          "2",
												Name:        "issue 2",
												BoardId:     "1",
												ColumnId:    func() *issues.IdColumn { var id issues.IdColumn = "1"; return &id }(),
												Order:       "1",
												WorkspaceId: "1",
												UserId:      "1",
											},
										},
									},
								}, {
									Column: issues.Column{
										Id:          "2",
										Name:        "Next up",
										Color:       "#f39c12",
										Hidden:      false,
										Order:       "11",
										Default:     false,
										WorkspaceId: "1",
										UserId:      "1",
									},
									Issues: []issues.IssueWithEmbedDatas{},
								},
							},
						}, {
							Board: issues.Board{
								Id:          "2",
								Name:        "fickle-client",
								WorkspaceId: "1",
								Archived:    true,
								UserId:      "1",
							},
							Columns: []issues.ColumnWithEmbedDatas{
								{
									Column: issues.Column{
										Id:          "1",
										Name:        "Backlog",
										Color:       "#2c3e50",
										Hidden:      false,
										Order:       "1",
										Default:     true,
										WorkspaceId: "1",
										UserId:      "1",
									},
								}, {
									Column: issues.Column{
										Id:          "2",
										Name:        "Next up",
										Color:       "#f39c12",
										Hidden:      false,
										Order:       "11",
										Default:     false,
										WorkspaceId: "1",
										UserId:      "1",
									},
								},
							},
							Workspace: issues.Workspace{},
						},
					},
				},
				{
					Workspace: issues.Workspace{
						Id:   "2",
						Name: "ws 2",
						Columns: []issues.Column{
							{
								Id:          "3",
								Name:        "Backlog",
								Color:       "#2c3e50",
								Hidden:      false,
								Order:       "1",
								Default:     true,
								WorkspaceId: "2",
								UserId:      "1",
							}, {
								Id:          "4",
								Name:        "Next up",
								Color:       "#f39c12",
								Hidden:      false,
								Order:       "11",
								Default:     false,
								WorkspaceId: "2",
								UserId:      "1",
							},
						},
						Archived: false,
						UserId:   "1",
					},
					Boards: nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := issues.NewService()
			got, err := b.FindWorkspaces(tt.args.r, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("boardService.FindWorkspaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("boardService.FindWorkspaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for h := 0; h < len(got); h++ {
				if len(got[h].Boards) != len(tt.want[h].Boards) {
					t.Errorf("boardService.FindWorkspace().Boards = %v, want %v", got[h], tt.want)
					return
				}
				for i := 0; i < len(got[h].Boards); i++ {
					g := got[h].Boards[i]
					w := tt.want[h].Boards[i]
					if len(g.Columns) != len(w.Columns) {
						t.Errorf("boardService.FindWorkspace().Columns = %v, want %v", g.Columns, w.Columns)
						return
					}
					for j := 0; j < len(g.Columns); j++ {
						if !reflect.DeepEqual(g.Columns[j].Column, w.Columns[j].Column) {
							t.Errorf("boardService.FindWorkspace().Columns[].Column = %v, want %v", g.Columns[j].Column, w.Columns[j].Column)
							return
						}
						if !reflect.DeepEqual(g.Columns[j].Workspace, w.Columns[j].Workspace) {
							t.Errorf("boardService.FindWorkspace().Column[].Workspace = %v, want %v", g.Columns[j].Workspace, w.Columns[j].Workspace)
							return
						}
						if len(g.Columns[j].Issues) != len(w.Columns[j].Issues) {
							for k := 0; k < len(g.Columns[j].Issues); k++ {
								if !reflect.DeepEqual(g.Columns[j].Issues[k], w.Columns[j].Issues[k]) {
									t.Errorf("boardService.FindWorkspace().Column[].Issues = %v, want %v", g.Columns[j].Issues[k], w.Columns[j].Issues[k])
									return
								}
							}
						}
					}
					if !reflect.DeepEqual(g.Board, w.Board) {
						t.Errorf("boardService.FindWorkspace().Boards[].Board = %v, want %v", g.Board, w.Board)
						return
					}
					if !reflect.DeepEqual(g.Workspace, w.Workspace) {
						t.Errorf("boardService.FindWorkspace().Workspace = %v, want %v", g.Workspace, w.Workspace)
						return
					}
				}
				if !reflect.DeepEqual(got[h].Workspace, tt.want[h].Workspace) {
					t.Errorf("boardService.FindWorkspace().Workspace = %v, want %v", got[h].Workspace, tt.want[h].Workspace)
				}
			}
		})
	}
}
