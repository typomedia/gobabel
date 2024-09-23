# gobabel - Babel CLI Transpiler

[![Go Report Card](https://goreportcard.com/badge/github.com/typomedia/gobabel)](https://goreportcard.com/report/github.com/typomedia/gobabel)
[![Go Reference](https://pkg.go.dev/badge/github.com/typomedia/gobabel.svg)](https://pkg.go.dev/github.com/typomedia/gobabel)
[![GitHub release](https://img.shields.io/github/release/typomedia/gobabel.svg)](https://github.com/typomedia/gobabel/releases/latest)
[![GitHub license](https://img.shields.io/github/license/typomedia/gobabel.svg)](https://github.com/typomedia/gobabel/blob/master/LICENSE)

This is a simple CLI to transpile JavaScript code using Babel. It is written in [Go](https://go.dev/) and uses the [goja-babel](https://github.com/jvatic/goja-babel) library.

## Motivation

For [lessgo](https://github.com/typomedia/lessgo) I needed a toolchain command to transpile JavaScript code using Babel **without the need** of installing the Node.js runtime.

## Install

    go install github.com/typomedia/gobabel@latest

# Flags

    -s, --source string   Source folder
    -t, --target string   Target folder
    -p, --preset string   Babel preset [env, react, flow] (default "env")
    -h, --help            help for gobabel

# Example

    gobabel -s less.js/packages/less/src/less/ -t dist/less/

## Help

    gobabel -h

## Build

    make

## Cross compile

    make cross

---
Copyright © 2024 Typomedia Foundation. All rights reserved.