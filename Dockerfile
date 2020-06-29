FROM golang:alpine as go-builder
WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
COPY ./server.go ./
RUN go build -o server .


FROM node:12 AS builder
WORKDIR /app
COPY . .
# RUN cd client && rm yarn.lock && npm install && npm run build && cd ..
RUN rm yarn.lock
RUN npm install
RUN npm run build


FROM node:12 AS production
WORKDIR /app
COPY --from=builder ./app/build .
COPY --from=go-builder ./app .
RUN ls
# EXPOSE 3000
CMD ["./server"]
