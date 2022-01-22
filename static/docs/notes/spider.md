关键字：crawler

爬虫系统要素

- 高匿代理池
- 验证码破解工具、用户池
- 浏览器渲染器 如 chromeless/puppteer/NightmareJS/Selenium/PhantomJS

- 爬虫核心（网站遍历）
- 爬虫任务管理器
- 日志队列，消息队列

- 内网数据存储，如 Hbase、TiDB、数据仓库
- 数据分析、文本解析、数据清洗


[无头浏览模式]
puppteer 似乎需要主机具有界面 和 Chromeium
https://www.zhihu.com/question/314668782


https://github.com/henrylee2cn/pholcus

无头浏览器底层是:
WebDriver / DevTools / devtools-protocol

golang技术栈推荐的是 chromedp 和 go-rod，后者提供了较完善的demo，且编程接口更直观

无头浏览器还是可以被反爬，因此还可以采用按键精灵方式，golang技术栈实现可以考虑 [robotgo](https://github.com/go-vgo/robotgo)

Autoit、AutoHotkey