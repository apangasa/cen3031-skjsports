import Home from "../Home";
import {useState, useEffect} from 'react'
import {useLocation} from "react-router-dom";
import WriteImage from "./WriteImage";
import WriteText from "./WriteText";

function Writer(props) {
    const [image, setImage] = useState(null)
    const [draftState, setDraftState] = useState({loading:true, objects:[]})
    let id = useLocation().state.id
    useEffect(() => {
        console.log("effect")
        if (!draftState.loading) {

        }
        else if (draftState.objects.length > 0) {
            setDraftState({loading:false, objects:draftState.objects})
        }
        else if (id==null) {
            setDraftState({loading:false, objects:[]});
        }
        else {
            fetch('http://localhost:5001/drafts/' + id)
                .then(response => response.json())
                .then(result => setDraftState({loading: false, objects:result.list}))
        }
        },[draftState])
    let length = draftState.objects.length
    if (draftState.loading) {
        console.log("loading")
        return (
            <div>
                Loading Writer
            </div>
        )
    }
    else if (draftState.objects.length > 0) {
        return (
            <div>
                {(() => {
                    const objects = [];

                    for (let i = 0; i<length; i++) {
                        if (draftState.objects[i].contentType=='img')
                        objects.push(<WriteImage imageProps={draftState.objects[i].id}></WriteImage>);
                        else if(draftState.objects[i].contentType=='text')
                            objects.push(<WriteText textProps={draftState.objects[i].text}></WriteText>)
                    }

                    return objects;
                })()}
            </div>
        )
    }
    else {
        console.log("empt")
        return (
            <div>
                <h2>Article Title: {id}</h2>
                <input type="file"  className="filetype" />
                <img alt="preview image" src={image}/>
            </div>
        )
    }
}

function AddObject() {
    return (
        <div>
            <button>
                Add Image
            </button>
            <button>
                Add Heading
            </button>
            <button>
                Add Text
            </button>
        </div>
    )
}


export default Writer;