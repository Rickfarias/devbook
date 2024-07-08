package modelos

import "time"

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID uint64 `json:"id,omitempty"` //omit empty
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"` 
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}