package main

import (
    "flag"
    
    "google.golang.org/protobuf/compiler/protogen"
)

func main() {
    var flags flag.FlagSet
    test := flags.String("test", "", "flags from command")
    opts := &protogen.Options{
        ParamFunc: flags.Set,
    }
    opts.Run(func(gen *protogen.Plugin) error {
        for _, f := range gen.Files {
            if !f.Generate {
                continue
            }
            generateFile(gen, f, test)
        }
        return nil
    })
}

// generateFile generates a _ascii.pb.go file containing gRPC service definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File, test *string) {
    filename := file.GeneratedFilenamePrefix + "_ascii.pb.go"
    g := gen.NewGeneratedFile(filename, file.GoImportPath)
    g.P("// Code generated by protoc-gen-go-ascii. DO NOT EDIT.")
    g.P("//", *test)
    g.P()
    g.P("package ", file.GoPackageName)
    g.P()
    g.P("func main() {")
    
    for _, srv := range file.Services {
        for _, method := range srv.Methods {
            if method.GoName == "Get" {
                g.P("// it's get")
            }
        }
    }
    
    g.P()
    g.P("}")
}
