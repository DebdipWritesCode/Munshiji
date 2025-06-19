package constants

const EmailSubject = `Welcome to Munshiji - Verify Your Email`

func CreateEmailBody(verificationLink string) string {
	return `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>` + EmailSubject + `</title>
</head>
<body style="font-family: Arial, sans-serif; background-color: #f9f9f9; padding: 40px; color: #333;">
	<div style="max-width: 600px; margin: auto; background: #ffffff; padding: 30px; border-radius: 8px; box-shadow: 0 2px 5px rgba(0,0,0,0.1);">
		<h2 style="color: #222; margin-bottom: 20px;">Welcome to Munshiji!</h2>
		<p style="font-size: 16px; line-height: 1.5;">Thank you for signing up. Please verify your email address to activate your account.</p>
		<p style="margin: 30px 0;">
			<a href="` + verificationLink + `" style="display: inline-block; padding: 12px 20px; background-color: #4CAF50; color: white; text-decoration: none; border-radius: 5px; font-weight: bold;">Verify Email</a>
		</p>
		<p style="font-size: 14px; color: #777;">If you didn’t request this, you can safely ignore this email.</p>
	</div>
</body>
</html>`
}
