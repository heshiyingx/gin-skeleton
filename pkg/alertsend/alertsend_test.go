package alertsend

import (
	"context"
	"testing"
)

func TestAlertSend(t *testing.T) {
	AlertSend(context.TODO(), "http://192.168.31.16:9093/api/v1/alerts")
}
