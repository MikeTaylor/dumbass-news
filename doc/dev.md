# Developer documentation: the `dumbass-news` source code


<!-- md2toc -l 2 dev.md -->
* [Overview](#overview)
* [The `catlogger` package](#the-catlogger-package)
* [The `main` package](#the-main-package)
    * [Main routine](#main-routine)
    * [Configuration](#configuration)
    * [Server](#server)
    * [Entry parsing](#entry-parsing)
    * [Transformation](#transformation)
    * [Tests](#tests)
* [Future directions](#future-directions)


## Overview

[The `dumbass-news` source code](../src) is organized into two packages: `main`, which contains almost all the code, and `catlogger` (in [the subdirectory of that name](../src/catlogger)), which contains only the very simple categorical-logger library. Compilation is controlled by a very simple [`Makefile`](../src/Makefile), which supports the targets `run`, `lint`, `test` and `clean` as well as the basic build.



## The `catlogger` package

This package provides a single object, `Logger`, via the exported `MakeLogger` factory function. The object's `log` method can then be used to log information in arbitrary categories. The library is intended to be usable in other projects, at least in principal, and is documented separately using [godoc comments](https://go.dev/blog/godoc).



## The `main` package


### Main routine

`main.go` provides a simple main function that parses command-line arguments, loads the named configuration file, creates a News Server (see below), and launches it.


### Configuration

`config.go` defines the configuration structure and the JSON file format that is used to load it, and provides the `readConfig` function used by the main routine. The configuration file format [is described separately](config.md).


### Server

`server.go` provides the `NewsServer` object, created by the `MakeNewsServer` factory function, which does all the work. The server routes incoming HTTP requests to render a hardwired HTML home-page, show the contents of a channel, or return a static file. When showing the contents of a channel, the channel name and transformation are extracted from the HTTP request path, and resolved respectively to an entry-parser object and a transformer object. Assuming both are recognised, the channel's source is read (usually via HTTP), entries extracted and transformed, and an HTML response written.

`http-error.go` provides a very simple data structure, created by the factory function `MakeHttpError`, which aggregates a Go Error object with an HTTP status code. This is used only in `server.go`.


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

Aside from improvements to functionality, the following work could usefully be done just to improve the clarity of the code:

* `server.go` feels a bit too broad in its responsibilities, and might usefully be broken up. For example, HTML generation might be better handled by its own source file.
* The `httpError` structure is not necessary. In `server.go,` the `getData` function should simply return an (entries, httpStatus, err) triple instead of an (entries, httpError) pair.
* See other plans in [the GitHub issue tracker](https://github.com/MikeTaylor/dumbass-news/issues).


