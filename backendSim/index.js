const express = require('express');
var cors = require('cors')
const app = express()

app.use(cors())
const PORT = 5001;

app.get('/articleList', (req, res)=>{
    res.status(200);
    res.send({
        list: [
            {
                imageID: '1',
                title: "EXAMPLE TITLE - My 2022/2023 Premier League Predictions: ",
                articleID: '01'
            },
            {
                imageID: '2',
                title: "EXAMPLE TITLE - Who Will Win Super Bowl 57?",
                articleID: '02'
            }
        ]

    });
});

app.get('/article/:articleID', (req, res)=>{
    res.status(200);
    if (req.params.articleID=='01') {
        res.send({
            list: [
                {
                    id: '1',
                    contentType: "img",
                },
                {
                    contentType: 'text',
                    text: "EXAMPLE TITLE - My 2022/2023 Premier League Predictions: ",
                },
                {
                    contentType: 'img',
                    id: "2",
                },
                {
                    contentType: 'text',
                    text: "EXAMPLE TITLE - Who Will Win Super Bowl 57?",
                }
            ]
        });
    }

});




app.get('/image/:imageID', (req, res)=>{
    console.log("request img:" + req.params.imageID)
    res.status(200);
    res.sendFile('/Users/sumeetjena/Documents/GitHub/cen3031-skjsports/backendSim/'+req.params.imageID+'.png')

});

app.listen(PORT, (error) =>{
        if(!error)
            console.log("Server is Successfully Running, and App is listening on port "+ PORT)
    else
        console.log("Error occurred, server can't start", error);
    }
);

