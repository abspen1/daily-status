package function

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	secret, _ := getDBSecret("redis-password")
	c, err := redis.Dial("tcp", "192.168.1.6:6379")
	if err != nil {
		log.Fatal(err)
	}

	response, err := c.Do("AUTH", secret)
	if err != nil {
		log.Fatal(err)
	}

	dailyStatus, err := redis.StringMap(client.Do("HGETALL", "daily_status"))
	if err != nil {
		log.Fatal(err)
	}

	str, _ := json.Marshal(dailyStatus)

	w.Write(string(str))
}

func getDBSecret(secretName string) (secretBytes []byte, err error) {
	// read from the openfaas secrets folder
	secretBytes, err = ioutil.ReadFile("/var/openfaas/secrets/" + secretName)
	if err != nil {
		// read from the original location for backwards compatibility with openfaas <= 0.8.2
		secretBytes, err = ioutil.ReadFile("/run/secrets/" + secretName)
	}

	return secretBytes, err
}
