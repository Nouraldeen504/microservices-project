from flask import Flask, jsonify
import requests

app = Flask(__name__)

@app.route('/api/rates', methods=['GET'])
def get_rates():
    # Fetch exchange rates from an external API
    response = requests.get('https://api.exchangerate-api.com/v4/latest/USD')
    data = response.json()
    return jsonify(data['rates'])

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)