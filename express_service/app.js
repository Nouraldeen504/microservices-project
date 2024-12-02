const express = require('express');
const axios = require('axios');
const path = require('path');
const app = express();
const PORT = 3000;

// Serve static files from the public directory
app.use(express.static(path.join(__dirname, 'public')));

// Route to fetch rates from Flask service
app.get('/api/rates', async (req, res) => {
    try {
        const ratesResponse = await axios.get('http://flask_service:5000/api/rates');
        res.json(ratesResponse.data);
    } catch (error) {
        res.status(500).json({ error: 'Error fetching rates' });
    }
});

// Route to convert currency using Go service
app.get('/api/convert', async (req, res) => {
    const { amount, from, to } = req.query;
    try {
        const convertResponse = await axios.get(`http://go_service:8080/api/convert?amount=${amount}&from=${from}&to=${to}`);
        res.json(convertResponse.data);
    } catch (error) {
        res.status(500).json({ error: 'Error converting currency' });
    }
});

app.listen(PORT, () => {
    console.log(`Express service running on http://localhost:${PORT}`);
});