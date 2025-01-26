const functions = require('@google-cloud/functions-framework');
const { v4: uuidv4 } = require('uuid');

functions.http('sendLineMessages', (req, res) => {
    const LINE_TOKEN = process.env.LINE_TOKEN;
    const LINE_USER_ID = process.env.LINE_USER_ID;

    fetch('https://api.line.me/v2/bot/message/push', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${LINE_TOKEN}`,
            'X-Line-Retry-Key': uuidv4(),
        },
        body: JSON.stringify({
            to: LINE_USER_ID,
            messages: req.body,
        }),
    }).then((r) => res.sendStatus(r.status));
});
