package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/ktr0731/go-fuzzyfinder"
)

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

var (
	mode     string
	version  string
	fzf_desc bool
	red      = color.New(color.FgRed).SprintFunc()
	yellow   = color.New(color.FgYellow).SprintFunc()
)

func getUrls(version string) map[string]string {
	baseUrl := "https://hadoop.apache.org/docs/" // note the r

	modes := map[string]string{
		"core": baseUrl + version + "/hadoop-project-dist/hadoop-common/core-default.xml",
		"hdfs": baseUrl + version + "/hadoop-project-dist/hadoop-hdfs/hdfs-default.xml",
		"yarn": baseUrl + version + "/hadoop-yarn/hadoop-yarn-common/yarn-default.xml",
	}

	return modes
}

func main() {
	const (
		defaultMode = "hdfs"
		usageMode   = `mode. decide which part of hdfs you want to configure and fuzzy find. Supported modes are: 
        - core [core-site.xml] 
        - yarn [yarn-site.xml]
        - hdfs [hdfs-site.xml]`
		defaultVersion = "current"
		usageVersion   = "The version of hadoop in format: rx.y.z (e.g. r3.3.6)"
		defaultSwitch  = false
		usageSwitch    = "Wether to search in configuration names or descriptions"
	)

	flag.StringVar(&mode, "mode", defaultMode, usageMode)
	flag.StringVar(&mode, "m", defaultMode, usageMode+"(shorthand)")
	flag.StringVar(&version, "version", defaultVersion, usageVersion)
	flag.StringVar(&version, "v", defaultVersion, usageVersion+"(shorthand)")
	flag.BoolVar(&fzf_desc, "switch", defaultSwitch, usageSwitch)
	flag.BoolVar(&fzf_desc, "s", defaultSwitch, usageSwitch+"(shorthand)")

	flag.Parse()

	//	fmt.Println(*mode)
	modes := getUrls(version)
	url, ok := modes[mode]
	if !ok {
		log.Fatal("wrong mode. type --help for usage")
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var config Configuration

	xml.Unmarshal(body, &config)

	idx, err := fuzzyfinder.FindMulti(
		config.Property,
		func(i int) string {
			if fzf_desc {
				return config.Property[i].Description
			}
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
		output := string(out)
		fmt.Println(red(strings.Replace(output, "&#xA;", "\n", -1)))
	}
}
