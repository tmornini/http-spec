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

### Docker image

    docker pull tmornini/http-spec

Please see https://hub.docker.com/r/tmornini/http-spec/ for instructions on how
to use the Docker image.

## Basic Usage

    http-spec [-prefix http(s)://host.name:port] path/to/\*.htsf

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

http-spec will send the first request and compare the expected response to the
actual response.

If the responses match, the next request/response pair is executed until the
end of the file.

If a response doesn't match, a error is logged to STDERR. Any response mismatch
or any error of any sort will cause http-spec to terinate with exit code 1 to
signal failure.

If a response matches, a success is logged to STDERR along with the response
time.

If all responses in all files match, http-spec terminates with exit code 0
to signal success.

When passed more than one file http-spec processes each file concurrently,
allowing for large spec suites to be completed quickly.

## Request-Only Mode

If no expected response follows a request, http-spec will make the request and
output the response formatted as an http-spec expected response ready to be
copied into an .htsf file. This allows for rapid, iterative request/response
development.

Request-only mode reports as a failure to prevent false-positives on incomplete
specs.

## Regexp Matching and Capture

Expected responses are parsed for regexp matchers that allow dynamic matching
and named capture for subsequent substitution within the file.

Matchers take the form:

    ⧆optional-name⧆mandatory-regexp⧆

If name is provided, the complete match of the regexp is assigned to the name,
making the matched text available for substitutions later in the file.

Matching makes it easy to match variable content items such as UUIDs and
authentication tokens.

* character is SQUARED ASTERISK (U+29C6)

## Built-in Date Matcher

⧆optional-name⧆:date⧆ is a special-case matcher that matches the RFC-822 date
format used by the HTTP 1.1 Date header.

## Substitution

Substitutions allows the re-use of previous regexp matches and take the form:

    ⧈name⧈

Substitutions are applied to both requests and responses within the same file.

* character is SQUARED SQUARE (U+29C8)

## TODO

* integrate Travis CI
* integrate Dockerhub
* improve testing dramatically :-(
