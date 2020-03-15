package pr

import (
	"reflect"
	"testing"
)

func TestPropertiesSimple(t *testing.T) {
	text := "ABC"
	number := 0.5
	products := []Product{
		Product{
			Id:    "1",
			Sku:   "1",
			Name:  "A",
			Price: 42.0,
		},
		Product{
			Id:    "2",
			Sku:   "2",
			Name:  "B",
			Price: 100.0,
		},
	}

	tests := map[string]struct {
		ref Properties
		run func(Properties)
	}{
		"revenue":  {Properties{Revenue: number}, func(p Properties) { p.SetRevenue(number) }},
		"currency": {Properties{Currency: text}, func(p Properties) { p.SetCurrency(text) }},
		"value":    {Properties{Value: number}, func(p Properties) { p.SetValue(number) }},
		"path":     {Properties{Path: text}, func(p Properties) { p.SetPath(text) }},
		"referrer": {Properties{Referrer: text}, func(p Properties) { p.SetReferrer(text) }},
		"title":    {Properties{Title: text}, func(p Properties) { p.SetTitle(text) }},
		"url":      {Properties{Url: text}, func(p Properties) { p.SetURL(text) }},
		"name":     {Properties{Name: text}, func(p Properties) { p.SetName(text) }},
		"category": {Properties{Category: text}, func(p Properties) { p.SetCategory(text) }},
		"sku":      {Properties{Sku: text}, func(p Properties) { p.SetSKU(text) }},
		"price":    {Properties{Price: number}, func(p Properties) { p.SetPrice(number) }},
		"id":       {Properties{Id: text}, func(p Properties) { p.SetProductId(text) }},
		"orderId":  {Properties{OrderId: text}, func(p Properties) { p.SetOrderId(text) }},
		"total":    {Properties{Total: number}, func(p Properties) { p.SetTotal(number) }},
		"subtotal": {Properties{Subtotal: number}, func(p Properties) { p.SetSubtotal(number) }},
		"shipping": {Properties{Shipping: number}, func(p Properties) { p.SetShipping(number) }},
		"tax":      {Properties{Tax: number}, func(p Properties) { p.SetTax(number) }},
		"discount": {Properties{Discount: number}, func(p Properties) { p.SetDiscount(number) }},
		"coupon":   {Properties{Coupon: text}, func(p Properties) { p.SetCoupon(text) }},
		"products": {Properties{Products: pointyProductSlice(products)}, func(p Properties) { p.SetProducts(products...) }},
		"repeat":   {Properties{Repeat: true}, func(p Properties) { p.SetRepeat(true) }},
	}

	for name, test := range tests {
		prop := NewProperties()
		test.run(prop)

		if !reflect.DeepEqual(prop, test.ref) {
			t.Errorf("%s: invalid properties produced: %#v\n", name, prop)
		}
	}
}

func TestPropertiesMulti(t *testing.T) {
	p0 := Properties{Title: "A", Value: 0.5}
	p1 := NewProperties().SetTitle("A").SetValue(0.5)

	if !reflect.DeepEqual(p0, p1) {
		t.Errorf("invalid properties produced by chained setters:\n- expected %#v\n- found: %#v", p0, p1)
	}

}
