package crud

import "github.com/ikiruhawk/food-store/internal/models"

func GetProducts() []models.Product {
	return []models.Product{
		{Id: 1, Name: "Potato", Category: "Vegetables", Price: 70.00, Description: "The potato is a starchy tuber that is cooked and eaten as a vegetable."},
		{Id: 2, Name: "Tomato", Category: "Vegetables", Price: 100.00, Description: "The tomato is a red, juicy fruit that is often used in salads or cooked as a sauce ingredient."},
		{Id: 3, Name: "Cheeze", Category: "Dairies", Price: 20.00, Description: "Cheese is a dairy product obtained from milk, it is used in cooking, salads, and as a standalone snack."},
		{Id: 4, Name: "Milk", Category: "Dairies", Price: 60.00, Description: "Milk is a white, nutritious drink produced by mammals and used in cooking, baking, and as a beverage."},
		{Id: 5, Name: "Apple", Category: "Fruits", Price: 110.00, Description: "The apple is a crisp, juicy fruit that is enjoyed raw or used in cooking and baking."},
		{Id: 6, Name: "Orange", Category: "Fruits", Price: 50.00, Description: "The orange is a round, juicy citrus fruit that is often eaten fresh or used in juices."},
		{Id: 7, Name: "Mango", Category: "Fruits", Price: 80.00, Description: "The mango is a sweet, fibrous tropical fruit that is enjoyed fresh or used in desserts and smoothies."},
		{Id: 8, Name: "Bread", Category: "Baking", Price: 120.00, Description: "Bread is a staple food made from flour, water, yeast or leavening agent, and sometimes other ingredients such as salt and sugar."},
		{Id: 9, Name: "Bun", Category: "Baking", Price: 40.00, Description: "A bun is a type of bread that is typically sweet and leavened, often served with butter or other spreads."},
		{Id: 10, Name: "Hotdog", Category: "Baking", Price: 10.00, Description: "A hot dog is a traditional American dish consisting of a fried sausage served in a cut bun."},
	}
}

func GetProductById(id int) *models.Product {
	prs := GetProducts()

	p := models.Product{}

	for _, v := range prs {
		if v.Id == id {
			p = v
			break
		}
	}

	return &p
}
