package token

import "time"

// managin token makers
type Maker interface {
	//create and sign new token for user
	CreateToken(username string, id int64, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
