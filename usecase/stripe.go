package usecase

import (
	"errors"
	"log"
	"os"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

func PaymentServiceByProductID(productID, currency, userID string) (string, error) {
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
		PaymentIntentData: &stripe.CheckoutSessionPaymentIntentDataParams{
			Metadata: map[string]string{
				"product_id": productID,
				"user_id":    userID,
			},
		},
		SuccessURL: stripe.String("http://localhost:5000/payment/success?session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String("http://localhost:5000/payment/cancel"),
	}

	s, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
		return "", err
	}

	return s.URL, nil
}

func GetStripeSessionDetails(sessionID string) (string, string, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	s, err := session.Get(sessionID, nil)
	if err != nil {
		log.Printf("session.Get: %v", err)
		return "", "", err
	}

	if s.PaymentStatus != stripe.CheckoutSessionPaymentStatusPaid {
		return "", "", errors.New("payment not completed")
	}

	// Get metadata from PaymentIntent
	if s.PaymentIntent == nil {
		return "", "", errors.New("no payment intent found in session")
	}

	pi, err := paymentintent.Get(s.PaymentIntent.ID, nil)
	if err != nil {
		log.Printf("paymentintent.Get: %v", err)
		return "", "", err
	}

	userID, exists := pi.Metadata["user_id"]
	if !exists {
		return "", "", errors.New("user ID not found in payment intent metadata")
	}

	productID, exists := pi.Metadata["product_id"]
	if !exists {
		return "", "", errors.New("product ID not found in payment intent metadata")
	}

	return userID, productID, nil
}
