package models

// type AutoGenerated struct {
// 	ID         string `json:"_id"`
// 	Rev        string `json:"_rev"`
// 	Attributes []struct {
// 		Key struct {
// 			Name string `json:"name"`
// 			ID   int    `json:"id"`
// 		} `json:"key"`
// 		Value struct {
// 			Name string `json:"name"`
// 			ID   int    `json:"id"`
// 		} `json:"value"`
// 		Starred bool `json:"starred"`
// 	} `json:"attributes"`
// 	Brand    string `json:"brand"`
// 	Category struct {
// 		L1   string `json:"L1"`
// 		L2   string `json:"L2"`
// 		L3   string `json:"L3"`
// 		Slug string `json:"Slug"`
// 	} `json:"category"`
// 	Cinsiyet          string `json:"cinsiyet"`
// 	Color             string `json:"color"`
// 	ColorVariantCount int    `json:"color_variant_count"`
// 	ColorVariants     []struct {
// 		Attributes []struct {
// 			Key struct {
// 				Name string `json:"name"`
// 				ID   int    `json:"id"`
// 			} `json:"key"`
// 			Value struct {
// 				Name string `json:"name"`
// 				ID   int    `json:"id"`
// 			} `json:"value"`
// 			Starred bool `json:"starred"`
// 		} `json:"attributes"`
// 		Brand        string `json:"brand"`
// 		Cinsiyet     string `json:"cinsiyet"`
// 		Color        string `json:"color"`
// 		Description  string `json:"description"`
// 		Descriptions []struct {
// 			Description string `json:"description"`
// 			Bold        bool   `json:"bold"`
// 		} `json:"descriptions"`
// 		Images              []string `json:"images"`
// 		Name                string   `json:"name"`
// 		NameWithProductCode string   `json:"name_with_product_code"`
// 		Price               struct {
// 			DiscountedPrice struct {
// 				Text  string  `json:"text"`
// 				Value float64 `json:"value"`
// 			} `json:"discountedPrice"`
// 			OriginalPrice struct {
// 				Text  string  `json:"text"`
// 				Value float64 `json:"value"`
// 			} `json:"originalPrice"`
// 		} `json:"price"`
// 		ProductCode      string `json:"product_code"`
// 		ProductGroupID   string `json:"product_group_id"`
// 		ProductNumber    string `json:"product_number"`
// 		ShortDescription string `json:"short_description"`
// 		SizeVariants     [][]struct {
// 			Price float64 `json:"Price"`
// 			Size  string  `json:"Size"`
// 		} `json:"size_variants"`
// 		Sku    string      `json:"sku"`
// 		Stock  interface{} `json:"stock"`
// 		URLKey string      `json:"url_key"`
// 	} `json:"color_variants"`
// 	Description  string `json:"description"`
// 	Descriptions []struct {
// 		Description string `json:"description"`
// 		Bold        bool   `json:"bold"`
// 	} `json:"descriptions"`
// 	Images              []string `json:"images"`
// 	Name                string   `json:"name"`
// 	NameWithProductCode string   `json:"name_with_product_code"`
// 	Price               struct {
// 		DiscountedPrice struct {
// 			Text  string  `json:"text"`
// 			Value float64 `json:"value"`
// 		} `json:"discountedPrice"`
// 		OriginalPrice struct {
// 			Text  string  `json:"text"`
// 			Value float64 `json:"value"`
// 		} `json:"originalPrice"`
// 	} `json:"price"`
// 	ProductCode      string `json:"product_code"`
// 	ProductGroupID   string `json:"product_group_id"`
// 	ProductNumber    string `json:"product_number"`
// 	ShortDescription string `json:"short_description"`
// 	SizeVariants     [][]struct {
// 		Price float64 `json:"Price"`
// 		Size  string  `json:"Size"`
// 	} `json:"size_variants"`
// 	Sku    string      `json:"sku"`
// 	Stock  interface{} `json:"stock"`
// 	URLKey string      `json:"url_key"`
// }