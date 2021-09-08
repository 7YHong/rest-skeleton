package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"rest-skeleton/core/di"
	"time"
)

type BodyLogWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ReqLogMiddleware(c *gin.Context) {
	blw := &BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	start := time.Now()

	reqInfo := getReqInfo(c)
	c.Set("reqInfo", &reqInfo)
	reqJson, _ := json.Marshal(reqInfo)
	//reqJson, _ := json.MarshalIndent(reqInfo, "", "    ")

	c.Next()

	di.Zap().Infof("Send HTTP response, Req uri: %s, \nReqInfo: %v, \nResp code: %v, \nTime: %v, \nRespbody: %v",
		c.Request.RequestURI, string(reqJson), c.Writer.Status(), time.Since(start), blw.Body.String())

}

type ReqInfo struct {
	Headers    map[string]string `json:"headers"`
	Method     string            `json:"method"`
	Params     gin.Params        `json:"params"`
	RemoteAddr string            `json:"remote_addr"`
	Form       map[string]string `json:"form"`
	Data       string            `json:"data"`
}

func getReqInfo(c *gin.Context) ReqInfo {
	c.Request.ParseMultipartForm(32 << 20)
	r := ReqInfo{
		Headers:    map[string]string{},
		Form:       map[string]string{},
		Params:     c.Params,
		RemoteAddr: c.ClientIP(),
		Method:     c.Request.Method,
	}
	for k, v := range c.Request.Header {
		r.Headers[k] = v[0]
	}
	for k, v := range c.Request.Form {
		r.Form[k] = v[0]
	}
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	r.Data = string(bodyBytes)

	return r
}
