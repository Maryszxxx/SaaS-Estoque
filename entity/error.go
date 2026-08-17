package entity

import "net/http"

type RestErr struct {
	Message string   `json:"message"` // mensagem geral do erro, legível para humanos
	Err     string   `json:"error"`   // "tipo" do erro, em formato de slug (ex: "bad_request")
	Code    int64    `json:"code"`    // código HTTP correspondente (400, 404, 500 etc)
	Causes  []Causes `json:"causes"`  // lista de causas específicas (útil em erros de validação)
}
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  []Causes{},
	}
}

func NewBadRequestValidationError(message string, Causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  Causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
		Causes:  []Causes{},
	}
}

// NewUnauthorizedError cria um erro 401 (Unauthorized),
// usado quando falta autenticação válida.
func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
		Causes:  []Causes{},
	}
}
