package schemas

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestBookPriceValidate(t *testing.T) {
	tests := []struct {
		name   string
		book   Book
		expect string
	}{
		{"Valid_name", Book{Title: "ThePartyIsOver", Price: 400.00, Code: "B12356"}, "“Code must start with BK followed by 6 digits (0-9)"},
	}

	g := NewGomegaWithT(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {})
		err := tt.book.Validate()
		if tt.expect == "" {
			g.Expect(err).To(BeNil())
		} else {
			g.Expect(err.Error()).To(Equal(tt.expect))
		}
	}
}
