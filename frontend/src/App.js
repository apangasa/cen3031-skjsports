<<<<<<< Updated upstream
import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Routes, Route} from 'react-router-dom'
import Home from './Home.js'
import Article from './Post/Article.js'
import Search from './Navbar/Searchbar';

=======
import logo from "./logo.svg";
import Button from "@mui/material/Button";
import "./App.css";
import Home from "./Home.js";
import Article from "./Post/Article.js";
import Search from "./Navbar/Searchbar";
import Draftboard from "./Writer/Draftboard";
import Writer from "./Writer/Writer";
import { BrowserRouter, MemoryRouter, Routes, Route } from "react-router-dom";
import Subscribe from "./Subscribe.js";
>>>>>>> Stashed changes

function App() {

  return (
    <>
<<<<<<< Updated upstream
    <Search />
    <BrowserRouter>
        <Routes>
            <Route path={''} element={<Home/>}/>
            <Route path={'article'} element={<Article/>}/>
=======
      <BrowserRouter>
        <Search />
        <Routes>
          <Route path={"/"} element={<Home />} />
          <Route path={"/article"} element={<Article />} />
          <Route path={"/write"} element={<Draftboard />} />
          <Route path={"/edit"} element={<Writer />} />
>>>>>>> Stashed changes
        </Routes>
      </BrowserRouter>
    </>
  );
  
}

export default App;
