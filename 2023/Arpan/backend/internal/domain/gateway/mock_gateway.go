// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package gateway

import (
	"context"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/todo"
	"sync"
)

// Ensure, that TodoCommandsGatewayMock does implement TodoCommandsGateway.
// If this is not the case, regenerate this file with moq.
var _ TodoCommandsGateway = &TodoCommandsGatewayMock{}

// TodoCommandsGatewayMock is a mock implementation of TodoCommandsGateway.
//
//	func TestSomethingThatUsesTodoCommandsGateway(t *testing.T) {
//
//		// make and configure a mocked TodoCommandsGateway
//		mockedTodoCommandsGateway := &TodoCommandsGatewayMock{
//			CreateFunc: func(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
//				panic("mock out the Create method")
//			},
//		}
//
//		// use mockedTodoCommandsGateway in code that requires TodoCommandsGateway
//		// and then make assertions.
//
//	}
type TodoCommandsGatewayMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// NewTodo is the newTodo argument value.
			NewTodo *todo.NewTodo
		}
	}
	lockCreate sync.RWMutex
}

// Create calls CreateFunc.
func (mock *TodoCommandsGatewayMock) Create(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
	if mock.CreateFunc == nil {
		panic("TodoCommandsGatewayMock.CreateFunc: method is nil but TodoCommandsGateway.Create was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		NewTodo *todo.NewTodo
	}{
		Ctx:     ctx,
		NewTodo: newTodo,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, newTodo)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedTodoCommandsGateway.CreateCalls())
func (mock *TodoCommandsGatewayMock) CreateCalls() []struct {
	Ctx     context.Context
	NewTodo *todo.NewTodo
} {
	var calls []struct {
		Ctx     context.Context
		NewTodo *todo.NewTodo
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}