import React, { Component } from "react";
import axios from 'axios';
import './Login.css';

class Health extends Component {
	constructor(props) {
		super(props);

		this.state = {
			health: {}
		};
	}

	componentDidMount() {
		axios.get("/api/v1/healthfull").then(response => {
			console.log("response: ", response.data);
			this.setState({
				health: response.data
			})
		})
	}

	render() {
		return (
			<div className="health-container">
				<div className="item-container">
					<span className="item-name">WebStatus</span>
					<span className="item-description">{this.state.health.WebStatus}</span>
				</div>
				<div className="item-container">
					<span className="item-name">LoginStatus</span>
					<span className="item-description">{this.state.health.LoginStatus}</span>
				</div>
				<div className="item-container">
					<span className="item-name">ShopStatus</span>
					<span className="item-description">{this.state.health.ShopStatus}</span>
				</div>
			</div>
		);
	}
}

export default Health;
