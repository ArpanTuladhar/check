// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package gateway

import (
	"context"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/domain/model/todo"
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
//			CreateTodoFunc: func(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
//				panic("mock out the CreateTodo method")
//			},
//		}
//
//		// use mockedTodoCommandsGateway in code that requires TodoCommandsGateway
//		// and then make assertions.
//
//	}
type TodoCommandsGatewayMock struct {
	// CreateTodoFunc mocks the CreateTodo method.
	CreateTodoFunc func(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateTodo holds details about calls to the CreateTodo method.
		CreateTodo []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// NewTodo is the newTodo argument value.
			NewTodo *todo.NewTodo
		}
	}
	lockCreateTodo sync.RWMutex
}

// CreateTodo calls CreateTodoFunc.
func (mock *TodoCommandsGatewayMock) CreateTodo(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
	if mock.CreateTodoFunc == nil {
		panic("TodoCommandsGatewayMock.CreateTodoFunc: method is nil but TodoCommandsGateway.CreateTodo was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		NewTodo *todo.NewTodo
	}{
		Ctx:     ctx,
		NewTodo: newTodo,
	}
	mock.lockCreateTodo.Lock()
	mock.calls.CreateTodo = append(mock.calls.CreateTodo, callInfo)
	mock.lockCreateTodo.Unlock()
	return mock.CreateTodoFunc(ctx, newTodo)
}

// CreateTodoCalls gets all the calls that were made to CreateTodo.
// Check the length with:
//
//	len(mockedTodoCommandsGateway.CreateTodoCalls())
func (mock *TodoCommandsGatewayMock) CreateTodoCalls() []struct {
	Ctx     context.Context
	NewTodo *todo.NewTodo
} {
	var calls []struct {
		Ctx     context.Context
		NewTodo *todo.NewTodo
	}
	mock.lockCreateTodo.RLock()
	calls = mock.calls.CreateTodo
	mock.lockCreateTodo.RUnlock()
	return calls
}

// Ensure, that BinderMock does implement Binder.
// If this is not the case, regenerate this file with moq.
var _ Binder = &BinderMock{}

// BinderMock is a mock implementation of Binder.
//
//	func TestSomethingThatUsesBinder(t *testing.T) {
//
//		// make and configure a mocked Binder
//		mockedBinder := &BinderMock{
//			BindFunc: func(contextMoqParam context.Context) context.Context {
//				panic("mock out the Bind method")
//			},
//		}
//
//		// use mockedBinder in code that requires Binder
//		// and then make assertions.
//
//	}
type BinderMock struct {
	// BindFunc mocks the Bind method.
	BindFunc func(contextMoqParam context.Context) context.Context

	// calls tracks calls to the methods.
	calls struct {
		// Bind holds details about calls to the Bind method.
		Bind []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
	}
	lockBind sync.RWMutex
}

// Bind calls BindFunc.
func (mock *BinderMock) Bind(contextMoqParam context.Context) context.Context {
	if mock.BindFunc == nil {
		panic("BinderMock.BindFunc: method is nil but Binder.Bind was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockBind.Lock()
	mock.calls.Bind = append(mock.calls.Bind, callInfo)
	mock.lockBind.Unlock()
	return mock.BindFunc(contextMoqParam)
}

// BindCalls gets all the calls that were made to Bind.
// Check the length with:
//
//	len(mockedBinder.BindCalls())
func (mock *BinderMock) BindCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockBind.RLock()
	calls = mock.calls.Bind
	mock.lockBind.RUnlock()
	return calls
}

// Ensure, that TransactorMock does implement Transactor.
// If this is not the case, regenerate this file with moq.
var _ Transactor = &TransactorMock{}

// TransactorMock is a mock implementation of Transactor.
//
//	func TestSomethingThatUsesTransactor(t *testing.T) {
//
//		// make and configure a mocked Transactor
//		mockedTransactor := &TransactorMock{
//			TransactionFunc: func(contextMoqParam context.Context, fn func(context.Context) error) error {
//				panic("mock out the Transaction method")
//			},
//		}
//
//		// use mockedTransactor in code that requires Transactor
//		// and then make assertions.
//
//	}
type TransactorMock struct {
	// TransactionFunc mocks the Transaction method.
	TransactionFunc func(contextMoqParam context.Context, fn func(context.Context) error) error

	// calls tracks calls to the methods.
	calls struct {
		// Transaction holds details about calls to the Transaction method.
		Transaction []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Fn is the fn argument value.
			Fn func(context.Context) error
		}
	}
	lockTransaction sync.RWMutex
}

// Transaction calls TransactionFunc.
func (mock *TransactorMock) Transaction(contextMoqParam context.Context, fn func(context.Context) error) error {
	if mock.TransactionFunc == nil {
		panic("TransactorMock.TransactionFunc: method is nil but Transactor.Transaction was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Fn              func(context.Context) error
	}{
		ContextMoqParam: contextMoqParam,
		Fn:              fn,
	}
	mock.lockTransaction.Lock()
	mock.calls.Transaction = append(mock.calls.Transaction, callInfo)
	mock.lockTransaction.Unlock()
	return mock.TransactionFunc(contextMoqParam, fn)
}

// TransactionCalls gets all the calls that were made to Transaction.
// Check the length with:
//
//	len(mockedTransactor.TransactionCalls())
func (mock *TransactorMock) TransactionCalls() []struct {
	ContextMoqParam context.Context
	Fn              func(context.Context) error
} {
	var calls []struct {
		ContextMoqParam context.Context
		Fn              func(context.Context) error
	}
	mock.lockTransaction.RLock()
	calls = mock.calls.Transaction
	mock.lockTransaction.RUnlock()
	return calls
}
