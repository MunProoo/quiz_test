package main

import (
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

type ExamResult struct {
	Name        string    `json:"name"`
	Score       int       `json:"score"`
	PopularName string    `json:"popularName"`
	Timestamp   time.Time `json:"timestamp"`
}

const (
	// SMTP 설정
	smtpHost     = "smtp.gmail.com"
	smtpPort     = 587
	senderEmail  = "mjy5178@gmail.com"   // 발신자 이메일 주소
	senderPasswd = "vofr tfmn dosz drqg" // Gmail 앱 비밀번호
	targetEmail  = "mjy5178@gmail.com"   // 결과를 받을 이메일 주소
)

func sendEmail(result ExamResult) error {
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", targetEmail)
	m.SetHeader("Subject", fmt.Sprintf("[취저 모의고사] %s님 : %d점", result.Name, result.Score))

	// 이메일 본문 작성
	body := retMailBody(result)

	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPasswd)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func retMailBody(result ExamResult) (body string) {
	body = fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<style>
			.email-container {
				max-width: 600px;
				margin: 0 auto;
				padding: 30px;
				font-family: 'Arial', sans-serif;
				background-color: #f8f9fa;
				border-radius: 10px;
			}
			.header {
				text-align: center;
				padding: 20px;
				background-color: #ffffff;
				border-radius: 8px;
				box-shadow: 0 2px 4px rgba(0,0,0,0.1);
				margin-bottom: 25px;
			}
			.header h2 {
				color: #1a73e8;
				margin: 0;
				font-size: 24px;
				border-bottom: 2px solid #e8eaed;
				padding-bottom: 15px;
			}
			.content {
				background-color: #ffffff;
				padding: 25px;
				border-radius: 8px;
				box-shadow: 0 2px 4px rgba(0,0,0,0.1);
			}
			.result-row {
				margin: 15px 0;
				padding: 10px;
				border-bottom: 1px solid #e8eaed;
			}
			.label {
				color: #5f6368;
				font-size: 16px;
				margin-right: 10px;
			}
			.value {
				color: #1a73e8;
				font-weight: bold;
				font-size: 16px;
			}
			.timestamp {
				text-align: right;
				color: #80868b;
				font-size: 14px;
				margin-top: 20px;
			}
		</style>
	</head>
	<body>
		<div class="email-container">
			<div class="header">
				<h2>📋 취저 모의고사 결과</h2>
			</div>
			<div class="content">
				<div class="result-row">
					<span class="label">응시자:</span>
					<span class="value">%s</span>
				</div>
				<div class="result-row">
					<span class="label">이 사람의 좋아요 픽:</span>
					<span class="value">%s</span>
				</div>
				<div class="result-row">
					<span class="label">점수:</span>
					<span class="value">%d점</span>
				</div>
				<div class="timestamp">
					제출 시간: %s
				</div>
			</div>
		</div>
	</body>
	</html>
	`, result.Name, result.PopularName, result.Score, result.Timestamp.Format("2006-01-02 15:04:05"))

	return
}
