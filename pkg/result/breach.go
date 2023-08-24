package result

// Breach provides a representation for different breach types.
type Breach interface{}

type BreachType string

const (
	// BreachTypeValue is a breach with a value.
	BreachTypeValue BreachType = "value"
	// BreachTypeKeyValue is a breach with a key and a value.
	BreachTypeKeyValue BreachType = "key-value"
	// BreachTypeKeyValues is a breach with a key and a list of values.
	BreachTypeKeyValues BreachType = "key-values"
)

// Simple breach with no key.
// Example:
//
//	"file foo.ext not found": file is the ValueLabel, foo.ext is the Value
type ValueBreach struct {
	BreachType
	CheckType     string
	CheckName     string
	Severity      string
	ValueLabel    string
	Value         string
	ExpectedValue string
}

// Breach with key and value.
// Example:
//
//	"wordpress detected at /path/to/app":
//	  - file could be the KeyLabel
//	  - /path/to/app is the Key
//	  - app could be the ValueLabel
//	  - wordpress is the Value
type KeyValueBreach struct {
	BreachType
	CheckType     string
	CheckName     string
	Severity      string
	KeyLabel      string
	Key           string
	ValueLabel    string
	Value         string
	ExpectedValue string
}

// Breach with key and list of values.
// Example:
//
//	"[site_administrator] disallowed permissions: [administer site configuration, import configuration]":
//	  - role could be the KeyLabel
//	  - site_administrator is the Key
//	  - permissions could be the ValueLabel
//	  - [administer site configuration, import configuration] are the Values
type KeyValuesBreach struct {
	BreachType
	CheckType  string
	CheckName  string
	Severity   string
	KeyLabel   string
	Key        string
	ValueLabel string
	Values     []string
}

func BreachSetCommonValues(bIfc *Breach, checkType string, checkName string, severity string) {
	if b, ok := (*bIfc).(ValueBreach); ok {
		b.BreachType = BreachTypeValue
		b.CheckType = checkType
		b.CheckName = checkName
		b.Severity = severity
		*bIfc = b
	} else if b, ok := (*bIfc).(KeyValueBreach); ok {
		b.BreachType = BreachTypeKeyValue
		b.CheckType = checkType
		b.CheckName = checkName
		b.Severity = severity
		*bIfc = b
	} else if b, ok := (*bIfc).(KeyValuesBreach); ok {
		b.BreachType = BreachTypeKeyValues
		b.CheckType = checkType
		b.CheckName = checkName
		b.Severity = severity
		*bIfc = b
	}
}

func BreachGetBreachType(bIfc Breach) BreachType {
	if _, ok := bIfc.(ValueBreach); ok {
		return BreachTypeValue
	} else if _, ok := bIfc.(KeyValueBreach); ok {
		return BreachTypeKeyValue
	} else if _, ok := bIfc.(KeyValuesBreach); ok {
		return BreachTypeKeyValues
	}
	return ""
}

func BreachGetCheckType(bIfc Breach) string {
	if b, ok := bIfc.(ValueBreach); ok {
		return b.CheckType
	} else if b, ok := bIfc.(KeyValueBreach); ok {
		return b.CheckType
	} else if b, ok := bIfc.(KeyValuesBreach); ok {
		return b.CheckType
	}
	return ""
}

func BreachGetCheckName(bIfc Breach) string {
	if b, ok := bIfc.(ValueBreach); ok {
		return b.CheckName
	} else if b, ok := bIfc.(KeyValueBreach); ok {
		return b.CheckName
	} else if b, ok := bIfc.(KeyValuesBreach); ok {
		return b.CheckName
	}
	return ""
}

func BreachGetSeverity(bIfc Breach) string {
	if b, ok := bIfc.(ValueBreach); ok {
		return b.Severity
	} else if b, ok := bIfc.(KeyValueBreach); ok {
		return b.Severity
	} else if b, ok := bIfc.(KeyValuesBreach); ok {
		return b.Severity
	}
	return ""
}

func BreachGetKeyLabel(bIfc Breach) string {
	if b, ok := bIfc.(KeyValueBreach); ok {
		return b.KeyLabel
	} else if b, ok := bIfc.(KeyValuesBreach); ok {
		return b.KeyLabel
	}
	return ""
}

func BreachGetKey(bIfc Breach) string {
	if b, ok := bIfc.(KeyValueBreach); ok {
		return b.Key
	} else if b, ok := bIfc.(KeyValuesBreach); ok {
		return b.Key
	}
	return ""
}

func BreachGetValueLabel(bIfc Breach) string {
	if b, ok := bIfc.(ValueBreach); ok {
		return b.ValueLabel
	} else if b, ok := bIfc.(KeyValueBreach); ok {
		return b.ValueLabel
	} else if b, ok := bIfc.(KeyValuesBreach); ok {
		return b.ValueLabel
	}
	return ""
}

func BreachGetValue(bIfc Breach) string {
	if b, ok := bIfc.(ValueBreach); ok {
		return b.Value
	} else if b, ok := bIfc.(KeyValueBreach); ok {
		return b.Value
	}
	return ""
}

func BreachGetValues(bIfc Breach) []string {
	if b, ok := bIfc.(KeyValuesBreach); ok {
		return b.Values
	}
	return []string(nil)
}

func BreachGetExpectedValue(bIfc Breach) string {
	if b, ok := bIfc.(ValueBreach); ok {
		return b.ExpectedValue
	} else if b, ok := bIfc.(KeyValueBreach); ok {
		return b.ExpectedValue
	}
	return ""
}