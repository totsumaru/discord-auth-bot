package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const webhookURL = "https://discord.com/api/webhooks/1136568735692497000/Wfav5O3GvlRTuLdXBlZ_I0LFZGKnx-0K3CFt7OM-HLUVjtZTa0PfFTzDkaED4n_xnHsK"

type Webhook struct {
	Content string `json:"content"`
}

func SendDiscord(err error) {
	msg := Webhook{
		Content: fmt.Sprintf("<@960104306151948328> %s", err.Error()),
	}

	jsonData, _ := json.Marshal(msg)

	_, _ = http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
}
