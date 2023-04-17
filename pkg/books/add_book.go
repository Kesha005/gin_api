package books

import(
	"net/http"
	"github.com/Kesha005/go_package/pkg/common/models"

)

type AddBookRequestBody struct{
	Title string `json:"title:`
	Author string `json:"author"`
	Description string `json:"description"`
}