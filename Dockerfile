FROM golang:1.14

WORKDIR /go/src/albert02lowis/alw-golang-redis-geoloc

COPY . .

RUN go build
RUN go install -i github.com/albert02lowis/alw-golang-redis-geoloc
#RUN echo $PATH
#RUN mv alw-golang-redis-geoloc.exe /go/bin

CMD ["alw-golang-redis-geoloc"]

EXPOSE 8080