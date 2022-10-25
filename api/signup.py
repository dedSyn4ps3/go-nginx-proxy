#!/usr/bin/env python3

import smtplib
from sys import argv
from email.message import EmailMessage
from email.headerregistry import Address

signup_email = argv[1]

msg = EmailMessage()
msg['Subject'] = "Email List Submission"
msg['From'] = Address("Info", "info", "awesome-it.com")
msg['To'] = Address("Some Client", "some_client22", "gmail.com")
msg.set_content("""\

A new user has signed up for our contact list:
<ul>
    <li><b> %s </b></li>
</ul>
""" % signup_email)

mail_server = smtplib.SMTP("some-smtp-server.com", port=587)
mail_server.starttls()

try:
    mail_server.login("some-user", "SomeSecretPassword!")
except Exception as e:
    print("Error while connecting to server: ", e)
    pass

mail_server.send_message(msg)