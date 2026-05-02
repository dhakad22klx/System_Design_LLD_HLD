package proxy

/*
Proxy is a structural design pattern that provides an object that acts
as a substitute for a real service object used by a client.
A proxy receives client requests, does some work (access control, caching, etc.)
and then passes the request to a service object.

The proxy object has the same interface as a service,
which makes it interchangeable with a real object when passed to a client.

A proxy helps when you want to:

- Delay creation or loading until it’s actually needed (lazy access).
- Restrict or control access (authentication, authorization, rate limiting).
- Add cross-cutting behavior like logging, caching, retries, or monitoring without changing the original class.


Depending on the use case, the Proxy may take different forms:

Virtual Proxy: Defers creation of the real object until it’s actually needed (lazy loading).
Protection Proxy: Performs permission checks before allowing access to certain operations.
Remote Proxy: Handles communication between local and remote objects over a network.
Caching Proxy: Caches expensive results and avoids repeated calls to the real subject.
Smart Proxy: Adds logging, reference counting, or monitoring before/after method calls.

*/

import "fmt"

func TestProxyPattern() {

	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
}
