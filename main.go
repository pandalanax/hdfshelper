package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/ktr0731/go-fuzzyfinder"
)

//type Configuration struct {
//	XMLName    xml.Name   `xml:"configuration"`
//	Properties []Property `xml:"configuration>property"`
//}
//
//type Property struct {
//	XMLname     xml.Name `xml:"property"`
//	Name        string   `xml:"name"`
//	Value       string   `xml:"value"`
//	Description string   `xml:"description"`
//}

// Configuration was generated 2024-02-25 13:50:16 by https://xml-to-go.github.io/ in Ukraine.
type Configuration struct {
	XMLName  xml.Name `xml:"configuration"`
	Property []struct {
		XMLName     xml.Name `xml:"property"`
		Name        string   `xml:"name"`
		Value       string   `xml:"value"`
		Description string   `xml:"description"`
	} `xml:"property"`
}

var modes = map[string]string{
	"core": "https://hadoop.apache.org/docs/current/hadoop-project-dist/hadoop-common/core-default.xml",
	"hdfs": "https://hadoop.apache.org/docs/current/hadoop-project-dist/hadoop-hdfs/hdfs-default.xml",
	"yarn": "https://hadoop.apache.org/docs/current/hadoop-yarn/hadoop-yarn-common/yarn-default.xml",
}
var mode string

func main() {
	const (
		defaultMode = "hdfs"
		usageMode   = `mode. decide which part of hdfs you want to configure and fuzzy find. Supported modes are: 
        - core [core-site.xml] 
        - yarn [yarn-site.xml]
        - hdfs [hdfs-site.xml]`
	)
	flag.StringVar(&mode, "mode", defaultMode, usageMode)
	flag.StringVar(&mode, "m", defaultMode, usageMode+"(shorthand)")

	flag.Parse()

	//	fmt.Println(*mode)
	url, ok := modes[mode]

	if !ok {
		log.Fatal("wrong mode. type --help for usage")
	}

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

	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	idx, err := fuzzyfinder.FindMulti(
		config.Property,
		func(i int) string {
			return config.Property[i].Name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Name: %s \nDefault: %s \n\nDescription: %s",
				red(config.Property[i].Name),
				yellow(config.Property[i].Value),
				yellow(config.Property[i].Description))
		}))
	if err != nil {
		log.Fatal(err)
	}

	for _, element := range idx {
		// print out the xml snippet for copy
		out, err := xml.MarshalIndent(config.Property[element], " ", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	}
}
