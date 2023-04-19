import { Link } from 'react-router-dom';
import {useLocation} from "react-router-dom";

function SearchResults() {
    //State

    let articles = useLocation()

    if (articles.state) {
        articles = articles.state.returnQuery
    }
    console.log(articles)

    const output = []

    if (articles!= null) {

        articles.forEach((i,x) => {
            output.push(
                <p>
                    <Link to={{pathname:'/article'}}
                          state={{articleID: i.id}}>
                        <img src={'http://localhost:8080/image/'+i.imageID} /> {i.title}
                    </Link>
                </p>)
        })
    }
        return(
            <>
                {articles ? output :
                    <>Loading!</>
                }
            </>

        )

}
export default SearchResults;