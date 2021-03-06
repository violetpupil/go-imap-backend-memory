// A memory backend.
package memory

import (
	"errors"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend"
)

type Backend struct {
	Users map[string]*User
}

func (be *Backend) Login(_ *imap.ConnInfo, username, password string) (backend.User, error) {
	user, ok := be.Users[username]
	if ok && user.Password == password {
		return user, nil
	}

	return nil, errors.New("Bad username or password")
}

func (be *Backend) NewUser(name, password string) (err error) {
	user := &User{Name: name, Password: password, mailboxes: make(map[string]*Mailbox)}
	err = user.CreateMailbox("INBOX")
	if err == nil {
		be.Users[user.Name] = user
	}
	return
}

func New() *Backend {
	user := &User{Name: "username", Password: "password"}

	body := "From: contact@example.org\r\n" +
		"To: contact@example.org\r\n" +
		"Subject: A little message, just for you\r\n" +
		"Date: Wed, 11 May 2016 14:31:59 +0000\r\n" +
		"Message-ID: <0000000@localhost/>\r\n" +
		"Content-Type: text/plain\r\n" +
		"\r\n" +
		"Hi there :)"

	user.mailboxes = map[string]*Mailbox{
		"INBOX": {
			name: "INBOX",
			user: user,
			Messages: []*Message{
				{
					Uid:   6,
					Date:  time.Now(),
					Flags: []string{"\\Seen"},
					Size:  uint32(len(body)),
					Body:  []byte(body),
				},
			},
		},
	}

	return &Backend{
		Users: map[string]*User{user.Name: user},
	}
}
