import React, { Component } from "react";
import axios from 'axios';
import './Login.css';

class Shop extends Component {
	constructor(props) {
		super(props);

		this.state = {
			items: []
		};
	}

	componentDidMount() {
		axios.get("/api/v1/shop").then(response => {
			console.log("response: ", response.data);
			this.setState({
				items: response.data
			})
		})
	}

	render() {
		const itemElements = [];

		this.state.items.forEach( item => {
			itemElements.push(
				<div className="item-container">
					<span className="item-title">{item.Name}</span>
					<span className="item-description">{item.Description}</span>
					<span className="item-price">{item.Price}</span>
				</div>);
		});

		return (
			<div className="shop-container">
				{ itemElements }
			</div>
		);
	}
}

export default Shop;
