import { Link } from 'react-router-dom';
import {useState, useEffect} from 'react'
import SubscribeForm from "./Subscribe";
function Home() {
    //State
    const [articles, setArticles] = useState(null);
    useEffect(() => {
                fetch('http://localhost:8080/article')
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
                        <img src={'http://localhost:8080/image/'+i.imageID} /> {i.title}
                    </Link>
                </p>)
        })
    }
    if (articles== null) {
        return(
            <>
            <>Loading!</>
        <SubscribeForm />
                </>

    )
    }
    else {
        return (
            <>
                <p> Welcome to home page! </p>
                {output}
                <SubscribeForm />

            </>
        )
    }


}
export default Home;