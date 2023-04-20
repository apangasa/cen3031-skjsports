import Home from "../Home";
import {useState, useEffect} from 'react'
import {useLocation} from "react-router-dom";
import WriteImage from "./WriteImage";
import WriteText from "./WriteText";
import axios from 'axios';
import setA from './auth'
function Writer(props) {


    const [image, setImage] = useState(null)
    const [draftState, setDraftState] = useState({loading:true, objects:[]})
    const token = localStorage.getItem("token");

    let id = useLocation().state.id
    const handleImageAdd = (i) => {
        console.log("text add2")
        console.log(i)
        let temp = draftState.objects
        console.log(draftState)
        temp.content.splice(i, 0, {contentType: 'img', 'img': ""})
        console.log(temp)

        setDraftState({loading:false, objects:temp})
    }
    function handleTextAdd(i)  {
        console.log("text add2")
        console.log(i)
        let temp = draftState.objects
        console.log(draftState)
        temp.content.splice(i, 0, {contentType: 'text', 'text': ""})
        console.log(temp)

        setDraftState({loading:false, objects:temp})
    }
    const handleSubmit = (i) => {
        let data = {'article_id':'4', 'author_email': draftState.objects.author_email, 'content': draftState.objects.content}
        console.log(data)
        fetch('http://localhost:8080/edit-draft', {
            method: "POST", // *GET, POST, PUT, DELETE, etc.
                referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
                body: JSON.stringify(data), // body data type must match "Content-Type" header
        });
    }

    useEffect(() => {
        console.log("effect")
        if (!draftState.loading) {

        }
        else if (draftState.objects.length > 0) {
        }
        else if (id==null) {
            setDraftState({loading:false, objects:[]});
        }
        else {
            fetch('http://localhost:8080/draft?id=4') // + id
                .then(response => response.json())
                .then(result => setDraftState({loading: false, objects:result}))
        }
        },[draftState])
    let length = draftState.objects
    if (length.content) {
        length = length.content.length
    }
    console.log("length")
    console.log(length)
    if (draftState.loading) {
        console.log("loading")
        return (
            <div>
                Loading Writer
            </div>
        )
    }


    else if (length > 0) {

        return (
            <div>
                    <p><h2>Article Title:
                        <input type="text"  className="title" value={draftState.objects.title}/> </h2> </p>
                {(() => {
                    const objects = [];

                    for (let i = 0; i<length; i++) {
                        objects.push(< AddObject count={i} handleImageAdd={handleImageAdd} handleTextAdd={handleTextAdd}/>)
                        if (draftState.objects.content[i].contentType=='img')
                        objects.push(<WriteImage imageProps={draftState.objects.content[i].id}></WriteImage>);
                        else if (draftState.objects.content[i].contentType=='text')
                            objects.push(<WriteText textProps={draftState.objects.content[i].text}></WriteText>)
                    }

                    return objects;
                })()}
                < AddObject count={length} handleImageAdd={handleImageAdd} handleTextAdd={handleTextAdd}/>
                <button onClick={handleSubmit} > Upload</button>

            </div>
        )
    }
    else {
        console.log("empt")
        return (
            <div>
                <p><h2>Article Title:
                    <input type="text"  className="title" value={draftState.objects.title}/> </h2> </p>
                < AddObject count={0} handleImageAdd={handleImageAdd} handleTextAdd={handleTextAdd}/>
            </div>
        )
    }
}

function AddObject(props) {
    console.log(props)
    let id = props.count
    const handleImageAdd = () => {
        props.handleImageAdd(id)
    }

    return (
        <div id={id}>
            <button onClick = {() => handleImageAdd(id)}>
                Add Image
            </button>
            <button onClick = {() => props.handleTextAdd(id)}>
                Add Text
            </button>
        </div>
    )
}



export default Writer;