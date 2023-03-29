import { Link } from 'react-router-dom';
import {useState, useEffect} from 'react'

function Home() {
    //State
    const [articles, setArticles] = useState(null);
    useEffect(() => {
                fetch('http://localhost:5001/articleList')
                    .then(response => response.json())
                    .then(result => setArticles(result.list))

        },[])

    const output = []

    if (articles!= null) {

        articles.forEach((i,x) => {
            output.push(
                <p>
                    <Link to={{pathname:'/article'}}
                    state={{articleID: i.articleID}}>
                        <img src={'http://localhost:5001/image/'+i.imageID} /> {i.title}
                    </Link>
                </p>)
        })
    }
    if (articles== null) {
        return(
            <>Loading!</>
        )
    }
    else {
        return (
            <>
                <p> Welcome to home page! </p>
                {output}
            </>
        )
    }


}
export default Home;