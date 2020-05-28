package datatransfers

type MailSuccessResponse struct {
	Message string `json:"message" bson:"message"`
	Id      string `json:"id" bson:"id"`
}

type MailParam struct {
	Subject   string
	Body      string
	Recipient string
}
