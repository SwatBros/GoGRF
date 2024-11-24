package enums

type Cargo_class int

const (
	ALL_NORMAL_CARGO_CLASSES Cargo_class = iota
	ALL_CARGO_CLASSES
	NO_CARGO_CLASS
	CC_SPECIAL
	CC_REFRIGERATED
	CC_POWDERIZED
	CC_PIECE_GOODS
	CC_PASSENGERS
	CC_OVERSIZED
	CC_NON_POURABLE
	CC_NEO_BULK
	CC_MAIL
	CC_LIQUID
	CC_HAZARDOUS
	CC_EXPRESS
	CC_COVERED
	CC_BULK
	CC_ARMOURED
)

func (c Cargo_class) String() string {
	switch c {
	case NO_CARGO_CLASS:
		return "NO_CARGO_CLASS"
	case CC_SPECIAL:
		return "CC_SPECIAL"
	case CC_REFRIGERATED:
		return "CC_REFRIGERATED"
	case CC_POWDERIZED:
		return "CC_POWDERIZED"
	case CC_PIECE_GOODS:
		return "CC_PIECE_GOODS"
	case CC_PASSENGERS:
		return "CC_PASSENGERS"
	case CC_OVERSIZED:
		return "CC_OVERSIZED"
	case CC_NON_POURABLE:
		return "CC_NON_POURABLE"
	case CC_NEO_BULK:
		return "CC_NEO_BULK"
	case CC_MAIL:
		return "CC_MAIL"
	case CC_LIQUID:
		return "CC_LIQUID"
	case CC_HAZARDOUS:
		return "CC_HAZARDOUS"
	case CC_EXPRESS:
		return "CC_EXPRESS"
	case CC_COVERED:
		return "CC_COVERED"
	case CC_BULK:
		return "CC_BULK"
	case CC_ARMOURED:
		return "CC_ARMOURED"
	case ALL_NORMAL_CARGO_CLASSES:
		return "ALL_NORMAL_CARGO_CLASSES"
	case ALL_CARGO_CLASSES:
		return "ALL_CARGO_CLASSES"
	default:
		return ""
	}
}

type Cargo_classes []Cargo_class

func (c Cargo_classes) IsBitmask() bool {
	return len(c) == 1 && c[0] < 2
}
