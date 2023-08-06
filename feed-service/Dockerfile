
FROM public.ecr.aws/z7v3g1r9/golang:v1.18-alpine as builder
LABEL maintainer="Abhishek Dwivedi <abhishek.dwivedi@mabel.jewelry>"
ARG PORT

RUN apk update && apk add --no-cache git
WORKDIR /app

COPY go.mod go.sum cmd ./
RUN go mod download 
RUN echo "PORT=$PORT" >> .env
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin 

FROM public.ecr.aws/z7v3g1r9/golang:v1.18-alpine
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/bin /app/.env ./ 

EXPOSE 8080

CMD ["./bin"]