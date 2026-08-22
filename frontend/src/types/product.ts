export interface Product {
  id: number
  name: string
  description: string
  price: number
  quantity: number
	categoryId: number
	categoryName: string
}

export interface ProductPayload {
  name: string
  description: string
  price: number
  quantity: number
  categoryId: number
}
