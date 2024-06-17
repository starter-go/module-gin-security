package main

import (
	"os"

	"github.com/starter-go/security-gin-gorm/modules/securitygingorm"
	"github.com/starter-go/starter"
)

func main() {
	m := securitygingorm.ModuleForTest()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
