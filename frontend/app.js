import {Route, Routes} from "react-router-dom"

function app() {
    return (
    <>
    <Navbar />
    <div className="container">
        <Routes>
            <Route path="/Featured" element={<Featured />} />
            <Route path="/The Agenda Podcast" element={<The Agenda Podcast />} />
            <Route path="/About" element={<About />} />
        </Routes>
    </div>
    </>
    )
}

export default app;