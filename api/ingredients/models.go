package ingredients

type IngredientType string

const (
	INGREDIENT_TYPE_VEG   IngredientType = "VEG"
	INGREDIENT_TYPE_FRUIT IngredientType = "FRUIT"
	INGREDIENT_TYPE_MEAT  IngredientType = "MEAT"
	INGREDIENT_TYPE_FISH  IngredientType = "FISH"
	INGREDIENT_TYPE_DAIRY IngredientType = "DAIRY"
	INGREDIENT_TYPE_HERB  IngredientType = "HERB"
	INGREDIENT_TYPE_SPICE IngredientType = "SPICE"
)

// swagger:model
type Ingredient struct {
	Id          string         `json:"id" validate:"omitempty,uuid4"`
	Name        string         `json:"name" binding:"required" validate:"required,min=0,max=50"`
	Type        IngredientType `json:"type" binding:"required" validate:"required,min=0,max=10"`
	DefaultUnit string         `json:"defaultUnit" validate:"omitempty,oneof=g mil cup tea table pinch"`
}
