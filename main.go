package main

import (
	"fmt"

	c "github.com/TurboHsu/Customa/drivers/config"
	db "github.com/TurboHsu/Customa/drivers/database"
)

func main() {
	config, err := c.Read("config.toml")
	if err.IsErr {
		panic(err.Msg)
	}
	err = db.Conn(config.Database.Addr, config.Database.User, config.Database.Passwd, config.Database.DatabaseName)
	if err.IsErr {
		panic(err.Msg)
	}

	//Useless test function.
	fmt.Println(db.Test())
}
