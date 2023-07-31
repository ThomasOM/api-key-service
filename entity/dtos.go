package entity

type GenerateKeyRequest struct {
	Owner string `json:"owner"`
}

type FindKeysRequest struct {
	Owner string `json:"owner"`
}

type AuthenticateRequest struct {
	Key   []byte `json:"key"`
}
