package messages

const (

	// commands
	GETMETRICCOMMAND  = "status"
	ADDSITECOMMAND    = "addurl"
	DELETESITECOMMAND = "deleteurl"
	HELPCOMMAND       = "help"
	UNKMOWNCOMMAND    = "Use /help to see available commands"

	// Errors
	URLNOTADDED     = "Error while adding Url"
	URLNOTDELETED   = "Error while deleting Url"
	CANNOTGETURLS   = "Error while getting urls"
	URLSNOTFOUND    = "You dont have urls, use /addurl for add url"
	URLNOTFOUND     = "This url not found"
	NOTURL          = "Incorrect URL"
	URLALREADYADDED = "This url already added"
	SERVERERROR     = "Unexpected error, please try later"

	// success
	URLADDED   = "Url successfully added"
	URLDELETED = "Url successfully deleted"

	// messages
	URLAWAILABLE   = "Available✅"
	URLUNAWAILABLE = "Unavailable❌"
	SENDDATA       = "Send me absolute url (example: https://google.com/)"
	HELP           = "Available commands: \n➖/addurl - adding url \n➖/deleteurl - deleting url \n➖/status - send status of your urls \n➖/help - sending this message"
)
