#!/usr/bin/python3

import smtplib
from sys import argv
from email.message import EmailMessage
from email.headerregistry import Address

name = argv[1]
email = argv[2]
phone = argv[3]


msg = EmailMessage()
msg['Subject'] = "New Contact Form"
msg['From'] = Address("Info", "info", "awesome-it.com")
msg['To'] = Address("Some Client", "some_client22", "gmail.com")
msg.set_content("""\
<h1>[+] New Contact Form Submited [+]</h1>

<h3>Name: %s</h3>
<h3>Email: %s</h3>
<h3>Phone: %s</h3>
""" % (name, email, phone))

mail_server = smtplib.SMTP("some-smtp-server.com", port=587)
mail_server.starttls()

try:
    mail_server.login("some-user", "SomeSecretPassword!")
except Exception as e:
    print("Error while connecting to server: ", e)
    pass

mail_server.send_message(msg)