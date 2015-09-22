# go-quickstart
Quick setup and running of Golang backend service on Google App Engine.

### Requirements
1. [Git](https://git-scm.com/)
2. [Go compiler](https://golang.org/dl/)
3. [Google App Engine SDK for Go](https://cloud.google.com/appengine/downloads)

### Quick steps
1. Install and the requirement softwares.
2. Make sure GOROOT and GOPATH is setup correctly.
3. Get the go-quickstart source code and its dependencies. By running the below commands

    ```
    go get github.com/iamjay/go-quickstart
    go get github.com/gorilla/context
    go get github.com/gorilla/mux
    go get github.com/dgrijalva/jwt-go
    ```
4. Run the app locally using Google App Engine SDK for Go.

    ```
    goapp serve appengine
    ```
5. Create a new project on [Google Cloud Platform](https://console.developers.google.com).
6. Change application name in appengine/app.yaml from go-quickstart to be your new project.
7. Deploy

    ```
    goapp deploy appengine
    ```
