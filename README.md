📄 Moisture Monitoring and Alerting System 🌱
📌 Overview 🚀
A simple and efficient Go program that monitors real-time soil moisture levels 📡 and automatically triggers a voice call alert 📞 using Twilio when moisture levels fall outside the safe range!
Perfect for smart farming, IoT gardens, and remote moisture tracking 🌾.

✨ Features 🎉
✅ Continuously monitors soil moisture via MQTT 📡

✅ Sends automatic voice call alerts using Twilio 📞

✅ Dynamic threshold updates through MQTT payloads ✏️

✅ Reliable MQTT reconnections handled smoothly 🔄

✅ Simple and modular codebase for easy customization 🛠️

🛠️ Tech Stack & Dependencies 💻
Built with:

🔹 Go – High-performance backend programming language 🚀

🔹 Paho MQTT – Lightweight MQTT client for Go 📡

🔹 Twilio Go SDK – For programmable voice call alerts 📞

🔹 JSON – For structured moisture data handling 📦

📌 Before Running the Program 💡
Ensure you have the following installed:

✅ Go (Check version: go version)

✅ Paho MQTT library:

sh
Copy
Edit
go get github.com/eclipse/paho.mqtt.golang
✅ Twilio Go SDK:

sh
Copy
Edit
go get github.com/twilio/twilio-go
⚙️ Configuration
-------------------->

📌 Modify the configuration variables in the code to include your own MQTT broker and Twilio details:

go
Copy
Edit
BrokerHost        = "your-mqtt-broker-host"
BrokerPort        = "your-mqtt-broker-port"
TwilioAccountSid  = "your-twilio-account-sid"
TwilioAuthToken   = "your-twilio-auth-token"
CallFrom          = "your-twilio-verified-number"
CallTo            = "recipient-phone-number"
📜 Example MQTT Payloads 📑
---------------------->

🎯 The program expects moisture values and threshold updates via MQTT in the following formats:


Action	Topic	Payload Example
🌱 Moisture Monitoring	moisture/monitor	{ "moist": 75 }
🛠️ Update Thresholds	moisture/update	{ "min": 30, "max": 80 }
⚠️ Notes 📝
-------------->

⚡ Always verify your MQTT broker connection before starting
⚡ Twilio phone numbers must be verified if using a trial account
⚡ Dynamic threshold updates help adjust moisture sensitivity easily
⚡ Reliable MQTT reconnect ensures continuous monitoring without manual resets

------>

🚀 Happy Monitoring! 🚀