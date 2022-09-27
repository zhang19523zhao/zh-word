package module

import (
	"database/sql"
	"fmt"
	"log"
)

type Result struct {
	Rname    string
	Rtime    string
	Rword    string
	Rchinese string
}

type Res []*Result

func Add(db *sql.DB, name, word, chinese string) {
	if name == "" || word == "" || chinese == "" {
		log.Println("name or word or chinese have a empty...")
	}
	//sqlAdd := fmt.Sprintf("insert into info(name,word,chinese) values(%s, %s, %s)", name, word, chinese)
	stmt, err := db.Prepare(`INSERT info (name, word, chinese) values (?,?,?)`)
	if err != nil {
		fmt.Println(err)
	}
	stmt.Exec(name, word, chinese)
}

func Query(db *sql.DB, name string) Res {
	if name == "" {
		log.Println("name have a empty...")
	}
	sqlT := fmt.Sprintf(`select * from info where name='%s'`, name)
	rows, err := db.Query(sqlT)
	RS := make([]*Result, 0, 20)
	for rows.Next() {
		var res Result
		err = rows.Scan(&res.Rname, &res.Rtime, &res.Rword, &res.Rchinese)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res.Rname, res.Rword)
		RS = append(RS, &res)
	}
	return RS
}
