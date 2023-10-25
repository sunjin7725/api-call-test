FROM golang

RUN mkdir /app
RUN mkdir -p $GOPATH/src/github.com/sunjin7725/api-call-test

COPY ../* $GOPATH/src/github.com/sunjin7725/api-call-test

WORKDIR $GOPATH/src/github.com/sunjin7725/api-call-test
RUN git clone https://github.com/sunjin7725/api-call-test

RUN go mod tidy && go build /app/app $GOPATH/src/github.com/sunjin7725/api-call-test/cmd
RUN chmod +x /app/app

CMD [ "/app/app" ]
