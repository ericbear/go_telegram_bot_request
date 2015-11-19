package message

type Response struct {
	Ok bool					`json:"ok"`
	Result []Result			`json:"result"`
}

type Result struct {
	UpdateId int			`json:"update_id"`
	Message ResultMessage	`json:"message"`
}

type ResultMessage struct {
	MessageId int			`json:"message_id"`
	From User				`json:"from"`
	Chat Chat				`json:"chat"`
	Date int64				`json:"date"`
	Text string				`json:"text"`
}

type User struct {
	Id int					`json:"id"`
	FirstName string		`json:"first_name"`
	LastName string			`json:"last_name"`
	Username string			`json:"username"`
}

type Chat struct {
	Id int					`json:"id"`
	FirstName string		`json:"first_name"`
	LastName string			`json:"last_name"`
	Username string			`json:"username"`
	Type string				`json:"type"`
}