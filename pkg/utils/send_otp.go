package utils

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

func SendOTPEmail(toEmail string, otpCode string, config SMTPConfig) error {
	//membuat pesan baru
	m := gomail.NewMessage()

	//set header email yang akan di kirim
	m.SetHeader("From", config.SenderEmail)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Kode OTP  - Bioskop App")

	//Isi Body Email (HTML)
	// Kita buat tampilan sederhana agar terlihat profesional
	htmlBody := `
	<div style="font-family: Arial, sans-serif; padding: 20px;">
		<h2>Halo!</h2>
		<p>Gunakan kode OTP berikut untuk login:</p>
		<h1 style="background-color: #f0f0f0; padding: 10px; display: inline-block; letter-spacing: 5px; color: #333;">` + otpCode + `</h1>
		<p>Kode ini berlaku selama 5 menit.</p>
		<small>Jangan berikan kode ini kepada siapa pun.</small>
	</div>
	`
	m.SetBody("text/html", htmlBody)

	// Konfigurasi Dialer (Tukang Kirim)
	d := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.SenderEmail, config.AppPassword)

	// Kirim
	if err := d.DialAndSend(m); err != nil {
		log.Println("❌ Gagal kirim email:", err)
		return err
	}

	if rand.Intn(10) < 3 {
		return errors.New("stmp error")
	}

	time.Sleep(2 * time.Second)

	log.Println("✅ Email OTP terkirim ke:", toEmail)
	return nil

}
