package http

import (
	"github.com/redis/go-redis/v9"
	breaker "github.com/sony/gobreaker"
	"sportbit.com/racechip/logger"
	"time"
)

type Client struct {
	*breaker.CircuitBreaker
	Config
	logger.AppLog
	RedisRequestBuilder
}

func New(c Config, log logger.AppLog) IRedisBackendService {
	var st breaker.Settings
	st.Name = "Redis"
	st.Timeout = 3 * time.Second
	st.ReadyToTrip = func(counts breaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}
	cb := breaker.NewCircuitBreaker(st)
	return &Client{CircuitBreaker: cb, Config: c, AppLog: log}
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       string
}

type Config struct {
	RedisCfg RedisConfig
}

func (s *Client) Set(input RedisRequestBuilder) error {
	defer func(t time.Time) {
		s.AppLog.Trace.Printf("%s", input.Req)
		s.AppLog.Trace.Printf("Redis.DoRequest elapsed time %.4f ms", float64(time.Since(t).Nanoseconds())/float64(time.Millisecond))

	}(time.Now())
	_, err := s.CircuitBreaker.Execute(func() (i interface{}, e error) {

		//rawBody, err := input.Req.GetBody()
		//if err != nil {
		//	return nil, err
		//}
		//body, err := io.ReadAll(rawBody)
		//if err != nil {
		//	return nil, err
		//}
		//s.Redis.Println("REQ Body:", string(body))
		//resp, err := input.client.Get(input.Req) //client.Do(req)
		//if err != nil {
		//	return nil, err
		//}
		//defer resp.Body.Close()
		//body, err = io.ReadAll(resp.Body)
		//if err != nil {
		//	return nil, err
		//}
		//s.Redis.Println("RES Body:", string(body))
		//if resp.StatusCode < 200 || resp.StatusCode > 299 {
		//	return body, fmt.Errorf("%d:%s", resp.StatusCode, http.StatusText(resp.StatusCode))
		//}
		return nil, nil
	})
	//if res != nil {
	//	errParse := json.Unmarshal(res.([]byte), input.responseModel)
	//	if errParse != nil {
	//		return errParse
	//	}
	//	if err != nil {
	//		return err
	//	}
	//}
	return err
}

type RedisRequestBuilder struct {
	Req           string
	client        *redis.Client
	responseModel interface{}
}

/*func (builder *RedisRequestBuilder) SetQueryParams(query map[string]string) *RedisRequestBuilder {
	if query != nil {
		q := builder.Req.URL.Query()
		for k, v := range query {
			q.Add(k, v)
		}
		builder.Req.URL.RawQuery = q.Encode()
	}
	return builder
}

func (builder *RedisRequestBuilder) SetBasicAuth(username string, pass string) *RedisRequestBuilder {
	builder.Req.SetBasicAuth(username, pass)
	return builder
}

func (builder *RedisRequestBuilder) SetNormalTransport() *RedisRequestBuilder {
	builder.client = &http.Client{}
	return builder
}

func (builder *RedisRequestBuilder) SetSecurityTransport() *RedisRequestBuilder {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	builder.client = &http.Client{
		Transport: transport,
	}
	return builder
}

func (builder *RedisRequestBuilder) SetAuth(token string) *RedisRequestBuilder {
	if token != "" {
		builder.Req.Header.Set("Authorization", token)
	}
	return builder
}

func (builder *RedisRequestBuilder) SetHeader(header map[string]string) *RedisRequestBuilder {
	if header != nil {
		for k, v := range header {
			builder.Req.Header.Set(k, v)
		}

	}
	return builder
}

func (builder *RedisRequestBuilder) SetContentType(contentType string) *RedisRequestBuilder {
	if contentType == "" {
		contentType = gin.MIMEJSON
	}
	builder.Req.Header.Set("Content-type", contentType)
	return builder
}

func (builder *RedisRequestBuilder) SetResponseModel(model interface{}) *RedisRequestBuilder {
	builder.responseModel = model
	return builder
}
*/
