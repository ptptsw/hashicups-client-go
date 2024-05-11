package hashicups

// Order -
type Food struct {
	ID    int         `json:"id,omitempty"`
	Items []FoodItem `json:"items,omitempty"`
}

// OrderItem -
type FoodItem struct {
	Name        string             `json:"name"`
	Price       float64            `json:"price"`
}

// // Coffee -
// type Coffee struct {
// 	ID          int                `json:"id"`
// 	Name        string             `json:"name"`
// 	Teaser      string             `json:"teaser"`
// 	Collection  string             `json:"collection"`
// 	Origin      string             `json:"origin"`
// 	Color       string             `json:"color"`
// 	Description string             `json:"description"`
// 	Price       float64            `json:"price"`
// 	Image       string             `json:"image"`
// 	Ingredient  []CoffeeIngredient `json:"ingredients"`
// }

// // Ingredient -
// type CoffeeIngredient struct {
// 	ID       int    `json:"ingredient_id"`
// 	Name     string `json:"name"`
// 	Quantity int    `json:"quantity"`
// 	Unit     string `json:"unit"`
// }

// // Ingredient -
// type Ingredient struct {
// 	ID       int    `json:"id"`
// 	Name     string `json:"name"`
// 	Quantity int    `json:"quantity"`
// 	Unit     string `json:"unit"`
// }
