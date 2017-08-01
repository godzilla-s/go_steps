package main


func spider_get() {
	spider := NewSpider()
	spider.SetURL("http://www.baidu.com").SetMethod("GET")
	spider.Get()
}

func spider_post() {
	spider := NewSpider()
	spider.SetUrl("https://www.zhihu.com/login/email").SetRefer("https://www.zhihu.com/").SetUa(boss.RandomUa())
	spider.SetFormData("email", "zuvakin2010").SetFormData("password", "zwj2010")
	
	spider.Post()
}

func main() {
	spider_get()	
}

