package handlercontrol

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/narawichsaphimarn/handlercontrol/models"
)

type httpRequest struct {
	TimeOut time.Duration
	Header  struct {
		ContentType   string
		Authorization struct {
			Token string
			Type  string
		}
		ContentEncoding string
		ContentLength   string
		From            string
		Host            string
		Date            string
		Cookie          string
		HTTP2Settings   string
		Origin          string
		Warning         string
		UserAgent       string
		XRequestID      string
	}
}

func NewHttpRequest(_value models.HttpRequestForm) *httpRequest {
	return &httpRequest{
		TimeOut: _value.TimeOut,
		Header: struct {
			ContentType   string
			Authorization struct {
				Token string
				Type  string
			}
			ContentEncoding string
			ContentLength   string
			From            string
			Host            string
			Date            string
			Cookie          string
			HTTP2Settings   string
			Origin          string
			Warning         string
			UserAgent       string
			XRequestID      string
		}{
			ContentType: _value.ContentType,
			Authorization: struct {
				Token string
				Type  string
			}{
				Token: _value.Authorization.Token,
				Type:  _value.Authorization.Type,
			},
			ContentEncoding: _value.ContentEncoding,
			ContentLength:   _value.ContentLength,
			From:            _value.From,
			Host:            _value.Host,
			Date:            _value.Date,
			Cookie:          _value.Cookie,
			HTTP2Settings:   _value.HTTP2Settings,
			Origin:          _value.Origin,
			Warning:         _value.Warning,
			UserAgent:       _value.UserAgent,
			XRequestID:      _value.XRequestID,
		},
	}
}

func (srv *httpRequest) ClientRequest(url string, method string, payload []byte) ([]byte, error) {
	var request *http.Request
	var err error
	var body *bytes.Buffer
	var timeout time.Duration
	if !checkTimeDurationEmpty(srv.TimeOut) {
		timeout = time.Duration(srv.TimeOut * time.Second)
	} else {
		// ?If parameter `timeout` is null set defualt 1 minut.
		timeout = time.Duration(60 * time.Second)
	}
	client := &http.Client{
		Timeout: timeout,
	}
	switch method {
	case GET:
		body = bytes.NewBuffer(nil)
	case POST:
		body = bytes.NewBuffer(payload)
	case PUT:
		body = bytes.NewBuffer(payload)
	case DELETE:
		body = bytes.NewBuffer(nil)
	default:
		return nil, fmt.Errorf("%s", "Error step set method becouse variable `method` is empty value.Please input `method` value to parameter.")
	}
	request, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("%s{%v}", "Error step wraps context background HTTP.Error msg : ", err)
	}
	if !checkStringEmpty(srv.Header.Authorization.Token) && !checkStringEmpty(srv.Header.Authorization.Type) {
		request.Header.Add(Authorization, strings.Join([]string{srv.Header.Authorization.Token, srv.Header.Authorization.Type}, VALUE_SPACE))
	} else {
		request.Header.Add(Authorization, srv.Header.Authorization.Token)
	}
	if !checkStringEmpty(srv.Header.ContentType) {
		request.Header.Add(ContentType, srv.Header.ContentType)
	}
	if !checkStringEmpty(srv.Header.ContentEncoding) {
		request.Header.Add(ContentEncoding, srv.Header.ContentEncoding)
	}
	if !checkStringEmpty(srv.Header.ContentLength) {
		request.Header.Add(ContentLength, srv.Header.ContentLength)
	}
	if !checkStringEmpty(srv.Header.Cookie) {
		request.Header.Add(Cookie, srv.Header.Cookie)
	}
	if !checkStringEmpty(srv.Header.Date) {
		request.Header.Add(Date, srv.Header.Date)
	}
	if !checkStringEmpty(srv.Header.From) {
		request.Header.Add(From, srv.Header.From)
	}
	if !checkStringEmpty(srv.Header.HTTP2Settings) {
		request.Header.Add(HTTP2Settings, srv.Header.HTTP2Settings)
	}
	if !checkStringEmpty(srv.Header.Host) {
		request.Header.Add(Host, srv.Header.Host)
	}
	if !checkStringEmpty(srv.Header.Origin) {
		request.Header.Add(Origin, srv.Header.Origin)
	}
	if !checkStringEmpty(srv.Header.UserAgent) {
		request.Header.Add(UserAgent, srv.Header.UserAgent)
	}
	if !checkStringEmpty(srv.Header.Warning) {
		request.Header.Add(Warning, srv.Header.Warning)
	}
	if !checkStringEmpty(srv.Header.XRequestID) {
		request.Header.Add(XRequestID, srv.Header.XRequestID)
	}

	request.Header.Set("Content-type", srv.Header.ContentType)
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%s{%v}", "Error step HTTP sent request.Error msg : ", err)
	}
	payloadBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("%s{%v}", "Error step read body.Error msg : ", err)
	}
	response.Body.Close()
	return payloadBody, nil
}

func checkStringEmpty(_value string) bool {
	return _value == ""
}

func checkTimeDurationEmpty(_value time.Duration) bool {
	return _value == 0
}
