package enum

const (
	ROLE_ADMIN = "admin"
	ROLE_USER  = "user"
)

type MqQueue string

func (r MqQueue) String() string {
	return string(r)
}

const (
	DefaultQueue  MqQueue = "default"
	TransferQueue MqQueue = "transfer"
)
