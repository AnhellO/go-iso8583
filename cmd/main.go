package main

import (
	"fmt"
	"os"

	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/padding"
	"github.com/moov-io/iso8583/prefix"
)

func main() {
	spec := &iso8583.MessageSpec{
		Fields: map[int]field.Field{
			0: field.NewString(&field.Spec{
				Length:      4,
				Description: "Message Type Indicator",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
			1: field.NewBitmap(&field.Spec{
				Description: "Bitmap",
				Enc:         encoding.BytesToASCIIHex,
				Pref:        prefix.Hex.Fixed,
			}),

			// Message fields:
			2: field.NewString(&field.Spec{
				Length:      19,
				Description: "Primary Account Number",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LL,
			}),
			3: field.NewNumeric(&field.Spec{
				Length:      6,
				Description: "Processing Code",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			4: field.NewString(&field.Spec{
				Length:      12,
				Description: "Transaction Amount",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
		},
	}

	// create message with defined spec
	message := iso8583.NewMessage(spec)

	// set message type indicator at field 0
	message.MTI("0100")

	// set all message fields you need as strings
	err := message.Field(2, "4242424242424242")
	if err != nil {
		panic(err)
	}

	err = message.Field(3, "123456")
	if err != nil {
		panic(err)
	}

	err = message.Field(4, "100")
	if err != nil {
		panic(err)
	}

	// generate binary representation of the message into rawMessage
	rawMessage, err := message.Pack()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Raw binary message: %+v", rawMessage)

	err = iso8583.Describe(message, os.Stdout)
	if err != nil {
		panic(err)
	}
}
