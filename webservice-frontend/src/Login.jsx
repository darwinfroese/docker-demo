import React, { Component } from "react";
import { Link } from "react-router-dom";
import './Login.css';

class Login extends Component {
	render () {
		return (
			<div className="login-container">
				<input className="username-tb" type="text" placeholder="username"></input>
				<input className="password-tb" type="password" placeholder="password"></input>
				<Link to="/shop" className="login-button">Login</Link>
			</div>
		);
	}
}

export default Login;
