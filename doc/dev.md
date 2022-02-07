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

`entry.go` defines the `Entry` structure, consisting of a headline and link (such as might be extracted from an RSS feed), and the EntryParser strcture -- a structure of a single method, `parse`. The idea is that many kinds of entry parser may be supported, each with its own parser implemented in a source file called `entry-*.go` which provides a single entry-parser object.

`entry-rss.go` is, at present, the only entry parser. As the name suggests, it parses an RSS document and returns a slice of entries.


### Transformation

`transformer.go` defines the `Transformer` structure -- a structure of a single method. `transform`. The idea is that many kinds of transformer may be supported, each implemented in a source file called `transformer-*.go` which provides a single transformer object. This source file also defined the `transformData` function which chooses which of the supported transformers to use and applies it to each of the entries in the supplied list.

`transformer-disemvowel.go` is the first and simplest transformer, written as a proof of concept, which simply removes all vowels from the headlines it is given.

`transformer-insert.go` is a transformer which inserts a specfied word (such as "dumbass" )into headlines before a word matching one in a specified list (such as a list of nouns that are not also verbs).


### Tests

`dumbass-news_test.go` is a very high-level test which launches a server and checks that it returns appropriate responses for a variety of HTTP requests, both invalid and valid.



## Future directions

Aside from improvements to functionality, the following work could usefully be done just to improve the clarity of the code:

* `server.go` feels a bit too broad in its responsibilities, and might usefully be broken up. For example, HTML generation might be better handled by its own source file.
* `entry.go` should provide a `parseData` function, analogous to `transformer.go`'s `transformData` function, which selects which of the candidate entry-parsers to use. This will move some more non-server-related code out of `server.go`.
* The `httpError` structure is not necessary. In `server.go,` the `getData` function should simply return an (entries, httpStatus, err) triple instead of an (entries, httpError) pair.
* See other plans in [the GitHub issue tracker](https://github.com/MikeTaylor/dumbass-news/issues).



