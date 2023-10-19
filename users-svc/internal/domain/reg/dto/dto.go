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

func NewCheckOutput(checked bool) CheckOutput {
	return CheckOutput{
		Checked: checked,
	}
}

type CheckRoleInput struct {
	ID int64
}

type CheckRoleOutput struct {
	Role   int
	Error  string
	Status int64
}

func NewCheckRoleOutput(role int, error string, status int64) CheckRoleOutput {
	return CheckRoleOutput{
		Role:   role,
		Error:  error,
		Status: status,
	}
}
