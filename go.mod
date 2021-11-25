module kumact.com/dc_scheduler

go 1.17

replace kumact.com/gosdk v0.0.0 => ../gosdk

replace iiwsm/sdk v0.0.0 => ../iiwsm/sdk

replace iiws_da/sdk v0.0.0 => ../iiws_da/sdk

replace kumact.com/dc_scheduler v0.0.0 => ./

require (
	gitea.com/lunny/tango v0.6.5
	github.com/streadway/amqp v1.0.0
	kumact.com/gosdk v0.0.0
)

require (
	gitea.com/lunny/log v0.0.0-20190322053110-01b5df579c4e // indirect
	gitea.com/tango/session v0.0.0-20201110080243-87f6e468e457 // indirect
	github.com/i-solveware/uuid v0.0.0-20160212161150-cd53251766d7 // indirect
	golang.org/x/net v0.0.0-20211111160137-58aab5ef257a // indirect
)
