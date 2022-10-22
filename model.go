package flowmail

import (
	"time"
)

type Object struct {
	// not implemented
}

type Disposition string

const DISPOSITION_ATTACHMENT Disposition = "attachment"
const DISPOSITION_INLINE Disposition = "inline"
const DISPOSITION_RELATED Disposition = "related"

type Attachment struct {
	Content     string      `json:"content,omitempty"`
	ContentId   string      `json:"contentId,omitempty"`   // Content-ID header (required for disposition related) Example: <part1.DE1D8F7E.E51807FF@flowmailer.com>
	ContentType string      `json:"contentType,omitempty"` // Examples: application/pdf, image/jpeg
	Disposition Disposition `json:"disposition,omitempty"` // Supported values include: attachment, inline and related, special value related should be used for images referenced in the HTML
	Filename    string      `json:"filename,omitempty"`
}
type SubmitMessage struct {
	Attachments              []Attachment `json:"attachments,omitempty"`
	Data                     Object       `json:"data,omitempty"`
	DeliveryNotificationType string       `json:"deliveryNotificationType,omitempty"`
	FlowSelector             string       `json:"flowSelector,omitempty"`
	HeaderFromAddress        string       `json:"headerFromAddress,omitempty"`
	HeaderFromName           string       `json:"headerFromName,omitempty"`
	HeaderToAddress          string       `json:"headerToAddress,omitempty"`
	HeaderToName             string       `json:"headerToName,omitempty"`
	Headers                  []Header     `json:"headers,omitempty"`
	Html                     string       `json:"html,omitempty"`
	MessageType              MessageType  `json:"messageType,omitempty"`
	Mimedata                 string       `json:"mimedata,omitempty"`
	RecipientAddress         string       `json:"recipientAddress,omitempty"`
	ScheduleAt               time.Time    `json:"scheduleAt,omitempty"`
	SenderAddress            string       `json:"senderAddress,omitempty"`
	Subject                  string       `json:"subject,omitempty"`
	Tags                     []string     `json:"tags,omitempty"`
	Text                     string       `json:"text,omitempty"`
}

type MessageEvent struct {
	Data           string `json:"data"`           // base64 Event data
	DeviceCategory string `json:"deviceCategory"` //
	//ExtraData              string    `json:"extraData"`      // Event data
	Id                     string    `json:"id"`       //Message event ID
	Inserted               time.Time `json:"inserted"` // Database insert date
	LinkName               string    `json:"linkName"`
	LinkTarget             string    `json:"linkTarget"`
	MessageId              string    `json:"messageId"`   // Message ID
	MessageTags            []string  `json:"messageTags"` // Message tags- Only filled for the GET /{account_id}/message_events api call when the parameter addmessagetags is true
	Mta                    string    `json:"mta"`         // mTA that reported this event
	OperatingSystem        string    `json:"operatingSystem"`
	OperatingSystemVersion string    `json:"operatingSystemVersion"`
	Received               time.Time `json:"received"` // Event date
	Referer                string    `json:"referer"`
	RemoteAddr             string    `json:"remoteAddr"`
	Snippet                string    `json:"snippet"`   // Bounce snippet or SMTP conversation snippet
	SubType                string    `json:"subType"`   // Bounce sub type
	Tag                    string    `json:"tag"`       // Custom event type
	EventType              string    `json:"eventType"` // Event type, must be CUSTOM
	UserAgent              string    `json:"userAgent"`
	UserAgentDisplayName   string    `json:"userAgentDisplayName"`
	UserAgentString        string    `json:"userAgentString"`
	UserAgentType          string    `json:"userAgentType"`
	UserAgentVersion       string    `json:"userAgentVersion"`
}

type ObjectDescription struct {
	Description string `json:"description"`
	Id          string `json:"id"`
}

type Address struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MessageType string

const (
	EMAIL  MessageType = "EMAIL"
	SMS    MessageType = "SMS"
	LETTER MessageType = "LETTER"
)

type Message struct {
	BackendDone        time.Time         `json:"backendDone"`        // The time flowmailer was done processing this message
	BackendStart       time.Time         `json:"backendStart"`       // The time flowmailer started processing this message
	Events             []MessageEvent    `json:"events"`             //Message events Ordered by received, new events first.
	Flow               ObjectDescription `json:"flow"`               // Flow this message was processed in
	From               string            `json:"from"`               // The email address in From email header
	FromAddress        Address           `json:"fromAddress"`        // The address in From email header
	HeadersIn          []Header          `json:"headersIn"`          // E-Mail headers of the submitted email message., Only applicable when messageType = EMAIL and addheaders parameter is true
	HeadersOut         []Header          `json:"headersOut"`         // Headers of the final e-mail. Only applicable when messageType = EMAIL and addheaders parameter is true
	Id                 string            `json:"id"`                 // Message id
	MessageDetailsLink string            `json:"messageDetailsLink"` // Link for the message details page. With resend button.
	MessageIdHeader    string            `json:"messageIdHeader"`    // Content of the Message-ID email header
	MessageType        MessageType       `json:"messageType"`        // Message type: EMAIL, SMS or LETTER
	OnlineLink         string            `json:"onlineLink"`         // Last online link
	RecipientAddress   string            `json:"recipientAddress"`   // Recipient address
	SenderAddress      string            `json:"senderAddress"`      // Sender address
	Source             ObjectDescription `json:"source"`             // Source system that submitted this message
	Status             string            `json:"status"`             // Current message status
	Subject            string            `json:"subject"`            // Message subject Only applicable when messageType = EMAIL
	Submitted          time.Time         `json:"submitted"`          // The time this message was submitted to flowmailer
	Tags               []string          `json:"tags"`               // Message tags, only available for api calls with addtags=true
	ToAddressList      []Address         `json:"toAdressList"`       // The recipients in the To email header
	TransactionId      string            `json:"transactionId"`      // The SMTP transaction id, returned with the SMTP 250 response
}

type OAuthTokenResponse struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
	Scope        string `json:"scope"`
	Token_type   string `json:"token_type"`
}

type MessageHold struct {
	BackendDone   time.Time         `json:"backendDone"`   // The time flowmailer was done processing this message
	Data          string            `json:"data"`          // MIME message data or text for SMS messages
	DataCoding    byte              `json:"dataCoding"`    // Only for SMS messages
	ErrorText     string            `json:"errorText"`     // Message error text
	Flow          ObjectDescription `json:"flow"`          // Flow this message was processed in
	MessageId     string            `json:"messageId"`     // Message id
	MessageType   MessageType       `json:"messageType"`   // Message type: EMAIL, SMS or LETTER
	Reason        string            `json:"reason"`        // Message processing failure reason
	Recipient     string            `json:"recipient"`     // Recipient address
	Sender        string            `json:"sender"`        // The email address in From email header
	Source        ObjectDescription `json:"source"`        // Source system that submitted this message
	Status        string            `json:"status"`        // Current message status
	Submitted     time.Time         `json:"submitted"`     // The time this message was submitted to flowmailer
	TransactionId string            `json:"transactionId"` // The SMTP transaction id, returned with the SMTP 250 response
}
