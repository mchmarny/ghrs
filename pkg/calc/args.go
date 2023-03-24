package calc

const (
	StateArg     = "state"
	KeyArg       = "key"
	OperationArg = "operation"
	ValueArg     = "value"

	ResultArg = "value"
)

func GetArgs() map[string]string {
	return map[string]string{
		StateArg:     "data.db",
		KeyArg:       "counter",
		OperationArg: "add",
		ValueArg:     "0",
	}
}
