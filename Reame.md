### 项目介绍

这是一个采用前后端分离开发的项目，前端采用 Vue 开发、后端采用 gin + gorm 开发。


### 前端部分

#### 项目运行

**由于涉及大量的 ES6/7 等新属性，node 需要 6.0 以上版本**

```shell
git clone git@github.com:ruanjiancheng/StudentManageSystem.git

cd /StudentManageSystem/student_client

npm install

npm run serve
```

#### 技术栈

- vuex
- router
- axios
- element ui
- sessionStorage

#### 项目介绍

采用 vue 2.0 开发,通过调用后端提供的数据接口实现数据的动态渲染. 项目默认端口号 8080

- 使用监视器,以及得益于 Mybatis 提供动态 SQL 功能,实现动态搜索功能
- 同时实现了基于前端和后端的数据分页功能
- 使用 router 配置路由,实现不同用户类型导航栏的动态渲染
- 使用 axios 异步加载后端数据
- 使用 element ui 实现表单的前端校验功能
- 使用 sessionStorage 实现登录拦截

#### 实现的功能

1. admin
   1. 实现对教师, 学生, 课程的 CRUD
   2. 实现对教师业务以及学生业务的全方位控制
2. teacher 
   1. 实现查询我开设的课程, 以及选择我课程的学生信息
   2. 对学生成绩的登陆
3. student
   1. 实现选课退课的功能
   2. 实现成绩查询的功能



### 后端部分

#### 项目运行

```shell
git clone git@github.com:Kafumio/Course_Selection_System.git

cd /Course_Selection_System/css_server

go run main.go
```

#### 技术栈

- gin
- gorm
- viper
- JDBC
- servlet

#### 项目介绍

采用 Restful 风格开发,采用 CrossOrigin 解决跨域问题. 采用 gorm 实现动态 SQL 的功能. 

项目启动在 9090 端口, 可以使用 YAML 文件配置

### 数据库设计

建表代码：

```sql
CREATE DATABASE css;

USE css;

CREATE TABLE `s` (
    `sid`      BIGINT AUTO_INCREMENT,
    `sname`    VARCHAR(30) NOT NULL,
    `password` VARCHAR(30) NOT NULL,
    PRIMARY KEY (`sid`)
);

CREATE TABLE `c` (
    `cid`     BIGINT AUTO_INCREMENT,
    `cname`   VARCHAR(30) NOT NULL,
    `ccredit` TINYINT,
    PRIMARY KEY (`cid`)
);

CREATE TABLE `t` (
    `tid`      BIGINT AUTO_INCREMENT,
    `tname`    VARCHAR(30) NOT NULL,
    `password` VARCHAR(30) NOT NULL,
    PRIMARY KEY (`tid`)
);

CREATE TABLE `ct` (
    `ctid`  BIGINT AUTO_INCREMENT,
    `cid`   BIGINT,
    `tid`   BIGINT,
    `term`  CHAR(30) NOT NULL,
		UNIQUE  INDEX(`cid` ,`tid`,`term`),
    FOREIGN KEY (`cid`) REFERENCES c(`cid`),
    FOREIGN KEY (`tid`) REFERENCES t(`tid`),
    PRIMARY KEY (`ctid`)
);

CREATE TABLE `sct` (
    `sctid` BIGINT AUTO_INCREMENT,
    `sid`   BIGINT,
    `cid`   BIGINT,
    `tid`   BIGINT,
    `grade` FLOAT,
    `term`  CHAR(30),
		UNIQUE  INDEX(`sid`, `cid` ,`tid`,`term`),
    FOREIGN KEY (`sid`) REFERENCES s(`sid`),
    FOREIGN KEY (`tid`) REFERENCES ct(`tid`),
    FOREIGN KEY (`cid`) REFERENCES ct(`cid`),
    PRIMARY KEY (`sctid`)
);


### 项目存在的问题

- 由于是第一次编写 Vue 项目, 代码复用做得并不是很好. 导致许多组件代码量巨大。
- 动态搜索导致前端频繁调用数据接口, 使得性能降低。