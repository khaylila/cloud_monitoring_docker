package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Article struct {
	ID              uint      `gorm:"primaryKey;column:id"`
	Title           string    `gorm:"size:255;not null;column:title"`
	Content         string    `gorm:"type:text;not null;column:content"`
	AuthorName      string    `gorm:"size:100;column:author_name"`
	Category        string    `gorm:"size:50;column:category"`
	PublicationDate time.Time `gorm:"column:publication_date;default:CURRENT_TIMESTAMP"`
	Status          string    `gorm:"size:50;default:draft;not null;column:status"`
	Slug            string    `gorm:"size:255;unique;not null;column:slug"`
}

func main() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	app := fiber.New()

	// dsn := "host=localhost user=admin password=admin dbname=admin port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	dsn := "host=postgres_db user=admin password=admin dbname=admin port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database!")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Error getting *sql.DB object:", err)
	}

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxIdleTime(1 * time.Minute)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	fmt.Println("Connected to database with pooling settings applied...")

	// Ambil jumlah koneksi dalam pool
	stats := sqlDB.Stats()

	fmt.Printf("Max Open Connections: %d\n", stats.MaxOpenConnections)
	fmt.Printf("Open Connections: %d\n", stats.OpenConnections)
	fmt.Printf("In Use Connections: %d\n", stats.InUse)
	fmt.Printf("Idle Connections: %d\n", stats.Idle)

	fmt.Println("Connecting to database...")

	// Anda bisa mulai menggunakan variabel 'db' di sini untuk query, migrasi, dll.
	// Contoh: db.AutoMigrate(&YourModel{})

	if len(os.Args) > 1 && os.Args[1] == "--migrate" {
		db.AutoMigrate(&Article{})
		// Cek apakah tabel sudah berisi data
		// Inisialisasi Migration
		if err := db.Debug().Exec(`
			INSERT INTO articles (title, content, author_name, category, status, slug) VALUES
			(
				'Pasar Saham Global Menguat di Tengah Ketidakpastian',
				'Indeks saham di seluruh dunia menunjukkan tren positif meskipun ada kekhawatiran tentang inflasi. Investor tampak optimis dengan laporan pendapatan perusahaan kuartal ini...',
				'Rina Amelia',
				'Ekonomi',
				'published',
				'pasar-saham-global-menguat-di-tengah-ketidakpastian'
			),
			(
				'Mengenal Manfaat Meditasi untuk Kesehatan Mental',
				'Di tengah kesibukan sehari-hari, kesehatan mental sering terabaikan. Meditasi adalah salah satu cara efektif untuk mengurangi stres dan meningkatkan fokus. Artikel ini membahas teknik dasar meditasi bagi pemula.',
				'Dr. Haryanto',
				'Kesehatan',
				'published',
				'mengenal-manfaat-meditasi-untuk-kesehatan-mental'
			),
			(
				'Review Film "Pengejar Senja": Sebuah Karya Sinematik yang Memukau',
				'Film terbaru dari sutradara ternama Bima Nugraha, "Pengejar Senja", berhasil memukau penonton dengan alur cerita yang kuat dan sinematografi yang indah. Film ini diprediksi akan meraih banyak penghargaan.',
				'Putri Anindita',
				'Hiburan',
				'published',
				'review-film-pengejar-senja'
			),
			(
				'Lima Destinasi Wisata Tersembunyi di Indonesia',
				'Selain Bali dan Lombok, Indonesia memiliki banyak surga tersembunyi yang menanti untuk dijelajahi. Dari pantai eksotis hingga pegunungan yang megah, berikut lima destinasi yang wajib Anda kunjungi.',
				'Ahmad Fauzi',
				'Wisata',
				'published',
				'lima-destinasi-wisata-tersembunyi-di-indonesia'
			),
			(
				'Perkembangan AI Diprediksi Akan Mengubah Dunia Kerja',
				'Kecerdasan buatan (AI) terus berkembang pesat. Para ahli memprediksi bahwa dalam satu dekade ke depan, AI akan mengambil alih banyak pekerjaan rutin, namun juga menciptakan profesi baru.',
				'Andi Wijaya',
				'Teknologi',
				'published',
				'perkembangan-ai-diprediksi-akan-mengubah-dunia-kerja'
			),
			(
				'Atlet Bulu Tangkis Indonesia Juarai All England',
				'Pasangan ganda putra Indonesia berhasil merebut gelar juara di turnamen All England setelah mengalahkan pasangan dari Tiongkok dalam pertandingan final yang sengit.',
				'Budi Santoso',
				'Olahraga',
				'published',
				'atlet-bulu-tangkis-indonesia-juarai-all-england'
			),
			(
				'Uji Coba Mobil Listrik Terbaru: Efisien dan Ramah Lingkungan',
				'Produsen otomotif terkemuka meluncurkan mobil listrik generasi terbaru dengan jarak tempuh lebih jauh dan waktu pengisian daya lebih singkat. Ini adalah langkah besar menuju transportasi berkelanjutan.',
				'Rian Pratama',
				'Otomotif',
				'draft',
				'uji-coba-mobil-listrik-terbaru-efisien'
			),
			(
				'Pentingnya Pendidikan Finansial Sejak Dini',
				'Mengajarkan anak-anak tentang manajemen keuangan sejak dini adalah investasi untuk masa depan mereka. Artikel ini membahas cara-cara sederhana untuk mengenalkan konsep uang kepada anak.',
				'Citra Lestari',
				'Pendidikan',
				'published',
				'pentingnya-pendidikan-finansial-sejak-dini'
			),
			(
				'Penemuan Spesies Baru di Laut Dalam Amazon',
				'Tim peneliti internasional mengumumkan penemuan lebih dari selusin spesies baru di perairan dalam Sungai Amazon. Penemuan ini menyoroti betapa kayanya keanekaragaman hayati yang belum terungkap.',
				'Dr. Silvia',
				'Sains',
				'published',
				'penemuan-spesies-baru-di-laut-dalam-amazon'
			),
			(
				'Tren Fashion Musim Gugur 2025: Warna Bumi dan Gaya Minimalis',
				'Para desainer meramalkan bahwa tren fashion untuk musim gugur mendatang akan didominasi oleh palet warna bumi seperti cokelat, terakota, dan hijau zaitun, dengan siluet yang simpel dan minimalis.',
				'Laura Basuki',
				'Gaya Hidup',
				'published',
				'tren-fashion-musim-gugur-2025'
			);
		`).Error; err != nil {
			log.Println("Error seeding initial data:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Database migrated successfully.")
		os.Exit(0) // Keluar setelah migration
	}

	app.Get("/", func(c *fiber.Ctx) error {

		rand := time.Now().UnixNano() % 10

		var article Article

		if err := db.Select("*").Where("id = ?", rand+1).First(&article).Error; err != nil {
			log.Println("Error fetching article:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":      "Error",
				"status_code": fiber.StatusInternalServerError,
				"message":     "Failed to fetch article",
			})
		}
		// return c.SendStatus(fiber.StatusOK)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":      "OK",
			"status_code": fiber.StatusOK,
			"data":        article,
		})
	})

	app.Listen(":8080")
}
