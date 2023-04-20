import { Link } from "react-router-dom";
import { useState, useEffect } from "react";
import SubscribeForm from "./Subscribe";
import { Container, Heading, Box, Text, Center } from "@chakra-ui/react";
function Home() {
  //State
  const [articles, setArticles] = useState(null);
  useEffect(() => {
    fetch("http://localhost:8080/articles")
      .then((response) => response.json())
      .then((result) => setArticles(result.results));
  }, []);

  const output = [];

  if (articles != null) {
    articles.forEach((i, x) => {
      output.push(
        <Box>
          <Text>
            <Link to={{ pathname: "/article" }} state={{ articleID: i.id }}>
              <img src={"http://localhost:8080/image/" + i.imageID} /> {i.title}
            </Link>
          </Text>
        </Box>
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
        <Container>
          <Center>
            <Heading my="20px" fontSize="xl">
              SKJ Sports
            </Heading>
          </Center>
          <Center>
            <Text color="blue.300" fontSize="xl">
              Welcome to SKJ Sports!
            </Text>
          </Center>
        </Container>
        {output}
        <SubscribeForm />
      </>
    );
  }
}
export default Home;
