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
                title: "sohil is cool",
                articleID: '01'
            },
            {
                imageID: '2',
                title: "sumeet isn't cool",
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
                    type: "img",
                },
                {
                    type: 'text',
                    text: "Some random text from article 1 part 1.",
                },
                {
                    type: 'img',
                    id: "2",
                },
                {
                    type: 'text',
                    text: "Some random text from article 1 part 2.",
                }
            ]
        });
    }

});




app.get('/image/:imageID', (req, res)=>{
    console.log("request img:" + req.params.imageID)
    res.status(200);
    res.sendFile('/Users/kolliparas/Documents/GitHub/cen3031-skjsports/backendSim/'+req.params.imageID+'.png')

});

app.listen(PORT, (error) =>{
        if(!error)
            console.log("Server is Successfully Running,and App is listening on port "+ PORT)
    else
        console.log("Error occurred, server can't start", error);
    }
);

