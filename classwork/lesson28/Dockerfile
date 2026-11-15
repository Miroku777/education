FROM golang:1.27
WORKDIR /src
COPY . /src
RUN go build -o /bin/myapp ./main.go
CMD ["/bin/myapp"]


# docker run -d -p 8080:8080 --name itsMy_project my_project
# docker run -d -p 8080:8080 -e VERSION=2.1.1 --name go-server-2 go-docker-server