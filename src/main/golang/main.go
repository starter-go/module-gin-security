package main

import (
	"os"

	"github.com/starter-go/security-gin-gorm/modules/securitygingorm"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(securitygingorm.Module())
	i.WithPanic(true).Run()
}
