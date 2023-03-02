import logo from './logo.svg';
import './App.css';
import { MemoryRouter, Routes, Route} from 'react-router-dom'
import Home from './Home.js'
import Article from './Post/Article.js'
import Search from './Navbar/Searchbar';


function App() {

  return (
    <>
    <Search />
    <MemoryRouter>
        <Routes>
            <Route path={''} element={<Home/>}/>
            <Route path={'article'} element={<Article/>}/>
        </Routes>
    </MemoryRouter>
    </>
  );
  
}

export default App;
