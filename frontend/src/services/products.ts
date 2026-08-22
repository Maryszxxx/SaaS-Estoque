import { isAxiosError } from 'axios'
import api from './api'
import type { Product, ProductPayload } from '../types/product'

type ApiProduct = Record<string, unknown>
function asNumber(value: unknown) { return typeof value === 'number' ? value : Number(value ?? 0) }
function toProduct(product: ApiProduct): Product { return { id: asNumber(product.ID ?? product.id), name: String(product.Name ?? product.name ?? ''), description: String(product.Description ?? product.description ?? ''), price: asNumber(product.Price ?? product.price), quantity: asNumber(product.Quantity ?? product.quantity), categoryId: asNumber(product.CategoryID ?? product.category_id ?? product.categoryID), categoryName: String(product.CategoryName ?? product.category_name ?? product.categoryName ?? '') } }
function toApiPayload(product: ProductPayload) { return { name: product.name, description: product.description, price: product.price, quantity: product.quantity, category_id: product.categoryId } }

export async function getProducts(): Promise<Product[]> { try { const { data } = await api.get<ApiProduct[]>('/products'); return data.map(toProduct) } catch (error) { if (isAxiosError(error) && error.response?.status === 404) return []; throw error } }
export async function getProduct(id: number) { const { data } = await api.get<ApiProduct>(`/products/${id}`); return toProduct(data) }
export async function createProduct(product: ProductPayload) { await api.post('/products', toApiPayload(product)) }
export async function patchProduct(id: number, product: Partial<ProductPayload>) { const { categoryId, ...fields } = product; await api.patch(`/products/${id}`, { ...fields, ...(categoryId === undefined ? {} : { category_id: categoryId }) }) }
export async function deleteProduct(id: number) { await api.delete(`/products/${id}`) }
export async function restoreProduct(id: number) { await api.patch(`/products/${id}/restore`) }
