# Configuration file format for `dumbass-news`


<!-- md2toc -l 2 config.md -->
* [Overview](#overview)
* [`logging`](#logging)
* [`listen`](#listen)
* [`channels`](#channels)
* [`transformations`](#transformations)


## Overview

A single configuration file provides all run-time information to `dumbass-news`. The example configuration file, [`etc/config.json`](../etc/config.json) illustrates all parts of the configuration.

The configuration file must be well-formed JSON, containing the following sections:


## `logging`

This section is used to configure the [categorical logger](../src/catlogger/catlogger.go) that the server uses. It provides a means to specify what information is logged. This section contains the following elements:

* `categories` (e.g. `"config,listen,path,parse"`): a comma-separated list of short strings, each of which names a category of log messages to be included in the logging output. See below for a list of categories used by `dumbass-news`.
* `prefix` (e.g. `"--"`): a string which, if specified, is emitted at the start of each logging line. This can be useful to determine which lines of output are from the logger and which from other sources.
* `timestamp`: either `true` or `false`, specifying whether or not a timestamp should be included on each log line.

The following logging categories are used by `dumbass-news`:

* `config` -- emits the compiled configuration when starting up, and the per-channel configuration when responding to each request.
* `listen` -- notes when starting to listen, stating what host and port are served, and when stopping.
* `path` -- logs each path that is requested
* `body` -- logs each HTTP body retrieved from an RSS server or other remote source. (Warning: very verbose)


## `listen`

This section specifies what host to serve and what port to listen on. It contains the following elements:

* `host` (e.g. `"localhost"`) -- the name of a host to serve, or and empty string to accept connections from anywhere.
* `port` (e.g. `12368`) -- an integer that is the port number to listen on.

When using these configuration values, the server can be accessed on http://localhost:12368/


## `channels`

A set of channels that the server is configured to provide. Each channel is a source of news, which can then be modified by a specified transformation (see below). Each channel is specified by a key (see "channel key" above) and represented by a JSON structure with the following elements:

* `type` (e.g. `"rss"`) -- the type of the channel, i.e. the fundamental choice of how data is obtained from it. At present the only supported type is "rss", which fetches and analyses [RSS feeds](https://en.wikipedia.org/wiki/RSS). Others may be supported in future.
* `url` (e.g. `"https://feeds.bbci.co.uk/news/rss.xml"`) -- the fetchable location of the news source (typically an RSS feed). These are often `http` URLs, but `file` URLs are also supported to load from local sources.
render
* `render` (e.g. `"description"`) -- **valid for RSS feeds only** (but since all feeds are currently RSS feeds that's not a problem). Specifies whether to use the `title` field from the feed, or the `description`, or (the default case) both of them together, separated by a dash.


## `transformations`

A set of transformations that the server is configured to provide. Each transformation is a way of modifying a news headline obtained from a channel (see above). Each transformation is specified by a key (see "transformation key" above) and represented by a JSON structure with the following elements:

* `type` (e.g. `"insert"`) -- the type of the transformation, i.e. the fundamental choice of what transformation is run. At present, two transformation types are supported: `disemvowel`, which simply removes all vowels from the headline; and `insert` which inserts words into it.
* `params` -- a JSON structure containing parameters that are specific to the transformation type.

For the `disemvowel` transformation type, no parameters are required.

For the `insert` transformation type, the following parameters are used:

* `text` (e.g. `"dumbass"`) -- the text to be inserted before a word in the headine.
* `anchorDataPath` (e.g. `"etc/words/nouns"`) -- the path to a file containing a list of words, one per line, which are candidates to have the text inserted ahead of them.

Note that `file` URLs used for RSS feeds and `anchorDataPath` paths are interpreted relative to the working directory.


