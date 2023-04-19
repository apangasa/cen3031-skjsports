import {setState} from 'react'
import AuthService from "auth.js";

function Login(props) {

    let [user, setUser] = setState({
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

}