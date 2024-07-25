package crud

import "github.com/ikiruhawk/food-store/internal/models"

func GetProducts() []models.Product {
	c, m := getCategories(), getManufacturers()
	return []models.Product{
		{Id: 1, Name: "Potato", Category: c[0], Manufacturer: m[0], Price: 7000, Amount: 20, Description: "The potato is a starchy tuber that is cooked and eaten as a vegetable."},
		{Id: 2, Name: "Tomato", Category: c[0], Manufacturer: m[0], Price: 10000, Amount: 2023, Description: "The tomato is a red, juicy fruit that is often used in salads or cooked as a sauce ingredient."},
		{Id: 3, Name: "Cheeze", Category: c[1], Manufacturer: m[0], Price: 2000, Amount: 2012, Description: "Cheese is a dairy product obtained from milk, it is used in cooking, salads, and as a standalone snack."},
		{Id: 4, Name: "Milk", Category: c[1], Manufacturer: m[0], Price: 6000, Amount: 2023, Description: "Milk is a white, nutritious drink produced by mammals and used in cooking, baking, and as a beverage."},
		{Id: 5, Name: "Apple", Category: c[2], Manufacturer: m[0], Price: 11000, Amount: 203, Description: "The apple is a crisp, juicy fruit that is enjoyed raw or used in cooking and baking."},
		{Id: 6, Name: "Orange", Category: c[2], Manufacturer: m[1], Price: 5000, Amount: 240, Description: "The orange is a round, juicy citrus fruit that is often eaten fresh or used in juices."},
		{Id: 7, Name: "Mango", Category: c[2], Manufacturer: m[1], Price: 8000, Amount: 24, Description: "The mango is a sweet, fibrous tropical fruit that is enjoyed fresh or used in desserts and smoothies."},
		{Id: 8, Name: "Bread", Category: c[3], Manufacturer: m[1], Price: 12000, Amount: 203, Description: "Bread is a staple food made from flour, water, yeast or leavening agent, and sometimes other ingredients such as salt and sugar."},
		{Id: 9, Name: "Bun", Category: c[3], Manufacturer: m[1], Price: 4000, Amount: 22, Description: "A bun is a type of bread that is typically sweet and leavened, often served with butter or other spreads."},
		{Id: 10, Name: "Hotdog", Category: c[3], Manufacturer: m[1], Price: 1000, Amount: 21, Description: "A hot dog is a traditional American dish consisting of a fried sausage served in a cut bun."},
	}
}

func getCategories() []models.Category {
	return []models.Category{
		{Id: 1, Name: "Vegetables"},
		{Id: 2, Name: "Dairies"},
		{Id: 3, Name: "Fruits"},
		{Id: 4, Name: "Baking"},
	}
}

func getManufacturers() []models.Manufacturer {
	return []models.Manufacturer{
		{Id: 1, Name: "PishaBogov"},
		{Id: 2, Name: "GoHavat"},
	}
}

func GetProductById(id int) *models.Product {
	p := models.Product{}
	prs := GetProducts()

	for _, v := range prs {
		if v.Id == id {
			p = v
			break
		}
	}

	return &p
}
