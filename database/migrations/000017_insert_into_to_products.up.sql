INSERT INTO categories (name) VALUES
('Vegetables'),
('Dairies'),
('Fruits'),
('Baking');

INSERT INTO manufacturers(name) VALUES
('PishaBogov'),
('GoHavat');

INSERT INTO products(category_id, manufacturer_id, name, price, amount, description) VALUES
(1, 1, 'Potato', 7000, 20, 'The potato is a starchy tuber that is cooked and eaten as a vegetable.'),
(1, 1, 'Tomato', 10000, 2023, 'The tomato is a red, juicy fruit that is often used in salads or cooked as a sauce ingredient.'),
(2, 1, 'Cheeze', 2000, 2012, 'Cheese is a dairy product obtained from milk, it is used in cooking, salads, and as a standalone snack.'),
(2, 1, 'Milk', 6000, 2023, 'Milk is a white, nutritious drink produced by mammals and used in cooking, baking, and as a beverage.'),
(3, 1, 'Apple', 11000, 203, 'The apple is a crisp, juicy fruit that is enjoyed raw or used in cooking and baking.'),
(3, 2, 'Orange', 5000, 240, 'The orange is a round, juicy citrus fruit that is often eaten fresh or used in juices.'),
(3, 2, 'Mango', 8000, 24, 'The mango is a sweet, fibrous tropical fruit that is enjoyed fresh or used in desserts and smoothies.'),
(4, 2, 'Bread', 12000, 203, 'Bread is a staple food made from flour, water, yeast or leavening agent, and sometimes other ingredients such as salt and sugar.'),
(4, 2, 'Bun', 4000, 22, 'A bun is a type of bread that is typically sweet and leavened, often served with butter or other spreads.'),
(4, 2, 'Hotdog', 1000, 21, 'A hot dog is a traditional American dish consisting of a fried sausage served in a cut bun.');
