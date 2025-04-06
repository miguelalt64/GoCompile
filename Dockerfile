# Etapa 1: Compilación
FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go mod download
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# Etapa 2: Imagen final distroless
FROM gcr.io/distroless/static-debian11:nonroot

WORKDIR /home/nonroot

COPY --from=builder /app/server .
COPY --from=builder /app/.env .
COPY --from=builder /app/config.json .

USER nonroot:nonroot

ENV ENDPOINT_SUMA=/adicionar
ENV ENDPOINT_RESTA=/restar

EXPOSE 8080

ENTRYPOINT ["./server"]
CMD ["rojo", "grande"]


#docker build -t go-app .
#docker run -p 8080:8080 go-app --name servergo

#curl http://172.17.0.2:8080/color

#curl -X POST http://172.17.0.2:8080/sumar -H "Content-Type: application/json" -d '{"a": 10, "b": 5}'
#curl -X POST http://172.17.0.2:8080/restar -H "Content-Type: application/json" -d '{"a": 10, "b": 3}'
#curl http://172.17.0.2:8080/props