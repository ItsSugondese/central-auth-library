package token_auth_type_enums

var TokenAuthType = newTokenAuthType()

func newTokenAuthType() *tokenAuthType {
	return &tokenAuthType{
		JWT:    "JWT",
		OAUTH:  "OAuth",
		PASETO: "Paseto",
	}
}

type tokenAuthType struct {
	JWT    string
	OAUTH  string
	PASETO string
}
