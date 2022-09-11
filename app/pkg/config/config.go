package config

import "os"

type Database struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type Gmail struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Credential string
}

type ReportSend struct {
	Subject string
	Body    string
}

type TestEnv struct {
	test string
}

type Jwt struct {
	Secret string
}

func LoadTest() *TestEnv {
	return &TestEnv{
		test: GetEnv("TEST", "Hello-World-Default"),
	}
}

func LoadDB() *Database {
	return &Database{
		Username: GetEnv("USER", "localhost"),
		Password: GetEnv("PASSWORD", "P@ssw0rd2"),
		Host:     GetEnv("HOST", "database-rungmod.sit.kmutt.ac.th"),
		Port:     GetEnv("PORT", "55013"),
		// Host:     GetEnv("HOST", "10.4.56.39"),
		// Port:     GetEnv("PORT", "3306"),
		Database: GetEnv("DATABASE", "project"),
	}
}

func LoadGmail() *Gmail {
	return &Gmail{
		Host:       GetEnv("MAILER_HOST", "smtp.gmail.com"),
		Port:       GetEnv("MAILER_PORT", "465"),
		Username:   GetEnv("MAILER_USERNAME", "rungmod.sit.kmutt@gmail.com"),
		Password:   GetEnv("MAILER_PASSWORD", "Project371@rungmod"),
		Credential: GetEnv("CREDENTIAL", "AIzaSyBvlioW5xWd9dl9w9ynxAOSYBLqJXc-AUU"),
	}
}

func LoadReportSend() *ReportSend {
	return &ReportSend{
		Subject: "Received",
		Body:    "รับเรื่องเรียบร้อยทางเรากำลังดำเนินการ",
	}
}

func LoadRegisCustomerSend() *ReportSend {
	return &ReportSend{
		Subject: "ขอบคุณสำหรับการสมัรสมาชิก กรุณายืนยันตัวตนด้านล่าง",
		Body:    "",
	}
}

func LoadJWTConfig() *Jwt {
	return &Jwt{
		Secret: GetEnv("SECRET", "abcdefghijkmn"),
	}
}

func GetEnv(key string, defaultVal string) string {
	if value, exits := os.LookupEnv(key); exits {
		return value
	}
	return defaultVal
}
