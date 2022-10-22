package flowmail

import "time"

type Flowmailer interface {
	// You do not need to call this directly. Login is called on expired sessions by each Method
	Login() error

	// Get stored messages
	GetMessages(from, until time.Time, rangemin, rangemax int) ([]Message, int, error)

	// Returns:
	//   - List of Messages
	//   - Count of Items in total (not only the pagination window)
	//   - nil if nor error occurs
	GetMessagesHeld(from, until time.Time, rangemin, rangemax int) ([]MessageHold, int, error)

	SubmitEmail(toEmail, toName, fromEmail, fromName, subject, textbody, htmlbody string, attachments []Attachment) error
}
