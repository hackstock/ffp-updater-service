package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// SmsAPIConfig defines authentication credentials for send SMS
type SmsAPIConfig struct {
	username string
	password string
	senderID string
}

// NewSmsAPIConfig returns a pointer to SmsApiConfig
func NewSmsAPIConfig(usr, pwd, sid string) *SmsAPIConfig {
	return &SmsAPIConfig{
		username: usr,
		password: pwd,
		senderID: sid,
	}
}

// SendSMS sends SMS to recipient
func SendSMS(apiCfg *SmsAPIConfig, to, msg string) error {
	endpoint := "https://infoline.nandiclient.com/AWA/campaigns/sendmsg"

	form := url.Values{}
	form.Set("username", apiCfg.username)
	form.Set("password", apiCfg.password)
	form.Set("from", apiCfg.senderID)
	form.Set("message", msg)
	form.Set("numbers", to)

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return fmt.Errorf("failed creating sms request for nandi : %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed sending sms to %s : %v", to, err)
	}
	io.Copy(ioutil.Discard, res.Body)

	return nil
}
