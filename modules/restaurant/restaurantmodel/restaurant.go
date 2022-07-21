package restaurantmodel

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"address" gorm:"column:addr"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr string  `json:"address" gorm:"column:addr"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

// CREATE TABLE `restaurants` (
// 	`id` int(11) NOT NULL AUTO_INCREMENT,
// 	`owner_id` int(11) NOT NULL,
// 	`name` varchar(50) NOT NULL,
// 	`addr` varchar(255) NOT NULL,
// 	`city_id` int(11) DEFAULT NULL,
// 	`lat` double DEFAULT NULL,
// 	`lng` double DEFAULT NULL,
// 	`cover` json NOT NULL,
// 	`logo` json NOT NULL,
// 	`shipping_fee_per_km` double DEFAULT '0',
// 	`status` int(11) NOT NULL DEFAULT '1',
// 	`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
// 	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 	PRIMARY KEY (`id`),
// 	KEY `owner_id` (`owner_id`) USING BTREE,
// 	KEY `city_id` (`city_id`) USING BTREE,
// 	KEY `status` (`status`) USING BTREE
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
