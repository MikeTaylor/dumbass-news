# Dumbass News - a web service to report dumbass news

Copyright (C) 2022 Mike Taylor <mike@miketaylor.org.uk>

Licensed under [the GNU General Public License (GPL) v3.0](https://www.gnu.org/licenses/gpl-3.0.html).


## Overview

This program runs a web server which harvests news headlines from various sources, and presents them to the user transformed in various ways. In its initial form, it will present [BBC news headlines](https://www.bbc.co.uk/news) with the adjective "dumbass" inserted in front of some nouns, as in
* PM warned about dumbass lockdown drinks, claims Cummings
* Ministers suffer defeats in Lords over dumbass crime bill
* French far-right candidate guilty of dumbass hate speech
* Woman completes dumbass bis to run length of New Zealand.

This is in part a response to how incredibly stupid nearly every news story seems to be at the moment (17 January 2022).

With time, I may add support for other news sources (probably only those that provide RSS feeds) and other transformations.


## Compilation

Dumbass News is written in [the Go programming language](https://go.dev/). Assuming you gave Go installed, you just need to run `make` in the `src` directory. A binary, `dumbass-news` will be created in that directory.


## Invocation

```src/dumbass-news etc/config.json```

The only command-line argument is the name of a JSON configuration file which specifies details such as what port to listen to, what categories of information to log, which news channels are supported and which transformations can be carried out on them.

See the sample configuration file [`etc/config.json`](etc/config.json).


## Warning! Lark's Vomit!

This exists primarily because I wanted to write something in Go as part of learning the language. It is probably bad code, and should not be studied or emulated. If anyone apart from me is childish enough to find it amusing, that's just a bonus.


