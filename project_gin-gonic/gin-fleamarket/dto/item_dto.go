package dto

type CreateItemInput struct {
	Name        string `json:"name" binding:"required,min=2"`
	Price       uint   `json:"price" binding:"required,min=1,max=999999"`
	Description string `json:"description"`
}

// 컬럼이 있는 경우에만 업데이트하기 위해 *를 붙임. *를 붙이지 않으면 해당 필드가 없는 경우에도 업데이트가 됨.
// binding:"omitnil"은 포인터 타입이 nil이 아닌 경우에만 유효성 검사를 수행한다. 즉, 해당 필드가 없는 경우에는 유효성 검사를 수행하지 않는다.
type UpdateItemInput struct {
	Name        *string `json:"name" binding:"omitnil,min=2"`
	Price       *uint   `json:"price" binding:"omitnil,min=1,max=999999"`
	Description *string `json:"description"`
	SoldOut     *bool   `json:"soldOut"`
}
