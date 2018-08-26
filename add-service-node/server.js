const express = require('express');
const bodyParser = require('body-parser');
const app = express();
app.use(bodyParser.json());

app.post('/add', (req, res) => {
    const { num1, num2 } = req.body;
    res.status(200).json({ sum: num1 + num2 });
});

app.listen(3000, () => console.log('Add service listening on port 3000!')); // eslint-disable-line
