<!-- https://github.com/RichardLitt/standard-readme -->

# RWM (Reids web manager)

[![](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/slowrookie/redis-web-manager/blob/master/LICENSE)
[![](https://github.com/slowrookie/redis-web-manager/actions/workflows/release.yml/badge.svg)](https://github.com/slowrookie/redis-web-manager/actions/workflows/release.yml)


RWM 是一款使用 React & Golang 开发的Web应用，用于管理Redis，支持多平台运行。

## 介绍
  - [下载与安装](#下载与安装)
  - [使用说明](#使用说明)
  - [项目结构](#项目结构)
  - [相关仓库](#相关仓库)
    - [前端](#前端)
    - [后端](#后端)
  - [维护者](#维护者)
  - [如何贡献](#如何贡献)
  - [使用许可](#使用许可)

## 下载与安装

项目使用 [Github Actions](https://github.com/slowrookie/redis-web-manager/actions) 编译发布，直接在 [Release](https://github.com/slowrookie/redis-web-manager/releases) 下载对应操作系统版本的压缩包，解压后即可运行。

## 使用说明

解压后执行执行 `RWM` 或者 `RWM.exe` 文件。
```sh
./RWM
```

服务启动后会自动调用操作系统默认浏览器，并访问 `http://localhost:9090`。

服务第一次启动后会产生两个文件，一个是`.config.json`用于存放配置信息。一个是`.connections.json`,用于保存维护的连接信息。

## 项目结构

- `api` 目录存放后端逻辑以及接口文件
- `web` 目录存放前端项目文件
- `.gitignore` 是 `Git` 用于忽略的配置文件。
- `.goreleaser.yml` 是 `goreleaser` 的配置文件。
 
注意：如果在 `main.go` 文件中 `go:embed web/build/*` 错误，请先在 `web` 目录执行 `npm install & npm run build` 安装和构建前端项目.

## 相关仓库

### 前端

- [microsoft/fluentui](https://github.com/microsoft/fluentui)

### 后端

- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [go-redis/redis](https://github.com/go-redis/redis)

## 维护者

[@slowrookie](https://github.com/slowrookie)

## 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/slowrookie/redis-web-manager/issues/new) 或者提交一个 Pull Request。

### 贡献者

感谢以下参与项目的人：
<a href="graphs/contributors"><img src="https://opencollective.com/standard-readme/contributors.svg?width=890&button=false" /></a>

## 使用许可

[MIT](LICENSE) © slowrookie