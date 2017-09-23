# COUP Challenge: Find number of Fleet Engineers

## Usage

This tool can be used in two ways.

### 1. Handle single request from STDIN:

Example 1:
```bash
$ echo '{"scooters": [15, 10], "C": 12, "P": 5}' | ./coup
{"fleet_engineers":3}
```

Example 2:
```bash
$ cat request.txt
{"scooters": [11, 15, 13], "C": 10, "P": 5}

$ ./coup < request.txt
{"fleet_engineers":7}
```

### 2. Listen for HTTP POST requests on a given address:

```bash
$ ./coup --listen :8080
Listening on :8080...
```

In a separate shell:
```bash
$ curl -XPOST -d '{"scooters":[15, 10], "C":12, "P":5}' localhost:8080
{"fleet_engineers":3}
```

## Hosted Version

A hosted version is available under `https://coup.cfapps.io`.

Example:
```bash
$ curl -XPOST -d '{"scooters":[11,15,13], "C": 10, "P": 5}' https://coup.cfapps.io
{"fleet_engineers":7}
```

## Download Binary

See [releases](https://github.com/st3v/coup/releases).

## Build from Source

Make sure you have Go installed. Refer to the [official docs](https://golang.org/doc/install) 
for more details on how to install Go.

```bash

$ go get github.com/st3v/coup

$ coup --help

```

## Run Tests

```bash
$ cd $GOPATH/src/github.com/st3v/coup

$ go test -v ./...
```

## Original Problem Statement

You are given a `[]int scooters`, which has as many elements as there are 
districts in Berlin that Coup operates in. For each `i`, `scooters[i]` is the 
number of scooters in that district (0-based index).

During a work day, scooters are maintained (batteries changed, cleaned, 
checked for damages) by the Fleet Manager (FM) and possibly other Fleet 
Engineers (FEs). Each FE, as well as the FM, can only maintain scooters in 
one district. Additionally, there is a limit on how many scooters a single 
FE may supervise: the FM is able to maintain up to `C` scooters, and a FE is 
able to maintain up to `P` scooters. Each scooter has to be maintained by some FE or the FM.

## How do we solve the problem?
Find the minimum number of FEs which are required to help the FM so that every scooter in 
each district of Berlin is maintained. Note that you may choose which district the FM should 
go to.

### Input / Output
As input you are given the `[]int scooters`, `int C` and `int P`.
Result should be `int` - the minimum number of FEs which are required to help the FM.

### Constraints
* `[]scooters` will contain between `1` and `100` elements.
* Each element in `scooters` will be between `0` and `1000`.
* `C` will be between `1` and `999`.
* `P` will be between `1` and `1000`.

### Examples
* input:
  ```
  { 
   scooters: [15, 10],
   C: 12,
   P: 5
  }
  ```
  
  expected output: 
  ```
  { fleet_engineers: 3 }
  ```

* input:
  ```
  { scooters: [11, 15, 13],
    C: 9,
    P: 5
  }
  ```

  expected output: 
  ```
  { fleet_engineers: 7 }
  ```

Please create an application (CLI or HTTP API) which solves this problem. Create a git 
repository and share the code with us through github or as a tar.gz. You can choose any 
programming language you are familiar with; java or ruby are preferred, though.
