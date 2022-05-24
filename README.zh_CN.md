<!-- https://github.com/RichardLitt/standard-readme -->

<img alt="Logo" style="float: right;" width="100" src="https://user-images.githubusercontent.com/1582077/158051511-e664578b-2d95-47ed-8371-0b782d7943d9.png">

# RedisWebManager

[![](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/slowrookie/redis-web-manager/blob/master/LICENSE)
[![](https://github.com/slowrookie/redis-web-manager/actions/workflows/release.yml/badge.svg)](https://github.com/slowrookie/redis-web-manager/actions/workflows/release.yml)
![](https://shields.io/github/v/release/slowrookie/redis-web-manager)
[![Docker Hub](https://img.shields.io/docker/pulls/slowrookie/redis-web-manager.svg)](https://hub.docker.com/r/slowrookie/redis-web-manager)


简体中文 | [English](README.md)

Redis Web Manager 是一款使用 React & Golang 开发的应用同时提供桌面版和Web版，用于管理Redis，支持多平台。

## 介绍
  - [项目截图](#项目截图)
  - [下载](#下载与安装)
  - [项目结构](#项目结构)
  - [相关仓库](#相关仓库)
  - [维护者](#维护者)
  - [如何贡献](#如何贡献)
  - [使用许可](#使用许可)

## 项目截图

<img style="text-align: center;" width="1136" alt="1" src="https://user-images.githubusercontent.com/1582077/165335000-556559a5-4c66-461e-b2b1-89d51f753cba.png">
<img style="text-align: center;" width="1136" alt="2" src="https://user-images.githubusercontent.com/1582077/165334999-43262777-ad6d-4bec-8ba8-5d16b3b041fc.png">
<img style="text-align: center;" width="1136" alt="3" src="https://user-images.githubusercontent.com/1582077/165334096-39063104-e56e-4d00-b40a-5e8865abe616.png">
<img style="text-align: center;" width="1136" alt="4" src="https://user-images.githubusercontent.com/1582077/165334167-7c8418ce-6d93-4a28-89fc-016b4732f7e5.png">
## 下载与安装

[Release](https://github.com/slowrookie/redis-web-manager/releases)

桌面应用包名：
  * `Windows`:  RedisWebManager_windows_amd64.exe
  * `Linux` : RedisWebManager_amd64.AppImage
  * `MacOS`: RedisWebManager.dmg

Web应用服务名:
  * `Windows`:
    * redis-web-manager_${version}-server.2_windows_amd64.tar.gz
    * redis-web-manager_${version}-server.2_windows_arm64.tar.gz
  * `Linux`:
    * redis-web-manager_${version}-server.2_linux_amd64.tar.gz
    * redis-web-manager_${version}-server.2_linux_arm64.tar.gz
  * `MacOS`:
    * redis-web-manager_${version}-server.2_darwin_amd64.tar.gz
    * redis-web-manager_${version}-server.2_darwin_arm64.tar.gz
  * `Docker`:
    * https://hub.docker.com/r/slowrookie/redis-web-manager
      ```shell
      docker run --rm -d  -p 63790:63790/tcp slowrookie/redis-web-manager:latest
      ```
  

## 项目结构

- `api` 主要功能逻辑
- `web` 前端项目文件
- `desktop` 桌面应用构建相关
- `server` 服务应用构建相关
 
## 相关仓库

- [microsoft/fluentui](https://github.com/microsoft/fluentui)
- [wails2](https://github.com/wailsapp/wails)
- [go-redis/redis](https://github.com/go-redis/redis)

## 维护者

[@slowrookie](https://github.com/slowrookie)

## 如何贡献

欢迎你的加入！[提一个 Issue](https://github.com/slowrookie/redis-web-manager/issues/new) 或者提交一个 Pull Request。

## 特别感谢

<p style="text-align: center">
   特别感谢 JetBrains 提供许可支持！<br/>
   <a href="https://jb.gg/OpenSourceSupport?from=RedisWebManager">
    <img alt="JetBrains" src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png" width="30%">
   </a>
</p>

## 使用许可

[MIT](LICENSE) © slowrookie
