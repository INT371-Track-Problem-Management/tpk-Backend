package constants

const (
	CHECK_FALSE = false
	CHECK_TRUE  = true
)

const (
	SUBJECT_EMAIL_SENDING_REPORT = "ส่งการรายงงานปัญหาเรียบร้อยแล้ว"
	SUBJECT_EMAIL_STATUS_REPORT  = "แจ้งเตือนการเปลี่ยนสถานะรายงาน"
	SUBJECT_EMAIL_END_REPORT     = "รายงานของคุณได้รับการแก้ไขเรียบร้อยแล้ว"
	SUBJECT_EMAIL_CANCEL_REPORT  = "ยกเลิกรายการ"

	BODY_EMAIL_SENDING_REPORT        = "กรุณารอระบบตรวจสอบและรอรับเรื่อง"
	BODY_EMAIL_APPROVE_REPORT        = "ระบบได้ตรวจสอบและทำการรับเรื่องเรียบร้อยแล้ว"
	BODY_EMAIL_ENAGAGE_REPORT        = "กรุณานัดวันเข้าซ่อมเพื่อทำการตรวจสอบและแก้ไขปัญหา"
	BODY_EMAIL_WAITE_REPORT          = "ระบบได้ทำการนัดวันเข้าซ่อมเรียร้อยกรุณารอวันดำเนินการและแก้ไข"
	BODY_EMAIL_PLAN_FIX_REPORT       = "แจ้งเตือนนัดเข้าซ่อม/แก้ไขปัญหา(ว/ป/ด)(เวลา)"
	BODY_EMAIL_POSTPONE_REPORT       = "(ชื่อนามสกุล)ได้ทำการเลื่อนนัดวันเข้าซ่อมจาก(วดป)(เวลา) เป็น (วดป)(เวลา)"
	BODY_EMAIL_WAITE_POSTPONE_REPORT = "ทำการส่งเรื่องเข้าระบบเรียบร้อย รอตรวจสอบและรอยืนยันการเลื่อนนัด"
	BODY_EMAIL_CANCEL_REPORT         = "ระบบได้ทำการยกเลิกการรายงานปัญหาเรียบร้อย"
	BODY_EMAIL_ENDJOB_REPORT         = "การรายงานปัญหาเสร็จสิ้น ระบบได้บันทึกผลการรีวิวเสร็จสิ้นขอบคุณที่ใช้บริการ"
)
