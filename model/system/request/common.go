package request

type IdDTO struct {
	Ids []int // id集合
}

type List struct {
	Offset int // 分页参数
	Limit  int `valid:"range(1|200)"` // 分页参数
}
