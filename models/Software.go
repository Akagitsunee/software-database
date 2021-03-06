package models

type Software struct {
	Id				string	 `json:"id" example:"1"`
	Name			string	 `json:"name" example:"Google Chrome" validate:"required"`
	Description		string	 `json:"description" example:"Web Browser"`
	Version			string	 `json:"version" example:"1.0" validate:"required"`
	License 		string	 `json:"license" example:"MIT"`
	Homepage		string	 `json:"homepage" example:"https://www.google.com/intl/de/chrome/"`
}
