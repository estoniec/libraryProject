package dto

type RegOutput struct {
	Err    string
	Status int64
}

func NewRegOutput(err string, status int64) RegOutput {
	return RegOutput{
		Err:    err,
		Status: status,
	}
}

type CheckInput struct {
	ID int64
}

type CheckOutput struct {
	Checked bool
}

func NewCheckInput(id int64) CheckInput {
	return CheckInput{
		ID: id,
	}
}

func NewCheckOutput(checked bool) CheckOutput {
	return CheckOutput{
		Checked: checked,
	}
}
