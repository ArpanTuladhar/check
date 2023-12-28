package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/handler/graph/model"
	"github.com/google/go-cmp/cmp"
)

func Test_Integration_CreateTodo(t *testing.T) {
	type query struct {
		Query string
	}

	type args struct {
		q query
	}

	type expected struct {
		todo       *model.Todo
		statusCode int
	}
	type createTodoResponse struct {
		CreateTodo *model.Todo `json:"createTodo"`
	}
	type createTodoResponseData struct {
		Data createTodoResponse `json:"data"`
	}

	tests := map[string]struct {
		args     args
		expected expected
	}{
		"create todo success": {
			args: args{
				q: query{
					Query: `
						mutation {
							createTodo(
								input: {
									text: "todo_text_1",
									userId: "user_id_1",
								}
							){
								id
								text
							}
						}
					`,
				},
			},
			expected: expected{
				todo: &model.Todo{
					ID:   "todo_id_1",
					Text: "test",
				},
				statusCode: 200,
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			body := bytes.Buffer{}
			if err := json.NewEncoder(&body).Encode(&tt.args.q); err != nil {
				panic(err)
			}
			recorder := DoGraphQLRequest(
				&body,
			)
			re, err := io.ReadAll(recorder.Result().Body)
			if err != nil {
				panic(err)
			}

			res := createTodoResponseData{}
			json.Unmarshal(re, &res)

			if recorder.Code != tt.expected.statusCode {
				t.Errorf("[integration test] Mutation { CreateTodo }v: actual statusCode = %v, expected statusCode = %v", recorder.Code, tt.expected.statusCode)
			}

			if diff := cmp.Diff(res.Data.CreateTodo, tt.expected.todo); diff != "" {
				t.Errorf("[integration test] query { CreateTodo } value is mismatch (-actual +expected):\n%s", diff)
			}
		})
	}

}