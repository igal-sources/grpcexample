package main

import (
	"context"
	"fmt"
)

// Server is the Logic handler for the server
// It has to fullfill the GRPC schema generated Interface
// In this case its only 1 function called Ping
type Server struct {
	UnimplementedProcedureServiceServer
}

// Ping fullfills the requirement for PingPong Server interface

func (s *Server) GetProcedures(ctx context.Context, in *GetProceduresRequest) (*GetProceduresResponse, error) {
	return &GetProceduresResponse{
		Procedures: s.loadProcedures(),
	}, nil
}
func (s *Server) CreateProcedure(ctx context.Context, in *CreateProcedureRequest) (*CreateProcedureResponse, error) {
	return nil, nil
}
func (s *Server) UpdateProcedure(ctx context.Context, in *UpdateProcedureRequest) (*UpdateProcedureResponse, error) {
	return nil, nil
}
func (s *Server) DeleteProcedure(ctx context.Context, in *DeleteProcedureRequest) (*DeleteProcedureResponse, error) {
	return nil, nil
}
func (s *Server) LockProcedure(ctx context.Context, in *LockProcedureRequest) (*LockProcedureResponse, error) {
	return nil, nil
}

func (s *Server) loadProcedures() []*ProcedureTemplate {
	arr := make([]*ProcedureTemplate, 0, 10)

	for i := 0; i < 10; i++ {
		arr = append(arr, &ProcedureTemplate{
			ProcedureId: fmt.Sprintf("%d", i),
			IsActive: true,
			Name:        fmt.Sprintf("Procedure %d", i),	
		})
	}

	return arr
}
