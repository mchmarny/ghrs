package calc

const (
	OperationUndefinedVal = "undefined"
	OperationAddVal       = "add"
	OperationSubVal       = "sub"
	OperationSetVal       = "set"
	OperationGetVal       = "get"
	OperationSaveVal      = "save"

	OperationUndefined = iota
	OperationAdd
	OperationSub
	OperationSet
	OperationGet
	OperationSave
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
	case OperationSave:
		return OperationSaveVal
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
	case OperationSaveVal:
		return OperationSave
	}
	return OperationUndefined
}
