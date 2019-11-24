import React from "react"
import {UserPanel} from "./UserPanel";
import _ from "lodash"

export class LoginPanel extends React.Component {
    constructor(props) {

        super(props)

        this.state = {
            username: '',
            password: '',
            status: null
        }
        this.loginHandler = this.loginHandler.bind(this)
    }

    render() {
        return (
            _.isEqual(this.state.status, "OK") ? <UserPanel/> : this.Login.call()
        );
    }

    Login = () => {
        return (
            <div>
                <div>Username: <input type="text" value={this.state.username} onChange={(e) => {
                    this.setState({username: e.target.value})
                }}/></div>
                <div>Password: <input type="text" value={this.state.password} onChange={(e) => {
                    this.setState({password: e.target.value})
                }}/></div>
                <div>
                    <button onClick={this.loginHandler}>Login</button>
                </div>
            </div>
        );
    };

    loginHandler() {
        console.log("Username:", this.state.username)
        console.log("Password:", this.state.password)

        if (_.isEqual("root", this.state.username) && _.isEqual("1234", this.state.password)) {
            this.setState({status: "OK"})
        }
    }

}