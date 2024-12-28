# Calculator API (Golang)

## Purpose

The purpose of this project is to learn how to set up an API with golang on the backend.

## Commands

### `/add`
- adds two numbers

### `/subtract`
- subtracts two numbers

### `/multiply`
- multiplies two numbers

### `/divide`
- divides two numbers
- returns error when dividing by zero

### `/sum`
- adds all numbers in an array

## Packages

`net/http` - This is the http package of the standard library which I use for routing and setting up an http server. If you want to know how to use this package for advanced routing, I have a video you can check out.

`encoding/json` - This package is used for encoding and decoding JSON from the request and to the response bodies.

`log/slog`- For structured logging

`github.com/rs/cors` - For cors, if you need it.