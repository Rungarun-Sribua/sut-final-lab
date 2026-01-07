package schemas

import (
	"github.com/asaskevich/govalidator"
)

type Book struct {
	Title string  `valid:"required~Title must in between 3 to 100 characters,range(3|100)~Title must in between 3 to 100 characters"`
	Price float64 `valid:"required~Price must be between 50 and 5000,range(50|5000)~Price must be between 50 and 5000"`
	Code  string  `valid:"matches(^(["BK"]\\[0-9]{6})$)~Code must start with BK followed by 6 digits (0-9)"`
}

func (b *Book) Validate() error {
	_, err := govalidator.ValidateStruct(b)
	return err
}
