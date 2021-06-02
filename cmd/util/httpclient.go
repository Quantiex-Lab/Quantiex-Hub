package util

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/golang/glog"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var netClient *http.Client
var httpOnce sync.Once
var ProxyIPPort string

func GetHTTPClient() *http.Client {
	httpOnce.Do(func() {
		var proxy = func(_ *http.Request) (*url.URL, error) {
			if ProxyIPPort == "" {
				return nil, nil
			}
			return url.Parse("http://" + ProxyIPPort)
		}

		var netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			Proxy: proxy,
		}

		netClient = &http.Client{
			Timeout:   time.Second * 5,
			Transport: netTransport,
		}

	})

	return netClient
}

func HttpGet(method string, url string, data string, cookies []*http.Cookie, headers map[string]string, httpClient *http.Client) (string, *http.Response, error) {
	var buf *bytes.Buffer
	if data != "" {
		var jsonstr = []byte(data)
		buf = bytes.NewBuffer(jsonstr)
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return "", nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		err := fmt.Errorf("HTTP get failed. err = %v, url = %s", err, url)
		return "", nil, err
	}

	var reader io.ReadCloser
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			err = fmt.Errorf("HTTP get failed. err = %v, url = %s", err, url)
			return "", nil, err
		}
	} else {
		reader = resp.Body
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		err := fmt.Errorf("HTTP read body failed. err = %v, url = %s", err, url)
		return "", nil, err
	}
	bodyStr := string(body)
	if len(bodyStr) > 50 {
		glog.V(5).Infof("HTTP get %s OK.\n Body is %s ...\n", url, bodyStr[0:50])
	} else {
		glog.V(5).Infof("HTTP get %s OK.\n Body is %s\n", url, bodyStr)
	}
	return bodyStr, resp, nil
}
