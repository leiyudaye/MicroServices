module UserInfoSvr

go 1.12

replace pb => ./pb

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/garyburd/redigo v1.6.2
	github.com/micro/micro/v3 v3.2.1
	golang.org/x/net v0.0.0-20210521195947-fe42d452be8f // indirect
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9 // indirect
	golang.org/x/sys v0.0.0-20210521203332-0cec03c779c1 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.36.1 // indirect
	google.golang.org/protobuf v1.26.0-rc.1 // indirect
	pb v0.0.0
)
