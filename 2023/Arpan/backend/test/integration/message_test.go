package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"testing"
)

func Test_Integration_CreateTodo(t *testing.T) {
	type query struct {
		Query string
	}

	type args struct {
		q query
	}

	type todo struct {
		ID     string `json:"id"`
		Text   string `json:"text"`
		UserID string `json:"userId"`
	}

	type expected struct {
		todo       todo
		statusCode int
	}

	type createTodoResponse struct {
		CreateTodo todo `json:"createTodo"`
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
								userId
							}
						}
					`,
				},
			},
			expected: expected{todo: todo{ID: "todo_id_1", Text: "test", UserID: "user_id_1"}, statusCode: 200},
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

			if !reflect.DeepEqual(res.Data.CreateTodo, tt.expected.todo) {
				t.Errorf("Mutation { CreateTodo } invalid Todo : actual todo = %v, expected todo = %v", res, tt.expected.todo)
			}
		})
	}
}
