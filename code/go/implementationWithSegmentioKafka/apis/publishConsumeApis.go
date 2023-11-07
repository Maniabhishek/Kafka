package apis

import (
	"encoding/json"
	"io"

	"github.com/Maniabhishek/Kafka/commons"
	"github.com/Maniabhishek/Kafka/internal/services"
	"github.com/labstack/echo/v4"
)

type messageController struct {
	mservice services.MInterface
}

func NewMessageController(mservice services.MInterface) *messageController {
	return &messageController{
		mservice: mservice,
	}
}

func (m *messageController) SendKafkaMessage(ctx echo.Context) error {
	body := ctx.Request().Body
	bytes, err := io.ReadAll(body)
	defer body.Close()
	if err != nil {
		return ctx.JSON(500, map[string]string{"message": "internal server error"})
	}

	var payload commons.KeyValue
	if err := json.Unmarshal(bytes, &payload); err != nil {
		return ctx.JSON(500, map[string]string{"message": "error in payload"})
	}

	key := payload.Key
	value := payload.Value

	data := []byte(value)

	if err := m.mservice.SendTokafka(key, data); err != nil {
		return ctx.JSON(500, map[string]string{"message": "error while sending message"})
	}
	return ctx.JSON(200, map[string]string{"key": key, "value": value})
}
