package jwt_token

type Payload struct {
	token string
}

func NewPayload(token string) (payload *Payload, err error) {
	payload.token = token
	return payload, nil
}
