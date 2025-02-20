package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var id = 0
type Contact struct {
	Name string
	Email string
	Id int
}
func newContact(name, email string) Contact {
	id++
	return Contact{Name: name, Email: email, Id: id}
}
type Contacts = []Contact

type Data struct {
	Contacts Contacts
}
func newData() Data {
	return Data{Contacts: Contacts{
		newContact("John", "jd@gmail.com"),
		newContact("Clara", "cd@gmail.com"),
	}}
}
func (d *Data) indexOf(id int) int {
	for i, contact := range d.Contacts {
		if contact.Id == id {
			return i
		}
	}
	return -1
}
func (d *Data) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}
func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form FormData
}
func newPage() Page {
	return Page{Data: newData(), Form: newFormData()}
}

func main() {
	page := newPage()
	router := gin.Default()

	router.LoadHTMLGlob("views/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", page) // Using index and not index.html because is identifying the block, not the file
	})

	router.POST("/contacts", func(ctx *gin.Context) {
		name := ctx.Request.FormValue("name")
		email := ctx.Request.FormValue("email")

		if page.Data.hasEmail(email) {
			formData := newFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Errors["email"] = "Email already exists"

			ctx.HTML(http.StatusUnprocessableEntity, "form", formData)
		}

		contact := newContact(name, email)
		page.Data.Contacts = append(page.Data.Contacts, contact)

		ctx.HTML(http.StatusOK, "form", newFormData())
		ctx.HTML(http.StatusOK, "oob-contact", contact)
	})

	router.DELETE("/contacts/:id", func(ctx *gin.Context) {
		idString := ctx.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			ctx.String(http.StatusBadRequest, "Invalid Id")
		}

		index := page.Data.indexOf(id)
		if index == -1 {
			ctx.String(http.StatusNotFound, "Contact not found")
		}
		page.Data.Contacts = append(page.Data.Contacts[:index], page.Data.Contacts[index+1:]...)

		ctx.Status(http.StatusOK)
	})

	router.Run(":8080")
}