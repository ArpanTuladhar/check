package service

import (
	"context"
	"errors"
	"testing"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/todo"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase/input"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase/output"
	"github.com/google/go-cmp/cmp"
)

func TestCreateTodo(t *testing.T) {
	type fields struct {
		ctx                     context.Context
		mockTodoCommandsGateway *gateway.TodoCommandsGatewayMock
		mockBinder              *gateway.BinderMock
		mockTransactor          *gateway.TransactorMock
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
					CreateFunc: func(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
						return &todo.Todo{
							ID:     "todo_id_1",
							Text:   "todo_text_1",
							UserID: 12345,
						}, nil
					},
				}

				f.mockBinder = &gateway.BinderMock{
					BindFunc: func(contextMoqParam context.Context) context.Context {
						return f.ctx
					},
				}

				f.mockTransactor = &gateway.TransactorMock{
					TransactionFunc: func(contextMoqParam context.Context, fn func(context.Context) error) error {
						return nil
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

		"create message failure (gateway error)": {
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
			creator := NewTodoCreator(f.mockBinder, f.mockTransactor, f.mockTodoCommandsGateway)
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
