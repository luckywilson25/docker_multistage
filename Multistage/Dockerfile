FROM golang:1.22.0-alpine AS base

FROM base AS builder
WORKDIR /app
COPY . .
RUN go get
RUN go build -o ./bin/app ./main.go

FROM base AS runner
COPY --from=builder /app/bin/app /app/bin/app
EXPOSE 8000
ENTRYPOINT [ "/app/bin/app" ]