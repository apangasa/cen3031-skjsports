import { useState } from "react";
import { Link } from "react-router-dom";
import { useEffect } from "react";
import { Navigate, useLocation } from "react-router-dom";
import { Input, Button } from "@chakra-ui/react";
function SearchBar() {
  let { pathname } = useLocation();
  const [returnQuery, setReturnQuery] = useState("");
  const [input, setInput] = useState("");

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
  let returnSize = returnQuery != "";
  console.log(returnSize);
  return (
    <>
      {returnSize && !pathname.includes("/search") ? (
        <Navigate to="/search" state={{ returnQuery }} replace={true} />
      ) : (
        <form action="/" method="GET">
          <label htmlFor="header-search">
            <span className="visually-hidden"></span>
          </label>
          <Input
            width="50%"
            variant="outline"
            onInput={(e) => setInput(e.target.value)}
            type="text"
            id="header-search"
            placeholder="Search: "
            name="s"
          />
          <Button width="50%" type="Search" onClick={handleClick}>
            Search
          </Button>
        </form>
      )}
    </>
  );
}

export default SearchBar;
