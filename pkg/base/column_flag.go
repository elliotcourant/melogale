package base

type ColumnFlag uint8

const (
	ColumnPrimaryKey ColumnFlag = 1
	ColumnIndexed    ColumnFlag = 2
	ColumnUnique     ColumnFlag = 4
)

func (i ColumnFlag) IsPrimaryKey() bool {
	return i&ColumnPrimaryKey != 0
}

func (i ColumnFlag) IsIndexed() bool {
	return i&ColumnIndexed != 0 || i.IsPrimaryKey()
}

func (i ColumnFlag) IsUnique() bool {
	return i&ColumnUnique != 0 || i.IsPrimaryKey()
}

func NewColumnFlag(flags ...ColumnFlag) ColumnFlag {
	if len(flags) == 0 {
		return 0
	}
	flag := flags[0]
	for i, f := range flags {
		if i == 0 {
			continue
		}
		flag |= f
	}
	return flag
}
