package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/push"
	gin_skeleton "gitlab.myshuju.top/heshiying/gin-skeleton"
	"gitlab.myshuju.top/heshiying/gin-skeleton/g"
	"go.uber.org/zap"
	"net/http"
	"regexp"
	"strconv"
	"sync/atomic"
	"time"
)

type ServerCollect struct {
	reqCnt *prometheus.CounterVec
	reqDur *prometheus.HistogramVec
	//connectionGauge *prometheus.GaugeVec
}

var (
	//Gauge = prometheus.NewCounterVec()
	pusher  *push.Pusher
	c       *ServerCollect
	counter int64
)

const (
	DefaultMetricPath = "/metrics"
)

func NewServerCollect(name string) *ServerCollect {
	reqCnt := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   "",
		Subsystem:   "",
		Name:        "requests_total",
		Help:        "How many HTTP requests processed, partitioned by status code and HTTP method.",
		ConstLabels: map[string]string{"svc": name},
	}, []string{"path", "code", "method"})

	reqDur := prometheus.NewHistogramVec(prometheus.HistogramOpts(prometheus.HistogramOpts{
		Namespace:   "",
		Subsystem:   "",
		Name:        "request_duration_seconds",
		Help:        "request_duration_seconds",
		ConstLabels: map[string]string{"svc": name},
		Buckets:     prometheus.LinearBuckets(0.5, 0.5, 20),
	}), []string{"path", "code", "method"})

	return &ServerCollect{
		reqCnt: reqCnt,
		reqDur: reqDur,
	}
}

func (s *ServerCollect) Describe(ch chan<- *prometheus.Desc) {
	s.reqDur.MetricVec.Describe(ch)
	s.reqCnt.MetricVec.Describe(ch)
}

func (s *ServerCollect) Collect(metrics chan<- prometheus.Metric) {
	s.reqDur.Collect(metrics)
	s.reqCnt.Collect(metrics)
}
func UsePromMiddleware(name string) gin.HandlerFunc {
	s := NewServerCollect(name)
	c = s
	gin_skeleton.RegisterServer(StartPromHttp)
	return useMiddleWithServiceCollect(s)
}

func useMiddleWithServiceCollect(s *ServerCollect) gin.HandlerFunc {
	return func(c *gin.Context) {
		atomic.AddInt64(&counter, 1)
		defer atomic.AddInt64(&counter, -1)
		if c.Request.URL.String() == DefaultMetricPath {
			c.Next()
			return
		}
		if c.Request.Method == http.MethodOptions || c.Request.Method == http.MethodHead {
			c.Next()
			return
		}

		start := time.Now()
		c.Next()
		code := strconv.Itoa(c.Writer.Status())
		dur := float64(time.Since(start)) / float64(time.Second)
		path := getPath(c)
		s.reqCnt.WithLabelValues(path, code, c.Request.Method).Inc()
		s.reqDur.WithLabelValues(path, code, c.Request.Method).Observe(dur)

	}
}

func UserPromGatewayMiddleware(name string, gatewayUrl string) gin.HandlerFunc {
	collect := NewServerCollect(name)
	c = collect
	pusher = push.New(gatewayUrl, "svc").Collector(collect)
	go func() {
		ticker := time.NewTicker(time.Second * 3)
		for range ticker.C {
			err := pusher.Push()
			if err != nil {
				g.Error("ServerCollect push err", zap.Error(err))
			}
		}
	}()

	ginMiddle := useMiddleWithServiceCollect(collect)

	return ginMiddle
}

func getPath(c *gin.Context) string {
	return c.FullPath()
}
func registerMetrics(name string) {
	prometheus.Register(&ServerCollect{})
}
func Init(url, jobName string, s *ServerCollect) {
	pusher = push.New(url, jobName)
	pusher.Collector(s)
}

func StartPromHttp() error {
	if c == nil {
		panic("请使用UsePromMiddleware应用中间件")
	}

	reg := prometheus.NewRegistry()
	// Add Go module build info.
	//reg.MustRegister(collectors.NewBuildInfoCollector())
	reg.MustRegister(c)
	reg.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
	))

	// Expose the registered metrics via HTTP.
	http.Handle(DefaultMetricPath, promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))
	fmt.Println("Hello world from new Go Collector!")
	if err := http.ListenAndServe(":9100", nil); err != nil {
		return err
	}
	return nil

}
