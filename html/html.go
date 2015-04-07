// Copyright (c) 2015, Vastech SA (PTY) LTD. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package html

import (
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"strings"
)

type html struct {
	*generator.Generator
	generator.PluginImports
	ioPkg      generator.Single
	reflectPkg generator.Single
	stringsPkg generator.Single
	jsonPkg    generator.Single
}

func New() *html {
	return &html{}
}

func (p *html) Name() string {
	return "html"
}

func (p *html) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *html) typeName(name string) string {
	return p.TypeName(p.ObjectNamed(name))
}

func (p *html) writeError() {
	p.P(`if err != nil {`)
	p.In()
	p.P(`if err == `, p.ioPkg.Use(), `.EOF {`)
	p.In()
	p.P(`return`)
	p.Out()
	p.P(`}`)
	p.P(`w.Write([]byte(err.Error()))`)
	p.P(`return`)
	p.Out()
	p.P(`}`)
}

func (p *html) w(s string) {
	p.P(`w.Write([]byte("`, s, `"))`)
}

func formable(msg *descriptor.DescriptorProto) bool {
	return true
}

func (p *html) getInputType(method *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto {
	fileDescriptorSet := p.AllFiles()
	inputs := strings.Split(method.GetInputType(), ".")
	packageName := inputs[1]
	messageName := inputs[2]
	msg := fileDescriptorSet.GetMessage(packageName, messageName)
	if msg == nil {
		p.Fail("could not find message ", method.GetInputType())
	}
	return msg
}

func (p *html) generateSets(servName string, method *descriptor.MethodDescriptorProto) {
	msg := p.getInputType(method)
	if !formable(msg) {
		return
	}
	p.P(`fieldnames := []string{`)
	p.In()
	for _, f := range msg.GetField() {
		p.P(`"`, f.GetName(), `",`)
	}
	p.Out()
	p.P(`}`)
	p.P(`fields := make([]string, 0, len(fieldnames))`)
	p.P(`for _, name := range fieldnames {`)
	p.In()
	p.P(`v := req.FormValue(name)`)
	p.P(`if len(v) > 0 {`)
	p.In()
	p.P(`someValue = true`)
	p.P(`fields = append(fields, "\"" + name + "\":" + v)`)
	p.Out()
	p.P(`}`)
	p.P(`if someValue {`)
	p.In()
	p.P(`s := "{" + `, p.stringsPkg.Use(), `.Join(fields, ",") + "}"`)
	p.P(`err := `, p.jsonPkg.Use(), `.Unmarshal([]byte(s), msg)`)
	p.writeError()
	p.Out()
	p.P(`}`)
	p.Out()
	p.P(`}`)
}

func (p *html) generateForm(servName string, method *descriptor.MethodDescriptorProto) {
	msg := p.getInputType(method)
	p.P(`s := "<form action=\"/`, servName, `/`, method.GetName(), `\" method=\"GET\">"`)
	p.P(`w.Write([]byte(s))`)
	p.In()
	p.w(`Json for ` + servName + `(` + method.GetInputType() + `):<br>`)
	if !formable(msg) {
		panic("I don't think it is complicated")
		p.w(`<input name=\"json\" type=\"text\"><br>`)
	} else {
		for _, f := range msg.GetField() {
			p.w(f.GetName() + `: <input name=\"` + f.GetName() + `\" type=\"text\"><br>`)
		}
	}
	p.w(`<input type=\"submit\" value=\"Submit\"/>`)
	p.Out()
	p.w(`</form>`)
}

func (p *html) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	httpPkg := p.NewImport("net/http")
	p.jsonPkg = p.NewImport("encoding/json")
	p.ioPkg = p.NewImport("io")
	netPkg := p.NewImport("net")
	contextPkg := p.NewImport("golang.org/x/net/context")
	p.reflectPkg = p.NewImport("reflect")
	p.stringsPkg = p.NewImport("strings")
	for _, s := range file.GetService() {
		origServName := s.GetName()
		servName := generator.CamelCase(origServName)
		p.P(`type html`, servName, ` struct {`)
		p.In()
		p.P(`client `, servName, `Client`)
		p.P(`stringer func(interface{}) ([]byte, error)`)
		p.P(`port string`)
		p.Out()
		p.P(`}`)

		p.P(`func NewHTML`, servName, `Server(client `, servName, `Client, stringer func(interface{}) ([]byte, error)) *html`, servName, ` {`)
		p.In()
		p.P(`if stringer == nil {`)
		p.In()
		p.P(`stringer = `, p.jsonPkg.Use(), `.Marshal`)
		p.Out()
		p.P(`}`)
		p.P(`return &html`, servName, `{client, stringer, ":8080"}`)
		p.Out()
		p.P(`}`)

		p.P(`func (this *html`, servName, `) Serve(addr string) error {`)
		p.In()
		for _, m := range s.GetMethod() {
			p.P(httpPkg.Use(), `.HandleFunc("/`, servName, `/`, m.GetName(), `", this.`, m.GetName(), `)`)
		}
		p.P(`_, port, err := `, netPkg.Use(), `.SplitHostPort(addr)`)
		p.P(`if err != nil {`)
		p.In()
		p.P(`return err`)
		p.Out()
		p.P(`}`)
		p.P(`this.port = port`)
		p.P(`return `, httpPkg.Use(), `.ListenAndServe(addr, nil)`)
		p.Out()
		p.P(`}`)

		for _, m := range s.GetMethod() {
			p.P(`func (this *html`, servName, `) `, m.GetName(), `(w `, httpPkg.Use(), `.ResponseWriter, req *`, httpPkg.Use(), `.Request) {`)
			p.In()
			p.w("<html>")
			p.w("<head>")
			p.w("<title>" + servName + " - " + m.GetName() + "</title>")
			p.w("</head>")
			p.P(`jsonString := req.FormValue("json")`)
			p.P(`someValue := false`)
			p.P(`msg := &`, p.typeName(m.GetInputType()), `{}`)
			p.P(`if len(jsonString) > 0 {`)
			p.In()
			p.P(`err := `, p.jsonPkg.Use(), `.Unmarshal([]byte(jsonString), msg)`)
			p.writeError()
			p.P(`someValue = true`)
			p.Out()
			p.P(`} else {`)
			p.In()
			p.generateSets(servName, m)
			p.Out()
			p.P(`}`)
			p.generateForm(servName, m)
			p.P(`if someValue {`)
			p.In()
			if !m.GetClientStreaming() {
				if !m.GetServerStreaming() {
					p.P(`reply, err := this.client.`, m.GetName(), `(`, contextPkg.Use(), `.Background(), msg)`)
					p.writeError()
					p.P(`out, err := this.stringer(reply)`)
					p.writeError()
					p.P(`w.Write(out)`)
				} else {
					p.P(`down, err := this.client.`, m.GetName(), `(`, contextPkg.Use(), `.Background(), msg)`)
					p.writeError()
					p.P(`for {`)
					p.In()
					p.P(`reply, err := down.Recv()`)
					p.writeError()
					p.P(`out, err := this.stringer(reply)`)
					p.writeError()
					p.w(`<p>`)
					p.P(`w.Write(out)`)
					p.w(`</p>`)
					p.P(`w.(`, httpPkg.Use(), `.Flusher).Flush()`)
					p.Out()
					p.P(`}`)
				}
			} else {
				if !m.GetServerStreaming() {
					p.P(`up, err := this.client.Upstream(`, contextPkg.Use(), `.Background())`)
					p.writeError()
					p.P(`err = up.Send(msg)`)
					p.writeError()
					p.P(`reply, err := up.CloseAndRecv()`)
					p.writeError()
					p.P(`out, err := this.stringer(reply)`)
					p.writeError()
					p.P(`w.Write(out)`)
				} else {
					p.P(`bidi, err := this.client.Bidi(`, contextPkg.Use(), `.Background())`)
					p.writeError()
					p.P(`err = bidi.Send(msg)`)
					p.writeError()
					p.P(`reply, err := bidi.Recv()`)
					p.writeError()
					p.P(`out, err := this.stringer(reply)`)
					p.writeError()
					p.P(`w.Write(out)`)
				}
			}
			p.Out()
			p.P(`}`)
			p.w("</html>")
			p.Out()
			p.P(`}`)
		}
	}
}
