module kumact.com/mqserver

go 1.17

replace kumact.com/gosdk v0.0.0 => ../gosdk

replace iiwsm/sdk v0.0.0 => ../iiwsm/sdk

replace iiws_da/sdk v0.0.0 => ../iiws_da/sdk

replace kumact.com/mqserver v0.0.0 => ./

require (
	github.com/streadway/amqp v1.0.0
	kumact.com/gosdk v0.0.0
)

require golang.org/x/net v0.0.0-20211111160137-58aab5ef257a // indirect
