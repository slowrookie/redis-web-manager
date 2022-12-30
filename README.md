<!-- https://github.com/RichardLitt/standard-readme -->

<img alt="Logo" align="right" style="float: right;" width="100" src="https://user-images.githubusercontent.com/1582077/158051511-e664578b-2d95-47ed-8371-0b782d7943d9.png">

# RedisWebManager 

[![](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/slowrookie/redis-web-manager/blob/master/LICENSE)
[![](https://github.com/slowrookie/redis-web-manager/actions/workflows/release.yml/badge.svg)](https://github.com/slowrookie/redis-web-manager/actions/workflows/release.yml)
![](https://shields.io/github/v/release/slowrookie/redis-web-manager)
[![Docker Hub](https://img.shields.io/docker/pulls/slowrookie/redis-web-manager.svg)](https://hub.docker.com/r/slowrookie/redis-web-manager)
![](https://visitor-badge.glitch.me/badge?page_id=slowrookie.redis-web-manager)


English | [简体中文](README.zh_CN.md)

Redis Web Manager is an application developed with React & Golang that provides both desktop and web versions for managing Redis and supports multiple platforms.

## Introduction
  - [Project Screenshot](#Project-screenshot)
  - [Download & Install](#Download--install)
  - [Project Structure](#Project-structure)
  - [Related Efforts](#Related-efforts)
  - [Maintainers](#Maintainers)
  - [Contributing](#Contributing)
  - [License](#License)

## Project Screenshot

<img style="text-align: center;" width="1136" alt="1" src="https://user-images.githubusercontent.com/1582077/165335000-556559a5-4c66-461e-b2b1-89d51f753cba.png">
<img style="text-align: center;" width="1136" alt="2" src="https://user-images.githubusercontent.com/1582077/165334999-43262777-ad6d-4bec-8ba8-5d16b3b041fc.png">
<img style="text-align: center;" width="1136" alt="3" src="https://user-images.githubusercontent.com/1582077/165334096-39063104-e56e-4d00-b40a-5e8865abe616.png">
<img style="text-align: center;" width="1136" alt="4" src="https://user-images.githubusercontent.com/1582077/165334167-7c8418ce-6d93-4a28-89fc-016b4732f7e5.png">

## Download & Install

[Release](https://github.com/slowrookie/redis-web-manager/releases)

Desktop：
* `Windows`:  RedisWebManager_windows_amd64.exe
* `Linux` : RedisWebManager_amd64.AppImage
* `MacOS`: RedisWebManager.dmg

Web Server:
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


## Project Structure

- `api` main function logic
- `web` frontend project
- `desktop` desktop build
- `server` web server build

## Related Efforts

- [microsoft/fluentui](https://github.com/microsoft/fluentui)
- [wails2](https://github.com/wailsapp/wails)
- [go-redis/redis](https://github.com/go-redis/redis)

## Maintainers

[@slowrookie](https://github.com/slowrookie)

## Contributing

Welcome to join us! [Open an issue](https://github.com/slowrookie/redis-web-manager/issues/new) or submit PRs.

## Special Thanks

<p align="center" style="text-align: center">
   Special thanks to JetBrains for licensing support！<br/>
   <a href="https://jb.gg/OpenSourceSupport?from=RedisWebManager">
    <img alt="JetBrains" src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png" width="30%">
   </a>
</p>

## License

[MIT](LICENSE) © slowrookie
