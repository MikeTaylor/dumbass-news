# Developer documentation: the `dumbass-news` source code


<!-- md2toc -l 2 overview.md -->
* [Overview](#overview)
* [The `catlogger` package](#the-catlogger-package)
* [The `main` package](#the-main-package)
    * [Compilation](#compilation)
    * [Main routine](#main-routine)
    * [Configuration](#configuration)
    * [Server](#server)
    * [Entry parsing](#entry-parsing)
    * [Transformation](#transformation)
    * [Tests](#tests)
* [Future directions](#future-directions)


## Overview

The `dumbass-news` source code is organized into two packages: `main`, which contains almost all the code, and `catlogger` (in [the subdirectory of that name](catlogger)), which contains only the very simple categorical-logger library. Compilation is controlled by a very simple [`Makefile`](Makefile), which supports the targets `run`, `lint`, `test` and `clean` as well as the basic build.



## The `catlogger` package

This package provides a single object, `Logger`, via the exported `MakeLogger` function. The object's `log` method can then be used to log information in arbitrary categories. The library is intended to be usable in other projects, at least in principal, and is documented separately using [godoc comments](https://go.dev/blog/godoc).



## The `main` package


### Main routine

`main.go` provides a simple main function that parses command-line arguments, loads the named configuration file, creates a News Server (see below), and launches it.


### Configuration

`config.go` defines the configuration structure and the JSON file format that is used to load it, and provides the `readConfig` function used by the main routine. The configuration file format is described in [the top-level README](../README.md).


### Server

XXX server.go
XXX http-error.go


### Entry parsing

XXX entry.go
XXX entry-rss.go


### Transformation

XXX transformer.go
XXX transformer-disemvowel.go
XXX transformer-insert.go


### Tests

XXX dumbass-news_test.go



## Future directions

XXX


