package routers

import (
	"github.com/MichaelYcCho/go-practice/pratice/follow_practice/controllors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/", controllors.UserRegistration)
}

func InitializeRouter() {
	r := mux.NewRouter()

	// gorilla mux 를 활용한 api 예시
	// r.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(handlers.AllowedHeaders([]string{
			"X-Requested-With",
			"Content-Type",
			"Authorization"}),
			handlers.AllowedMethods([]string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"HEAD",
				"OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))

}
