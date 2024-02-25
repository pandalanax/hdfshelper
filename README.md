# hdfshelper
[![Go](https://github.com/pandalanax/hdfshelper/actions/workflows/go.yml/badge.svg)](https://github.com/pandalanax/hdfshelper/actions/workflows/go.yml)

Do you find yourself constantly searching for config options in [Apache Hadoop](https://hadoop.apache.org/)? Then this is for you.
Fuzzy find your way to the configuration option you're searching for. By default it pulls the options for the latest version aka `current`.

It prints out the xml ready to be pasted into the configuration file.



## Usage
search all options from `hdfs-site.xml` on version `r3.3.4`:
```bash
hdfshelper -m hdfs -v r3.3.4
```

copy the output to clipboard (macOS):
```bash
hdfshelper -m core | pbcopy
```

You can select more than one option to print with `<TAB>`.
![screenshot of fuzzyfind](images/hdfshelper.png?raw=true "Screenshot with added option with <TAB>")

```bash
Usage of hdfshelper:
  -m string
    	mode. decide which part of hdfs you want to configure and fuzzy find. Supported modes are:
    	        - core [core-site.xml]
    	        - yarn [yarn-site.xml]
    	        - hdfs [hdfs-site.xml](shorthand) (default "hdfs")
  -mode string
    	mode. decide which part of hdfs you want to configure and fuzzy find. Supported modes are:
    	        - core [core-site.xml]
    	        - yarn [yarn-site.xml]
    	        - hdfs [hdfs-site.xml] (default "hdfs")
  -v string
    	The version of hadoop in format: rx.y.z (e.g. r3.3.6)(shorthand) (default "current")
  -version string
    	The version of hadoop in format: rx.y.z (e.g. r3.3.6) (default "current")
```

## Building
```bash
go build 
```
or 
```bash
nix build
```

## Disclaimer
I use this little program for learning go. Bugs, crashes and everything that sucks are to be expected.

