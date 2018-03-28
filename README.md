[![Build Status](https://travis-ci.org/tmornini/http-spec.svg?branch=master)](https://travis-ci.org/tmornini/http-spec)

# http-spec
Template-based HTTP request/response specification tool

As an API developer who loves HTTP and test driven development, I wanted a very lightweight, language-agnostic way to test APIs.

http-spec intends to allow development of explicit and fast spec suites.

If you have code that executes behind an HTTP API, then it is ideal to test
that code via HTTP -- anything else is missing the point!

Stated another way: APIs are contracts executed via HTTP. http-spec allows
API contracts to be written in and verified by HTTP.

While low-level tests may still be required for edge cases such as networking
exceptions and intermittent service retry, for instance, the *vast* majority
of the code behind the API should be exercisable via HTTP requests.

A significant advantage of this method is that it requires you consider every
byte of your API's input and output, including request and response headers.
There is much to be learned from that experience!

## Installation

### Go developers

    go get github.com/tmornini/http-spec

### Mac OS X developers

Download a binary executable from:

    https://github.com/tmornini/http-spec/releases

Don't forget to put it somewhere in your PATH, and chmod 755 it!

### Options

```
-http-retry-delay duration
    delay between failed HTTP requests (default 250ms)
-max-http-attempts int
    maximum number of attempts per HTTP request (default 20)
-skip-tls-verification
    skip TLS verification (hostname mismatch, self-signed certifications, etc.)
```

### Docker image

    https://hub.docker.com/r/tmornini/http-spec/

This image is not intended to be used directly, but as the base to build your
own custom http-spec container atop.

Just COPY your .htsf files and /run-http-specs executable to the image.

/run-http-specs allows you to orchestrate the test invocation order, timing,
and prefix handling when testing microservices within a docker-compose cluster.

### Dockerfile example

```
FROM tmornini/http-spec
MAINTAINER Tom Mornini <tmornini@incentivenetworks.com>

COPY run-http-specs /run-http-specs

COPY *.htsf /
```

## Basic Usage

    http-spec path/to/\*.htsf

## HTSF (Hypertext Specification Format)

HTSF is, first and foremost, modeled after curl -v output.

It consists of a HTTP request, including the request line, HTTP headers, mandatory empty line, and optional body. The Content-Length header is set
automatically.

Each line of the request is prefixed by "> " -- though ">" is allowed for the
blank line.

An empty line separates the request from the expected response.

The expected response is an HTTP response, including the status line,
followed by HTTP headers *sorted in ASCII order*, a mandatory empty line, and
an optional body.

Each line of the expected response is prepended by "< " though "<" is acceptable
for the blank line.

http-spec will send the request and compare the expected response to the actual
response.

If the responses match, the next request/response pair is executed until the
end of the file.

If a response doesn't match, a error is logged to STDERR. Any response mismatch
or any error of any sort will cause http-spec to terminate with exit code 1 to
signal failure.

If a response matches, a success is logged to STDOUT along with the response
time.

If all responses in all files match, http-spec terminates with exit code 0
to signal success.

When passed more than one file http-spec processes each file concurrently,
allowing for large spec suites to be run quickly.

## Request-Only Mode

If no expected response follows a request, http-spec will make the request and
output the response formatted as an http-spec expected response ready to be
copied into an .htsf file. This allows for rapid, iterative request/response
development.

Request-only mode reports as a failure to prevent false-positives on incomplete
specs.

## Regexp Matchers and Capture

Expected responses are parsed for regexp matchers that allow dynamic matching
and named capture for subsequent substitution within the file.

Matchers take the form:

    ⧆optional-name⧆mandatory-regexp⧆

If name is provided, the complete match of the regexp is assigned to the name,
making the matched text available for substitutions later in the file.

Matching makes it easy to match variable content items such as UUIDs and
authentication tokens.

* character is SQUARED ASTERISK (U+29C6)

## Built-in Matchers

⧆optional-name⧆:date⧆ is a matcher for RFC-822 dates used by the HTTP 1.1 Date
header.

⧆optional-name⧆:uuid⧆ is a matcher for [RFC-4122](https://tools.ietf.org/html/rfc4122) UUIDs.

⧆optional-name⧆:b62:22⧆ is a matcher for 22 base 62 characters sometimes used
for 128+ bit UUIDs.

⧆optional-name⧆:iso8601:µs:z⧆ is a matcher for ISO 8601 format timestamps
with microsecond resolution and zulu (Z) timezone.

## Built-in Substitutes

    ⧈YYYY-MM-DD⧈ is a substitute for today's date

## Delayed Requests

If you need to delay between one request (and it's associated response) and the next,
you can use a separating line that begins with "+ " followed by a sequence of decimal
numbers, each with optional fraction and a unit suffix, such as "300ms" or "1.5s".
Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"

This is useful when making a request that triggers an asyncronous operation where you need
to wait enough time to allow the asyncronous operation to complete before sending the next
request.

## Substitution

Substitution allows the re-use of previous regexp matches and take the form:

    ⧈name⧈

Substitutions are applied to requests and responses within the same file.

* character is SQUARED SQUARE (U+29C8)

## TODO

* integrate Travis CI
* improve testing dramatically :-(
