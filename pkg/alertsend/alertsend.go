package alertsend

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/prometheus/common/model"
	"net/http"
	"time"
)

func AlertSend(ctx context.Context, alertMUrl string) {
	// 发送内容参考：https://prometheus.io/docs/alerting/latest/clients/
	labels := model.LabelSet{}
	labels["alertname"] = "报警测试"
	labels["group"] = "abc"
	labels["severity"] = "2"
	labels["job"] = "svc_alert"

	anno := model.LabelSet{}
	anno["value"] = "80"
	alerts := make([]*model.Alert, 0, 1)
	alert := &model.Alert{
		Labels:       labels,
		Annotations:  anno,
		StartsAt:     time.Time{},
		EndsAt:       time.Time{},
		GeneratorURL: "http://xxxxx:9090", //查看报警的连接
	}
	alerts = append(alerts, alert)
	marshal, err := json.Marshal(alerts)
	if err != nil {
		return
	}
	request, err := http.NewRequest("POST", alertMUrl, bytes.NewBuffer(marshal))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

}
