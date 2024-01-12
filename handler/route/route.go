package route

import (
	"HackFest/config/mdtrans"
	"HackFest/database"
	"HackFest/handler"
	"HackFest/middleware"
	"HackFest/repository"
	"HackFest/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initHandler(db *gorm.DB) (*handler.UserHandler, *handler.CourseHandler, *handler.TransactionHandler) {
	mdt := mdtrans.NewMdtDriver()
	ctr := repository.NewCategoryRepository(db)

	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur, ctr)
	uh := handler.NewUserHandler(us)

	cr := repository.NewCourseRepository(db)
	cs := service.NewCourseService(cr)
	ch := handler.NewCourseHandler(cs, us)

	cur := repository.NewCourseUserRepository(db)

	tr := repository.NewTransactionRepository(db)
	ts := service.NewTransactionService(tr, ur, cr, cur, &mdt)
	th := handler.NewTransactionHandler(ts)

	return uh, ch, th
}

func Route(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Content-Type", "application/json")
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	})

	userHandler, courseHandler, transactionHandler := initHandler(database.InitDb())

	user := r.Group("/users")
	course := r.Group("/courses")
	transaction := r.Group("/transactions")

	user.GET("/profile", middleware.AuthMiddleware(), userHandler.GetProfile)
	user.PATCH("/update-profile", userHandler.UpdateUser)
	user.POST("/register", userHandler.CreateUser)

	course.GET("/all", courseHandler.FindAll)
	course.GET("/:id", courseHandler.FindByID)
	course.POST("/create", courseHandler.Create)

	transaction.GET("/transactions-by-user", middleware.AuthMiddleware(), transactionHandler.FindByUserID)
	transaction.GET("/:id", transactionHandler.FindByID)
	transaction.POST("/charge", middleware.AuthMiddleware(), transactionHandler.Create) //auth
	transaction.POST("/update", transactionHandler.Update)
}
