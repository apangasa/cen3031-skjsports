import Home from "../Home";
import {useState, useEffect} from 'react'
import {useLocation} from "react-router-dom";
function WriteImage(props) {
    const [imageState, setImageState] = useState({image: (props.imageProps ? 'http://localhost:5001/image/' + props.imageProps : ""), local: false} )

    const onImageChange = (event) => {
        if (event.target.files && event.target.files[0]) {
            setImageState({image:URL.createObjectURL(event.target.files[0]),local:true});
        }
    }

    return (
        <div>
            <input type="file" onChange={onImageChange} className="filetype" />
            Preview:
            <img src={imageState.image} />
         </div>
    )
}

export default WriteImage;