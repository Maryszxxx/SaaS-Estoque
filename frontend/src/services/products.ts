import { isAxiosError } from 'axios'
import api from './api'
import type { Product } from '../types/product'

type ApiProduct = Record<string, unknown>

function asNumber(value: unknown) {
  return typeof value === 'number' ? value : Number(value ?? 0)
}

function toProduct(product: ApiProduct): Product {
  return {
    id: asNumber(product.ID ?? product.id),
    name: String(product.Name ?? product.name ?? ''),
    description: String(product.Description ?? product.description ?? ''),
    price: asNumber(product.Price ?? product.price),
    quantity: asNumber(product.Quantity ?? product.quantity),
    categoryId: asNumber(product.CategoryID ?? product.category_id ?? product.categoryID),
  }
}

export async function getProducts(): Promise<Product[]> {
  try {
    const { data } = await api.get<ApiProduct[]>('/products')
    return data.map(toProduct)
  } catch (error) {
    if (isAxiosError(error) && error.response?.status === 404) return []
    throw error
  }
}
