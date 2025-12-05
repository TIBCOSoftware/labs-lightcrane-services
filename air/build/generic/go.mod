module main

go 1.24.0

require (
	github.com/TIBCOSoftware/labs-air-contrib/activity/accumulator v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-air-contrib/activity/aggregate v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-air-contrib/activity/dataembedder v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-air-contrib/activity/dataselector v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-air-contrib/activity/error v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-air-contrib/activity/graphbuilder/builder v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-air-contrib/activity/log v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-air-contrib/activity/mqtt v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-air-contrib/activity/rules v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-air-contrib/activity/sqlinsert v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-air-contrib/function/air v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-air-contrib/trigger/kafka v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-air-contrib/trigger/mqtt v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-air-contrib/trigger/notificationlistener v0.0.0-20220812152750-89bbdfbc9f29
	github.com/TIBCOSoftware/labs-lightcrane-contrib/activity/exec v0.0.0-20220811184959-5f34ebac1524
	github.com/TIBCOSoftware/labs-lightcrane-contrib/activity/filereader v0.0.0-20220811184959-5f34ebac1524
	github.com/TIBCOSoftware/labs-lightcrane-contrib/activity/filewriter v0.0.0-20220811184959-5f34ebac1524
	github.com/TIBCOSoftware/labs-lightcrane-contrib/activity/httpclient v0.0.0-20220811184959-5f34ebac1524
	github.com/TIBCOSoftware/labs-lightcrane-contrib/activity/jsondeserializer v0.0.0-20220811184959-5f34ebac1524
	github.com/TIBCOSoftware/labs-lightcrane-contrib/activity/jsonserializer v0.0.0-20220811184959-5f34ebac1524
	github.com/TIBCOSoftware/labs-lightcrane-contrib/activity/mapping v0.0.0-20220811184959-5f34ebac1524
	github.com/TIBCOSoftware/labs-lightcrane-contrib/activity/objectmaker v0.0.0-20220811184959-5f34ebac1524
	github.com/TIBCOSoftware/labs-lightcrane-contrib/activity/objectserializer v0.0.0-20220811184959-5f34ebac1524
	github.com/TIBCOSoftware/labs-lightcrane-contrib/activity/pipecoupler v0.0.0-20220811184959-5f34ebac1524
	github.com/TIBCOSoftware/labs-lightcrane-contrib/activity/tablemutate v0.0.0-20220811184959-5f34ebac1524
	github.com/lib/pq v1.10.4
	github.com/mmussett/flogo-components v0.0.0-20210325110539-1e2ee32ab7c8
	github.com/project-flogo/contrib/activity/actreturn v0.10.0
	github.com/project-flogo/contrib/function/array v0.10.0
	github.com/project-flogo/contrib/function/coerce v0.10.0
	github.com/project-flogo/contrib/function/datetime v0.10.0
	github.com/project-flogo/contrib/function/number v0.10.0
	github.com/project-flogo/contrib/function/string v0.10.0
	github.com/project-flogo/contrib/function/utils v0.10.0
	github.com/project-flogo/contrib/trigger/rest v0.10.0
	github.com/project-flogo/contrib/trigger/timer v0.10.0
	github.com/project-flogo/core v1.6.1
	github.com/project-flogo/flow v1.6.1
	github.com/project-flogo/flow/activity/subflow v0.10.0
	github.com/project-flogo/legacybridge v0.9.4
)

require (
	github.com/Shopify/sarama v1.26.1 // indirect
	github.com/TIBCOSoftware/flogo-lib v0.5.8 // indirect
	github.com/TIBCOSoftware/labs-air-contrib/common/graphbuilder/model v0.0.0-20220812152750-89bbdfbc9f29 // indirect
	github.com/TIBCOSoftware/labs-air-contrib/common/notification v0.0.0-20220812152750-89bbdfbc9f29 // indirect
	github.com/TIBCOSoftware/labs-air-contrib/common/rules v0.0.0-20220812152750-89bbdfbc9f29 // indirect
	github.com/TIBCOSoftware/labs-air-contrib/common/util v0.0.0-20220812152750-89bbdfbc9f29 // indirect
	github.com/TIBCOSoftware/labs-air-contrib/connector/graph v0.0.0-20220812152750-89bbdfbc9f29 // indirect
	github.com/TIBCOSoftware/labs-lightcrane-contrib/common/exec v0.0.0-20220811184959-5f34ebac1524 // indirect
	github.com/TIBCOSoftware/labs-lightcrane-contrib/common/keywordreplace v0.0.0-20220811184959-5f34ebac1524 // indirect
	github.com/TIBCOSoftware/labs-lightcrane-contrib/common/objectbuilder v0.0.0-20220811184959-5f34ebac1524 // indirect
	github.com/TIBCOSoftware/labs-lightcrane-contrib/common/table v0.0.0-20220811184959-5f34ebac1524 // indirect
	github.com/TIBCOSoftware/labs-lightcrane-contrib/common/util v0.0.0-20220811184959-5f34ebac1524 // indirect
	github.com/TIBCOSoftware/labs-lightcrane-contrib/connector/simpletable v0.0.0-20220811184959-5f34ebac1524 // indirect
	github.com/araddon/dateparse v0.0.0-20190622164848-0fb0a474d195 // indirect
	github.com/carlescere/scheduler v0.0.0-20170109141437-ee74d2f83d82 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/eclipse/paho.mqtt.golang v1.3.5 // indirect
	github.com/edgexfoundry/go-mod-messaging/v2 v2.2.0 // indirect
	github.com/fxamacker/cbor/v2 v2.4.0 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-redsync/redsync/v4 v4.14.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/jcmturner/gofork v1.0.0 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/onsi/gomega v1.36.2 // indirect
	github.com/pierrec/lz4 v2.5.2+incompatible // indirect
	github.com/project-flogo/rules v0.1.4 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0 // indirect
	github.com/sirupsen/logrus v1.8.3 // indirect
	github.com/streadway/amqp v1.0.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/tkuchiki/go-timezone v0.1.3 // indirect
	github.com/tkuchiki/parsetime v0.0.0-20160517140325-bece749572a8 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.44.0 // indirect
	google.golang.org/protobuf v1.36.1 // indirect
	gopkg.in/jcmturner/aescts.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/dnsutils.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/gokrb5.v7 v7.5.0 // indirect
	gopkg.in/jcmturner/rpc.v1 v1.1.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
