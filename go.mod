module github.com/portworx/pxc

go 1.15

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/cheynewallace/tabby v1.1.1
	github.com/dustin/go-humanize v1.0.0
	github.com/gizak/termui/v3 v3.1.0
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/libopenstorage/openstorage-sdk-clients v0.109.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	golang.org/x/sys v0.0.0-20220624220833-87e55d714810 // indirect
	google.golang.org/grpc v1.36.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/cli-runtime v0.20.4
	k8s.io/client-go v0.20.4
)

replace github.com/grpc-ecosystem/grpc-gateway v1.9.0 => github.com/grpc-ecosystem/grpc-gateway v1.16.0
