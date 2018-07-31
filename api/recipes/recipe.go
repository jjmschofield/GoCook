package recipes

import "net/url"

type Recipe struct{
	Id string
	Name string
	Url string
}

func (recipe *Recipe) IsValid() (bool, string) {

	isValidName, validationMessage := recipe.nameIsValid()
	if !isValidName {
		return false, validationMessage
	}

	isValidUrl, validationMessage := recipe.urlIsValid()
	if !isValidUrl {
		return false, validationMessage
	}

	return true, ""
}

func (recipe *Recipe) nameIsValid() (bool, string){
	if len(recipe.Name) < 3 {
		return false, "Name is not set or is too short - minimum 3 characters required"
	}

	if len(recipe.Name) > 50 {
		return false, "Name is too long - maximum 50 characters allowed"
	}

	return true, ""
}

func (recipe *Recipe) urlIsValid() (bool, string){
	if len(recipe.Url) > 0 {
		_, urlParseError := url.ParseRequestURI(recipe.Url)

		if urlParseError != nil {
			return false, "Url is not valid"
		}
	}

	return true, ""
}