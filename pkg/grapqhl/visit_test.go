package graphql

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	gophql "github.com/shuttl-io/gophql/pkg"
	"github.com/stretchr/testify/assert"
)

func TestParseReturnsProperly(t *testing.T) {
	assert := assert.New(t)
	r := ioutil.NopCloser(bytes.NewReader([]byte("query myquery { myfield }")))
	ast, err := Parse(r)
	assert.NotNil(ast)
	assert.NoError(err)
}

type printerVisitor struct {
}

func (p *printerVisitor) VisitField(field *gophql.FieldNode) error {
	fmt.Println(field.Name)
	return nil
}

func TestCanVisit(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte("query myquery { myfield { query } }")))
	ast, err := Parse(r)
	if err != nil {
		panic(err)
	}
	defer ast.Free()
	v := &printerVisitor{}
	visitor, err := gophql.NewVisitor(v)
	if err != nil {
		panic(err)
	}
	ast.Accept(visitor)
}

type v struct {
	visited []string
}

func (s *v) VisitField(node *gophql.FieldNode) error {
	s.visited = append(s.visited, node.Name)
	return fmt.Errorf("fail")
}

func TestAcceptsEndsEarlyWhenVisitorReturnsAnError(t *testing.T) {
	assert := assert.New(t)
	r := ioutil.NopCloser(bytes.NewReader([]byte("query myquery { myfield { query } }")))
	ast, err := Parse(r)
	if err != nil {
		panic(err)
	}

	defer ast.Free()
	v := &v{}
	visitor, err := gophql.NewVisitor(v)
	if !assert.NoError(err) {
		t.Fatal()
	}
	err = ast.Accept(visitor)
	assert.Error(err)
	assert.Len(v.visited, 1)
	assert.Equal("myfield", v.visited[0])
}

func TestThrowsAnError(t *testing.T) {
	assert := assert.New(t)
	r := ioutil.NopCloser(bytes.NewReader([]byte("query myquery oops { myfield }")))
	ast, err := Parse(r)
	assert.Nil(ast)
	assert.Error(err)
}
