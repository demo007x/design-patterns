package main

type Application struct{}

func (a *Application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "ok"
	}
	if url == "/create/user" && method == "POST" {
		return 201, "user created"
	}
	return 404, "Not OK"
}
