import { useState } from "react";
import Input from "./form/Input";
import { useNavigate, useOutletContext } from "react-router-dom";

const Login = () =>{
    const [email, setEmail] = useState("");
    const [passwd, setPasswd] = useState("");

    const {setJwtToken} = useOutletContext();
    const {setAlertClassName} = useOutletContext();
    const {setAlertMessage} = useOutletContext();
 
    const navigate = useNavigate();

    const handleSubmit = (event) =>{
        event.preventDefault();
        console.log("email/pass", email, passwd)
        if (email === "admin@example.com"){
            setJwtToken("ads")
            setAlertClassName("d-none")
            setAlertMessage("")
            navigate("/")
        }else{
            setAlertClassName("alert-danger")
            setAlertMessage("Invalid credentianal")
        }
    }
    return (
            <div>
                <h2>Login</h2>
                <hr />
                <form onSubmit={handleSubmit}>
                    <Input
                    title = "Email Address"
                    type = "email"
                    className = "form-control"
                    name = "email"
                    autoComplete = "email-new"
                    onChange ={(event)=> setEmail(event.target.value)}
                    ></Input>
                    <Input
                    title = "Password"
                    type = "password"
                    className = "form-control"
                    name = "password"
                    autoComplete = "password-new"
                    onChange ={(event)=> setPasswd(event.target.value)}
                    ></Input>
                    <hr />
                    <input
                        type = "submit"
                        className="btn btn-primary"
                        value= "Login"
                    />
                </form>
            </div>
    )
}

export default Login;