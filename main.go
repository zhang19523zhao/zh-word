package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhang19523zhao/zh-word/module"
	"html/template"
	"log"
	"net/http"
)

const (
	//sqlTask = "select count(*) num, today, name from info group by name order by num desc;"
	sqlTask = `select a.name, a.today ,   a.num,        b.tnum from         (         select             name, today, count(*) as num         from             info         group by name ) as a left join ( select count(*) tnum, today, name  from info  where to_days(today) = to_days(now())  group by name order by tnum )as b on a.name = b.name order by num desc;`
)

type Task struct {
	Num  int
	Tnum int
	Name string
	Time string
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset-utf8mb4)", "root", "zhanghaodemima19", "localhost", 3306, "zhword")
	//fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	//测试数据库连接
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	addr := ":8888"

	//显示任务列表
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//从mysql 获取name time num
		tasks := make([]Task, 0, 20)
		rows, err := db.Query(sqlTask)
		for rows.Next() {
			var task Task
			err = rows.Scan(&task.Name, &task.Time, &task.Num, &task.Tnum)
			if err != nil {
				fmt.Println(err)
			}
			tasks = append(tasks, task)
		}
		//fmt.Println(tasks)
		tpl := template.Must(template.New("tpl").ParseFiles("vieiws/info2.html"))
		tpl.ExecuteTemplate(w, "info2.html", tasks)
	})

	//增加单词
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.PostFormValue("name")
			word := r.PostFormValue("word")
			chinese := r.PostFormValue("chinese")
			module.Add(db, name, word, chinese)
			//http.Redirect(w, r, "/", 302)
		}
		tpl := template.Must(template.New("tpl").ParseFiles("vieiws/add.html"))
		tpl.ExecuteTemplate(w, "add.html", nil)
	})

	//结果
	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		var rs module.Res
		if r.Method == http.MethodPost {
			name := r.PostFormValue("name")

			rs = module.Query(db, name)
			fmt.Println(rs)
			//http.Redirect(w, r, "/", 302)
		}
		tpl := template.Must(template.New("tpl").ParseFiles("vieiws/result.html"))
		tpl.ExecuteTemplate(w, "result.html", rs)
	})

	//查询
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.New("tpl").ParseFiles("vieiws/query.html"))
		tpl.ExecuteTemplate(w, "query.html", nil)
	})

	http.ListenAndServe(addr, nil)
}
