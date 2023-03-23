package calc

const (
	DataKey    = "data"
	CounterKey = "counter"
	ActionKey  = "action"
	ValueKey   = "value"

	ResultKey = "result"
)

func GetArgs() map[string]string {
	return map[string]string{
		DataKey:    "data.db",
		CounterKey: "counter",
		ActionKey:  "add",
		ValueKey:   "0",
	}
}
