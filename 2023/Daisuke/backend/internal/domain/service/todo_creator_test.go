package service

import (
	"context"
	"testing"

	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/domain/model/todo"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/usecase/input"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/usecase/output"
	"github.com/google/go-cmp/cmp"
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

			if len(f.mockTodoCommandsGateway.CreateTodoCalls()) != 1 {
				t.Error("one creator.CreateTodo expected")
			}

			if diff := cmp.Diff(out, tt.expected.todo); diff != "" {
				t.Errorf("usecase.CreateTodo not correct: (-actual +expected):\n%s", diff)
			}

		})
	}

}
