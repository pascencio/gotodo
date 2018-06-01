# Simple TODO application for learn Golang

## Getting Started

This instructions are available only for Linux or Mac OSX

### Environment

#### Golang

1. For installation, you must follow the instructions from [Golang Web Page](https://golang.org/doc/install).
2. Make a source code directory on your home: `mkdir -p $HOME/golang`
3. Set environment variables in your .bashrc file: `echo "export GOROOT=\"$(which go)\"" >> ~/.bashrc && echo "export GOPATH=\"\$HOME/golang\"" >> ~/.bashrc && echo "export PATH=\"\$PATH:\$GOROOT/bin:\$GOPATH/bin\"" >> ~/.bashrc && source ~/.bashrc`
4. Install [dep](https://golang.github.io/dep/docs/installation.html) for dependency management.
5. Create project directory: `mkdir -p $HOME/golang/src/github.com/pascencio`
6. Clone the project: `git clone https://github.com/pascencio/gotodo.git $HOME/golang/src/github.com/pascencio/gotodo`
7. Download dependencies (Execute this on root dir of project): `dep ensure`

#### MongoDB

1. Install [Docker](https://www.docker.com/community-edition#/download)
2. Run mongodb container: `docker run -it --rm -p 27017:27017 -e MONGODB_USERNAME=gtdusr -e MONGODB_PASSWORD=supersecret -e MONGODB_DATABASE=gotodo bitnami/mongodb`

### Run the TODO application

On the root of the project ejecute this command:

```shell
go run main.go
```

After that you can see this on your console:

```shell
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v3.3.5
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080
```

Then, if everything it's ok you can access to the app on this URL: http://localhost:8080/todo

## Dependencies

- Logrus: [Github Site](https://github.com/sirupsen/logrus)
- Mgo (Fork created by globalsign): [Github Site](https://github.com/globalsign/mgo)
- Echo: [Github Site](https://github.com/labstack/echo)
- Viper: [Github Site](https://github.com/spf13/viper)

## Next Steps - Version 1

- [x] Get Stable Version
- [x] Write Getting Started
- [ ] Include [Swag](https://github.com/swaggo/swag) and [Govalidator](https://github.com/asaskevich/govalidator)
- [ ] Implement Docker Environment
- [ ] Implement Dependencies Injection
- [ ] Implement Unit Testing
- [ ] Implement CI using Travis
- [ ] Make a Release

## Wishes

- Implement Integration Test
- Create a framewok using this expirience
- Reciebe help from the community