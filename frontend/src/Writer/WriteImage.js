import Home from "../Home";
import {useState, useEffect} from 'react'
import {useLocation} from "react-router-dom";
function WriteImage(props) {
    const [image, setImage] = useState(null)

    const onImageChange = (event) => {
        if (event.target.files && event.target.files[0]) {
            setImage(URL.createObjectURL(event.target.files[0]));
        }
    }

    return (
        <div>
            <input type="file" onChange={onImageChange} className="filetype" />
            <img alt="preview image" src={image}/>
        </div>
    )
}

export default WriteImage;