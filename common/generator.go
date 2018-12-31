package common

import "fmt"

// EnvGenerator is used to generate strings that could be eval to set env
type EnvGenerator interface {
	GenerateEnv() string
}

// OutputEnv takes a generator end output its evaluable env string
// All subcommands should finish by invoking this function
func OutputEnv(g EnvGenerator) {
	str := g.GenerateEnv()
	fmt.Println(str)
}
