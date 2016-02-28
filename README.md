libpostal-go
============

The package provides Go bindings for libpostal, a C library for fast international street address parsing and normalization.<br />
All the binding code has automatically been generated with rules defined in [postal.yml](/postal.yml).

### Usage

```
$ go get github.com/xlab/libpostal-go/postal
```

### Demo

There is an example parser command line utility. You will need to setup your [libpostal](https://github.com/openvenues/libpostal#installation) first, follow the installation instructions in the official distribution.

```
$ pkg-config --libs libpostal
-L/usr/local/lib -lpostal

$ go get github.com/xlab/libpostal-go/cmd/postal-parser

$ postal-parser Champ de Mars, 5 Avenue Anatole France, 75007 Paris, France
main.go:40: address components: 6
main.go:52:
╭──────────────────────────────────────╮
│           libpostal parse            │
├──────────────┬───────────────────────┤
│ road         │ champ de mars         │
│ house_number │ 5                     │
│ road         │ avenue anatole france │
│ postcode     │ 75007                 │
│ city         │ paris                 │
│ country      │ france                │
╰──────────────┴───────────────────────╯
```

### Rebuilding the package

You will need to get the [cgogen](https://git.io/cgogen) tool installed first.

```
$ git clone https://github.com/xlab/libpostal-go && cd libpostal-go
$ make clean
$ make
```
