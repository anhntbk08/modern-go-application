package todov1beta1

import (
	"context"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type TodoListKitServer struct {
	*UnimplementedTodoListServer

	CreateTodoHandler kitgrpc.Handler
	ListTodosHandler  kitgrpc.Handler
	MarkAsDoneHandler kitgrpc.Handler
}

func (s TodoListKitServer) CreateTodo(ctx context.Context, req *CreateTodoRequest) (*CreateTodoResponse, error) {
	_, resp, err := s.CreateTodoHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*CreateTodoResponse), nil
}
func (s TodoListKitServer) ListTodos(ctx context.Context, req *ListTodosRequest) (*ListTodosResponse, error) {
	_, resp, err := s.ListTodosHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*ListTodosResponse), nil
}
func (s TodoListKitServer) MarkAsDone(ctx context.Context, req *MarkAsDoneRequest) (*MarkAsDoneResponse, error) {
	_, resp, err := s.MarkAsDoneHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*MarkAsDoneResponse), nil
}
