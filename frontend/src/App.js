import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Routes, Route} from 'react-router-dom'
import Home from './Home.js'
import Article from './Post/Article.js'
function App() {
  return (
      <>
    <BrowserRouter>
        <Routes>
            <Route path={''} element={<Home/>}/>
            <Route path={'article'} element={<Article/>}/>
        </Routes>
    </BrowserRouter>

          </>
  );
}

export default App;
