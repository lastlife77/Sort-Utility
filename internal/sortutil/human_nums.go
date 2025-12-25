package sortutil

type humanNum int

const (
	b humanNum = iota //byte
	k                 //Kilobyte
	m                 //Megabyte
	g                 //Gigabyte
	t                 //Terabyte
	p                 //Petabyte
	e                 //Exabyte
	z                 //Zettabyte
	y                 //Yottabyte
)

var humanNums = []struct {
	name  string
	value humanNum
}{
	{"b", b},
	{"k", k},
	{"m", m},
	{"g", g},
	{"t", t},
	{"p", p},
	{"e", e},
	{"z", z},
	{"y", y},
}
