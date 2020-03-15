package pr

import "github.com/fuziontech/pr-idl/pb"

// This type is used to represent properties in messages that support it.
// It is a free-form object so the application can set any value it sees fit but
// a few helper method are defined to make it easier to instantiate properties with
// common fields.
// Here's a quick example of how this type is meant to be used:
//
//	analytics.Page{
//		UserId: "0123456789",
//		Properties: analytics.NewProperties()
//			.SetRevenue(10.0)
//			.SetCurrency("USD"),
//	}
//
type Properties pb.Properties

func NewProperties() Properties {
	return Properties{}
}

func (p Properties) SetRevenue(revenue float64) Properties {
	p.Revenue = revenue
	return p
}

func (p Properties) SetCurrency(currency string) Properties {
	p.Currency = currency
	return p
}

func (p Properties) SetValue(value float64) Properties {
	p.Value = value
	return p
}

func (p Properties) SetPath(path string) Properties {
	p.Path = path
	return p
}

func (p Properties) SetReferrer(referrer string) Properties {
	p.Referrer = referrer
	return p
}

func (p Properties) SetTitle(title string) Properties {
	p.Title = title
	return p
}

func (p Properties) SetURL(url string) Properties {
	p.Url = url
	return p
}

func (p Properties) SetName(name string) Properties {
	p.Name = name
	return p
}

func (p Properties) SetCategory(category string) Properties {
	p.Category = category
	return p
}

func (p Properties) SetSKU(sku string) Properties {
	p.Sku = sku
	return p
}

func (p Properties) SetPrice(price float64) Properties {
	p.Price = price
	return p
}

func (p Properties) SetProductId(id string) Properties {
	p.ProductId = id
	return p
}

func (p Properties) SetOrderId(id string) Properties {
	p.OrderId = p.OrderId
	return p
}

func (p Properties) SetTotal(total float64) Properties {
	p.Total = total
	return p
}

func (p Properties) SetSubtotal(subtotal float64) Properties {
	p.Subtotal = subtotal
	return p
}

func (p Properties) SetShipping(shipping float64) Properties {
	p.Shipping = shipping
	return p
}

func (p Properties) SetTax(tax float64) Properties {
	p.Tax = tax
	return p
}

func (p Properties) SetDiscount(discount float64) Properties {
	p.Discount = discount
	return p
}

func (p Properties) SetCoupon(coupon string) Properties {
	p.Coupon = coupon
	return p
}

func (p Properties) SetProducts(products ...Product) Properties {
	pointyProducts := pointyProductSlice(products)
	p.Products = pointyProducts
	return p
}

func (p Properties) SetRepeat(repeat bool) Properties {
	p.Repeat = repeat
	return p
}

func (p Properties) Set(name string, value string) Properties {
	p.ExtraFields[name] = value
	return p
}

func (p Properties) ToPB() *pb.Properties {
	out := pb.Properties(p)
	return &out
}

type Product pb.Product

func (p Product) ToPB() *pb.Product {
	out := pb.Product(p)
	return &out
}

func pointyProductSlice(in []Product) []*pb.Product {
	var out []*pb.Product
	for _, v := range in {
		pbp := pb.Product(v)
		out = append(out, &pbp)
	}
	return out
}
