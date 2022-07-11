package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/TurboHsu/Customa/structs"
)

var Config structs.Config

func Read(fileAddr string) structs.ErrMsg {
	var ret structs.Config
	var error structs.ErrMsg = structs.ErrMsg{IsErr: false, Msg: ""}

	//Reads the config file
	file, err := os.Open(fileAddr)
	if err != nil {
		error.IsErr = true
		error.Msg = err.Error()
		return error
	}
	defer file.Close()
	cfg, err := ioutil.ReadAll(file)
	if err != nil {
		error.IsErr = true
		error.Msg = err.Error()
		return error
	}

	//Unmarshal the config file.
	err = toml.Unmarshal(cfg, &ret)
	if err != nil {
		error.IsErr = true
		error.Msg = err.Error()
		return error
	}
	Config = ret
	return error
}
