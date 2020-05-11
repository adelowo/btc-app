# BTC Query app

![Build, run all tests and deploy Docker image](https://github.com/adelowo/queryapp/workflows/Build,%20run%20all%20tests%20and%20deploy%20Docker%20image/badge.svg)

This is a simple app that fetches BTC price in USD from Coindesk and converts to
Naira - while performing basic arithmetics.

### How to run

- Docker:

  ```
  $ docker pull adelowo/queryapp
  $ docker run -p 9080:8080 adelowo/queryapp
  ```
  ```


  ```

> Visit `http://localhost:9080/graphqiql`


- Building from source:

  ```
  $ go get github.com/adelowo/queryapp
  $ go build -o app ./cmd
  $ ./app
  ```


> By default, the server runs on `8080`. You can add your preferred port as
`PORT` environmental value such as `export PORT=2000` and the server will be
available in port 2000


### Testing, CI and CD

This repo is automatically deployed as a Docker Image via Github Actions.
Github Action also acts as a CI tool. Please take a look at https://github.com/adelowo/queryapp/actions


### Miscellaneous

The app is deployed at https://btc-exchangerate-app.herokuapp.com/graphiql

