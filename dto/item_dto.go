package dto

// DTOは「Data Transfer Object」の略で、異なるシステム間、またはアプリケーションの異なる層間でデータを転送するために使用されるオブジェクト
// Go言語の構造体タグは、Javaのアノテーションやマーカーインターフェースに似た役割を果たし、メタデータを提供し、実行時に特定の動作を指示するために使用される
type CreateItemInuput struct{
	Name string `json:"name" binding:"required,min=2"`
	Price uint `json:"price" binding:"required,min=1,max=999999"`
	Description string `json:"description"`
}

type UpdateItemInput struct{
	Name *string `json:"Name" binding:"omitnil,min=2"`
	Price *uint `json:"price" binding:"omitnil,min=1,max=999999"`
	Description *string `json:"description"`
	SoldOut *bool `json:"soldOut"`
}