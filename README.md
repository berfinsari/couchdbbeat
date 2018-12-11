# Couchdbbeat

Couchdbbeat is an [Elastic Beat](https://www.elastic.co/products/beats) that collects metrics from Couchdb and indexes them into Elasticsearch or Logstash.

## Description
> [CouchDB](http://couchdb.apache.org/) is an open source NoSQL database based on common standards to facilitate Web accessibility and compatibility with a variety of devices.

## Configuration

Adjust the `couchdbbeat.yml` configuration file to your needs.

### `period`
Defines how often to read statistics. Default to `30` s.

### `port`
Defines the couchdb port serviced. Default to `:5984`

### `host`
Host name of ElasticSearch. Default to `localhost`


## Document Example

<pre>

  "couchdb": {
    "server": {
      "httpd": {
        "clients_requesting_changes": {
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of clients for continuous _changes",
          "current": 0,
          "sum": 0,
          "mean": 0
        },
        "temporary_view_reads": {
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of temporary view reads",
          "current": 0,
          "sum": 0,
          "mean": 0
        },
        "requests": {
          "stddev": 0.214,
          "min": 0,
          "max": 12,
          "description": "number of HTTP requests",
          "current": 352,
          "sum": 352,
          "mean": 0.023
        },
        "viewReads": {
          "description": "number of view reads",
          "current": 0,
          "sum": 0,
          "mean": 0,
          "stddev": 0,
          "min": 0,
          "max": 0
        },
        "bulk_requests": {
          "sum": 0,
          "mean": 0,
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of bulk requests",
          "current": 0
        }
      },
      "httpd_request_methods": {
        "GET": {
          "min": 0,
          "max": 12,
          "description": "number of HTTP GET requests",
          "current": 352,
          "sum": 352,
          "mean": 0.023,
          "stddev": 0.214
        },
        "PUT": {
          "mean": 0,
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of HTTP PUT requests",
          "current": 0,
          "sum": 0
        },
        "COPY": {
          "description": "number of HTTP COPY requests",
          "current": 0,
          "sum": 0,
          "mean": 0,
          "stddev": 0,
          "min": 0,
          "max": 0
        },
        "HEAD": {
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of HTTP HEAD requests",
          "current": 0,
          "sum": 0,
          "mean": 0
        },
        "POST": {
          "description": "number of HTTP POST requests",
          "current": 0,
          "sum": 0,
          "mean": 0,
          "stddev": 0,
          "min": 0,
          "max": 0
        },
        "DELETE": {
          "current": 0,
          "sum": 0,
          "mean": 0,
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of HTTP DELETE requests"
        }
      },
      "httpd_status_codes": {
        "202": {
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of HTTP 202 Accepted responses",
          "current": 0,
          "sum": 0,
          "mean": 0
        },
        "301": {
          "description": "number of HTTP 301 Moved Permanently responses",
          "current": 1,
          "sum": 1,
          "mean": 0,
          "stddev": 0.008,
          "min": 0,
          "max": 1
        },
        "403": {
          "sum": 0,
          "mean": 0,
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of HTTP 403 Forbidden responses",
          "current": 0
        },
        "404": {
          "current": 0,
          "sum": 0,
          "mean": 0,
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of HTTP 404 Not Found responses"
        },
        "400": {
          "min": 0,
          "max": 1,
          "description": "number of HTTP 400 Bad Request responses",
          "current": 3,
          "sum": 3,
          "mean": 0,
          "stddev": 0.016
        },
        "401": {
          "sum": 0,
          "mean": 0,
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of HTTP 401 Unauthorized responses",
          "current": 0
        },
        "200": {
          "mean": 0.021,
          "stddev": 0.154,
          "min": 0,
          "max": 5,
          "description": "number of HTTP 200 OK responses",
          "current": 320,
          "sum": 320
        },
        "201": {
          "min": 0,
          "max": 0,
          "description": "number of HTTP 201 Created responses",
          "current": 0,
          "sum": 0,
          "mean": 0,
          "stddev": 0
        },
        "409": {
          "max": 0,
          "description": "number of HTTP 409 Conflict responses",
          "current": 0,
          "sum": 0,
          "mean": 0,
          "stddev": 0,
          "min": 0
        },
        "500": {
          "max": 0,
          "description": "number of HTTP 500 Internal Server Error responses",
          "current": 0,
          "sum": 0,
          "mean": 0,
          "stddev": 0,
          "min": 0
        },
        "405": {
          "min": 0,
          "max": 0,
          "description": "number of HTTP 405 Method Not Allowed responses",
          "current": 0,
          "sum": 0,
          "mean": 0,
          "stddev": 0
        },
        "412": {
          "min": 0,
          "max": 0,
          "description": "number of HTTP 412 Precondition Failed responses",
          "current": 0,
          "sum": 0,
          "mean": 0,
          "stddev": 0
        },
        "304": {
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of HTTP 304 Not Modified responses",
          "current": 0,
          "sum": 0,
          "mean": 0
        }
      },
      "couchdb": {
        "request_time": {
          "stddev": 1.574,
          "min": 1,
          "max": 22,
          "description": "length of a request inside CouchDB without MochiWeb",
          "current": 980.594,
          "sum": 980.594,
          "mean": 3.113
        },
        "database_reads": {
          "description": "number of times a document was read from a database",
          "current": 0,
          "sum": 0,
          "mean": 0,
          "stddev": 0,
          "min": 0,
          "max": 0
        },
        "auth_cache_hits": {
          "max": 0,
          "description": "number of authentication cache misses",
          "current": 0,
          "sum": 0,
          "mean": 0,
          "stddev": 0,
          "min": 0
        },
        "open_os_files": {
          "min": 0,
          "max": 0,
          "description": "number of file descriptors CouchDB has open",
          "current": 0,
          "sum": 0,
          "mean": 0,
          "stddev": 0
        },
        "database_writes": {
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of times a database was changed",
          "current": 0,
          "sum": 0,
          "mean": 0
        },
        "open_databases": {
          "sum": 0,
          "mean": 0,
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of open databases",
          "current": 0
        },
        "auth_cache_misses": {
          "stddev": 0,
          "min": 0,
          "max": 0,
          "description": "number of authentication cache misses",
          "current": 0,
          "sum": 0,
          "mean": 0
        }
      }
    }
  },
    "type":"couchdbbeat"
}
</pre>

## Getting Started with Couchdbbeat

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/berfinsari/couchdbbeat`

### Requirements

* [Golang](https://golang.org/dl/) 1.11

### Init Project
To get running with Couchdbbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Couchdbbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/berfinsari/couchdbbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Couchdbbeat run the command below. This will generate a binary
in the same directory with the name couchdbbeat.

```
make
```


### Run

To run Couchdbbeat with debugging output enabled, run:

```
./couchdbbeat -c couchdbbeat.yml -e -d "*"
```


### Test

To test Couchdbbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  Couchdbbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Couchdbbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/berfinsari/couchdbbeat
git clone https://github.com/berfinsari/couchdbbeat ${GOPATH}/src/github.com/berfinsari/couchdbbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.

## License
Covered under the Apache License, Version 2.0
Copyright (c) 2018 Berfin Sarı
