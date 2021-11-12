module mqserver

go 1.17

replace kumact.com/gosdk v0.0.0 => ../gosdk

replace iiwsm/sdk v0.0.0 => ../iiwsm/sdk

replace iiws_da/sdk v0.0.0 => ../iiws_da/sdk

require (
	github.com/streadway/amqp v1.0.0
	kumact.com/gosdk v0.0.0
)
