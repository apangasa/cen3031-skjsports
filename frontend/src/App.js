import logo from './logo.svg';
import './App.css';
import Home from './Home.js'
import Article from './Post/Article.js'
import Search from './Navbar/Searchbar';
import Draftboard from "./Writer/Draftboard";
import Writer from "./Writer/Writer";
import SearchResults from "./Navbar/SearchResults";

import { BrowserRouter, MemoryRouter, Routes, Route } from "react-router-dom";
import Subscribe from "./Subscribe.js";

function App() {
  return (
    <>
    <BrowserRouter>
        <Search />

        <Routes>

            <Route path={'/search'} element={<SearchResults />}/>
            <Route path={'/'} element={<Home/>}/>
            <Route path={'/article'} element={<Article/>}/>
            <Route path={'/write'} element={<Draftboard/>}/>
            <Route path={'/edit'} element={<Writer/>}/>

        </Routes>
    </BrowserRouter>
    </>
  );
}

export default App;
