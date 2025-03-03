package principles

import (
	"fmt"
	"time"
)

/*
- A class should have only one reason to change
- Any class or component in our code should only have one functionality
- Everything in the class should be related to just one goal.
*/

type Book struct {
	Title  string
	Author string
	Price  float32
	Year   time.Time
	ISBN   string
}

type InvoiceDetails struct {
	Book
	Quantity   int
	Discount   float32
	Tax        float32
	TotalPrice float32
}

type InvoicePrinter struct {
	InvoiceDetails
}

type InvoiceStorage struct {
	InvoiceDetails
}

func (b Book) ShowBookDetails() {
	fmt.Printf("Title: %s\nAuthor: %s\nPrice: %.2f\nYear: %d\nISBN: %s\n", b.Title, b.Author, b.Price, b.Year.Year(), b.ISBN)
}

func (invoceDetail InvoiceDetails) ShowInvoiceDetails() {
	fmt.Printf("Title: %s\nAuthor: %s\nPrice: %.2f\nYear: %d\nISBN: %s\nQuantity: %d\nDiscount: %.2f\nTax: %.2f\nTotal Price: %.2f\n", invoceDetail.Title, invoceDetail.Author, invoceDetail.Price, invoceDetail.Year.Year(), invoceDetail.ISBN, invoceDetail.Quantity, invoceDetail.Discount, invoceDetail.Tax, invoceDetail.TotalPrice)
}

func (invoiceStorage InvoiceStorage) SaveInvoiceToDB() {
	fmt.Println("Saving Invoice...")
	invoiceStorage.ShowInvoiceDetails()
}

func (invoicePrinter InvoicePrinter) PrintInvoice() {
	fmt.Println("Printing Invoice...")
	invoicePrinter.ShowInvoiceDetails()
}

func SRP() {
	fmt.Println("From SRP Package!!!")
	book := Book{
		Title:  "The Alchemist",
		Author: "Paul",
		Price:  10.50,
		Year:   time.Date(1988, time.April, 25, 0, 0, 0, 0, time.UTC),
		ISBN:   "978-0062315007",
	}
	book.ShowBookDetails()

	invoice := InvoiceDetails{
		Book:       book,
		Quantity:   2,
		Discount:   0.5,
		Tax:        0.5,
		TotalPrice: 20.50,
	}
	invoice.ShowInvoiceDetails()

	invoiceStorage := InvoiceStorage{
		InvoiceDetails: invoice,
	}
	invoiceStorage.SaveInvoiceToDB()

	invoicePrinter := InvoicePrinter{
		InvoiceDetails: invoice,
	}
	invoicePrinter.PrintInvoice()

}
