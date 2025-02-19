package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	Name string
	Email string
}
func newContact(name, email string) Contact {
	return Contact{Name: name, Email: email}
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

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", page) // Using index and not index.html because is identifying the block, not the file
	})

	router.POST("/contacts", func(c *gin.Context) {
		name := c.Request.FormValue("name")
		email := c.Request.FormValue("email")

		if page.Data.hasEmail(email) {
			formData := newFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Errors["email"] = "Email already exists"

			c.HTML(http.StatusBadRequest, "form", formData)
		}

		page.Data.Contacts = append(page.Data.Contacts, newContact(name, email))

		c.HTML(http.StatusOK, "display", page)
	})

	router.Run(":8080")
}