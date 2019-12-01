# httpst

[![Build Status](https://travis-ci.org/piotrpersona/httpst.svg?branch=master)](https://travis-ci.org/piotrpersona/httpst)

Display HTTP Status Code description

## Install

```console
curl -fsSL -o /usr/local/bin/httpst \
    https://github.com/piotrpersona/httpst/releases/download/latest/httpst-darwin-amd64 \
&& chmod +x /usr/local/bin/httpst
```

## Usage

Display single code info

```console
$ httpst 200
200 OK
```

or multiple codes

```console
$ httpst 200 201
200 OK
201 Created
```

with long description

```console
$ httpst 500 -l
500 Internal Server Error

The server encountered an unexpected condition which prevented it
from fulfilling the request.
```

or with group info

```console
$ httpst 200 -g
200 OK

2xx Successful 2xx
This class of status code indicates that the client's request was
successfully received, understood, and accepted.⏎
```
