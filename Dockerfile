from golang:1.12.13

ENV REPO_URL=mysvc
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL
ENV WORKPATH=$APP_PATH/src

COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go get
RUN go build -o mysvc .

EXPOSE 8080

CMD ["./mysvc"]
