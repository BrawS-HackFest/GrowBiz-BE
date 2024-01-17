package handler

import (
	"HackFest/service"
)

type ArticleHandler struct {
	articleService service.ArticleService
}

func NewArticleHandler(article service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		articleService: article,
	}
}

//func (ah *ArticleHandler) Create(c *gin.Context) {
//	a1 := models.Article{
//		Model: gorm.Model{},
//		Title: "Memajukan UMKM Lokal: Pendampingan dalam  menginovasi Memperluas Produk",
//		Pict:  "https://kpkmxnicmhvpqmxywspm.supabase.co/storage/v1/object/public/picture/art1.png?t=2024-01-17T05%3A45%3A04.415Z",
//		Description: "Pemilik UMKM Keripik Tahu Barokah  bernama Bapak Mustofa. Usaha ini di jalankan oleh keluarga sendiri tanpa ada bantuan/pekerja lain, di kerjakan sendiri oleh bapak, ibu dan anak. Sebelum usaha keripik tahu ini di jalankan, " +
//			"keluarga Bapak Mustofa hanya menjual atau membuat tahu pong yaitu tahu putih yang di goreng saja tanpa di olah lagi. Tahu pong ini biasa di gunakan oleh kosumen atau pelanggan untuk membuat tahu bakso, tahu isi dan lain-lain. " +
//			"Setelah adanya usaha tahu pong,pada tahun  2019 Bapak Mustofa dan keluarga mempunyai ide untuk membuka usaha lain yaitu dengan membuat produk Keripik Tahu. \n\n" +
//			"Bedasarkan wawancara dengan bapak Mustofa selaku pemilik UMKM Tahu Barokah beliau merasa kesulitan dalam memasarkan dan memperluas pasar dari produknya. \"Kami hanya bisa memasarkan berdasarkan pesanan saja untuk kita memperluas pemasarannya kami masih merasa kesulitan\". Ujar mustofa.",
//		Comments: nil,
//	}
//}
