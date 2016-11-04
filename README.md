# http-spec
Template-based HTTP request/response specification tool

As an API developer who loves HTTP and test driven development, I wanted a very lightweight way to test APIs and could not find anything that felt "just right."

I wanted a direct, explicit, flexible and fast way to develop and execute HTTP level testing against those APIs.

If you have code that executes behind an HTTP API, then it would be ideal to
test that code via HTTP, right?

Stated another way: Good APIs are contracts executed via HTTP. Why not test
contract conformance via HTTP requests and responses? Anything else sort of
misses the point.

While low-level tests may still be required for edge cases such as networking
exceptions and intermittent service retry, for instance, the *vast* majority
of the code behind the API should be exercisable via HTTP requests.

A not insignificant advantage of this method is that it requires you stare at
your API's input and output, including request and response headers. There is
much to be learned from that experience!

## installation

### for Go developers

    go get github.com/tmornini/http-spec

### for Mac OS X developers

Download a binary executable from:

    https://github.com/tmornini/http-spec/releases

Don't forget to put it somewhere in your PATH, and chmod 755 it!

### For Windows and Linux developers

Let me know if you'd like a binary and I'll add your platform to future.

## basic usage

    http-spec [-prefix http(s)://host.name:port] path/to/\*.htsf

## hypertext specification format (HTSF)

HTSF is, first and foremost, modeled after curl -v output.

It consists of a *complete* HTTP request, including the request line, HTTP
headers, mandatory empty line, and optional body.

Each line of the request is prefixed by "> " -- though ">" is allowed for the
blank line.

An empty line separates the request from the expected response.

The expected response is a *complete* HTTP response, including the status line,
followed by HTTP headers sort *in ASCII order*, a mandatory empty line, and an
optional body.

Each line of the expected response is prepended by "< " though "<" is acceptable
for the blank line.

http-spec will perform each request, then compare the expected response.

If the response matches, the next request/response pair is executed until the
end of the file.

If a response doesn't match, http-spec exits with an explanation on STDERR and
exit code 1.

If all responses match, http-spec exists silently with exit code 0.

If passed more than one file, including the globbed example above, http-spec
will process each file concurrently, allowing for large spec runs to be
completed quickly.

## request-only non-pairs

If no expected response follows a request, http-spec will make the request and
output the response formatted as an http-spec expected response allowing for
rapid, iterative request/response development.

Request only non-pairs are reported as a failure to prevent false-positive test
results.

## regexp matching

Expected responses are parsed for regexp matchers that allow dynamic matching
and named capture for subsequent substition, scoped to the file, and take the
form:

    ⧆optional-name⧆mandatory-regexp⧆

That character is SQUARED ASTERISK (U+29C6)

If name is provided, the complete match of the regexp is assigned to the name,
making the matched text available for substitions later in the file.

Matching makes it easy to match variable content items such as UUIDs and
authentication tokens.

## built-in matchers

⧆optional-name⧆:date⧆ is a special-case matcher that matches the RFC-822 date
format used by the HTTP 1.1 Date header.

## substition

Substitions allow the re-use of previous regexp matches and take the form:

    ⧈name⧈

That character is SQUARED SQUARE (U+29C8)

Substition allows re-use of response data in requests within the same file.

## TODO

* improve testing dramatically :-(
