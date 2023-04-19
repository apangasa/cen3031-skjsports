import { useState } from "react";

function SearchBar() {
  const [returnQuery, setReturnQuery] = useState("");
  const [input, setInput] = useState("");

<<<<<<< Updated upstream
    const [returnQuery, setReturnQuery] = useState("");
    const [input, setInput] = useState("");

    function handleSubmit(e) {
        e.preventDefault();
        var requestOptions = {
            method: 'GET',
            redirect: 'follow'
        };
          
          fetch("http://localhost8080:/search?search=" + input, requestOptions)
          .then(response => response.json())
          .then(result => setReturnQuery(returnQuery))       
          .catch(error => console.log('error', error));
          
    const output = []
    console.log(returnQuery)
    for (let i in returnQuery) {
        if (returnQuery[i].contentType == 'img') {
            output.push(<img src={'http://localhost:8080/'+ returnQuery[i].id} />)
        }
        else {
            output.push(<p>{returnQuery[i].text}</p>)
        }
    }
    console.log(output)
    }

    
    return (

        <form action="/" method="GET">
=======
  const handleClick = (e) => {
    e.preventDefault();
    async function returnArticles() {
      const response = await fetch(
        "http://localhost:8080/search?search=" + input
      );
      const json = await response.json();
      setReturnQuery(json.results);
    }
    returnArticles();
  };

  let articleList = [];

  if (returnQuery != "") {
    console.log(returnQuery);
    articleList = returnQuery.map((article) => <li>{article.title}</li>);
  }

  return (
    <>
      {articleList}

      <form action="/" method="GET">
>>>>>>> Stashed changes
        <label htmlFor="header-search">
          <span className="visually-hidden"></span>
        </label>
        <input
          onInput={(e) => setInput(e.target.value)}
          type="text"
          id="header-search"
          placeholder="Search: "
          name="s"
        />
<<<<<<< Updated upstream
        <button type="Search" onClick={handleSubmit}>Search</button>
    </form>

    )

    
};
=======
        <button type="Search" onClick={handleClick}>
          Search
        </button>
      </form>
    </>
  );
}
>>>>>>> Stashed changes

export default SearchBar;
