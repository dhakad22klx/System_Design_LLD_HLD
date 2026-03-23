package main


import (
	"fmt"
)


type DatabaseClient struct {
	maxConnections int
	retryAttempts  int
}

func NewDatabaseClient(maxConnections, retryAttempts int) *DatabaseClient {
	return &DatabaseClient{maxConnections: maxConnections, retryAttempts: retryAttempts}
}

// Clean public API: the caller's view
func (d *DatabaseClient) Connect(host string, port int) {
	d.openSocket(host, port)
	d.authenticate()
	d.initializeConnectionPool()
}

func (d *DatabaseClient) Query(sql string) string {
	parsedQuery := d.parseQuery(sql)
	return d.executeWithRetry(parsedQuery)
}

// Hidden complexity: unexported methods (that is Abstraction)
func (d *DatabaseClient) openSocket(host string, port int) {}
func (d *DatabaseClient) authenticate()                    {
	fmt.Println("Authenticated")
}
func (d *DatabaseClient) initializeConnectionPool()        {}
func (d *DatabaseClient) parseQuery(sql string) string     { return sql }
func (d *DatabaseClient) executeWithRetry(query string) string {
	// Retry logic hidden from caller
	return "Query : "+ query + "Result"
}

// func main() {

// 	var dbClient *DatabaseClient = NewDatabaseClient(10, 3)
// 	dbClient.Connect("localhost", 8080)
// }
