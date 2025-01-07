package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendEmail(to, subject, body string) error {
	url := "http://localhost:8080/send_email"

	fmt.Println("ok")

	payload := map[string]string{
		"to":      to,
		"subject": subject,
		"body":    body,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err)
	}

	fmt.Println(jsonPayload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("email service returned status: %s", resp.Status)
	}
	return nil
}
