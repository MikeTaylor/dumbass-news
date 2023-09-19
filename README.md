# Dumbass News - a web service to report dumbass news

Copyright (C) 2022 Mike Taylor <mike@miketaylor.org.uk>

Licensed under [the GNU General Public License (GPL) v3.0](LICENSE).


## Overview

This program runs a web server which harvests news headlines from various sources, and presents them to the user transformed in various ways. In its canonical manifestation, it presents [BBC news headlines](https://www.bbc.co.uk/news) with the adjective "dumbass" inserted in front of some nouns, as in
* PM warned about dumbass lockdown drinks, claims Cummings
* Ministers suffer defeats in Lords over dumbass crime bill
* French far-right candidate guilty of dumbass hate speech
* Woman completes dumbass bid to run length of New Zealand.

This is in part a response to how incredibly stupid nearly every news story seems to be at the moment (17 January 2022).

Other news sources that provide RSS feeds are supported, if added to the configuration file. In time, I may add support for other kinds of source. I may also add support for other transformations.


## Compilation

Dumbass News is written in [the Go programming language](https://go.dev/). Assuming you gave Go installed, you just need to run `make` in the `src` directory. A binary, `dumbass-news` will be created in that directory. Running it will start a web-server running on the host and port specified in the configuration file (see below).


## Invocation

```src/dumbass-news etc/config.json```

The only command-line argument is the name of [a JSON configuration file](etc/config.json) which specifies details such as what port to listen to, what categories of information to log, which news channels are supported and which transformations can be carried out on them (see below).

Apart from the home page, the server provides pages at URLs of the form `http://HOST:PORT/CHANNEL/TRANSFORMATION`, where `CHANNEL` is a channel key from the configuration file (see below) and `TRANSFORMATION` is a transformation key. For example, http://localhost:12368/bbc/dumbass serves the BBC news channel with the Dumbass transformation.


## Further documentation

* [_Configuration file format for `dumbass-news`_](doc/config.md).
* [_Developer documentation: the `dumbass-news` source code_](doc/dev.md).


## Warning! Lark's Vomit!

This exists primarily because I wanted to write something in Go as part of learning the language. It is probably bad code, and should not be studied or emulated. If anyone apart from me is childish enough to find it amusing, that's just a bonus.


