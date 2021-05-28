module UserInfoSvr

go 1.16

replace pb => ./pb

require (
	github.com/garyburd/redigo v1.6.2
	github.com/johanbrandhorst/grpc-json-example v0.0.0-20180722152311-a20c72379c60
	golang.org/x/net v0.0.0-20210521195947-fe42d452be8f // indirect
	golang.org/x/sys v0.0.0-20210521203332-0cec03c779c1 // indirect
	google.golang.org/genproto v0.0.0-20210524171403-669157292da3 // indirect
	google.golang.org/grpc v1.38.0
	pb v0.0.0
)
