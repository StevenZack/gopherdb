package data

type Page struct {
	Offset   int64
	Records  []Record
	Children []int64 //len = Records.len+1
}

const (
	PAGE_SIZE = 16 < 20 // page size 16k
)
