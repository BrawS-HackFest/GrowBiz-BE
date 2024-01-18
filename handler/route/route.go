package route

import (
	"HackFest/config/mdtrans"
	"HackFest/database"
	"HackFest/handler"
	"HackFest/middleware"
	"HackFest/repository"
	"HackFest/repository/article"
	"HackFest/repository/courses"
	"HackFest/service"
	article2 "HackFest/service/article"
	courses2 "HackFest/service/courses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initHandler(db *gorm.DB) (*handler.UserHandler, *handler.CourseHandler, *handler.TransactionHandler,
	*handler.ArticleHandler, *handler.CommentHandler) {
	mdt := mdtrans.NewMdtDriver()
	ctr := repository.NewCategoryRepository(db)

	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur, ctr)
	uh := handler.NewUserHandler(us)

	cur := courses.NewCourseUserRepository(db)
	cus := courses2.NewCourseUserService(cur)

	cr := courses.NewCourseRepository(db)
	cs := courses2.NewCourseService(cr)
	ch := handler.NewCourseHandler(cs, us, cus)

	tr := repository.NewTransactionRepository(db)
	ts := service.NewTransactionService(tr, ur, cr, cur, &mdt)
	th := handler.NewTransactionHandler(ts)

	ar := article.NewArticleRepository(db)
	as := article2.NewArticleService(ar, ur)
	ah := handler.NewArticleHandler(as)

	cmr := article.NewCommentRepository(db)
	cms := article2.NewCommentService(cmr)
	cmh := handler.NewCommentHandler(cms)

	return uh, ch, th, ah, cmh
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

	userHandler, courseHandler, transactionHandler,
		articleHandler, commentHandler := initHandler(database.InitDb())

	user := r.Group("/users")
	course := r.Group("/courses")
	transaction := r.Group("/transactions")
	articles := r.Group("/article")

	user.GET("/profile", middleware.AuthMiddleware(), userHandler.GetProfile)
	user.PATCH("/update-profile", userHandler.UpdateUser)
	user.POST("/register", userHandler.CreateUser)

	course.GET("/all", courseHandler.FindAll)
	course.GET("/:id", courseHandler.FindByID)
	course.GET("/courses-by-user", middleware.AuthMiddleware(), courseHandler.GetCoursesByID)
	course.POST("/create", courseHandler.Create)

	transaction.GET("/transactions-by-user", middleware.AuthMiddleware(), transactionHandler.FindByUserID)
	transaction.GET("/:id", transactionHandler.FindByID)
	transaction.POST("/charge", middleware.AuthMiddleware(), transactionHandler.Create) //auth
	transaction.POST("/update", transactionHandler.Update)

	articles.GET("/all", articleHandler.FindAll)
	articles.GET("/:id", articleHandler.FindByID)
	articles.GET("/comments-by-article/:id", commentHandler.FindByArticleID)
	articles.GET("/comments/:id", commentHandler.FindByID)
	articles.POST("/create", articleHandler.Create)
	articles.POST("/:articleID/create-comment", middleware.AuthMiddleware(), commentHandler.Create)
	articles.PATCH("/update-comment", middleware.AuthMiddleware(), commentHandler.Update)
	articles.DELETE("/delete-comment", middleware.AuthMiddleware(), commentHandler.Delete)

}
