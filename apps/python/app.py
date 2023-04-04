from flask import Flask

app = Flask(__name__)


@app.route("/")
def hello():
    return '<h1 style="color: dodgerblue">Hello world from Python</h1>'
