package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

var (
	// 1->产品，2->前端，3->后端，4->安卓，5->设计
	groupMap = map[uint8]string{1: "产品", 2: "前端", 3: "后端", 4: "安卓", 5: "设计"}

	// db
	username = ""
	password = ""
	addr     = ""
	DBName   = "workbench"

	// mail
	mailUser = ""
	authCode = ""
)

type Item struct {
	GroupId uint8
	Total   uint32
}

/*
环境变量

export REMINDER_DB_ADDR=
export REMINDER_DB_USERNAME=
export REMINDER_DB_PASSWORD=

export REMINDER_MAIL_USER=
export REMINDER_MAIL_AUTH=
*/

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("REMINDER")
	addr = viper.GetString("DB_ADDR")
	username = viper.GetString("DB_USERNAME")
	password = viper.GetString("DB_PASSWORD")
	mailUser = viper.GetString("MAIL_USER")
	authCode = viper.GetString("MAIL_AUTH")
}

func main() {
	DB.Init()
	// Email()
	// return

	c := cron.New()

	c.AddFunc("*/10 * * * * *", Email)
	// c.AddFunc("*/10 * * * * *", func() { Email() })
	// c.AddFunc("*/5 * * * * *", func() { fmt.Println("hello") })

	c.Start()
	defer c.Stop()

	for {
	}
}

func Email() {
	// var to = []string{"1027319981@qq.com", "shdwzhang@163.com"}
	// var to = []string{"shdwzhang@163.com"}
	var to = []string{"654957943@qq.com"}
	var body string
	var subject = "工作台进度每周汇总"

	items, err := Query()
	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range items {
		body += fmt.Sprintf("%s： %d\n", groupMap[item.GroupId], item.Total)
	}

	fmt.Println(body)

	if err := SendMail(mailUser, authCode, to, subject, body); err != nil {
		log.Println(err)
	}
	// if err := SendGoMail(user, auth, to, subject, body); err != nil {
	// 	log.Println(err)
	// }

	log.Println("Send mail OK")
}

// select users.group_id, count(group_id) from status inner join users on users.id = status.user_id
// where status.id < 100 and users.group_id is not null group by users.group_id

// SELECT * FROM status WHERE DATE_SUB(curdate(), INTERVAL 6 DAY) <= DATE(time)\G

// select users.group_id, count(group_id) from status inner join users on users.id = status.user_id
// where date_sub(curdate(), interval 6 day) <= date(time) and users.group_id is not null group by users.group_id

func Query() ([]Item, error) {
	query := DB.Self.Table("status").Select("users.group_id, count(*) as total").
		Joins("left join users on users.id = status.user_id").
		Where("date_sub(curdate(), interval 6 day) <= date(status.time) and users.group_id is not null").
		Group("users.group_id")

	var items []Item
	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}
	fmt.Println(items)

	return items, nil
}
