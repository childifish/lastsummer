# last Summer网盘🌞

### 实现了

- 文件上传 

  单个/多个文件

- 文件下载

  普普通通的文件下载

- 登录注册

  用到了路由分组和中间件，对中间件的payload里的敏感字段进行了md5加密

- 文件分享

  分享下载链接

### 还实现了

- 二维码分享链接

  

### Docker😡

29号中午基本就写完了，部署在服务器上试了一下，感觉差不多了就开始准备docker

然后一转眼就晚上了，现在22：54

docker还没整明白，一下午就多写个二维码分享链接

docker课上完之后我是安装好了的，但今天一打开就

```
docker: An error occurred trying to connect: Post http://%2F%2F.%2Fpipe%2Fdocker_engine/v1.24/contai
```

然后搜到了

```
cd "C:\Program Files\Docker\Docker"
./DockerCli.exe -SwitchDaemon
```

结果完全8行，

然后觉得是自己系统的问题，一直用的家庭中文版，之前安装的时候临时改了注册表

之后删了本地的docker转用Ubuntu服务器

...（省略500字在服务器上的折腾）

然后大概卡在了部署，我之前一直是在本地编译好了再用Xftp拖进服务器，一直没在linux环境下编译过，害怕出现各种导包的问题，转头去升级到了win10专业版，试图在本地完成docker打包部署

之后转用win10，年幼的lcyh并不知道，自己即将面临怎样的地狱

顺利地下载安装好了Docker Desktop，配置好了阿里云给的加速后，成功下好了Nginx和Mysql的镜像

又瞎折腾了半天

在网上找了很久的教程，都比较抽像，大多都是tomcat和spring boot在Docker上的使用，我不由感叹起golang在国内的冷清

最终，我还是找到了相应的教程，写好了Dockerfile之后开始docker build，经过漫长的下载后给我当头一棒的是

```go
go: github.com/dgrijalva/jwt-go@v3.2.0+incompatible: Get "https://proxy.golang.org/github.com/dgrijalva/jwt-go/@v/v3.2.0+incompatible.mod": dial tcp 172.217.27.145:443: connect: connection refused

```

okok,这个`jwt-go`库在拉下来的时候我就有种不详的预感，其他的版本号大多都是`v1.x.x`，就它是`v3.2.0`，并且当时也是我手动go get而不是gomod直接拉下来的，而且它后面跟着一个`+incompatible`

然后百度必应谷歌轮番搜索，硬是没找到跟我相同错误的，我恍然大悟--应该是proxy代理的问题，因为

> ```
> https://proxy.golang.org/github.com/dgrijalva/jwt-go/@v/v3.2.0+incompatible.mod
> ```

`proxy.golang.org`明显不是我之前设置的`goproxy.cn`

重新设置proxy，然后

> ```
> warning: go env -w GOPROXY=... does not override conflicting OS environment variable
> ```

不能覆盖OS级别的环境变量，csdn上也没找到点阳间的解决方法，然后卡了很就，不知怎么就go env看了下，发现环境变量里`goproxy = https://goproxy.cn`，喜出望外，觉得马上就脱离苦海了，结果：

```go
go: github.com/dgrijalva/jwt-go@v3.2.0+incompatible: Get "https://proxy.golang.org/github.com/dgrijalva/jwt-go/@v/v3.2.0+incompatible.mod": dial tcp 172.217.27.145:443: connect: connection refused

```

和刚才一模一样的错误

至此，我的心中再无悲喜，只想睡觉



### 最后😥

最后还是记录一下 想了一点点但因为和docker抗争没时间写的进阶功能吧

- 一次性快传：这个指的似乎是一次性上传多个文件（吗？）如果没有理解错的话，用到gin框架的MultipartForm()可以很轻松的做到

  如果指的多线程上传...暂时没什么思路

- 下载限速：用户的下载不在同一个路由组里，写一个限流器，不是vip用户的下载路由每隔几毫秒会休眠一毫秒之类的

- 加密分享链接：这个我问了学长，但他好像有事没回复我..

  如果是把url的一部分加密的话，用sha256就能加密解密

  如果是需要输入分享码之类的才能下载，保存的话，可以用一张分享表保存文件路径，分享人，过期时间和key，然后设置中间件

  

  

### 接口文档🍭

https://web.postman.co/collections/9499756-cd765c69-7796-4b57-a0c5-4869f0cb73bf?version=latest&workspace=3f2a82ee-bd11-4225-87cb-48343ee6e772



###### 被docker磨平了棱角

###### 希望天堂没有docker

