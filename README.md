Human-Like LinkedIn Automation System

A Go-based Browser Automation Proof of Concept

📌 Overview

This repository contains a Go-based LinkedIn automation proof-of-concept built using the Rod browser automation library. The project demonstrates advanced browser automation techniques, realistic human-like interaction simulation, and anti-detection strategies implemented using a clean, modular Go architecture.

The system is designed to programmatically interact with LinkedIn in a way that closely mimics real user behavior, focusing on realism, timing variability, and structured automation design.


⚠️ Critical Disclaimer

-Educational & Technical Evaluation Purpose Only
-This project is created strictly for educational and technical assessment
-It demonstrates automation concepts in a controlled environment
-Automated interaction with LinkedIn may violate LinkedIn’s Terms of Service
-This code must not be used in production environments


🎯 Project Objective

-The primary objective of this assignment is to demonstrate:
-Advanced browser automation using Go
-Human-like behavior simulation to reduce automation signatures
-Clean and maintainable modular architecture
-Awareness and handling of bot-detection mechanisms
-Secure configuration and credential management
-This project evaluates both technical correctness and engineering quality.


🧩 Project Architecture

.
├── auth/                   # Authentication & session handling
├── search/                 # Profile search logic
├── connection/             # Connection request automation
├── messaging/              # Messaging logic
├── stealth/                # Human-like behavior simulation
├── config/                 # Environment & configuration management
├── utils/                  # Browser helpers & logging
├── database/               # Local state storage
├── cmd/
│ └── rod_example/         # Application entry point
├── .env.example           # Environment variable template
├── .gitignore
├── go.mod
└── README.md

⚙️ Environment Setup

-.env.example
LINKEDIN_EMAIL=your_email_here
LINKEDIN_PASSWORD=your_password_here
HEADLESS=false
SLOW_MODE_MS=800
MAX_DAILY_CONNECTIONS=15
LOG_LEVEL=info

⚙️ Setup Instructions

- Go 1.20 or higher
- Google Chrome (latest version)
- 
- A LinkedIn account (for testing purposes)
-Clone the Repository
git clone https://github.com/your-username/Human-Like-LinkedIn-Automation-System.git
cd Human-Like-LinkedIn-Automation-System

▶️ How to Run
-go mod tidy
-go run ./cmd/rod_example

▶️ **Watch the demo video:**  
https://drive.google.com/drive/folders/1vCbOEdirFY4WFf8q6FPWk0DbIWP7NVL4
