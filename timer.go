package main

import (
    "github.com/robfig/cron"
    "log"
)

func main() {
    i := 0
    c := cron.New()
    spec := "*/5 * * * * ?"
    c.AddFunc(spec, func() {
        i++
        log.Println("cron running:", i)
    })
    c.Start()
    
    select{}
}


//go mod tidy
/*每隔5秒执行一次：*/5 * * * * ?

            每隔1分钟执行一次：0 */1 * * * ?

            每天23点执行一次：0 0 23 * * ?

            每天凌晨1点执行一次：0 0 1 * * ?

            每月1号凌晨1点执行一次：0 0 1 1 * ?

            在26分、29分、33分执行一次：0 26,29,33 * * * ?

            每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?*/
