FROM golang:1.18-bullseye

WORKDIR /app

RUN apt-get update
RUN apt-get install -y vim git

RUN git config --global user.name "SPO Web Team"
RUN git config --global user.email "pas@iitk.ac.in"

RUN git clone https://github.com/spo-iitk/ras-cdn.git .

RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o server

CMD [ "/app/server" ]
