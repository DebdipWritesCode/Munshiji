package mail

import (
	"testing"

	"github.com/DebdipWritesCode/Munshiji/backend/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping email test in short mode")
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(
		config.EmailSenderName,
		config.EmailSenderAddress,
		config.EmailSenderPassword,
	)

	subject := "Test Email"
	content := "<h1>This is a test email</h1>"

	to := []string{"debdipmukherjee52@gmail.com"}
	attachments := []string{"../.gitignore"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachments)
	require.NoError(t, err, "Failed to send email with Gmail sender")
}
