package calc

const (
	OperationUndefinedVal = "undefined"
	OperationAddVal       = "add"
	OperationSubVal       = "sub"
	OperationSetVal       = "set"
	OperationGetVal       = "get"

	OperationUndefined = iota
	OperationAdd
	OperationSub
	OperationSet
	OperationGet
)

type Operation int64

func (o Operation) String() string {
	switch o {
	case OperationAdd:
		return OperationAddVal
	case OperationSub:
		return OperationSubVal
	case OperationSet:
		return OperationSetVal
	case OperationGet:
		return OperationGetVal
	}
	return OperationUndefinedVal
}

func ParseOperation(s string) Operation {
	switch s {
	case OperationAddVal:
		return OperationAdd
	case OperationSubVal:
		return OperationSub
	case OperationSetVal:
		return OperationSet
	case OperationGetVal:
		return OperationGet
	}
	return OperationUndefined
}
