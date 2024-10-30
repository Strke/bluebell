# bluebell:基于gin实现的博客论坛后端系统

---

## 流程图总览
![1](./images/1.png)

&emsp;&emsp;本项目采用CLD分层结构实现了博客论坛后端系统的五个业务：注册业务、登录业务、社区业务、帖子业务、投票业务。

## 注册业务
<div align="center"> <img src="./images/2.png"/> </div>

* 1、首先客户端向服务端发送一个POST请求，POST请求中带有用户名和密码，并进行参数校验
* 2、然后通过用户名判断用户是否已经存在，若已存在则返回用户已存在错误
* 3、如果没问题的话则通过snowFlake库生成分布式ID，通过md5库对密码进行加密，然后将用户名、加密后的密码、分布式ID组成**用户记录**
* 4、将用户记录插入数据库


## 登录业务
<div align="center"> <img src="./images/3.png"/> </div>

* 1、客户端向服务端发送POST请求，请求中带有用户名和密码，并进行参数校验
* 2、通过查询数据库，判断数据库中是否存在用户且用户名和密码是否匹配
* 3、若查询到且匹配，则进行jwt的Token生成，否则，返回错误
* 4、将生成的Token返回给客户端，用于后续的鉴权

### 鉴权
在生成了AccessToken和RefreshToken并发送给客户端之后，现在客户端访问路由时需要带上Token来完成服务端的认证。

整个会话的流程如下：
* 客户端访问需要认证的接口时，携带AccessToken
* 如果AccessToken没有过期，则服务端鉴权后返回给客户端需要的数据
* 如果AccessToken过期，则客户端使用RefreshToken向服务端的刷新接口申请新的AccessToken
* 如果RefreshToken没有过期，则下发新的AccessToken
* 如果RefreshToken过期，则需要用户重新登录来获取新的AccessToken和RefreshToken


## 社区业务



## 帖子业务



## 投票业务


