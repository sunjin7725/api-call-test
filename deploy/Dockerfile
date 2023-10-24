FROM golang

RUN mkdir /app

COPY ../ /app
RUN go tidy && go build /app/app /app/cmd
RUN chmod +x /app/app

CMD [ "/app/app" ]
