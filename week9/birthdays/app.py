import os

from flask import Flask, flash, jsonify, redirect, render_template, request, session
from cs50 import SQL

# Configure application
app = Flask(__name__)

# Ensure templates are auto-reloaded
app.config["TEMPLATES_AUTO_RELOAD"] = True

# Configure CS50 Library to use SQLite database
db = SQL("sqlite:///birthdays.db")


@app.after_request
def after_request(response):
    """Ensure responses aren't cached"""
    response.headers["Cache-Control"] = "no-cache, no-store, must-revalidate"
    response.headers["Expires"] = 0
    response.headers["Pragma"] = "no-cache"
    return response


@app.route("/", methods=["GET", "POST"])
def index():
    if request.method == "POST":
        name = request.form.get("name")
        month = request.form.get("month")
        day = request.form.get("day")

        if is_valid_birthday(name, month, day):
            db.execute("INSERT INTO birthdays (name, month, day) VALUES (?, ?, ?)", name, month, day)

        return redirect("/")

    else:
        birthdays = db.execute("SELECT * FROM birthdays")

        return render_template("index.html", birthdays=birthdays)

def is_valid_birthday(name: str, month: str, day: str) -> bool:
    if len(name) < 2:
        return False

    try:
        m = int(month)
        if m < 1 or m > 12:
            return False
        d = int(day)
        if d < 1 or d > 31:
            return False
    except ValueError:
        return False

    return True

app.run(host="0.0.0.0", port=5000, debug=True)
