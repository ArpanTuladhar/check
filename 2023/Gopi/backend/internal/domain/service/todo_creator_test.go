package service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/domain/model/todo"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/usecase/input"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/usecase/output"
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
		"create message success": {
			prepare: func(f *fields) {
				f.mockTodoCommandsGateway = &gateway.TodoCommandsGatewayMock{
					CreateTodoFunc: func(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
						return &todo.Todo{
							ID:   "todo_id_1",
							Text: "todo_text_1",
						}, nil
					},
				}
			},
			args: args{
				text: "todo",
			},
			expected: expected{
				todo: &output.TodoCreator{
					ID:   "todo_id_1",
					Text: "todo_text_1",
				},
			},
			wantErr: false,
		},

		"create message failure (gateway error)": {
			prepare: func(f *fields) {
				f.mockTodoCommandsGateway = &gateway.TodoCommandsGatewayMock{
					CreateTodoFunc: func(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
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
				t.Errorf("todoCreator.CreateTodo error %v, wantErr %v", err, tt.wantErr)
			}

			if len(f.mockTodoCommandsGateway.CreateTodoCalls()) != 1 {
				t.Error("Unexpected number of calls to CreateTodo. Got calls, expected 1")
			}

			if diff := cmp.Diff(out, tt.expected.todo); diff != "" {
				t.Errorf("todoCreator.CreateTodo returns mismatch values : (-actual +expected):\n%s", diff)
			}

		})
	}

}
