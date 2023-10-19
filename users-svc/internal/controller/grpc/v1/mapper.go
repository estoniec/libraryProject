package v1

import (
	pb "github.com/estoniec/libraryProject/contracts/gen/go/users"
	"registration-svc/internal/domain/reg/dto"
	"registration-svc/internal/domain/reg/model"
)

func NewRegInput(req *pb.RegRequest) model.User {
	return model.CreateUser(req.GetID(), req.GetPhone(), req.GetUsername(), req.GetClass())
}

func NewRegOutput(output dto.RegOutput) *pb.RegResponse {
	return &pb.RegResponse{
		Error:  output.Err,
		Status: output.Status,
	}
}

func NewCheckInput(req *pb.CheckRequest) dto.CheckInput {
	return dto.CheckInput{
		ID: req.ID,
	}
}

func NewCheckOutput(output dto.CheckOutput) *pb.CheckResponse {
	return &pb.CheckResponse{
		Checked: output.Checked,
	}
}

func NewCheckRoleInput(req *pb.CheckRoleRequest) dto.CheckRoleInput {
	return dto.CheckRoleInput{
		ID: req.ID,
	}
}

func NewCheckRoleOutput(output dto.CheckRoleOutput) *pb.CheckRoleResponse {
	return &pb.CheckRoleResponse{
		Role:   int64(output.Role),
		Error:  output.Error,
		Status: output.Status,
	}
}
