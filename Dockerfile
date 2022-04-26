# Dockerfile

FROM cosmtrek/air

RUN mkdir -p /app/dist
COPY ./ /app/
WORKDIR /app/

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["air"]