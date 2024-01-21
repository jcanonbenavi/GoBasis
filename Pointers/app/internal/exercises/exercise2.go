package exercises

type Product interface {
	GetPrice() (price float64)
}

const (
	SmallProduct  = "SmallProduct"
	MediumProduct = "MediumProduct"
	LargeProduct  = "LargeProduct"
)

func Factory(productType string, price float64) Product {
	switch productType {
	case SmallProduct:
		return NewSmall(price)
	case MediumProduct:
		return NewMedium(price)
	case LargeProduct:
		return NewLarge(price)
	}
	return nil
}

type Small struct {
	price float64
}

func NewSmall(price float64) *Small {
	return &Small{price: price}
}

func (s *Small) GetPrice() (price float64) {
	price = s.price
	return
}

type Medium struct {
	price float64
}

func NewMedium(price float64) *Medium {
	return &Medium{price: price}
}

func (m *Medium) GetPrice() (price float64) {
	price = (m.price * 0.3) + m.price
	return
}

type Large struct {
	price float64
}

func NewLarge(price float64) *Large {
	return &Large{price: price}
}

func (l *Large) GetPrice() (price float64) {
	price = (l.price * 0.6) + l.price
	return
}
