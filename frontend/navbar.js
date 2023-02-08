import { Link } from "react-router-dom"

export default function Navbar() {
    return (
    <nav className="nav">
        <Link to="/" className="site-title">
            SKJ Sports
        </Link>
        <ul>
            <li>
                <Link to="/Featured">Featured </Link>
            </li>
            <li>
                <Link to="/The Agenda Podcast">Podcast </Link>
            </li>
            <li>
                <Link to="/About">About </Link>
            </li>
        </ul>
    </nav>
    )
}