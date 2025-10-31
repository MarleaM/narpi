package config

//I THINK THIS WORKS NOW
import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	DB  DB
	API API
}

type DB struct {
	URI        string
	Database   string
}

type API struct {
	Port string
}

func LoadConfig() App {
	//grab all of the env variables we need
	port, ret := os.LookupEnv("PORT")
	if !ret {
		go gracefulShutdown("PORT not set")
	}

	mongoURI, ret := os.LookupEnv("MONGO_URI")
	if !ret {
		go gracefulShutdown("MONGO_URI not set")
	}

	mongoDBName, ret := os.LookupEnv("MONGO_DBNAME")
	if !ret {
		go gracefulShutdown("MONGO_DBNAME not set")
	}
	return App{
		DB: DB{
			URI:      mongoURI,
			Database: mongoDBName,
		},
		API: API{
			Port: port,
		},
	}
}

func gracefulShutdown(reason string) {
	fmt.Println("Shutting Down: ", reason)
	_, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()
}