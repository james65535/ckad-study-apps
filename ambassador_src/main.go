package main

import (
	"errors"
	"github.com/mediocregopher/radix.v2/pool"
	"log"
	"net/http"
	"os"
)

var db *pool.Pool

func init() {
	var redisAddress string
	if os.Getenv("SERVER") != "" {
		redisAddress = os.Getenv("SERVER")
	}
	var err error
	log.Println(redisAddress)
	db, err = pool.New("tcp", redisAddress, 10)
	if err != nil {
		log.Panic(err)
	}
}

var ErrNoUsers = errors.New("Users: no users found")

//localhost:8080/setuser?id=1&greeting=hi&name=bob
func setUser(w http.ResponseWriter, r *http.Request) {
	log.Println("setUser visitor:" + r.RemoteAddr)
	id := "user:" + r.URL.Query().Get("id")
	greeting := r.URL.Query().Get("greeting")
	name := r.URL.Query().Get("name")
	log.Println("Added:", greeting, name)

	// store record to redis
	result := db.Cmd("HSET", id, "greeting", greeting, "name", name)
	if result.Err != nil{
		log.Fatal(result.Err)
	}
	w.Write([]byte("User received"))
}

//localhost:8080/getuser?id=1
func getUser(w http.ResponseWriter, r *http.Request) {
	log.Println("setUser visitor:" + r.RemoteAddr)
	id := r.URL.Query().Get("id")
	reply, err := db.Cmd("HGETALL", "user:"+id).Map()
	if err != nil {
		log.Fatal(err)
	} else if len(reply) == 0 {
		log.Println(ErrNoUsers)
		w.Write([]byte("No user found by that ID"))
		return
	}

	message := reply["greeting"] + " " + reply["name"]
	w.Write([]byte(message))
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("/ visitor:" + r.RemoteAddr)
	w.Write([]byte("redisWebApp Home"))
}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/setuser", setUser)
	http.HandleFunc("/getuser", getUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}