import Home from "../Home";
import {useState, useEffect} from 'react'
import {Link} from "react-router-dom";
import AuthService from './auth.js';
import Login from './Login'
function Draftboard(props) {
    const [data, setData] = useState(null);
    console.log("DRAFTBOARD")
    useEffect(() => {
        fetch('http://localhost:8080/draftBoard/')
            .then(response => response.json())
            .then(result => setData(result.list))
    },[])
    const output = []
    console.log(data)
    for (let i in data) {
        {
            output.push(
            <p>
                <Link to={"/edit"}
                      state={{id: data[i].id}}>
                    {data[i].title}</Link>
            </p>
            )
        }
    }
    console.log(output)
    let user = AuthService.getCurrentUser()
    console.log("user")
    console.log(user)
    if (!user) {
        console.log("user")
        return (
            <Login />
        )
    }
    if (output.length == 0) {
        return(
            <>Loading!</>
        )
    }
    else {
        return (
            <>
                <h1>Drafts</h1>
                {output}

                <Link to={"/edit"}
                      state={{id:null}}>
                    Create New Article
                </Link>
            </>
        )
    }
}

export default Draftboard;