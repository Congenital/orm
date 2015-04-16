package contrast

var SqlType = map[string]string{
	"int8":     "tinyint",
	"int16":    "smaillint",
	"int32":    "mediumint",
	"int":      "int",
	"int64":    "bigint",
	"float":    "float",
	"double":   "double",
	"double64": "real",
	"time":     "datetime",
	"uint64":   "timestamp",
	"string":   "varchar(767)",
	"[]byte":   "blob",
}

var (
	Tags = []string{
		"p",
		"f",
		"u",
		"n",
		"c",
		"d",
	}

	Constraint = map[string]string{
		"p": "primary key",
		"f": "foreign key",
		"u": "unique",
		"n": "not null",
		"c": "check",
		"d": "default",
	}

	Tags_length = len(Tags)
)
