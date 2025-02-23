package main

type User struct {
	name string //Private
	Age  uint   //Publice
}

func (u *User) SetName(name string) { //(u *User) *set value* สามารถเข้าถึงโครวสร้าง User ได้เข้าถึงฟิลด์  pointer receiver ช่วยให้ฟังก์ชันสามารถแก้ไขค่าของฟิลด์ในโครงสร้าง User ได้โดยตรง
	u.name = name
}

func (u User) GetName() string { //(u User) สำเนาไม่ใช่โครงสร้างต้นฉบับ *get value*
	return u.name
}
