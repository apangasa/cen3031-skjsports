<<<<<<< Updated upstream
import { Link } from 'react-router-dom';
import {useState, useEffect} from 'react'
=======
import { Link } from "react-router-dom";
import { Button, ButtonGroup } from "@mui/material";
import { useState, useEffect } from "react";
import SubscribeForm from "./Subscribe";
>>>>>>> Stashed changes

function Home() {
  //State
  const [articles, setArticles] = useState(null);
  useEffect(() => {
    fetch("http://localhost:5001/articleList")
      .then((response) => response.json())
      .then((result) => setArticles(result.list));
  }, []);

  const output = [];

<<<<<<< Updated upstream
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
                Welcome to home page!
                {output}
            </>
        )
    }

=======
  if (articles != null) {
    articles.forEach((i, x) => {
      output.push(
        <p>
          <Link
            to={{ pathname: "/article" }}
            state={{ articleID: i.articleID }}
          >
            <img src={"http://localhost:5001/image/" + i.imageID} /> {i.title}
          </Link>
        </p>
      );
    });
  }
  if (articles == null) {
    return (
      <>
        <>Loading!</>
        <SubscribeForm />
      </>
    );
  } else {
    return (
      <>
        <p> SKJ Sports </p>
>>>>>>> Stashed changes

        {output}
        <SubscribeForm />
      </>
    );
  }
}
export default Home;
