import React from "react";
import clientConfig from "../config/client"

export class UserPanel extends React.Component {

    constructor(props) {
        super(props)

        this.state = {
            user: {
                name: '',
                department: '',
                email: '',
                username: '',
                password: ''
            },
            roaster: []
        }
    }

    GetUser(user) {
        return {
            name: user.name,
            department: user.department,
            email: user.email,
            username: user.username
        };
    }

    componentDidMount() {
        console.log("componentDidMount--->")
        clientConfig.get('users').then((res) => {
            console.log(res)

            const {roaster} = {...this.state}
            res.data.map(usr => {
                roaster.push(this.GetUser(usr))
                console.log(roaster)
            })
            this.setState({roaster: roaster})
        }).catch((err) => {
            console.log(err)
        })
    }

    render() {
        return (
            <div>
                <div>{this.UserTable.call()}</div>
                <div><br/>
                    <hr/>
                    <br/></div>
                <div>{this.UserForm.call()}</div>
            </div>
        );
    }


    UserForm = () => {
        return (
            <div>

                <div align="justify"> Name: <input type="text" value={this.state.user.name} onChange={(e) => {
                    const {user} = {...this.state}
                    const current = user
                    current['name'] = e.target.value
                    this.setState({user: current})
                }}/></div>

                <div align="justify"> Department: <input type="text" value={this.state.department} onChange={(e) => {
                    const {user} = {...this.state}
                    const current = user
                    current['department'] = e.target.value
                    this.setState({user: current})
                }}/></div>

                <div align="justify"> E-mail: <input type="text" value={this.state.email} onChange={(e) => {
                    const {user} = {...this.state}
                    const current = user
                    current['email'] = e.target.value
                    this.setState({user: current})
                }}/></div>

                <div align="justify"> Username: <input type="text" value={this.state.username} onChange={(e) => {
                    const {user} = {...this.state}
                    const current = user
                    current['username'] = e.target.value
                    this.setState({user: current})
                }}/></div>

                <div align="justify"> Password: <input type="password" value={this.state.password} onChange={(e) => {
                    const {user} = {...this.state}
                    const current = user
                    current['password'] = e.target.value
                    this.setState({user: current})
                }}/></div>


                <div>
                    <button onClick={this.registerHandler}>Register</button>
                </div>

            </div>
        );
    }

    UserTable = () => {
        return (
            <div>
                <table border="1">
                    <tbody>
                    <tr>
                        <td>Name</td>
                        <td>Department</td>
                        <td>E-Mail</td>
                        <td>Username</td>
                    </tr>
                    {
                        this.state.roaster.map((usr) => {
                            return (
                                <tr>
                                    <td>{usr.name}</td>
                                    <td>{usr.department}</td>
                                    <td>{usr.email}</td>
                                    <td>{usr.username}</td>
                                </tr>
                            );
                        })
                    }
                    </tbody>
                </table>
            </div>
        );
    }

    registerHandler = () => {

        const {roaster} = {...this.state}

        roaster.push(this.state.user)
        clientConfig.post("users", this.state.user).then((res) => {
            console.log(res)
        }).catch((err) => {
            console.log(err)
        })
        console.log(roaster)
        this.setState({roaster: roaster})
    }
}
