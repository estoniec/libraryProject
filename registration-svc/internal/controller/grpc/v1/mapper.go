package v1

import (
	pb "github.com/estoniec/automaticLibrary/contracts/gen/go/registration"
	"project11/registration-svc/internal/domain/reg/dto"
	"project11/registration-svc/internal/domain/reg/model"
)

type CheckInput struct {
	ID int64
}

func NewRegInput(req *pb.RegRequest) model.User {
	return model.CreateUser(req.GetID(), req.GetPhone(), req.GetUsername(), req.GetClass())
}

func NewRegOutput(output dto.RegOutput) *pb.RegResponse {
	return &pb.RegResponse{
		Error:  output.Err,
		Status: output.Status,
	}
}

func NewCheckInput(req *pb.CheckRequest) CheckInput {
	return CheckInput{
		ID: req.ID,
	}
}

func NewCheckOutput(output dto.CheckOutput) *pb.CheckResponse {
	return &pb.CheckResponse{
		Checked: output.Checked,
	}
}
