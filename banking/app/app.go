package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/neoyewchuan/RestDevGo/banking/domain"
	"github.com/neoyewchuan/RestDevGo/banking/logger"
	"github.com/neoyewchuan/RestDevGo/banking/service"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASS",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Error(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func Start() {

	sanityCheck()
	// define own multiplexer
	router := mux.NewRouter()

	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDB(dbClient)
	accountReposityDb := domain.NewAccountRepositoryDB(dbClient)
	transactionRepositoryDb := domain.NewTransactionRepositoryDB(dbClient)
	// wiring
	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountReposityDb)}
	th := TransactionHandler{service.NewTransactionService(transactionRepositoryDb)}
	// define route
	// GET
	router.HandleFunc("/customers", ch.getAllCustomers)
		.Methods(http.MethodGet)
		.Name("GetAllCustomers")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer)
		.Methods(http.MethodGet)
		.Name("GetCustomer")

	// POST
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.newAccount)
		.Methods(http.MethodPost)
		.Name("NewAccount")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/transaction", th.newTransaction)
		.Methods(http.MethodPost)
		.Name("NewTransaction")

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())
	// starting server
	// Run the program with the following command:
	// SERVER_ADDRESS=localhost SERVER_PORT=8000 go run main.go
	// Getting environment variable
	serverAddr := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", serverAddr, serverPort), router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbConnectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dbConnectStr)
	if err != nil {
		panic(err)
	}
	// See important settings section
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
