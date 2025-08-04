package usecase

import (
	"log"
	"os"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

func PaymentServiceByProductID(productID, currency string) (string, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	if currency == "" {
		currency = "usd"
	}

	product, err := GetProductByID(productID)
	if err != nil {
		return "", err
	}

	amountInCents := int(product.Price * 100)

	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(currency),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(product.Name),
					},
					UnitAmount: stripe.Int64(int64(amountInCents)),
				},
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String("http://localhost:5000/payment/success"),
		CancelURL:  stripe.String("http://localhost:5000/payment/cancel"),
	}

	s, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
		return "", err
	}

	return s.URL, nil
}
