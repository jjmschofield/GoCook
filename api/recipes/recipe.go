package recipes

type Recipe struct {
	Id   string `json:"id" validate:"omitempty,uuid4"`
	Name string `json:"name" binding:"required" validate:"required,min=0,max=50"`
	Url  string `json:"url" validate:"omitempty,uri"`
}
