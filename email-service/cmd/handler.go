package main

import (
	"encoding/json"
	"fmt"
	"net/smtp"
)

type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func sendRegisterEmail(message []byte) {
	var payload EmailPayload
	_ = json.Unmarshal(message, &payload)

	fmt.Println(payload.To)
	fmt.Println(payload.Subject)
	fmt.Println(payload.Content)

	// cfg, err := loadEnv()

	// fmt.Println(cfg, err)

	sendGmail(payload.To, payload.Subject, payload.Content)
}

func sendGmail(to, subject, message string) error {
	cfg, err := loadEnv()

	if err != nil {
		return err
	}

	auth := smtp.PlainAuth(
		"",
		cfg.FromMail,  // 送信に使うアカウント
		cfg.GmailPass, // アカウントのパスワード or アプリケーションパスワード
		"smtp.gmail.com",
	)

	content := fmt.Sprintf("To: %s\r\nSubject:%s\r\n\r\n%s", to, subject, message)

	return smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		cfg.FromMail, // 送信元
		[]string{to}, // 送信先
		[]byte(content),
	)
}
