FROM golang:alpine

RUN apk --no-cache add ca-certificates bash curl git openssh

WORKDIR /go/src/github.com/naufalridho/tax-calculator
ADD . .

RUN go install -v github.com/naufalridho/tax-calculator

COPY ./bin/wait-for-it.sh /bin/wait-for-it.sh
RUN chmod +x /bin/wait-for-it.sh

CMD wait-for-it.sh postgresql:5432 -- tax-calculator