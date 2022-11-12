package config

import "os"

type Database struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type Mailer struct {
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

type Jwt struct {
	Secret string
}

type API_KEY_MAILER struct {
	API string
}

type PART_MEDIA struct {
	Path string
}

func LoadTest() *TestEnv {
	return &TestEnv{
		test: GetEnv("TEST", "Hello-World-Default"),
	}
}

func LoadDB() *Database {
	return &Database{
		Username: GetEnv("USER", "localhostDev"),
		Password: GetEnv("PASSWORD", "P@ssw0rd2"),
		// Host:     GetEnv("HOST", "database-rungmod.sit.kmutt.ac.th"),
		// Port:     GetEnv("PORT", "55013"),
		Host: GetEnv("HOST", "10.4.56.39"),
		Port: GetEnv("PORT", "3306"),
		Database: GetEnv("DATABASE", "rungmodDev"),
		// Database: GetEnv("DATABASE", "rungmodDev"),
	}
}

func LoadMailerStruct() *Mailer {
	return &Mailer{
		Host:     GetEnv("MAILER_HOST", "smtp.mailersend.net"),
		Port:     GetEnv("MAILER_PORT", "587"),
		Username: GetEnv("MAILER_USERNAME", "MS_Vw9hXV@rungmod.com"),
		Password: GetEnv("MAILER_PASSWORD", "KmeWcjP2XgaGp3a7"),
	}
}

func LoadMailer() *API_KEY_MAILER {
	return &API_KEY_MAILER{
		API: GetEnv("MAILER_KEY", "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiZmZhZTY0MWY4ZTQ2MjU4OGNjZWY4ZTk0ZjZmOTcwMGVmMTBiOTc3ZTM0N2VkZDY5NjJmM2E0ZWYzNjVhODMyMzM1YmUyNDA4MzQyMjUwZDciLCJpYXQiOjE2NjYwMDM3NDcuNDk0MjI1LCJuYmYiOjE2NjYwMDM3NDcuNDk0MjI4LCJleHAiOjQ4MjE2NzczNDcuNDg3MzI4LCJzdWIiOiI0MjA3NSIsInNjb3BlcyI6WyJlbWFpbF9mdWxsIiwiZG9tYWluc19mdWxsIiwiYWN0aXZpdHlfZnVsbCIsImFuYWx5dGljc19mdWxsIiwidG9rZW5zX2Z1bGwiLCJ3ZWJob29rc19mdWxsIiwidGVtcGxhdGVzX2Z1bGwiLCJzdXBwcmVzc2lvbnNfZnVsbCIsInNtc19mdWxsIiwiZW1haWxfdmVyaWZpY2F0aW9uX2Z1bGwiXX0.NC-VU_4XzOFg-9I5n6OxBRm3ozoL0ZV9eiOp1KewymEY-JYWpiGqcciivSxz0PnHMS7m5-g2B3wLVMTz1RPDThx1ZUxOUIMjxMMD1kpfzidfDmyzUkUKMHJ_l_47KQOsCNcjP_IrtL1lIt3zjltXESSCxducPtV1AylLmfQnLYLtHOlKexaE2PojGY3kZNnUecc9CtMqx1NxBDNG3z9VEUpNkfhq7t_bq3IbZK8PtPWS26-B6m6weswiVx4Rj9ZefgTIPKR4-N1gf_VgNUwIm_mlW-OlL4c6gGp4zX3PsoJSW50TQmKsaH4v1Mfh7Vor2V9a69W7_mhdCRtEcJFeMz591NzA7iNx-bd3KdPOyxNwK5H5pyeyOZek_CBRtA7quRhmZ7DVmio2z89F_iO3ON7MxPLxkCUp2wmgQy0hUtABEKZAy7CCz5t7JQ3wvoeOQc6YxmmYngGYithdVlpO6Aqeax9Jtw-EXRhmypNEW5-d8K_ThZOhQfsEBCt3WJf8CQySDOe8eCeJOi07YaC138W7QrM6s8NYSghh3KMxLzdjPlcTi1PSRdiqaxlgpEDp0bXvSl62ZKIceJNKvsbOPk0ou-asGmHmKsdaiaEW_xpwNwqoM97TBy-C5uqpHQUcL_Mmphe3z-NqUb1QnOGbc7YsKv2D1dMyHYW8iCUMEsk"),
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

func LoadPathMedia() *PART_MEDIA {
	return &PART_MEDIA{
		Path: GetEnv("PATH_MEDIA", "../../images/"),
	}
}

func GetEnv(key string, defaultVal string) string {
	if value, exits := os.LookupEnv(key); exits {
		return value
	}
	return defaultVal
}
