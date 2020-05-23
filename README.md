## Kr
[![sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?logo=sourcegraph)](https://sourcegraph.com/github.com/alimy/kr)

Kr provide some develop help toolkits.

### Feature
* support generate grpc style template project use ```kr new -d examples -s grpc``` command.
* support generate [gin](https://github.com/go-gonic/gin) style template project use ```kr new -d examples -s gin``` command.
* support generate [go-chi](https://github.com/go-chi/chi) style template project use ```kr new -d examples -s chi``` command.
* support generate [mux](https://github.com/gorilla/mux) style template project use ```kr new -d examples -s mux``` command.
* support generate [httprouter](https://github.com/julienschmidt/httprouter) style template project use ```kr new -d examples -s httprouter``` command.
* support generate [echo](https://github.com/labstack/echo) style template project use ```kr new -d examples -s echo``` command.
* support generate [macaron](https://github.com/go-macaron/macaron) style template project use ```kr new -d examples -s macaron``` command.
* support generate [iris](https://github.com/kataras/iris) style template project use ```kr new -d examples -s iris``` command.
### Usage
```bash
% go get github.com/alimy/kr@latest
% mkdir -p $GOPATH/src/github.com/alimy
% cd $GOPATH/src/github.com/alimy
% kr new -d examples -s grpc -p github.com/alimy/examples
```
