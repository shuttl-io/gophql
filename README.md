# GophQL

A Go Binding of libgraphql

## Installation

This is meant to be a used like any other Go module. You just need to run `go get github.com/shuttl-io/gophql` to install this. However this is also making the assumption that you have libgraphqlparser already installed. If you don't you can [follow the instructions on their repo](https://github.com/graphql/libgraphqlparser). This library was built with version 0.7.0 of this library.

## Usage

You must implement a visitor to visit a node, then you simply do this:

```go
package main

import (
    "fmt"

    "github.com/shuttl-io/gophql"
)

type MyVisitor struct {}

func (m *MyVisitor) VisitField(node gophql.FieldNode) error {
    fmt.Println(node.Name)
    return nil
}

// implement other VisitorInterfaces or use gophql.NewVisitor

func main() {
    ast, err := gophql.Parse("query main { field }")
    if err != nil {
        panic(err)
    }
    defer ast.Free() //Prevents memory leaks
    ast.Accept(myVisitor) // Will print "field"
}
```
