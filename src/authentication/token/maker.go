package token

type Maker interface {
	// CreateToken creates a new token for a specific phone and duration
	CreateToken(phone string) (*Payload, string, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
