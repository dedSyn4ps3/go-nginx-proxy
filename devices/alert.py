#!/usr/bin/python3

import smtplib
from sys import argv
from email.message import EmailMessage
from email.headerregistry import Address

triggered_device = argv[1]


msg = EmailMessage()
msg['Subject'] = "Device Alert"
msg['From'] = Address("Info", "info", "awesome-it.com")
msg['To'] = Address("Some Client", "some_client22", "gmail.com")
msg.set_content("""\
<h1>[!] %s Alert [!]</h1>

<h3>** Check device dashboard **</h3>


This alert has been sent because it meets the criteria of a moderate/severe 
resource issue per your guidelines. If you feel this alert has been made in error, 
or should be adjusted, please reach out to our support staff at support@awesomeinc.com
""" % triggered_device)

mail_server = smtplib.SMTP("some-smtp-server.com", port=587)
mail_server.starttls()

try:
    mail_server.login("some-user", "SomeSecretPassword!")
except Exception as e:
    print("Error while connecting to server: ", e)
    pass

mail_server.send_message(msg)