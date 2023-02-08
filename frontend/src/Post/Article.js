import Home from "../Home";
import {useState, useEffect} from 'react'
import {useLocation} from "react-router-dom";
function Article(props) {
    const [data, setData] = useState(null);
    let articleID = useLocation().state
    console.log(articleID)
    if (articleID.articleID) {
        articleID = articleID.articleID
    }
    useEffect(() => {
            fetch('http://localhost:5001/article/' + articleID)
                .then(response => response.json())
                .then(result => setData(result.list))
        },[])
    const output = []
    console.log(data)
    for (let i in data) {
        if (data[i].contentType == 'img') {
            output.push(<img src={'http://localhost:5001/image/'+data[i].id} />)
        }
        else {
            output.push(<p>{data[i].text}</p>)
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