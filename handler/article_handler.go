package handler

import (
	"HackFest/models"
	"HackFest/service"
	"HackFest/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ArticleHandler struct {
	articleService service.ArticleService
}

func NewArticleHandler(article service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		articleService: article,
	}
}

func (ah *ArticleHandler) Create(c *gin.Context) {
	a1 := models.Article{
		Model: gorm.Model{},
		Title: "Memajukan UMKM Lokal: Pendampingan dalam  menginovasi Memperluas Produk",
		Pict:  "https://kpkmxnicmhvpqmxywspm.supabase.co/storage/v1/object/public/picture/art1.png?t=2024-01-17T05%3A45%3A04.415Z",
		Description: "Pemilik UMKM Keripik Tahu Barokah  bernama Bapak Mustofa. Usaha ini di jalankan oleh keluarga sendiri tanpa ada bantuan/pekerja lain, di kerjakan sendiri oleh bapak, ibu dan anak. Sebelum usaha keripik tahu ini di jalankan, " +
			"keluarga Bapak Mustofa hanya menjual atau membuat tahu pong yaitu tahu putih yang di goreng saja tanpa di olah lagi. Tahu pong ini biasa di gunakan oleh kosumen atau pelanggan untuk membuat tahu bakso, tahu isi dan lain-lain. " +
			"Setelah adanya usaha tahu pong,pada tahun  2019 Bapak Mustofa dan keluarga mempunyai ide untuk membuka usaha lain yaitu dengan membuat produk Keripik Tahu. \n\n" +
			"Bedasarkan wawancara dengan bapak Mustofa selaku pemilik UMKM Tahu Barokah beliau merasa kesulitan dalam memasarkan dan memperluas pasar dari produknya. \"Kami hanya bisa memasarkan berdasarkan pesanan saja untuk kita memperluas pemasarannya kami masih merasa kesulitan\". Ujar mustofa.",
		Comments: nil,
	}

	a2 := models.Article{
		Model: gorm.Model{},
		Title: "Tips & Trick Membuat Kopi Sederhana",
		Pict:  "https://kpkmxnicmhvpqmxywspm.supabase.co/storage/v1/object/public/picture/art4.png",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. " +
			"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat " +
			"cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		Comments: nil,
	}

	a3 := models.Article{
		Model: gorm.Model{},
		Title: "Latte Art adalah Seni",
		Pict:  "https://kpkmxnicmhvpqmxywspm.supabase.co/storage/v1/object/public/picture/art3.png",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. " +
			"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. " +
			"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		Comments: nil,
	}

	a4 := models.Article{
		Model: gorm.Model{},
		Title: "Bisnis Kopi Menjanjikan Kenangan",
		Pict:  "https://kpkmxnicmhvpqmxywspm.supabase.co/storage/v1/object/public/picture/art2.png",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. " +
			"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. " +
			"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		Comments: nil,
	}

	if err := ah.articleService.Create(a1); err != nil {
		utils.HttpInternalError(c, "Can't create article", err)
		return
	}
	if err := ah.articleService.Create(a2); err != nil {
		utils.HttpInternalError(c, "Can't create article", err)
		return
	}
	if err := ah.articleService.Create(a3); err != nil {
		utils.HttpInternalError(c, "Can't create article", err)
		return
	}
	if err := ah.articleService.Create(a4); err != nil {
		utils.HttpInternalError(c, "Can't create article", err)
		return
	}
}

func (ah *ArticleHandler) FindAll(c *gin.Context) {

	articles, err := ah.articleService.FindAll()
	if err != nil {
		utils.HttpInternalError(c, "Can't get article", err)
		return
	}
	utils.HttpSuccess(c, "Success get articles", articles)
}

func (ah *ArticleHandler) FindByID(c *gin.Context) {

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	result, err := ah.articleService.FindByID(uint(id))
	if err != nil {
		utils.HttpInternalError(c, "Can't get article", err)
		return
	}
	utils.HttpSuccess(c, "Success get article", result)
}