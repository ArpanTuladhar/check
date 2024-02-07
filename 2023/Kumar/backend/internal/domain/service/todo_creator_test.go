package service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/domain/model/todo"
	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/usecase/input"
	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/usecase/output"
)

func TestCreateTodo(t *testing.T) {
	type fields struct {
		mockTodoCommandsGateway *gateway.TodoCommandsGatewayMock
	}

	type args struct {
		text string
	}

	type expected struct {
		todo *output.TodoCreator
	}

	tests := map[string]struct {
		prepare  func(f *fields)
		args     args
		expected expected
		wantErr  bool
	}{
		"create todo success": {
			prepare: func(f *fields) {
				f.mockTodoCommandsGateway = &gateway.TodoCommandsGatewayMock{
					CreateFunc: func(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
						return &todo.Todo{
							ID:     "todo_id_1",
							Text:   "todo_text_1",
							UserID: 12345,
						}, nil
					},
				}
			},
			args: args{
				text: "todo",
			},
			expected: expected{
				todo: &output.TodoCreator{
					ID:     "todo_id_1",
					Text:   "todo_text_1",
					UserID: 12345,
				},
			},
			wantErr: false,
		},
		"create todo failure (gateway error)": {
			prepare: func(f *fields) {
				f.mockTodoCommandsGateway = &gateway.TodoCommandsGatewayMock{
					CreateFunc: func(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
						return nil, errors.New("this is an error")
					},
				}
			},
			args: args{
				text: "error",
			},
			expected: expected{},
			wantErr:  true,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			f := fields{
				mockTodoCommandsGateway: nil,
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}
			creator := NewTodoCreator(f.mockTodoCommandsGateway)
			out, err := creator.CreateTodo(
				context.Background(),
				&input.TodoCreator{
					Text: tt.args.text,
				},
			)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.CreateTodo error %v, wantErr %v", err, tt.wantErr)
			}

			if len(f.mockTodoCommandsGateway.CreateCalls()) != 1 {
				t.Error("one creator.CreateTodo expected")
			}

			if diff := cmp.Diff(out, tt.expected.todo); diff != "" {
				t.Errorf("usecase.CreateTodo not correct: (-actual +expected):\n%s", diff)
			}

		})
	}

}
