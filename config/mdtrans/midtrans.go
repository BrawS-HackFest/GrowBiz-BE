package mdtrans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"os"
	"strconv"
)

type MdtClient struct {
	c coreapi.Client
}

func NewMdtDriver() MdtClient {
	return MdtClient{c: coreapi.Client{}}
}

func (c *MdtClient) CreateOrder(id string, amount int, prodId uint, productName string, custEmail string,
	custName string, bank string) (*coreapi.ChargeResponse, error) {
	c.c.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  id,
			GrossAmt: int64(amount),
		},
		Items: &[]midtrans.ItemDetails{
			midtrans.ItemDetails{
				ID:    strconv.Itoa(int(prodId)),
				Name:  productName,
				Price: int64(amount),
				Qty:   1,
			},
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: custName,
			Email: custEmail,
		},
	}
	if bank == "bca" {
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{Bank: midtrans.BankBca}
	} else if bank == "bri" {
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{Bank: midtrans.BankBri}
	} else if bank == "bni" {
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{Bank: midtrans.BankBni}
	}
	coreApiRes, _ := c.c.ChargeTransaction(chargeReq)
	return coreApiRes, nil
}

func (c *MdtClient) NotifHandler(id string) (*coreapi.TransactionStatusResponse, error) {
	c.c.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	response, err := c.c.CheckTransaction(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}
