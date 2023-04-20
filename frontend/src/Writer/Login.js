import {useState, useEffect} from 'react'
import AuthService from "./auth.js";

function Login(props) {

    let [user, setUser] = useState({
        username: "",
        password: "",
        loading: false,
        message: ""
    })

    const onChangeUsername = (e) =>
    {
        setUser({
            username: e.target.value,
            password: user.password,
            loading: false,
            message: user.message
        });
    }
    const onChangePassword = (e) => {
        setUser({
            username: user.username,
            password: e.target.value,
            loading: false,
            message: user.message        });
    }

    const handleLogin = (e) => {
        e.preventDefault();


        console.log("login")
        AuthService.login(user.username, user.password).then(
            () => {
                this.props.router.navigate("/write");
                window.location.reload();
            },
            error => {
                const resMessage =
                    (error.response &&
                        error.response.data &&
                        error.response.data.message) ||
                    error.message ||
                    error.toString();

                setUser({
                    username: user.username,
                    password: user.password,
                    loading: false,
                    message: resMessage
                });
            }
        );

        setUser({
            message: "",
            loading: true
        });
    }
    return (
        <div className="col-md-12">
            <div className="card card-container">

                <button
                    onClick={handleLogin}

                > login </button>
                    <div className="form-group">
                        <label htmlFor="username">Username</label>
                        <input
                            type="text"
                            className="form-control"
                            name="username"
                            onChange={onChangeUsername}
                        />
                    </div>
                    <div className="form-group">
                        <label htmlFor="password">Password</label>
                        <input
                            type="password"
                            className="form-control"
                            name="password"
                            onChange={onChangePassword}
                        />
                    </div>


                    {user.message && (
                        <div className="form-group">
                            <div className="alert alert-danger" role="alert">
                                {user.message}
                            </div>
                        </div>
                        )}
            </div>
        </div>
    );

}




export default Login;