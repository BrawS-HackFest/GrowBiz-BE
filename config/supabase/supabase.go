package supabase

import (
	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"os"
)

func InitSupabase() *supabasestorageuploader.Client {
	supClient := supabasestorageuploader.New(
		os.Getenv("SUPA_URL"),
		os.Getenv("SUPA_KEY"),
		"picture",
	)
	return supClient
}
