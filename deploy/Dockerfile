FROM golang

RUN mkdir /app
RUN mkdir -p $GOPATH/src/github.com/sunjin7725/api-call-test

COPY ./ $GOPATH/src/github.com/sunjin7725/api-call-test

WORKDIR $GOPATH/src/github.com/sunjin7725/api-call-test

CMD ["/bin/bash"]

# RUN go mod tidy && go build /app/app.exe $GOPATH/src/github.com/sunjin7725/api-call-test/cmd
# RUN chmod +x /app/app.exe

# CMD [ "/app/app.exe" ]
