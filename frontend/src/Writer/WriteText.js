import Home from "../Home";
import {useState, useEffect} from 'react'
import { useLocation } from 'react-router-dom';

function WriteText(props) {
    const [text, setText] = useState(props.textProps)
    const onTextChange = (event) => {
        if (event.target.value) {
            setText(event.target.value);
        }
    }

    return (
        <div>
            <input type="text" onChange={onTextChange} className="textType" value={text} style={{width: "80%"}}/>
        </div>
    )
}

export default WriteText;