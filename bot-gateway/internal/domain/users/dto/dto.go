package usersService

type RegOutput struct {
	Error  string
	Status int64
}

type CheckOutput struct {
	Checked bool
}

func NewRegOutput(error string, status int64) RegOutput {
	return RegOutput{
		Error:  error,
		Status: status,
	}
}

func NewCheckOutput(checked bool) CheckOutput {
	return CheckOutput{
		Checked: checked,
	}
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
