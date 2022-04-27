from flask import Flask

app = Flask(__name__)

@app.route("/", methods=['GET'])
def happy():
    return "Hello from the other side", 200

app.run(host='0.0.0.0', port=8082)