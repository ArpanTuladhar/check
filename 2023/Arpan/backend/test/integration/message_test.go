package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Integration_CreateTodo(t *testing.T) {
	type query struct {
		Query string
	}

	type args struct {
		q query
	}

	type user struct {
		Id string `json:"id"`
	}

	type todo struct {
		ID   string `json:"id"`
		Text string `json:"text"`
		User user   `json:"user"`
	}

	type expected struct {
		Todo       todo
		StatusCode int
	}
	type createTodoResponse struct {
		CreateTodo todo `json:"createTodo"`
	}
	type createTodoResponseData struct {
		Data createTodoResponse `json:"data"`
	}

	tests := map[string]struct {
		Args     args
		Expected expected
	}{
		"create todo success": {
			Args: args{
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
								user {
									id
								}
							}
						}
					`,
				},
			},
			Expected: expected{Todo: todo{ID: "todo_id_1", Text: "todo_text_1", User: user{Id: "12345"}}, StatusCode: 200},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			body := bytes.Buffer{}
			if err := json.NewEncoder(&body).Encode(&tt.Args.q); err != nil {
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

			if recorder.Code != tt.Expected.StatusCode {
				t.Errorf("[integration test] Mutation { CreateTodo }: actual statusCode = %v, expected statusCode = %v", recorder.Code, tt.Expected.StatusCode)
			}

			if diff := cmp.Diff(res.Data.CreateTodo, tt.Expected.Todo); diff != "" {
				t.Errorf("[integration test] query { CreateTodo } value is mismatch (-actual +expected):\n%s", diff)
			}
		})
	}
}
