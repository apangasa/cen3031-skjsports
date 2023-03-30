import { useState } from "react";
import { Link } from "react-router-dom";
import { useEffect } from "react";

function SearchBar () {

    const [returnQuery, setReturnQuery] = useState("");
    const [input, setInput] = useState("");

    const handleClick = (e) => {
        e.preventDefault();
        async function returnArticles() {
            const response = await fetch("http://localhost:8080/search?search=" + input);
            const json = await response.json();
            setReturnQuery(json.results);
        }
        returnArticles();
    
    }
    
    let articleList = [];

    if (returnQuery != "") {
        console.log(returnQuery);
        articleList = returnQuery.map((article) =>
            <li>{article.title}</li>
        );
    }
 
    return (

        <>

        {articleList}
    
        <form action="/" method="GET">
        <label htmlFor="header-search">
            <span className="visually-hidden"></span>
        </label>
        <input
        onInput={e => setInput(e.target.value)}
            type="text"
            id="header-search"
            placeholder="Search: "
            name="s" 
        />
        <button type="Search" onClick={handleClick}>Search</button>
    </form>
    </>
    )
}; 

export default SearchBar;
