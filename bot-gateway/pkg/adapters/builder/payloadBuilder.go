package builder

type Payload map[string]any

func NewPayload() Payload {
	return make(Payload)
}

func (p Payload) SetCommand(command string) Payload {
	p["command"] = command
	return p
}

func (p Payload) AddPayload(k string, v any) Payload {
	p[k] = v
	return p
}
