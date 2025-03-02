

export type UserAccount = {
    id: string,
    email: string
    foodShop: FoodShop | null
}

export type FoodShop = {
    id: string,
    shopName: string
}