<!-- https://github.com/RichardLitt/standard-readme -->

<img align="right" width="100" src="https://user-images.githubusercontent.com/1582077/131063491-58e25690-6180-4e2b-9c8e-801649b2423e.png">

# RWM (Redis web manager) 

[![](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/slowrookie/redis-web-manager/blob/master/LICENSE)
[![](https://github.com/slowrookie/redis-web-manager/actions/workflows/release.yml/badge.svg)](https://github.com/slowrookie/redis-web-manager/actions/workflows/release.yml)
[![Docker Hub](https://img.shields.io/docker/pulls/slowrookie/redis-web-manager.svg)](https://hub.docker.com/r/slowrookie/redis-web-manager)


English | [简体中文](README.zh_CN.md)

RWM is a web application developed with React & Golang, used to manage Redis, and supports multi-platform operation.
## Introduction
  - [Project Screenshot](#Project-screenshot)
  - [Download & Install](#Download--install)
  - [Useage](#Useage)
  - [Project Structure](#Project-structure)
  - [Related Efforts](#Related-efforts)
  - [Maintainers](#Maintainers)
  - [Contributing](#Contributing)
  - [License](#License)

## Project Screenshot
![Project Screenshot](https://user-images.githubusercontent.com/1582077/131060729-54eeef49-9a16-4f72-8ca7-2dee2ba9a33e.jpg)


## Download & Install

`Windows`, `Linux`, `MacOS`：
  [Release](https://github.com/slowrookie/redis-web-manager/releases)

`Docker`: 
  ```sh 
    docker push slowrookie/redis-web-manager:latest 
  ````

## Useage

`Windows`, `Linux`, `MacOS`：
After decompressing the downloaded file (`*.tar.gz`), execute the `RWM` or `RWM.exe` file.

```sh
./RWM
```

`Docker`:
```sh
docker run --rm -d  -p 8080:8080/tcp slowrookie/redis-web-manager:latest
```

After the service starts, it will automatically call the default browser of the operating system and visit `http://127.0.0.1:8080`。

After the service is started for the first time, two files will be generated, one is `.config.json` to store configuration information. One is `.connections.json`, which is used to save the maintenance connection information.

## Project Structure

- `api` The directory stores the back-end logic and interface files
- `web` The Directory to store front-end project files
- `.goreleaser.yml` is `goreleaser` config file
 
Note: If there is an error in `go:embed web/build/*` in the `main.go` file, please execute `npm install & npm run build` in the `web` directory to install and build the front-end project.

## Related Efforts

- [microsoft/fluentui](https://github.com/microsoft/fluentui)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [go-redis/redis](https://github.com/go-redis/redis)

## Maintainers

[@slowrookie](https://github.com/slowrookie)

## Contributing

Feel free to dive in! [Open an issue](https://github.com/slowrookie/redis-web-manager/issues/new) or submit PRs.

## License

[MIT](LICENSE) © slowrookie
