
## 返回值说明
返回值为一个json字符,有四部分：
- Code 状态码，只有200表示成功
- Msg 状态码对应的信息
- err 服务端出现的异常
- data 主要响应的数据

## 接口说明
**协议：** ``HTTP``

**API HOST：** ``101.42.38.110:9999/api/v1``

# 登录注册模块

## 发送验证码
**接口：** ``/user/code``

**请求方式：** ``POST``

**请求参数：** 

| 参数名   | 类型     | 备注  |
|-------|--------|-----|
| email | string | 邮箱  |

**成功：**
```json
{
    "data": "发送成功",
    "code": 200,
    "msg": "ok",
    "err": null
}
```
**失败：**
```json
{
    "data": null,
    "code": 10001,
    "msg": "邮箱格式不正确",
    "err": null
}
```


## 注册账号
**接口：** ``/user/register``

**请求方式：** ``POST``

**请求参数：**

| 参数名      | 类型     | 备注                                               |
|----------|--------|--------------------------------------------------|
| email    | string | 邮箱                                               |
 | code     | int    | 验证码                                              |
 | password | string | 密码<br/>要求：<br/>1.长度为8到16<br/>2.必须包含数字和字母，不能有其它字符 |

**成功：**
```json
{
  "data": "注册成功",
  "code": 200,
  "msg": "ok",
  "err": null
}
```
**失败：**
```json
{
  "data": null,
  "code": 10003,
  "msg": "验证码错误",
  "err": null
}
```
```json
{
    "data": null,
    "code": 10004,
    "msg": "该邮箱已经注册",
    "err": null
}
```
## 密码登录
**接口：** ``/user/login/password``

**请求方式：** ``POST``

**请求参数：**

| 参数名      | 类型     | 备注  |
|----------|--------|-----|
| email    | string | 邮箱  |
| password | string | 密码  |


**成功：**
```json
{
 "data": {
  "email": "2985496686@qq.com",
  "nickName": "含羞草9d222589001",
  "sex": 2,
  "icon": "http://rnyrwpase.bkt.clouddn.com/default.jpg",
  "token": "b7615494-8d8d-11ed-947c-38f3ab2900a7"
 },
 "code": 200,
 "msg": "ok",
 "err": null
}
```
**data 说明：**

| 字段名      | 类型     | 备注                                                                     |
|----------|--------|------------------------------------------------------------------------|
| nickName | string | 昵称<br/>注册时会生成一个默认昵称，并且是唯一的                                             |
| sex      | int    | 性别<br/> 0代表女 <br/>   1代表男 <br/> 2代表未知(默认就是未知)                          |
| icon     | string | 头像的链接，注册时会生成一个默认头像                                                     |
| token    | string | 用于身份验证，当用户登录后，以后每次请求都需要将token携带上<br/>当用户连续30天不使用app，token将会过期，用户需要重新登录 |




## 验证码登录

**接口：** ``/user/login/code``

**请求方式：** ``POST``

**请求参数：**

| 参数名   | 类型     | 备注  |
|-------|--------|-----|
| email | string | 邮箱  |
| code  | int    | 验证码 |


**成功：**
```json
{
 "data": {
  "email": "2985496686@qq.com",
  "nickName": "含羞草9d222589001",
  "sex": 2,
  "icon": "http://rnyrwpase.bkt.clouddn.com/default.jpg",
  "token": "b7615494-8d8d-11ed-947c-38f3ab2900a7"
 },
 "code": 200,
 "msg": "ok",
 "err": null
}
```
**失败：**
```json
{
  "data": null,
  "code": 10005,
  "msg": "账号或密码不正确",
  "err": null
}
```


**失败：**
```json
{
  "data": null,
  "code": 10003,
  "msg": "验证码错误",
  "err": null
}
```

# 签到模块


















