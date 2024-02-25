package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Configuration struct {
	XMLName    xml.Name   `xml:"configuration"`
	Properties []Property `xml:"property"`
}

type Property struct {
	XMLname     xml.Name `xml:"property"`
	Name        string   `xml:"name"`
	Value       string   `xml:"value"`
	Description string   `xml:"description"`
}

func (p Property) String() string {
	return fmt.Sprintf("name=%v value=%v description=%v", p.Name, p.Value, p.Description)
}

func main() {
	fmt.Println("Hello flake")

	url := "https://hadoop.apache.org/docs/r2.4.1/hadoop-project-dist/hadoop-hdfs/hdfs-default.xml"

	resp, err := http.Get(url)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our xmlFile so that we can parse it later on
	// We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var config Configuration

	xml.Unmarshal(body, &config)

	for i := 0; i < len(config.Properties); i++ {
		fmt.Println("Name: " + config.Properties[i].Name)
		fmt.Println("Value: " + config.Properties[i].Value)
		fmt.Println("Desc: " + config.Properties[i].Description)
	}
}
