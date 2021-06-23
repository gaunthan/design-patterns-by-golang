package templatemethod

import "fmt"

type Builder interface {
	Build()
	CheckSyntax()
	Compile()
	Link()
}

type builderImpl struct {
	builder Builder
}

func (p *builderImpl) Build() {
	p.builder.CheckSyntax()
	p.builder.Compile()
	p.builder.Link()
}

type CPP11Builder struct {
	builderImpl
}

func NewCPP11Builder() Builder {
	b := &CPP11Builder{}
	b.builder = b
	return b
}

func (*CPP11Builder) CheckSyntax() {
	fmt.Println("checking code syntax by C++11 standard")
}

func (*CPP11Builder) Compile() {
	fmt.Println("compiling code by GCC 4.8")
}

func (*CPP11Builder) Link() {
	fmt.Println("linking code by GNU linker")
}

type CPP20Builder struct {
	builderImpl
}

func NewCPP20Builder() Builder {
	b := &CPP20Builder{}
	b.builder = b
	return b
}

func (*CPP20Builder) CheckSyntax() {
	fmt.Println("checking code syntax by C++20 standard")
}

func (*CPP20Builder) Compile() {
	fmt.Println("compiling code by GCC 8")
}

func (*CPP20Builder) Link() {
	fmt.Println("linking code by GNU linker")
}
