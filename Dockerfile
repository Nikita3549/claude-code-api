FROM golang:tip-alpine3.24 as build
WORKDIR /opt/api

COPY go.mod go.sum ./
COPY ./internal ./internal
COPY ./cmd ./cmd

RUN go build ./cmd/main.go

FROM node:22-slim
WORKDIR /opt/api

RUN npm install -g @anthropic-ai/claude-code
COPY ./claude-config/.claude /root/.claude
COPY ./claude-config/.claude.json /root/.claude.json

COPY --from=build /opt/api/main /opt/api/main

CMD ["./main"]
