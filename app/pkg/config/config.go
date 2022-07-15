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
	Host     string
	Port     string
	Username string
	Password string
}

type ReportSend struct {
	Subject string
	Body    string
}

type TestEnv struct {
	test string
}

func LoadTest() *TestEnv {
	return &TestEnv{
		test: GetEnv("TEST", "Hello-World-Default"),
	}
}

func LoadDB() *Database {
	return &Database{
		Username: GetEnv("USER", "Dev"),
		Password: GetEnv("PASSWORD", "P@ssw0rd2"),
		Host:     GetEnv("HOST", "172.18.0.2"),
		Port:     GetEnv("PORT", "3306"),
		Database: GetEnv("DATABASE", "project"),
	}
}

func LoadGmail() *Gmail {
	return &Gmail{
		Host:     "smtp.gmail.com",
		Port:     "465",
		Username: "rungmod.sit.kmutt@gmail.com",
		Password: "Project371@rungmod",
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

func GetEnv(key string, defaultVal string) string {
	if value, exits := os.LookupEnv(key); exits {
		return value
	}
	return defaultVal
}
