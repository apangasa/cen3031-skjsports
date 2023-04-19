import Home from "../Home";
import {useState, useEffect} from 'react'
import {useLocation} from "react-router-dom";
import WriteImage from "./WriteImage";
import WriteText from "./WriteText";
import axios from 'axios';
function Writer(props) {


    const [image, setImage] = useState(null)
    const [draftState, setDraftState] = useState({loading:true, objects:[]})
    const token = localStorage.getItem("token");
    if (token) {
        setAuthToken(token);
    }
    let id = useLocation().state.id
    const handleImageAdd = (i) => {
        let temp = draftState.objects
        temp.splice(i, 0, {contentType: 'img'})
        setDraftState({loading:false, objects:temp})
    }
    function handleTextAdd(i)  {
        console.log("text add2")
        console.log(i)
        let temp = draftState.objects
        console.log(draftState)
        temp.splice(i, 0, {contentType: 'text'})
        setDraftState({loading:false, objects:temp})
    }
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
                <button onClick={() => {console.log("upload")}} > Upload</button>
                {(() => {
                    const objects = [];

                    for (let i = 0; i<length; i++) {
                        objects.push(< AddObject count={i} handleImageAdd={handleImageAdd} handleTextAdd={handleTextAdd}/>)
                        if (draftState.objects[i].contentType=='img')
                        objects.push(<WriteImage imageProps={draftState.objects[i].id}></WriteImage>);
                        else if(draftState.objects[i].contentType=='text')
                            objects.push(<WriteText textProps={draftState.objects[i].text}></WriteText>)
                    }

                    return objects;
                })()}
                < AddObject count={length} handleImageAdd={handleImageAdd} handleTextAdd={handleTextAdd}/>
            </div>
        )
    }
    else {
        console.log("empt")
        return (
            <div>
                <h2>Article Title: {id}
                    <input type="text"  className="title" /> </h2>
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

const setAuthToken = token => {
    if (token) {
        axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
    }
    else
        delete axios.defaults.headers.common["Authorization"];
}

function handleLogin(email, pass) {
//reqres registered sample user
    const loginPayload = {
        email: 'eve.holt@reqres.in',
        password: 'cityslicka'
    }

    axios.post("https://reqres.in/api/login", loginPayload)
        .then(response => {
            //get token from response
            const token = response.data.token;

            //set JWT token to local
            localStorage.setItem("token", token);

            //set token to axios common header
            setAuthToken(token);
    })
}
export default Writer;