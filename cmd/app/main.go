
package main

import (
    "log"
    "github.com/ilyakaznacheev/cleanenv"

    "github.com/danilasar/golang-crud/app"
    "github.com/danilasar/golang-crud/config"
)

func main() {
    var conf config.Config 
    err := cleanenv.ReadConfig("./config/server.yml", &conf)
    if err != nil {
        log.Fatal("Config not found at ./config/server.yml", err)
    }
    app.Run(&conf)
}
