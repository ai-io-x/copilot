version: '3.8'

services:
  client:
    build:
      context: ./client
      dockerfile: Dockerfile
    ports:
      - "3000:3000"
    depends_on:
      - rest-server

  rest-server:
    build:
      context: ./rest-server
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    depends_on:
      - db-server

  db-server:
    build:
      context: ./db-server
      dockerfile: Dockerfile
    ports:
      - "5432:5432"