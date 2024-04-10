package rabbitmq

type EmailDTO struct {
	Subject string `json:"subject"`
	To      string `json:"to"`
	Message string `json:"message"`
}

func (eDTO *EmailDTO) GetTo() string {
	return eDTO.To
}
func (eDTO *EmailDTO) GetMessage() string {
	return eDTO.Message
}
