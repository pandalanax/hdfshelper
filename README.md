# hdfshelper
[![Go](https://github.com/pandalanax/hdfshelper/actions/workflows/go.yml/badge.svg)](https://github.com/pandalanax/hdfshelper/actions/workflows/go.yml)

Do you find yourself constantly searching for config options in [Apache Hadoop](https://hadoop.apache.org/), this is for you.
Fuzzy find your way to the configuration option you're searching for.




## Usage
search all options from `hdfs-site.xml`:
```bash
hdfshelper -m hdfs
```

copy the output to clipboard (macOS):
```bash
hdfshelper -m core | pbcopy
```

You can select more than one option to print with `<TAB>`.

## Supports
Currently the options below are supported. 

 - yarn-site.xml
 - core-site.xml
 - hdfs-site.xml

## Disclaimer
I use this little program for learning go. Bugs, crashes and everything thats sucks are to be expected.
