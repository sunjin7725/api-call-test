FROM golang

RUN mkdir /app
RUN mkdir -p $GOPATH/src/github.com/sunjin7725/api-call-test

COPY ./ $GOPATH/src/github.com/sunjin7725/api-call-test

WORKDIR $GOPATH/src/github.com/sunjin7725/api-call-test

RUN go mod tidy
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g ./cmd/app.go -o ./docs
RUN go build -o /app/app $GOPATH/src/github.com/sunjin7725/api-call-test/cmd
RUN chmod +x /app/app

CMD [ "/app/app" ]
