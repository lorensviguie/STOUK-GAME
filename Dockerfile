FROM golang:latest

RUN apt-get update && apt-get install -y nodejs npm

WORKDIR /go/src/app

COPY . .

RUN npm install -D tailwindcss && npx tailwindcss init

CMD ["bash", "-c", "tail -f /dev/null"]


