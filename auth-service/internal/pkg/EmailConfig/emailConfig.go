package emailconfig

import (
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

// HTML shablonni fayldan o'qib olish
func readHTMLTemplate(templatePath string) (string, error) {
	// Joriy ishchi katalogni topish
	execPath, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("katalogni o'qib bo'lmadi: %v", err)
	}

	// HTML faylning to'liq yo'lini tuzish
	fullPath := filepath.Join(execPath, templatePath)

	// Faylni o'qish
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("HTML shablonni o'qib bo'lmadi: %v", err)
	}

	return string(content), nil
}

func SendCode(sendToEmail string, recoveryCode string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	from := "bekzodnematov.pk@gmail.com"
	password := "urpq hnxt nriq pyis"
	to := []string{sendToEmail}

	// HTML shablonni fayldan o'qib olish
	htmlTemplate, err := readHTMLTemplate("./internal/pkg/EmailConfig/emailShablon.html")
	if err != nil {
		return fmt.Errorf("HTML shablonni yuklab bo'lmadi: %v", err)
	}

	// Recovery codeni HTML shablonga joylashtirish
	htmlBody := strings.Replace(htmlTemplate, "{{RECOVERY_CODE}}", recoveryCode, -1)

	// Xabarni tayyorlash
	subject := "Subject: Hisobingizni qayta tiklash uchun tasdiqlash kodi\r\n"
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	msg := []byte(subject + mime + htmlBody)

	// Autentifikatsiya
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Xabarni yuborish
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
