import Home from "../Home";
import { useState, useEffect } from "react";
import { useLocation } from "react-router-dom";
import { Container, Heading, Box, Text, Center } from "@chakra-ui/react";
function Article(props) {
  console.log("article");
  const [data, setData] = useState({ content: [] });
  let articleID = useLocation().state;
  console.log(articleID);
  if (articleID.articleID) {
    articleID = articleID.articleID;
  }
  useEffect(() => {
    fetch("http://localhost:8080/article?id=" + articleID)
      .then((response) => response.json())
      .then((result) => setData(result));
  }, []);
  const output = [];
  console.log(data);
  for (let i = 0; i < data.content.length; i++) {
    if (data.content[i].contentType == "img") {
      output.push(
        <Center>
          <img src={"http://localhost:8080/image/" + data.content[i].id} />
        </Center>
      );
    } else {
      output.push(
        <Center>
          <Text>{data.content[i].text}</Text>{" "}
        </Center>
      );
    }
  }
  console.log(output);

  if (output.length == 0) {
    return <>Loading!</>;
  } else {
    return (
      <>
        <Center>
          <Heading> {data.title} </Heading>{" "}
        </Center>

        {output}
      </>
    );
  }
}

export default Article;
