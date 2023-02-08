import React from "react"
import ReactDOM from "react-dom/client";
import app from "./app"
import "./style.css"
import {BrowserRouter} from "react-router-dom"

const root = ReactDOM.createRoot(document.getElementById("root"))
root.render(
    <React.StrictMode>
        <BrowserRouter>
        <app />
        </BrowserRouter>
    </React.StrictMode>
)
