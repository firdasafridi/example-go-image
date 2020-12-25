FROM golang:1.14.4

COPY hello /app/hello

CMD ["go", "run", "/app/hello/main.go"] 
