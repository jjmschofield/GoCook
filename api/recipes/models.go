package recipes

type Recipe struct {
	Id           string       `json:"id" validate:"omitempty,uuid4"`
	Name         string       `json:"name" binding:"required" validate:"required,min=3,max=50"`
	Description  string       `json:"description" validate:"omitempty,min=0,max=2000"`
	Time         Time         `json:"time"`
	Yield        int          `json:"yield" validate:"omitempty,min=0"`
	Steps        []Step       `json:"steps" binding:"required" validate:"required,dive"`
	Ingredients  []Ingredient `json:"ingredients" binding:"required" validate:"required,dive"`
	Tags         []string     `json:"tags" binding:"required" validate:"required,dive,min=1,max=15"`
	Url          string       `json:"url" validate:"omitempty,uri"`
	ImgUrl       string       `json:"imgUrl" validate:"omitempty,uri"`
	Owner        string       `json:"owner"`
	Contributors []string	  `json:"contributors" binding:"required" validate:"required,dive,uuid4"`
}

type Ingredient struct {
	Name string  `json:"name" binding:"required" validate:"required,min=0,max=50"`
	Qty  float32 `json:"qty" validate:"omitempty,min=0,max=10000"`
	Unit string  `json:"unit" validate:"omitempty,oneof=g mil cup tea table pinch"`
}

type Step struct {
	Name        string `json:"name" binding:"required" validate:"required,min=0,max=50"`
	Description string `json:"description" validate:"omitempty,min=0,max=2000"`
	Time        Time   `json:"time"`
}

type Time struct {
	PrepSec int `json:"prepSec" validate:"omitempty,min=0"`
	CookSec int `json:"cookSec" validate:"omitempty,min=0"`
}
