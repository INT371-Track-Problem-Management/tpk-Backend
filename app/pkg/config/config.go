package config

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

func LoadDB() *Database {
	return &Database{
		Username: "dev",
		Password: "P@ssw0rd2",
		Host:     "172.18.0.2",
		Port:     "3306",
		Database: "project",
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
