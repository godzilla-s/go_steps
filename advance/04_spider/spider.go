package main

import (
	"log"
	"io/ioutil"
	"net/http"
	"net/url"
	"net/http/cookiejar"
)

type Config struct {
    Url         string
    Method      string
    Header      http.Header
    Data        url.Values
    Body        []byte
}

type Spider struct {
    *Config
    Client      *http.Client
    Request     *http.Request
    Response    *http.Response
}

func NewJar() *cookiejar.Jar {
    cookieJar, _ := cookiejar.New(nil)
    return cookieJar
}

// 创建Spider
func NewSpider() *Spider {
    sp := new(Spider)
    sp.Config = new(Config)
    sp.Header = http.Header{}
    sp.Data = url.Values{}
    sp.Body = []byte{}

	//创建 Client
    sp.Client = &http.Client{
        CheckRedirect: func(r *http.Request, v []*http.Request) error {
            log.Printf("-- Redirect: %v\n", r.URL)
            return nil
        },
        Jar: NewJar(),
    }
    return sp
}

// 链式的数据初始化 <<==
func (c *Config) SetURL(url string) *Config {
    c.Url = url
    return c
}

func (c *Config) SetMethod(method string) *Config {
    c.Method = method
    return c
}

//设置表单数据
func (c *Config) SetFormParam(key, value string) {
	c.Data.Set(key, value)
	return c
}

func (c *Config) SetRefer(refer string) *Config {
	c.Header.Set("Referer", refer)
	return c
}

func (c *Config) SetUa(ua string) *Config {
	c.Header.Set("User-Agent", ua)
	return c
}

func (c *Config) Clear() *Config {
	c.Data = url.Values{}
	c.Body = []byte{}

	return c
}

func (sp *Spider) Get() {
    request, err := http.NewRequest("GET", sp.Url, nil)
    if err != nil {
        log.Fatal(err)
    }

    response, err := sp.Client.Do(request)
    if err != nil {
        log.Println("Client Do: ", err)
        return
    }

    defer response.Body.Close()

    body, _ := ioutil.ReadAll(response.Body)
    log.Println(string(body[:500]))
}

func (sp *Spider) Post() {
	var request = &http.Request{}
	
	//post 数据
	if sp.Data != nil {
		pstr := ioutil.NopCloser(strings.NewReader(sp.Data.Encode()))
		request := http.NewRequest("POST", sp.Url, pstr)
	} else {
		request := http.NewRequest("POST", sp.Url, nil)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//request.Header.Set("Content-Type", "application/json")
	//request.Header.Set("Content-Type", "text/xml")
	//request.Header.Set("Content-Type", "multipart/form-data")

	sp.Request = request

	response, err := sp.Client.Do(request)
	if err != nil {
		log.Println("Post Client Do: ", err)
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	return
}

