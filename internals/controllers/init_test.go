package controllers

import (
	pointer_db "bricks-core/internals/pkg/db"
	"context"
	"fmt"
	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	r "github.com/moemoe89/integration-test-golang/repository"
	"github.com/moemoe89/integration-test-golang/repository/postgres"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"os"
	"testing"
)

func init() {
	// initialize the router
	router := gin.Default()
	// Handlers for testing
	router.GET("/applications", ReadApplications)
	router.POST("/applications", CreateApplication)
	// Setup the router
	unitTest.SetRouter(router)
	newLog := log.New(os.Stdout, "", log.Llongfile|log.Ldate|log.Ltime)
	unitTest.SetLog(newLog)
}

var (
	repo r.Repository
)

var (
	user     = "postgres"
	password = "secret"
	db       = "postgres"
	port     = "5433"
	dialect  = "postgres"
	dsn      = "postgres://%s:%s@localhost:%s/%s?sslmode=disable"
	idleConn = 25
	maxConn  = 25
)

// TestMain will automatically run before all test and hook back in // after all test is done
func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12.3",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + db,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err.Error())
	}

	dsn = fmt.Sprintf(dsn, user, password, port, db)
	if err = pool.Retry(func() error {
		repo, err = postgres.NewRepository(dialect, dsn, idleConn, maxConn)
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err.Error())
	}

	defer func() {
		repo.Close()
	}()

	err = repo.Drop()
	if err != nil {
		panic(err)
	}

	err = repo.Up()
	if err != nil {
		panic(err)
	}

	// create a local connection
	pointer_db.NewConnection()

	createSchema(pointer_db.Conn)
	// THIS IS KEY, we replace the DB pointer to the mock one we
	// recently build
	log.Println("Everything above here run before ALL test")
	// Run test suites
	exitVal := m.Run()
	log.Println("Everything below run after ALL test")
	// we can do clean up code here

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	pointer_db.CloseConnection()

	os.Exit(exitVal)
}

// Schema in the mock DB still needs to be created
func createSchema(db *pgxpool.Pool) error {
	//tx, _ := db.Begin(context.Background())
	_, err := db.Exec(context.Background(), `CREATE TABLE applications (
			"id" SERIAL PRIMARY KEY NOT NULL,
			"name" VARCHAR(45) NOT NULL
);`)
	//err = tx.Commit(context.Background())

	if err != nil {
		return err
	}

	return nil
}
