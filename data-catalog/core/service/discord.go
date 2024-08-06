package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type WebhookMessage struct {
	Content string `json:"content"`
}

func SendDiscordMessage(flag bool, path string, err error) error {
	url := os.Getenv("DISCORD_WEBHOOK_URL")
	if url == "" {
		return fmt.Errorf("DISCORD_WEBHOOK_URL 환경 변수가 설정되지 않았습니다")
	}

	message := WebhookMessage{
		Content: "",
	}

	basicMsg := "[S3 업로드 알림]\n새로운 파일이 업로드 되었습니다.\n\n* 경로: "
	if flag {
		message.Content = basicMsg + path + "\n* 상태 : 정상적인 경로 입니다."
	} else {
		errorMessage := err.Error()
		message.Content = basicMsg + path + "\n* 상태 : 잘못된 경로입니다. 오류 메시지를 확인해주세요.\n* 오류 : " + errorMessage
	}

	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("JSON 직렬화 오류: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("HTTP 요청 생성 오류: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP 요청 오류: %v", err)
	}
	defer resp.Body.Close()

	// fmt.Println("메시지가 성공적으로 전송되었습니다.")
	return nil
}
