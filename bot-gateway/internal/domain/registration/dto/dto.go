package dto

type RegInput struct {
	Phone    string
	Username string
	Class    string
	ID       int64
}

type RegOutput struct {
	Error  string
	Status int64
}

type CheckInput struct {
	ID int64
}

type CheckOutput struct {
	Checked bool
}

func NewRegInput(phone string, username string, class string, ID int64) RegInput {
	return RegInput{
		Phone:    phone,
		Username: username,
		Class:    class,
		ID:       ID,
	}
}

func NewRegOutput(error string, status int64) RegOutput {
	return RegOutput{
		Error:  error,
		Status: status,
	}
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
