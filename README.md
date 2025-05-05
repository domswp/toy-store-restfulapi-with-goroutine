# Go Routine Toy Store API

A simple RESTful API built in Golang — now with goroutines

## 🎯 Purpose

This project is a personal learning sandbox for:
- Understanding how RESTful APIs work in Go
- Exploring goroutines for handling concurrent processes
- Practicing basic struct, routing, and JSON handling

If you're learning Go, this repo might help you see how the basics come together.

## 🛠️ What’s Used

- `net/http` — for setting up routes and server
- `encoding/json` — to encode/decode JSON data
- Goroutines — to handle concurrent logic
- Edited with Neovim (because terminal = peace)

## 🧸 API Endpoints

| Method | Endpoint     | Function               |
|--------|--------------|------------------------|
| GET    | /toys        | List all toys          |
| GET    | /toys/{id}   | Get a toy by ID        |
| POST   | /toys        | Add a new toy          |
| PUT    | /toys/{id}   | Update toy by ID       |
| DELETE | /toys/{id}   | Delete toy by ID       |

## ⚡ Example Request

```bash
curl -X POST http://localhost:8080/toys \
  -H "Content-Type: application/json" \
  -d '{"id":1,"name":"Freya Bot","price":15000}'
```


## 🌀 Why Goroutines?

Sometimes your server needs to multitask. Goroutines help handle requests without blocking the whole system. In this project, they simulate async behavior when creating a toy.


## 📌 Note

If you find it helpful, awesome. If not, no worries — I’m just figuring things out.

## Output

still same like this repo (https://github.com/domswp/Toko-mainan-RESTful-APIs)

## Depedencies

for this repo you must install this depedencies, using this command
```bash
go get github.com/gorilla/mux
```
