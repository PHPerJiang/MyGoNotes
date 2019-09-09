package my_log

import "log"

func Log_test()  {
	log.SetPrefix("Trace：")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	//println写到标准日志记录器中
	log.Println("log.println将错误信息写到标准日志记录器中")

	//fatalln在调用println之后会接着调用os.exit(1)
	log.Fatalln("fatalln在调用println之后会接着调用os.exit(1)")

	//panicln在调用println之后会调用panic
	log.Panicln("panicln在调用println之后会调用panic")

}
