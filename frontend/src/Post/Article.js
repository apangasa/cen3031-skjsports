import Home from "../Home";
import {useState, useEffect} from 'react'
import {useLocation} from "react-router-dom";
function Article(props) {
    console.log("article")
    const [data, setData] = useState({'content': []});
    let articleID = useLocation().state
    console.log(articleID)
    if (articleID.articleID) {
        articleID = articleID.articleID
    }
    useEffect(() => {
            fetch('http://localhost:8080/article?id=' + articleID)
                .then(response => response.json())
                .then(result => setData(result))
        },[])
    const output = []
    console.log(data)
    for (let i = 0; i<data.content.length; i++) {
        if (data.content[i].contentType == 'img') {
            output.push(<img src={'http://localhost:8080/image/'+data.content[i].id} />)
        }
        else {
            output.push(<p>{data.content[i].text}</p>)
        }
    }
    console.log(output)

    if (output.length == 0) {
        return(
            <>Loading!</>
        )
    }
    else {
        return (
            <>
                {output}
            </>
        )
    }


}

export default Article;