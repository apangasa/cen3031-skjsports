import Home from "../Home";
import {useState, useEffect} from 'react'
import {useLocation} from "react-router-dom";
function WriteImage(props) {
    const [text, setText] = useState("")
    useEffect(() => {
        fetch('http://localhost:5001/draftBoard/')
            .then(response => response.json())
            .then(result => setText(result.txt))
    },[])
    const onTextChange = (event) => {
        if (event.target.value) {
            setText(event.target.value);
        }
    }

    return (
        <div>
            <input type="text" onChange={onTextChange} className="textType" />
        </div>
    )
}

export default WriteImage;