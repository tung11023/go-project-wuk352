package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "utils"
)

func main() {
    configData, _ := ioutil.ReadFile("config/config.yaml")
    var config struct{ DataPath string }
    yaml.Unmarshal(configData, &config)
    data := utils.ProcessData(config.DataPath)
    fmt.Printf("Processed %d records
", len(data))
}
