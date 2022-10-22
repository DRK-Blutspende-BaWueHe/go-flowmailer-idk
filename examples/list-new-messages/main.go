package main

import (
	"fmt"
	"go-flowmail"
	"time"
)

func main() {
	fm := flowmail.New(6377, "870605337621aa15a8645cc3eb80e595b1c67d2a", "eec45fcc32a696e64c0f249a6c3c161dd1a8a80f")

	/*
			err := fm.Login()
			if err != nil {
				fmt.Println(err.Error())
			}

		messages, maxitems, err := fm.GetMessages(time.Now().Add(-24*time.Hour), time.Now(), 0, 10)
		if err != nil {
			fmt.Println(err.Error())
		}
	*/

	/*
		attachments := make([]Attachment, 0)
		var attachment Attachment
		attachment.Filename = "etwas.txt"
		attachment.Content = b64.StdEncoding.EncodeToString([]byte("Guten Tag"))
		// attachment.ContentType = "application/octet-stream"
		// attachment.ContentId = "abcdefg"
		// attachment.Disposition = DISPOSITION_ATTACHMENT
		attachments = append(attachments, attachment)

		err := fm.SubmitEmail("kuhr@posteo.de", "stephan Kuhr", "s.kuhr@blutspende.de", "Stephan Kuhr", "Subject is this", "textbody", "htmlbody", attachments)
		if err != nil {
			fmt.Println(err.Error())
		}*/

	err := fm.Login()
	if err != nil {
		fmt.Println(err.Error())
	}

	messages, maxitems, err := fm.GetMessagesHeld(time.Now().Add(-24*time.Hour), time.Now(), 0, 10)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, m := range messages {
		fmt.Printf("%+v\n", m)
	}

	fmt.Printf("%d messages total", maxitems)
}
