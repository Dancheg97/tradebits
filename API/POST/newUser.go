package POST

import "net/http"

type newUserRequest struct {
	SenderPublicKey []byte `json:"SenderPublicKey"`
	MessageKey      []byte `json:"MessageKey"`
	Image           []byte `json:"Image"`
	SenderSign      []byte `json:"SenderSign"`
}

func NewUserRequest(w http.ResponseWriter, r *http.Request) {
	
}
