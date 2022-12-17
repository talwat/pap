#!/bin/sh

GOPROXY=proxy.golang.org go list -m "github.com/talwat/pap@v$1"