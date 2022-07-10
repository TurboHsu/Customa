package database

import (
	"database/sql"
	"fmt"

	"github.com/TurboHsu/Customa/structs"
	_ "github.com/go-sql-driver/mysql"
)

//Make a variable for whole package.
var db *sql.DB

//Connects to the database.
func Conn(addr string, usr string, passwd string, DB string) structs.ErrMsg {
	var err error
	var ret structs.ErrMsg = structs.ErrMsg{IsErr: false, Msg: ""}
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", usr, passwd, addr, DB))
	if err != nil {
		ret.IsErr = true
		ret.Msg = err.Error()
		return ret
	} else {
		return ret
	}
}

//Temporarily test function.
func Test() string {
	r, _ := db.Query("select * from untitled_table where id = 1;")
	rr, _ := r.Columns()
	var ret string
	for i := 0; i < len(rr); i++ {
		ret = fmt.Sprintf("%s %s", ret, rr[i])
	}
	return ret
}
