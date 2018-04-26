NEWS--每日golang最新新闻获取
===========================


[TOC]



### 一、如何使用

#### 1.1、一次性输出

```python
git clone https://github.com/zhuima/daily_news.git
cd news
glide install
make
./bin/news
```

#### 1.2、定时任务

```python
00 08 * * * flock -xn /tmp/news.lock -c '/usr/local/bin/news 2>&1 >> /tmp/news_`/bin/date +\%Y\%m\%d`.log'
```


### 二、实现方式

- goquery爬取网页并提取自己想要的信息
- iconv-go 解决网页乱码问题
    - 现在使用github.com/axgle/mahonia实现
- 收集TOP10信息，然后通过钉钉发送到相关群组


### 三、注意事项

- 如何通过类元素定位相应的html标签来获取自己想要的数据
- google的标签貌似是来回换的
- 钉钉通知是markdown模式，这样的话就可以直接点击跳转，不用再粘贴复制了


### 四、参(照)考(抄)文档


[golang80行代码钉钉群机器人舆情监控](https://segmentfault.com/a/1190000013241676)

[Go实现的抓取google最近24小时新闻并发送邮件](http://blog.chenmiao.cf/2018/01/11/go%E5%AE%9E%E7%8E%B0%E7%9A%84%E6%8A%93%E5%8F%96google%E6%9C%80%E8%BF%9124%E5%B0%8F%E6%97%B6%E6%96%B0%E9%97%BB%E5%B9%B6%E5%8F%91%E9%80%81%E9%82%AE%E4%BB%B6/)

[中文乱码问题解决](http://www.phperz.com/article/17/1103/352014.html)

[钉钉机器人订阅百度新闻v2](https://gist.github.com/mojocn/9b18db2c99b01e49ce6afbbb2322e07a)

[YAML support for the Go language](https://github.com/go-yaml/yaml)

[汽车之家爬虫](https://github.com/go-crawler/car-prices/blob/master/downloader/download.go)

### 五、结果大概样式

```python
1 - [ericchiang (Eric Chiang) · GitHub](https://github.com/ericchiang)

2 - [kubernetes/pkg/kubectl at master · kubernetes/kubernetes · GitHub](https://github.com/kubernetes/kubernetes/tree/master/pkg/kubectl)

3 - [kooper/clientset.go at master · spotahome/kooper · GitHub](https://github.com/spotahome/kooper/blob/.../k8s.../clientset.go)

4 - [k8s与caas--容器云caas平台的落地实践- k8s小店- SegmentFault 思否](https://segmentfault.com/a/1190000013855767)

5 - [开源项目- Go语言中文网- Golang中文社区](https://studygolang.com/projects)

6 - [技术晨读- Go语言中文网- Golang中文社区](https://studygolang.com/readings)

7 - [对一段Go语言代码输出结果的简要分析| Tony Bai](https://tonybai.com/.../the-analysis-of-output-results-of-a-go-code-snippet/)

8 - [Create versatile Microservices in Golang - part 8 of 10 part series](https://ewanvalentine.io/microservices-in-golang-part-8/)

9 - [deployment - Go libraries and apps - GolangLibs.com](https://golanglibs.com/top?page=76&q=deployment)

10 - [repo/gentoo.git - Official Gentoo ebuild repository (formerly known ...](https://gitweb.gentoo.org/repo/gentoo.git/commit/?id...)
```


### 六、TODO

- [ ] 多关键字搜索，不同的板块的信息
   
- [ ] 配置文件和代码分离, 引入yaml解析
- 历史信息持久化
    - 使用redis来实现存储
- DashBoard展示，每天来了花十分钟阅读下今日时政相关信息
    - vue + iris
- channel使用
- 容器化


### 七、注意事项

- Makefile命令要以tab为开头，要注意
