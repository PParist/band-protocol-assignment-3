# Band-Protocol-Assignment-3 Transaction Broadcasting and Monitoring Client

## Overview
ในโจทย์ข้อที่ 3 นี้ เป็นการพัฒนาโมดูล Broadcasting and Monitoring ของธุรกรรม ซึ่งได้ทำการ Implement โดยใช้ภาษา Go และได้ทำการออกแบบด้วย Hexagonal Architecture ซึ่งเป็นการแยกส่วนของระบบออกจากกัน โดยมุ่งเน้นไปที่การออกแบบและสร้างระบบที่มีความพร้อมใช้งานสูงและง่ายต่อการทดสอบ

## Document
 [Module Spac](https://docs.google.com/document/d/1w71E__7h67h4-Iur-JQaDMmRyU0Y0_vKDA3zy6aBBpw/edit?usp=sharing)

## Requirements
- Go 1.21.5+

## Getting Started
### Installation and Running the Project
1. Clone the project from GitHub:
   ```sh
   git clone https://github.com/PParist/band-protocol-assignment-3.git
2. Installation Go:
   ```sh
   Download and install https://go.dev/doc/install
3. Running the Project:
   ```sh
   cd assignment_3
   go run .
4. Test:
   ```sh
   1.ใส่ Input ตามลำดับ Symbol,Price,Timestamp
   2.ใส่ hash ที่ได้รับทาง console
   2.พิม yes เพื่อดำเนินการต่อเมื่อได้รับ status PENDING

