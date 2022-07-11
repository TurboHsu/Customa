package main

import (
	"fmt"

	c "github.com/TurboHsu/Customa/drivers/config"
	db "github.com/TurboHsu/Customa/drivers/database"
	l "github.com/TurboHsu/Customa/drivers/log"
)

func main() {
	//Reads the config
	err := c.Read("config.toml")
	if err.IsErr {
		panic(err.Msg)
	}

	err = db.Conn(c.Config.Database.Addr, c.Config.Database.User, c.Config.Database.Passwd, c.Config.Database.DatabaseName)
	if err.IsErr {
		l.Error(err.Msg)
	}

	l.Info("Customa started")
	//Useless test function.
	fmt.Println(db.Test())
}
