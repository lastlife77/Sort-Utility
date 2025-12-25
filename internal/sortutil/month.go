package sortutil

type month int

const (
	jan month = 1 + iota
	feb
	mar
	apr
	may
	jun
	jul
	aug
	sep
	oct
	nov
	dec
)

var months = []struct {
	name  string
	value month
}{
	{"jan", jan},
	{"feb", feb},
	{"mar", mar},
	{"apr", apr},
	{"may", may},
	{"jun", jun},
	{"jul", jul},
	{"aug", aug},
	{"sep", sep},
	{"oct", oct},
	{"nov", nov},
	{"dec", dec},
}
